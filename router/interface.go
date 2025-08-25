/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-13 10:55:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-08-25 15:26:19
 * @Description: routerName
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/conf"
)

// RouterHTTP is the main interface for managing HTTP routes, services, and middleware.
// It provides methods to register routes, associate services, apply middleware, and start the HTTP server.
type RouterHTTP interface {
	// RegisterHTTP registers an HTTP route with the given route name and RouterRegisterHTTP interface.
	// It returns the RouterHTTP interface to allow method chaining.
	RegisterHTTP(routerRegister RouterRegisterHTTP) RouterHTTP

	// RegisterHTTPService registers an HTTP service with the given route name and RouterHTTPService interface.
	// It returns the RouterHTTP interface to allow method chaining.
	RegisterHTTPService(routerHTTPService RouterHTTPService) RouterHTTP

	// UseHTTP registers middleware with the given middleware name and RouterUseHTTP interface.
	// It returns the RouterHTTP interface to allow method chaining.
	UseHTTP(routerUse RouterUseHTTP) RouterHTTP

	// RunHTTP starts the HTTP server using the provided configuration.
	// It sets up the Gin engine, registers middleware, routes, and services, and then starts the server.
	RunHTTP(conf conf.Conf) error
}

// RouterRegisterHTTP is an interface for registering HTTP routes and associating services with them.
// It provides methods to register routes with a Gin router and associate services with those routes.
type RouterRegisterHTTP interface {
	RouterName() string
	// RegisterHTTP registers an HTTP route with the given Gin router.
	// This method is responsible for defining the route's path, HTTP method, and handler function.
	RegisterHTTP(ginRouter gin.IRouter)

	// RegisterHTTPService associates an HTTP service with the registered route.
	// This method is used to link a service (e.g., business logic) to a specific route.
	RegisterHTTPService(RouterHTTPService RouterHTTPService)
}

// RouterUseHTTP is an interface for defining middleware functions to be used by the HTTP router.
// It provides a method to return a Gin middleware handler function.
type RouterUseHTTP interface {
	UseName() string
	// UseHTTP returns a Gin middleware handler function.
	// This function is applied to the Gin engine to process requests before they reach the route handlers.
	UseHTTP(ginContext *gin.Context)
}

// RouterHTTPService is an interface for defining HTTP services associated with specific routes.
// It provides a method to enforce the implementation of essential service functionality.
type RouterHTTPService interface {
	// MustRouterHTTPService is a placeholder method to enforce the implementation of essential service functionality.
	// This method ensures that all services implement the required behavior.
	MustRouterHTTPService()
	RouterName() string
}
