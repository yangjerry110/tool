/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 14:19:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:11:35
 * @Description: The router package provides functionality for managing the default router engine.
 * It includes functions to set, retrieve, and initialize the default router engine.
 */
package router

// defaultRouterEnginee holds the default router engine instance.
// It is used globally to manage the router engine across the application.
var defaultRouterEnginee router

// SetRouterEnginee sets the default router engine to the provided routerEnginee instance.
// It also returns the newly set router engine for convenience.
//
// Parameters:
//   - routerEnginee: The router engine instance to be set as the default.
//
// Returns:
//   - router: The newly set router engine instance.
func SetRouterEnginee(routerEnginee router) router {
	defaultRouterEnginee = routerEnginee
	return defaultRouterEnginee
}

// routerEnginee retrieves the currently set default router engine.
// If no router engine is set, it returns nil.
//
// Returns:
//   - router: The currently set default router engine instance.
func routerEnginee() router {
	return defaultRouterEnginee
}

// HttpRouterEnginee initializes and sets the default router engine to an instance of the `http` struct.
// It then returns the default router engine for immediate use.
//
// Returns:
//   - router: The default router engine instance (initialized as `http`).
func HttpRouterEnginee() router {
	SetRouterEnginee(&http{})
	return routerEnginee()
}
