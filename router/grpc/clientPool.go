/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:13:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:39:04
 * @Description:  client pool
 */
package grpc

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/yangjerry110/tool/router/grpc/internal/errors"
	"github.com/yangjerry110/tool/toolerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// clientPool struct represents a pool of gRPC connections.
// It manages multiple connections to a gRPC service, including connection creation,
// health checking, and retrieval of available connections.
type clientPool struct {
	// serviceName is the name of the gRPC service to which the connections are made.
	serviceName string
	// endpoints is a slice of strings representing the endpoints of the gRPC service.
	endpoints []string
	// dialTimeout specifies the maximum time to wait when establishing a connection.
	dialTimeout time.Duration
	// healthCheckInterval determines how often the health of the connections is checked.
	healthCheckInterval time.Duration
	// poolSize represents the maximum number of connections in the pool.
	poolSize int
	// balancerPolicy is the load balancing policy to be used for the connections.
	balancerPolicy string
	// dialOptions are the options used when dialing to establish a gRPC connection.
	dialOptions []grpc.DialOption
	// conns is a slice of pointers to gRPC.ClientConn, holding the actual connections in the pool.
	conns []*grpc.ClientConn
	// mu is a read-write mutex used to synchronize access to the conns slice and other shared resources.
	mu sync.RWMutex
	// stopChan is a channel used to signal the health check goroutine to stop.
	stopChan chan struct{}
}

// GetConn retrieves an active and healthy connection from the client pool.
// It first locks the mutex for reading, then checks if there are any available connections.
// If there are connections, it iterates through them to find one in the "Ready" state.
// If a suitable connection is found, it is returned. Otherwise, an error is returned.
//
// # Returns
// - A pointer to a gRPC.ClientConn if an active and healthy connection is found, or nil and an error if not.
func (p *clientPool) GetConn() (*grpc.ClientConn, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if len(p.conns) == 0 {
		return nil, errors.ErrNoClientPools
	}
	for _, conn := range p.conns {
		if conn.GetState() == connectivity.Ready {
			return conn, nil
		}
	}
	return nil, errors.ErrNoHealthyClientPools
}

// initConnections initializes the connections in the client pool.
// It creates a new slice to hold the connections with the specified pool size.
// Then, it attempts to create each connection using the createConn method.
// If an error occurs during the creation of a connection, it closes all existing connections
// and returns the error. If all connections are successfully created, it returns nil.
//
// # Returns
// - nil if all connections are successfully initialized, or an error if there was an issue.
func (p *clientPool) initConnections() error {
	p.conns = make([]*grpc.ClientConn, 0, p.poolSize)
	for i := 0; i < p.poolSize; i++ {
		conn, err := p.createConn()
		if err != nil {
			p.Close()
			return toolerrors.WithFields("index", i).WithFields("ErrMsg", err).NewError(errors.ErrInitClientPoolFailed)
		}
		p.conns = append(p.conns, conn)
	}
	return nil
}

// createConn creates a single gRPC connection.
// It constructs the service configuration string with the specified load balancing policy.
// Then, it creates the target address for the connection.
// It combines the default dial options with the custom dial options provided for the client pool.
// After that, it attempts to establish a connection within the specified dial timeout.
// If the connection is successfully established, it checks the health of the connection using checkHealth.
// If the health check passes, the connection is returned. Otherwise, the connection is closed and an error is returned.
//
// # Returns
// - A pointer to a gRPC.ClientConn if the connection is successfully created and healthy, or nil and an error if not.
func (p *clientPool) createConn() (*grpc.ClientConn, error) {
	svcCfg := fmt.Sprintf(`{"loadBalancingPolicy":"%s"}`, p.balancerPolicy)
	target := fmt.Sprintf("%s:///%s", staticScheme, p.serviceName)
	opts := append([]grpc.DialOption{
		grpc.WithDefaultServiceConfig(svcCfg),
		grpc.WithResolvers(&staticResolver{endpoints: p.endpoints}),
		grpc.WithBlock(),
	}, p.dialOptions...)

	ctx, cancel := context.WithTimeout(context.Background(), p.dialTimeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, err
	}
	if err := p.checkHealth(conn); err != nil {
		_ = conn.Close()
		return nil, err
	}
	return conn, nil
}

// checkHealth checks the health of a given gRPC connection.
// It creates a context with a timeout for the health check.
// Then, it creates a HealthClient using the connection and sends a health check request
// for the specified service name. If the response indicates that the service is not in the "Serving" state
// or if there was an error during the health check, it returns an error. Otherwise, it returns nil.
//
// # Arguments
// - `conn`: A pointer to a gRPC.ClientConn whose health is to be checked.
//
// # Returns
// - nil if the health check passes, or an error if the health check fails.
func (p *clientPool) checkHealth(conn *grpc.ClientConn) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultHealthCheckTimeout)
	defer cancel()
	client := healthpb.NewHealthClient(conn)
	resp, err := client.Check(ctx, &healthpb.HealthCheckRequest{Service: p.serviceName})
	if err != nil || resp.Status != healthpb.HealthCheckResponse_SERVING {
		return toolerrors.WithFields("service", p.serviceName).WithFields("ErrMsg", err).NewError(errors.ErrCheckHealthFailed)
	}
	return nil
}

// startHealthCheck starts a goroutine that periodically checks the health of all connections in the pool.
// It uses a ticker to trigger the health check at the specified interval.
// When the ticker fires, it calls the refreshConnections method to check and potentially refresh the connections.
// The goroutine continues running until it receives a signal on the stopChan channel.
func (p *clientPool) startHealthCheck() {
	ticker := time.NewTicker(p.healthCheckInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			p.refreshConnections()
		case <-p.stopChan:
			return
		}
	}
}

// refreshConnections checks the health of all connections in the pool.
// It first locks the mutex for writing to ensure exclusive access to the conns slice.
// Then, it iterates through each connection and checks its health using the checkHealth method.
// If a connection is found to be unhealthy, it closes the connection, attempts to create a new one,
// and replaces the old connection with the new one in the conns slice. If there is an error creating the new connection,
// it logs the error and continues with the next connection.
func (p *clientPool) refreshConnections() {
	p.mu.Lock()
	defer p.mu.Unlock()
	for i, conn := range p.conns {
		if err := p.checkHealth(conn); err != nil {
			log.Printf("[clientpool] conn %d unhealthy: %v", i, err)
			_ = conn.Close()
			newConn, err := p.createConn()
			if err != nil {
				log.Printf("[clientpool] failed to recreate conn %d: %v", i, err)
				continue
			}
			p.conns[i] = newConn
		}
	}
}

// Close closes all connections in the client pool.
// It first locks the mutex for writing. Then, it checks if the stopChan channel has already been closed.
// If not, it closes the stopChan channel. After that, it iterates through all the connections and closes them.
// Finally, it sets the conns slice to nil to indicate that there are no more connections in the pool.
func (p *clientPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()
	select {
	case <-p.stopChan:
		return
	default:
		close(p.stopChan)
	}
	for _, conn := range p.conns {
		_ = conn.Close()
	}
	p.conns = nil
}
