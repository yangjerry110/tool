/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-14 18:11:34
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-14 18:26:58
 * @Description: gRPC client implementation
 */
package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

// client struct is used to manage the gRPC client connection pool.
// It ensures that the connection pool is initialized only once and provides methods to get the pool.
type client struct {
	// instance holds the reference to the client pool.
	instance *clientPool
	// once is used to ensure that the initialization of the client pool happens only once.
	once sync.Once
	// poolMu is a read-write mutex to protect concurrent access to the client pool during initialization.
	poolMu sync.RWMutex
}

// grpcClient is a singleton instance of the client struct, used to manage the gRPC client pool.
var grpcClient = &client{}

// GetGrpcClientPool is a public function that serves as an entry point to get the gRPC client pool.
// It calls the internal method getGrpcClientPool of the grpcClient instance.
// ctx: The context used for the operation, which can be used for cancellation or timeout.
// cfg: The configuration for the gRPC client pool.
// Returns a pointer to the client pool and an error if any issues occur during the process.
func GetGrpcClientPool(ctx context.Context, cfg *Config) (*clientPool, error) {
	return grpcClient.getGrpcClientPool(ctx, cfg)
}

// getGrpcClientPool initializes the gRPC client pool if it hasn't been initialized yet.
// It sets default values for the configuration, creates connections to all endpoints, and performs a health check on the pool.
// ctx: The context used for the operation, which can be used for cancellation or timeout.
// cfg: The configuration for the gRPC client pool.
// Returns a pointer to the client pool and an error if any issues occur during the process.
func (c *client) getGrpcClientPool(ctx context.Context, cfg *Config) (*clientPool, error) {
	// Variable to hold any initialization errors.
	var initErr error

	// Use sync.Once to ensure that the initialization code runs only once.
	c.once.Do(func() {
		// Lock the mutex to prevent concurrent access during initialization.
		c.poolMu.Lock()
		// Unlock the mutex when the function returns.
		defer c.poolMu.Unlock()

		// Set default values for the configuration if they are not provided.
		c.setDefaults(cfg)

		// Initialize a slice to hold the gRPC client connections.
		clients := make([]*grpc.ClientConn, 0, len(cfg.Endpoints))
		// Iterate over each endpoint in the configuration.
		for _, endpoint := range cfg.Endpoints {
			// Create a gRPC connection to the current endpoint.
			conn, err := c.createConnection(ctx, cfg, endpoint)
			if err != nil {
				// If there is an error creating the connection, record the error.
				initErr = fmt.Errorf("failed to create connection to %s: %w", endpoint, err)
				// Close all the connections that have been created so far.
				for _, c := range clients {
					c.Close()
				}
				return
			}
			// Add the newly created connection to the slice of clients.
			clients = append(clients, conn)
		}

		// Create a new client pool with the initialized connections and the configuration.
		c.instance = &clientPool{
			clients: clients,
			config:  cfg,
		}
	})

	// If there was an initialization error, return it.
	if initErr != nil {
		return nil, initErr
	}

	// Perform a health check on the client pool.
	if err := c.instance.healthCheck(ctx); err != nil {
		return nil, fmt.Errorf("connection health check failed: %w", err)
	}

	// Return the initialized client pool.
	return c.instance, nil
}

// createConnection creates a single gRPC connection to the specified endpoint.
// It sets a timeout for the connection attempt and configures the load balancing and health check settings.
// ctx: The context used for the operation, which can be used for cancellation or timeout.
// cfg: The configuration for the gRPC client.
// endpoint: The target endpoint to connect to.
// Returns a pointer to the gRPC client connection and an error if any issues occur during the process.
func (c *client) createConnection(ctx context.Context, cfg *Config, endpoint string) (*grpc.ClientConn, error) {
	// Create a new context with a timeout for the connection attempt.
	dialCtx, cancel := context.WithTimeout(ctx, cfg.DialTimeout)
	// Cancel the context when the function returns to release resources.
	defer cancel()

	// Construct the service configuration JSON string with load balancing and health check settings.
	serviceConfig := fmt.Sprintf(`{
		 "loadBalancingPolicy": "%s",
		 "healthCheckConfig": {
			 "serviceName": "%s"
		 }
	 }`, cfg.LBPolicy, cfg.ServiceName)

	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))

	// Create a new gRPC client connection with the specified options.
	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig),
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(dialCtx, "tcp", addr)
		}),
	)
	if err != nil {
		// If there is an error creating the connection, return a formatted error message.
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	// Return the newly created gRPC client connection.
	return conn, nil
}

// setDefaults sets default values for the configuration if they are not provided.
// cfg: The configuration for the gRPC client.
func (c *client) setDefaults(cfg *Config) {
	// Set the dial timeout to 30 seconds if it is not provided.
	if cfg.DialTimeout == 0 {
		cfg.DialTimeout = 30 * time.Second
	}
	// Set the load balancing policy to "round_robin" if it is not provided.
	if cfg.LBPolicy == "" {
		cfg.LBPolicy = "round_robin"
	}
	// Set the maximum number of retries to 3 if it is not provided.
	if cfg.MaxRetries == 0 {
		cfg.MaxRetries = 3
	}
}
