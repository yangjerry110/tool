/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 16:51:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:02:37
 * @Description: The ping package provides functionality for registering a simple health check endpoint (`/ping`) in the HTTP router.
 * This endpoint is used to verify that the server is running and responding to requests.
 */
package router

import (
	netHttp "net/http"

	"github.com/gin-gonic/gin"
)

// ping struct is responsible for registering the `/ping` endpoint in the HTTP router.
// It implements the RouterRegister interface to provide route registration functionality.
type ping struct{}

// registerHTTP registers the `/ping` endpoint in the provided gin.IRouter instance.
// When the `/ping` endpoint is accessed via a GET request, it responds with "success" and an HTTP 200 status code.
//
// Parameters:
//   - ginRouter: The gin.IRouter instance where the `/ping` endpoint will be registered.
func (p *ping) registerHTTP(ginRouter gin.IRouter) {
	ginRouter.GET("/ping", func(ctx *gin.Context) {
		ctx.String(netHttp.StatusOK, "success")
	})
}

// registerService is a placeholder method to satisfy the RouterRegister interface.
// It does not perform any operations since the `ping` struct does not require service registration.
//
// Parameters:
//   - service: The RouterRegisterHttpService implementation (not used in this context).
func (p *ping) registerService(service RouterRegisterHttpService) {
	// No implementation required for the `ping` struct.
}
