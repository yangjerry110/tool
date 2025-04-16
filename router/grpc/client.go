/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:12:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:40:07
 * @Description:
 */
package grpc

import (
	"time"

	"github.com/yangjerry110/tool/router/grpc/internal/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

// Default values for various parameters of the client pool.
// defaultPoolSize defines the default number of connections in the pool.
const (
	defaultPoolSize = 3
	// defaultHealthCheckInterval sets the default interval for health checks.
	defaultHealthCheckInterval = 30 * time.Second
	// defaultDialTimeout specifies the default timeout for establishing a connection.
	defaultDialTimeout = 5 * time.Second
	// defaultHealthCheckTimeout sets the default timeout for performing a health check.
	defaultHealthCheckTimeout = 2 * time.Second
)

// New is a public function that creates a new instance of the ClientPool.
// It takes the service name, a slice of endpoints, and optional configuration options.
// If any error occurs during the creation process, it returns an error.
// Otherwise, it returns a pointer to the newly created clientPool.
//
// # Arguments
// - `serviceName`: A string representing the name of the gRPC service.
// - `endpoints`: A slice of strings representing the endpoints of the gRPC service.
// - `opts`: Variadic parameter of type Option, which are used to configure the client pool.
//
// # Returns
// - A pointer to the newly created clientPool if successful, or nil and an error if there was an issue.
func New(serviceName string, endpoints []string, opts ...Option) (*clientPool, error) {
	return newClientPool(serviceName, endpoints, opts...)
}

// newClientPool is an internal function that actually creates the client pool.
// It first validates that at least one endpoint is provided.
// Then, it initializes a clientPool struct with default values and any values set by the provided options.
// After that, it attempts to initialize the connections in the pool.
// If the connection initialization is successful, it starts the health check goroutine.
//
// # Arguments
// - `serviceName`: A string representing the name of the gRPC service.
// - `endpoints`: A slice of strings representing the endpoints of the gRPC service.
// - `opts`: Variadic parameter of type Option, which are used to configure the client pool.
//
// # Returns
// - A pointer to the newly created clientPool if successful, or nil and an error if there was an issue.
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
