/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 16:54:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 18:52:44
 * @Description: The swagger package provides functionality for serving Swagger API documentation.
 * It registers an endpoint to serve the Swagger UI and the OpenAPI specification file.
 */
package router

import (
	netHttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
)

// swagger struct is responsible for registering the Swagger API documentation endpoint.
// It implements the RouterRegister interface to provide route registration functionality.
type swagger struct{}

// registerHTTP registers the Swagger API documentation endpoint in the provided gin.IRouter instance.
// It maps the `/api/apidoc/*any` route to the `apidoc` handler function.
//
// Parameters:
//   - ginRouter: The gin.IRouter instance where the Swagger endpoint will be registered.
func (s *swagger) RegisterHTTP(ginRouter gin.IRouter) {
	ginRouter.GET("/api/apidoc/*any", s.apidoc)
}

// registerService is a placeholder method to satisfy the RouterRegister interface.
// It does not perform any operations since the `swagger` struct does not require service registration.
//
// Parameters:
//   - service: The RouterRegisterHttpService implementation (not used in this context).
func (s *swagger) RegisterService(service RouterRegisterHttpService) {
	// No implementation required for the `swagger` struct.
}

// apidoc handles requests to the Swagger API documentation endpoint.
// It serves either the Swagger UI (index.html) or the OpenAPI specification file (swagger.json).
//
// Parameters:
//   - ctx: The gin.Context object representing the HTTP request and response.
func (s *swagger) apidoc(ctx *gin.Context) {
	// Extract the "any" parameter from the URL path
	any := ctx.Param("any")

	// Handle different cases based on the value of the "any" parameter
	switch any {
	case "/index.html":
		// Serve the Swagger UI (index.html)
		var doc = []byte(`
		  <!DOCTYPE html>
		  <html>
		  <head>
			  <title>ReDoc</title>
			  <meta charset="utf-8"/>
			  <meta name="viewport" content="width=device-width, initial-scale=1">
			  <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
			  <style>
			  body {
				  margin: 0;
				  padding: 0;
			  }
			  </style>
		  </head>
		  <body>
			  <redoc spec-url='swagger.json'></redoc>
			  <script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
		  </body>
		  </html>
		  `)
		ctx.Data(200, "text/html; charset=utf-8", doc)
	case "/swagger.json":
		// Serve the OpenAPI specification file (swagger.json)
		doc, err := swag.ReadDoc()
		if err != nil {
			ctx.String(netHttp.StatusInternalServerError, err.Error())
			return
		}
		ctx.Writer.Write([]byte(doc))
	}
}
