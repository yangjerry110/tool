package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/router/internal/config"
)

// httpRouter struct is responsible for managing HTTP routes, services, and middleware.
// It contains maps to store route registrations, HTTP services, and middleware.
type httpRouter struct {
	// routerRegisterMap stores the HTTP route registrations.
	// The key is the route name, and the value is the RouterRegisterHTTP interface.
	routerRegisterMap map[string]RouterRegisterHTTP

	// routerHTTPServiceMap stores the HTTP services associated with specific routes.
	// The key is the route name, and the value is the RouterHTTPService interface.
	routerHTTPServiceMap map[string]RouterHTTPService

	// RouterUseHTTPMap stores the middleware functions to be used by the HTTP router.
	// The key is the middleware name, and the value is the RouterUseHTTP interface.
	RouterUseHTTPMap map[string]RouterUseHTTP
}

/**
 * @description: Registers an HTTP route with the given route name and RouterRegisterHTTP interface.
 * @param {string} routerName - The name of the route to register.
 * @param {RouterRegisterHTTP} routerRegister - The RouterRegisterHTTP interface to register.
 * @author: Jerry.Yang
 * @date: 2025-03-12 16:36:26
 * @return {RouterHTTP} - Returns the RouterHTTP interface to allow method chaining.
 */
func (h *httpRouter) RegisterHTTP(routerRegister RouterRegisterHTTP) RouterHTTP {
	// Initialize the routerRegisterMap if it is nil.
	if h.routerHTTPServiceMap == nil {
		h.routerRegisterMap = make(map[string]RouterRegisterHTTP)
	}

	// Store the RouterRegisterHTTP interface in the routerRegisterMap with the route name as the key.
	h.routerRegisterMap[routerRegister.RouterName()] = routerRegister

	// Return the RouterHTTP interface to allow method chaining.
	return h
}

/**
 * @description: Registers an HTTP service with the given route name and RouterHTTPService interface.
 * @param {string} routerName - The name of the route to associate the service with.
 * @param {RouterHTTPService} routerHTTPService - The RouterHTTPService interface to register.
 * @author: Jerry.Yang
 * @date: 2025-03-12 16:36:26
 * @return {RouterHTTP} - Returns the RouterHTTP interface to allow method chaining.
 */
func (h *httpRouter) RegisterHTTPService(routerHTTPService RouterHTTPService) RouterHTTP {
	// Initialize the routerHTTPServiceMap if it is nil.
	if h.routerHTTPServiceMap == nil {
		h.routerHTTPServiceMap = make(map[string]RouterHTTPService)
	}

	// Store the RouterHTTPService interface in the routerHTTPServiceMap with the route name as the key.
	h.routerHTTPServiceMap[routerHTTPService.RouterName()] = routerHTTPService

	// Return the RouterHTTP interface to allow method chaining.
	return h
}

/**
 * @description: Registers middleware with the given middleware name and RouterUseHTTP interface.
 * @param {string} useName - The name of the middleware to register.
 * @param {RouterUseHTTP} routerUse - The RouterUseHTTP interface to register.
 * @author: Jerry.Yang
 * @date: 2025-03-12 16:36:26
 * @return {RouterHTTP} - Returns the RouterHTTP interface to allow method chaining.
 */
func (h *httpRouter) UseHTTP(routerUse RouterUseHTTP) RouterHTTP {
	// Initialize the RouterUseHTTPMap if it is nil.
	if h.RouterUseHTTPMap == nil {
		h.RouterUseHTTPMap = make(map[string]RouterUseHTTP)
	}

	// Store the RouterUseHTTP interface in the RouterUseHTTPMap with the middleware name as the key.
	h.RouterUseHTTPMap[routerUse.UseName()] = routerUse

	// Return the RouterHTTP interface to allow method chaining.
	return h
}

/**
 * @description: Starts the HTTP server using the provided configuration.
 * It sets up the Gin engine, registers middleware, routes, and services, and then starts the server.
 * @param {conf.Conf} httpConf - The configuration to use for starting the server.
 * @author: Jerry.Yang
 * @date: 2025-03-12 16:36:26
 * @return {error} - Returns an error if the server fails to start, otherwise nil.
 */
func (h *httpRouter) RunHTTP(httpConf conf.Conf) error {
	// If there are no registered routes, return nil to indicate that there is nothing to run.
	if len(h.routerRegisterMap) == 0 {
		return nil
	}

	// Register default routes for "ping" and "swagger".
	h.RegisterHTTP(&ping{})
	h.RegisterHTTP(&swagger{})

	// Create a new Gin engine with default middleware (logger and recovery).
	ginEngine := gin.Default()

	// If there are registered middleware, apply them to the Gin engine.
	if len(h.RouterUseHTTPMap) != 0 {
		for _, useHttp := range h.RouterUseHTTPMap {
			ginEngine.Use(useHttp.UseHTTP())
		}
	}

	// Register all routes and their associated services with the Gin engine.
	for routerName, routerRegister := range h.routerRegisterMap {
		routerRegister.RegisterHTTP(ginEngine)
		httpService, isOk := h.routerHTTPServiceMap[routerName]
		if isOk {
			routerRegister.RegisterHTTPService(httpService)
		}
	}

	// Create and set the configuration using the provided conf.Conf interface.
	err := conf.CreateConf(httpConf).SetConfig()
	if err != nil {
		return err
	}

	// Start the Gin engine and listen on the address specified in the configuration.
	return ginEngine.Run(config.HttpConf.Addr)
}
