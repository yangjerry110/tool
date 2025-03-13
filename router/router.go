/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-13 10:55:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-13 14:37:48
 * @Description:
 */
package router

import (
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/router/internal/config"
)

/**
 * @description: Sets the HTTP router configuration by creating a new configuration instance
 * for the HTTP router. This function is used to initialize and return the configuration
 * required for the HTTP router.
 * @author: Jerry.Yang
 * @date: 2025-03-12 16:36:26
 * @return {conf.Conf} - Returns the HTTP router configuration.
 */
func SetHTTPRouterConfig() conf.Conf {
	// Create and return a new configuration instance for the HTTP router.
	return conf.CreateConf(&config.HttpRouterConfig{})
}

/**
 * @description: Registers an HTTP route with the given route name and RouterRegisterHTTP interface.
 * This function initializes a new HTTP router engine and registers the route using the provided
 * RouterRegisterHTTP interface. It returns the RouterHTTP interface to allow method chaining.
 * @param {string} routerName - The name of the route to register.
 * @param {RouterRegisterHTTP} routerRegister - The RouterRegisterHTTP interface to register.
 * @author: Jerry.Yang
 * @date: 2025-03-12 16:36:26
 * @return {RouterHTTP} - Returns the RouterHTTP interface to allow method chaining.
 */
func RegisterHTTP(routerRegister RouterRegisterHTTP) RouterHTTP {
	// Initialize a new HTTP router engine.
	routerEnginee := &httpRouter{}

	// Register the route using the provided RouterRegisterHTTP interface.
	return routerEnginee.RegisterHTTP(routerRegister)
}
