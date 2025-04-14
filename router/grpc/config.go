/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-14 18:10:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-14 18:22:56
 * @Description: config
 */
package grpc

import "time"

// Config struct is used to hold the configuration settings for a gRPC client or server.
// It encapsulates various parameters that are essential for establishing and managing gRPC connections.
type Config struct {
	// ServiceName represents the name of the gRPC service.
	// This name can be used for identification, logging, and service discovery purposes.
	// For example, in a microservices architecture, different services might have unique names to distinguish them.
	ServiceName string

	// Endpoints is a slice of strings that contains the addresses of the gRPC service endpoints.
	// Each endpoint represents a location where the gRPC service can be reached.
	// For a client, these are the servers it can connect to; for a server, it might be used in a multi - instance setup.
	Endpoints []string

	// DialTimeout defines the maximum amount of time the gRPC client will wait to establish a connection to the server.
	// If the connection cannot be established within this time, an error will be returned.
	// It is specified as a time.Duration type, which allows for flexible time unit specification (e.g., seconds, milliseconds).
	DialTimeout time.Duration

	// LBPolicy specifies the load - balancing policy to be used by the gRPC client.
	// Different policies can be chosen based on the requirements, such as round - robin, least - loaded, etc.
	// This helps in distributing the client requests evenly among multiple server instances.
	LBPolicy string

	// MaxRetries indicates the maximum number of times the gRPC client will retry a failed request.
	// In case of transient errors (e.g., network glitches), retrying the request can increase the chances of success.
	// Once the maximum number of retries is reached, the client will stop retrying and return an error.
	MaxRetries int
}
