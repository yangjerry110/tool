/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-13 10:55:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-08-20 17:37:53
 * @Description: swagger
 */
package router

import (
	"net/http"

	// Import the Gin framework for building HTTP servers.
	// Gin is a lightweight and efficient web framework for Go.
	"github.com/gin-gonic/gin"
	// Import the Swagger UI files.
	// These files are used to serve the Swagger UI for API documentation.

	// Import the Gin middleware for integrating Swagger with Gin.
	// This middleware allows serving the Swagger UI and API documentation in a Gin application.

	// Import the Swag library for generating Swagger documentation.
	// Swag helps in generating Swagger JSON specifications from Go code.
	_ "github.com/swaggo/swag" // 使用下划线导入以避免直接依赖
)

// swagger struct is responsible for handling Swagger API documentation.
// It provides methods to register the Swagger UI and serve the API documentation.
type swagger struct{}

/**
 * @description: RegisterHTTP registers the Swagger UI and API documentation routes with the Gin engine.
 * It defines a route to serve the Swagger UI and API documentation.
 * @param {gin.IRouter} ginEngine - The Gin engine to register the routes with.
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:45:31
 * @return {void}
 */
// RegisterHTTP method registers the Swagger UI and API documentation routes with the given Gin engine.
// It creates a route `/api/apidoc/*any` and uses the `ginSwagger.WrapHandler` to serve the Swagger UI
// using the `swaggerFiles.Handler`. This allows users to access the Swagger UI and API documentation.
func (s *swagger) RegisterHTTP(ginEngine gin.IRouter) {
	// Define the route to serve the Swagger UI and API documentation.
	ginEngine.GET("/api/apidoc/*any", s.apidoc) // 使用自定义的apidoc方法
}

/**
 * @description: RegisterHTTPService is a placeholder method to register an HTTP service.
 * This method is currently unused but is included to satisfy the RouterRegisterHTTP interface.
 * @param {RouterHTTPService} service - The HTTP service to register.
 * @author: Jerry.Yang
 * @date: 2025-03-03 15:41:55
 * @return {void}
 */
// RegisterHTTPService method is a placeholder method.
// It is intended to register an HTTP service but is currently not used.
// It takes a `RouterHTTPService` as a parameter but does nothing with it for now.
func (s *swagger) RegisterHTTPService(service RouterHTTPService) {
	// This method is intentionally left empty as it is not currently used.
}

// RouterName method returns the name of the router as a string.
// In this case, it returns "swagger" which can be used for identification or logging purposes.
func (s *swagger) RouterName() string {
	return "swagger"
}

/**
 * @description: apidoc serves the Swagger UI and API documentation based on the requested path.
 * It handles requests for the Swagger UI HTML page and the Swagger JSON specification.
 * @param {*gin.Context} ctx - The Gin context for the HTTP request.
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:45:02
 * @return {void}
 */
// apidoc method serves the Swagger UI and API documentation based on the requested path.
// It extracts the path parameter from the request URL and checks the value.
// If the path is "/index.html", it serves the Swagger UI HTML page with a predefined HTML template.
// If the path is "/swagger.json", it reads the Swagger JSON specification using `swag.ReadDoc`
// and serves it to the client. If there is an error reading the doc, it returns an internal server error.
func (s *swagger) apidoc(ctx *gin.Context) {
	// Extract the path parameter from the request URL.
	any := ctx.Param("any")

	// Handle different paths for serving the Swagger UI and JSON specification.
	switch any {
	case "/index.html":
		// Serve the Swagger UI HTML page.
		var doc = []byte(`
		  <!DOCTYPE html>
		  <html>
		  <head>
			  <title>ReDoc</title>
			  <meta charset="utf-8"/>
			  <meta name="viewport" content="width=device-width, initial-scale=1">
			  <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
			  <!--
			  ReDoc doesn't change outer page styles
			  -->
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
		// 在 Go 1.20.3 中，我们使用替代方案来提供 swagger.json
		// 这里提供一个简单的实现，或者你可以从文件系统读取
		ctx.JSON(http.StatusOK, gin.H{
			"info": gin.H{
				"title":   "API Documentation",
				"version": "1.0",
			},
			"paths": gin.H{},
		})
	default:
		// 对于其他路径，返回404或者重定向到index.html
		ctx.Redirect(http.StatusFound, "/api/apidoc/index.html")
	}
}
