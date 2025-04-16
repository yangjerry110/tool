/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:15:07
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:26:39
 * @Description:
 */
package grpc

import (
	"time"

	"google.golang.org/grpc"
)

// Option is a function type that takes a pointer to a clientPool struct as its parameter.
// It is used to configure the clientPool by setting specific properties.
type Option func(*clientPool)

// WithPoolSize is a function that returns an Option.
// This Option is used to set the size of the connection pool in the clientPool struct.
// If the provided size is greater than 0, it updates the poolSize field of the clientPool.
//
// # Arguments
// - `size`: An integer representing the desired size of the connection pool.
//
// # Returns
// An Option function that can be used to configure the clientPool.
func WithPoolSize(size int) Option {
	return func(p *clientPool) {
		if size > 0 {
			p.poolSize = size
		}
	}
}

// WithDialTimeout is a function that returns an Option.
// This Option is used to set the dial timeout duration in the clientPool struct.
// If the provided timeout duration is greater than 0, it updates the dialTimeout field of the clientPool.
//
// # Arguments
// - `timeout`: A time.Duration representing the timeout for establishing a connection.
//
// # Returns
// An Option function that can be used to configure the clientPool.
func WithDialTimeout(timeout time.Duration) Option {
	return func(p *clientPool) {
		if timeout > 0 {
			p.dialTimeout = timeout
		}
	}
}

// WithHealthCheckInterval is a function that returns an Option.
// This Option is used to set the interval for health checks in the clientPool struct.
// If the provided interval duration is greater than 0, it updates the healthCheckInterval field of the clientPool.
//
// # Arguments
// - `interval`: A time.Duration representing the interval between health checks.
//
// # Returns
// An Option function that can be used to configure the clientPool.
func WithHealthCheckInterval(interval time.Duration) Option {
	return func(p *clientPool) {
		if interval > 0 {
			p.healthCheckInterval = interval
		}
	}
}

// WithBalancerPolicy is a function that returns an Option.
// This Option is used to set the load balancing policy in the clientPool struct.
// If the provided policy string is not empty, it updates the balancerPolicy field of the clientPool.
//
// # Arguments
// - `policy`: A string representing the load balancing policy (e.g., "roundrobin").
//
// # Returns
// An Option function that can be used to configure the clientPool.
func WithBalancerPolicy(policy string) Option {
	return func(p *clientPool) {
		if policy != "" {
			p.balancerPolicy = policy
		}
	}
}

// WithDialOptions is a function that returns an Option.
// This Option is used to append additional gRPC dial options to the existing dial options in the clientPool struct.
// It takes variadic gRPC.DialOption values and appends them to the dialOptions slice of the clientPool.
//
// # Arguments
// - `opts`: Variadic parameter of type grpc.DialOption representing additional options for dialing a gRPC connection.
//
// # Returns
// An Option function that can be used to configure the clientPool.
func WithDialOptions(opts ...grpc.DialOption) Option {
	return func(p *clientPool) {
		p.dialOptions = append(p.dialOptions, opts...)
	}
}
