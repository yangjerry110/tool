/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 14:15:44
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 18:57:32
 * @Description: The router package provides core functionalities for routing, including route registration, middleware usage, and route execution.
 * By defining interfaces, it enables flexible extension and decoupling of the routing module.
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/conf"
)

// router interface defines the core functionalities for routing, including route registration, middleware usage, and route execution.
type router interface {
	// register registers a route by accepting an implementation of the RouterRegister interface.
	//
	// Parameters:
	//   - routerRegister: The RouterRegister implementation to be registered.
	//
	// Returns:
	//   - error: An error if any issue occurs during registration.
	register(routerRegister RouterRegister) error

	// use applies middleware by accepting an implementation of the RouterUse interface.
	//
	// Parameters:
	//   - routerUse: The RouterUse implementation to be applied.
	//
	// Returns:
	//   - error: An error if any issue occurs during middleware application.
	use(routerUse RouterUse) error

	// run starts the routing engine by accepting a configuration object of type conf.Conf.
	//
	// Parameters:
	//   - conf: The configuration object used to set up the routing engine.
	//
	// Returns:
	//   - error: An error if any issue occurs during route execution.
	run(conf conf.Conf) error
}

// RouterRegister interface defines functionalities for route registration, including HTTP route registration and service registration.
type RouterRegister interface {
	// registerHTTP registers HTTP routes by accepting an implementation of the gin.IRouter interface.
	//
	// Parameters:
	//   - gin: The gin.IRouter implementation used to register HTTP routes.
	RegisterHTTP(gin gin.IRouter)

	// registerService registers services by accepting an implementation of the RouterRegisterHttpService interface.
	//
	// Parameters:
	//   - service: The RouterRegisterHttpService implementation to be registered.
	RegisterService(service RouterRegisterHttpService)
}

// RouterUse interface defines functionalities for middleware usage.
type RouterUse interface {
	// useHTTP returns a gin.HandlerFunc, which is a middleware handler function for HTTP routes.
	//
	// Returns:
	//   - gin.HandlerFunc: The middleware handler function.
	UseHTTP() gin.HandlerFunc
}

// RouterRegisterHttpService interface defines functionalities for HTTP service registration.
type RouterRegisterHttpService interface {
	// mustRouterRegisterHttpService enforces the implementation of HTTP service registration logic.
	// This method is typically used to ensure that implementing classes must include HTTP service registration logic.
	MustRouterRegisterHttpService()
}
