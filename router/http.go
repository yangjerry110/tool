/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 14:21:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 19:05:25
 * @Description: http router package provides the concrete implementation of HTTP routing, including route registration, middleware usage, and route execution.
 * The `http` struct implements the `router` interface, supporting flexible route configuration and extension.
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/router/internal/config"
)

// http struct implements the `router` interface and is used to manage the registration and execution of HTTP routes.
type http struct {
	routerRegisters []RouterRegister // Stores all registered routes
	routerUses      []RouterUse      // Stores all used middleware
}

// register registers a route by appending the provided `RouterRegister` implementation to the `routerRegisters` list.
//
// Parameters:
//   - routerRegister: The `RouterRegister` implementation to be registered.
//
// Returns:
//   - error: An error if any issue occurs during registration.
func (h *http) register(routerRegister RouterRegister) error {
	if len(h.routerRegisters) == 0 {
		h.routerRegisters = make([]RouterRegister, 0)
	}
	h.routerRegisters = append(h.routerRegisters, routerRegister)
	return nil
}

// use registers a middleware by appending the provided `RouterUse` implementation to the `routerUses` list.
//
// Parameters:
//   - routerUse: The `RouterUse` implementation to be registered.
//
// Returns:
//   - error: An error if any issue occurs during middleware registration.
func (h *http) use(routerUse RouterUse) error {
	if len(h.routerUses) == 0 {
		h.routerUses = make([]RouterUse, 0)
	}
	h.routerUses = append(h.routerUses, routerUse)
	return nil
}

// run starts the HTTP routing engine using the provided configuration `httpConf`.
//
// Parameters:
//   - httpConf: The configuration object used to set up the HTTP server.
//
// Returns:
//   - error: An error if any issue occurs during route execution.
func (h *http) run(httpConf conf.Conf) error {
	// Set up the configuration
	err := conf.CreateConf(httpConf).SetConfig()
	if err != nil {
		return err
	}

	// Register default routes (e.g., ping and swagger)
	h.register(&ping{})
	h.register(&swagger{})

	// If no routes are registered, return immediately
	if len(h.routerRegisters) == 0 {
		return nil
	}

	// Create a new Gin engine instance
	ginEngine := gin.Default()

	// Register all routes
	for _, routerRegister := range h.routerRegisters {
		routerRegister.RegisterHTTP(ginEngine)
	}

	// Apply all middleware
	if len(h.routerUses) != 0 {
		for _, routerUse := range h.routerUses {
			ginEngine.Use(routerUse.UseHTTP())
		}
	}

	// Start the HTTP server using the address from the configuration
	ginEngine.Run(config.HttpConf.Addr)
	return nil
}
