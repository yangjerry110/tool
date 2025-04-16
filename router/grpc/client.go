/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:12:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 15:42:54
 * @Description:
 */
package grpc

import (
	"log"
	"sync"
	"time"

	// Import custom error handling package
	// This package likely contains custom error types and handling logic for the gRPC client pool.
	"github.com/yangjerry110/tool/router/grpc/internal/errors"
	"github.com/yangjerry110/tool/toolerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

// Default values for various parameters of the client pool.
const (
	// defaultPoolSize defines the default number of connections in the pool.
	defaultPoolSize = 3
	// defaultHealthCheckInterval sets the default interval for health checks.
	defaultHealthCheckInterval = 30 * time.Second
	// defaultDialTimeout specifies the default timeout for establishing a connection.
	defaultDialTimeout = 5 * time.Second
	// defaultHealthCheckTimeout sets the default timeout for performing a health check.
	defaultHealthCheckTimeout = 2 * time.Second
)

// clientPoolMap is a global variable that stores multiple client pool instances.
// It uses sync.Map to provide concurrent-safe access to the client pool instances,
// with the service name as the key and the corresponding poolEntry as the value.
var clientPoolMap sync.Map

// poolEntry is a struct used to store the initialization status and instance of a client pool.
// It contains a sync.Once object to ensure that the client pool is initialized only once,
// an instance pointer to the client pool, and an error variable to store any errors during initialization.
type poolEntry struct {
	// once ensures that the initialization of the client pool is done only once.
	once sync.Once
	// instance holds the actual client pool instance.
	instance *clientPool
	// err stores any error that occurred during the initialization of the client pool.
	err error
}

// Init initializes the client pool for a specific service.
// It uses the sync.Map to store and manage the client pool instances for different services.
// The sync.Once mechanism in poolEntry ensures that each service's client pool is initialized only once.
// If an error occurs during initialization, it logs the error message.
//
// @param serviceName: A string representing the name of the gRPC service.
// @param endpoints: A slice of strings representing the endpoints of the gRPC service.
// @param opts: Variadic parameter of type Option, which are used to configure the client pool.
func Init(serviceName string, endpoints []string, opts ...Option) {
	// Load or store a poolEntry for the given service name in the clientPoolMap.
	entry, _ := clientPoolMap.LoadOrStore(serviceName, &poolEntry{})
	// Type assertion to get the poolEntry.
	e := entry.(*poolEntry)

	// Ensure that the initialization of the client pool is done only once.
	e.once.Do(func() {
		// Create a new client pool instance.
		instance, err := newClientPool(serviceName, endpoints, opts...)
		if err != nil {
			// Log the error if the initialization fails.
			log.Printf("grpcpool init failed for %s: %v", serviceName, err)
			e.err = err
			return
		}
		// Store the successfully initialized client pool instance.
		e.instance = instance
	})
}

// GetClientPool retrieves the client pool instance for a specific service.
// It first checks if the client pool for the given service exists in the clientPoolMap.
// If it exists, it checks if there was an initialization error or if the instance is nil.
// If there are no issues, it returns the client pool instance; otherwise, it returns an appropriate error.
//
// @param serviceName: A string representing the name of the gRPC service.
// @return: A pointer to the clientPool if it exists and was initialized successfully,
//
//	or nil and an error if there was an issue with initialization or retrieval.
func GetClientPool(serviceName string) (*clientPool, error) {
	// Try to load the poolEntry for the given service name from the clientPoolMap.
	entry, ok := clientPoolMap.Load(serviceName)
	if !ok {
		// Return an error if the service name is not found in the clientPoolMap.
		return nil, toolerrors.WithFields("serviceName", serviceName).NewError(errors.ErrInitClientPoolFailed)
	}
	// Type assertion to get the poolEntry.
	e := entry.(*poolEntry)
	if e.err != nil {
		// Return the initialization error if it exists.
		return nil, e.err
	}
	if e.instance == nil {
		// Return an error if the client pool instance is nil.
		return nil, toolerrors.WithFields("serviceName", serviceName).NewError(errors.ErrInitClientPoolFailedNoInstance)
	}
	// Return the client pool instance if everything is okay.
	return e.instance, nil
}

// newClientPool is an internal function that actually creates the client pool.
// It first validates that at least one endpoint is provided.
// Then, it initializes a clientPool struct with default values and any values set by the provided options.
// After that, it attempts to initialize the connections in the pool.
// If the connection initialization is successful, it starts the health check goroutine.
//
// @param serviceName: A string representing the name of the gRPC service.
// @param endpoints: A slice of strings representing the endpoints of the gRPC service.
// @param opts: Variadic parameter of type Option, which are used to configure the client pool.
// @return: A pointer to the newly created clientPool if successful, or nil and an error if there was an issue.
func newClientPool(serviceName string, endpoints []string, opts ...Option) (*clientPool, error) {
	// Check if at least one endpoint is provided.
	if len(endpoints) == 0 {
		return nil, errors.ErrNoEndpoints
	}

	// Initialize a new clientPool struct with default values.
	p := &clientPool{
		serviceName:         serviceName,
		endpoints:           endpoints,
		dialTimeout:         defaultDialTimeout,
		healthCheckInterval: defaultHealthCheckInterval,
		poolSize:            defaultPoolSize,
		balancerPolicy:      roundrobin.Name,
		stopChan:            make(chan struct{}),
		dialOptions: []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	}

	// Apply the provided options to the client pool.
	for _, opt := range opts {
		opt(p)
	}

	// Initialize the connections in the pool.
	if err := p.initConnections(); err != nil {
		return nil, err
	}
	// Start the health check goroutine to monitor the connections.
	go p.startHealthCheck()
	return p, nil
}
