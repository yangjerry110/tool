/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 14:19:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:11:02
 * @Description: The router package provides utility functions for registering HTTP routes, applying middleware, and running the HTTP server.
 * These functions simplify the interaction with the underlying HTTP router engine.
 */
package router

import "github.com/yangjerry110/tool/conf"

// RegisterRouter registers a route by accepting a RouterRegister implementation.
// It delegates the registration to the underlying HTTP router engine.
//
// Parameters:
//   - routerRegister: The RouterRegister implementation to be registered.
//
// Returns:
//   - RouterRegister: The registered RouterRegister implementation.
func RegisterRouter(routerRegister RouterRegister) RouterRegister {
	routerEnginee().register(routerRegister)
	return routerRegister
}

// UseRouter applies middleware by accepting a RouterUse implementation.
// It delegates the middleware application to the underlying HTTP router engine.
//
// Parameters:
//   - routerUse: The RouterUse implementation to be applied.
//
// Returns:
//   - RouterUse: The applied RouterUse implementation.
func UseRouter(routerUse RouterUse) RouterUse {
	routerEnginee().use(routerUse)
	return routerUse
}

// Run starts the HTTP server using the provided configuration.
// It delegates the server startup to the underlying HTTP router engine.
//
// Parameters:
//   - conf: The configuration object used to set up the HTTP server.
//
// Returns:
//   - error: An error if any issue occurs during server startup.
func Run(conf conf.Conf) error {
	return routerEnginee().run(conf)
}
