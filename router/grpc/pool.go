/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-14 18:15:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-14 18:25:39
 * @Description:
 */
package grpc

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// clientPool is a struct that manages a pool of gRPC client connections.
// It provides methods to get a client connection, close all connections, and perform health checks on the connections.
// The pool uses a simple round - robin algorithm to distribute requests among the clients.
type clientPool struct {
	// clients is a slice that stores all the gRPC client connections in the pool.
	clients []*grpc.ClientConn
	// config holds the configuration settings for the gRPC client connections, such as service name, endpoints, etc.
	config *Config
	// current is an index used for the round - robin algorithm to keep track of the next client to be used.
	current int
	// mu is a read - write mutex used to ensure thread - safety when accessing or modifying the client pool.
	mu sync.RWMutex
}

// GetClient retrieves an available gRPC client connection from the pool using a round - robin algorithm.
// It locks the mutex to ensure thread - safety when accessing the client pool.
// If there are no available clients in the pool, it returns an error.
// Otherwise, it updates the current index to point to the next client in the pool and returns that client.
func (p *clientPool) GetClient() (*grpc.ClientConn, error) {
	// Lock the mutex to prevent concurrent access to the client pool.
	p.mu.Lock()
	// Unlock the mutex when the function returns.
	defer p.mu.Unlock()

	// Check if there are no available clients in the pool.
	if len(p.clients) == 0 {
		return nil, errors.New("no available gRPC clients")
	}

	// Use a simple round - robin algorithm to select the next client.
	p.current = (p.current + 1) % len(p.clients)
	return p.clients[p.current], nil
}

// Close closes all the gRPC client connections in the pool.
// It locks the mutex to ensure thread - safety when accessing the client pool.
// It iterates over all the clients in the pool and attempts to close each connection.
// If any connection fails to close, it records the error.
// After closing all connections, it resets the singleton instance so that new connections can be created later.
// If there are any errors during the closing process, it returns a combined error message.
func (p *clientPool) Close() error {
	// Lock the mutex to prevent concurrent access to the client pool.
	p.mu.Lock()
	// Unlock the mutex when the function returns.
	defer p.mu.Unlock()

	// A slice to store any errors that occur during the closing process.
	var errs []error
	// Iterate over all the clients in the pool.
	for _, client := range p.clients {
		// Attempt to close the client connection.
		if err := client.Close(); err != nil {
			// If the closing fails, record the error.
			errs = append(errs, err)
		}
	}

	// Reset the singleton instance so that new connections can be created later.
	grpcClient.poolMu.Lock()
	defer grpcClient.poolMu.Unlock()
	grpcClient.instance = nil
	grpcClient.once = sync.Once{}

	// Check if there are any errors during the closing process.
	if len(errs) > 0 {
		return fmt.Errorf("errors closing connections: %v", errs)
	}
	return nil
}

// healthCheck checks the health status of all the gRPC client connections in the pool.
// It reads - locks the mutex to allow concurrent read access to the client pool.
// It iterates over all the clients in the pool and checks their connection state.
// If a client is in the Ready or Idle state, it is considered healthy and added to the healthyClients slice.
// If a client is in an unhealthy state, it attempts to recreate the connection.
// If the re - creation fails, it records the error.
// After checking all clients, it updates the client pool with the healthy clients.
// If there are no healthy clients available, it returns an error.
func (p *clientPool) healthCheck(ctx context.Context) error {
	// Read - lock the mutex to allow concurrent read access to the client pool.
	p.mu.RLock()
	// Unlock the mutex when the function returns.
	defer p.mu.RUnlock()

	// A slice to store the healthy gRPC client connections.
	var healthyClients []*grpc.ClientConn
	// A slice to store any errors that occur during the health check.
	var errs []error

	// Iterate over all the clients in the pool.
	for _, client := range p.clients {
		// Get the current state of the client connection.
		state := client.GetState()
		// Check if the client is in the Ready or Idle state.
		if state == connectivity.Ready || state == connectivity.Idle {
			// If the client is healthy, add it to the healthyClients slice.
			healthyClients = append(healthyClients, client)
		} else {
			// If the client is unhealthy, attempt to recreate the connection.
			newClient, err := grpcClient.createConnection(ctx, p.config, p.config.Endpoints[p.current])
			if err != nil {
				// If the re - creation fails, record the error.
				errs = append(errs, err)
				continue
			}
			// Add the newly created client to the healthyClients slice.
			healthyClients = append(healthyClients, newClient)
			// Close the old, unhealthy client connection.
			client.Close()
		}
	}

	// Check if there are no healthy clients available.
	if len(healthyClients) == 0 {
		return fmt.Errorf("no healthy connections available: %v", errs)
	}

	// Update the client pool with the healthy clients.
	p.clients = healthyClients
	return nil
}
