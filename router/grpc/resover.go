/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:14:34
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:27:29
 * @Description:
 */
package grpc

import "google.golang.org/grpc/resolver"

// staticScheme is a constant string representing the scheme for the static resolver.
// It is used to identify the type of resolver in the gRPC connection setup.
const staticScheme = "static"

// staticResolver struct is a custom resolver implementation for gRPC.
// It resolves endpoints to addresses in a static manner, meaning it uses a fixed list of endpoints.
type staticResolver struct {
	// endpoints is a slice of strings that holds the list of endpoints to be resolved.
	endpoints []string
}

// Scheme returns the scheme of the resolver, which in this case is the staticScheme.
// This method is part of the resolver.Resolver interface and is used by gRPC to identify the resolver type.
//
// # Returns
// A string representing the resolver scheme ("static" in this implementation).
func (r *staticResolver) Scheme() string {
	return staticScheme
}

// Build is a method that builds the resolver.
// It takes a resolver.Target, a resolver.ClientConn, and resolver.BuildOptions as parameters.
// It creates a slice of resolver.Address from the endpoints list and updates the state of the ClientConn with these addresses.
// Finally, it returns a new instance of nopResolver and nil error, indicating a successful build.
//
// # Arguments
// - `target`: A resolver.Target representing the target of the resolution.
// - `cc`: A resolver.ClientConn used to update the state with the resolved addresses.
// - `_`: resolver.BuildOptions (unused in this implementation, hence the blank identifier).
//
// # Returns
// A resolver.Resolver instance (in this case, a nopResolver) and an error (nil if successful).
func (r *staticResolver) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	addrs := make([]resolver.Address, len(r.endpoints))
	for i, ep := range r.endpoints {
		addrs[i] = resolver.Address{Addr: ep}
	}
	_ = cc.UpdateState(resolver.State{Addresses: addrs})
	return &nopResolver{}, nil
}

// nopResolver struct is a no-op resolver.
// It implements the resolver.Resolver interface but has empty implementations for its methods.
// It is used as a placeholder or a default resolver in the context of the staticResolver's Build method.
type nopResolver struct{}

// ResolveNow is a method of the resolver.Resolver interface.
// In the nopResolver implementation, it is a no-op and does nothing.
// It is called to trigger an immediate resolution, but in this case, there is no actual action taken.
//
// # Arguments
// - `_`: resolver.ResolveNowOptions (unused, hence the blank identifier).
func (r *nopResolver) ResolveNow(_ resolver.ResolveNowOptions) {}

// Close is a method of the resolver.Resolver interface.
// In the nopResolver implementation, it is a no-op and does nothing.
// It is called to close the resolver, but since there is nothing to close in this no-op resolver, it has an empty implementation.
func (r *nopResolver) Close() {}
