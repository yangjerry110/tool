/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-13 10:55:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-08-25 15:43:51
 * @Description: swagger
 */
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
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
func (s *swagger) RegisterHTTP(ginEngine gin.IRouter) {
	// Define the route to serve the Swagger UI and API documentation.
	ginEngine.GET("/swagger", s.apidoc)
	ginEngine.GET("/swagger.json", s.json)
}

/**
 * @description: RegisterHTTPService is a placeholder method to register an HTTP service.
 * This method is currently unused but is included to satisfy the RouterRegisterHTTP interface.
 * @param {RouterHTTPService} service - The HTTP service to register.
 * @author: Jerry.Yang
 * @date: 2025-03-03 15:41:55
 * @return {void}
 */
func (s *swagger) RegisterHTTPService(service RouterHTTPService) {
	// This method is intentionally left empty as it is not currently used.
}

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
func (s *swagger) apidoc(ctx *gin.Context) {
	// Extract the path parameter from the request URL.
	// any := ctx.Param("any")

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
			 <script src="https://cdn.jsdelivr.net/npm/redoc@2.1.5/bundles/redoc.standalone.js"> </script>
		 </body>
		 </html>
		 `)
	ctx.Data(200, "text/html; charset=utf-8", doc)

	// Handle different paths for serving the Swagger UI and JSON specification.
	// switch any {
	// case "/index.html":
	// 	// Serve the Swagger UI HTML page.
	// 	var doc = []byte(`
	// 	 <!DOCTYPE html>
	// 	 <html>
	// 	 <head>
	// 		 <title>ReDoc</title>
	// 		 <meta charset="utf-8"/>
	// 		 <meta name="viewport" content="width=device-width, initial-scale=1">
	// 		 <link href="https://fonts.googleapis.com/css?family=Montserrat:300,400,700|Roboto:300,400,700" rel="stylesheet">
	// 		 <!--
	// 		 ReDoc doesn't change outer page styles
	// 		 -->
	// 		 <style>
	// 		 body {
	// 			 margin: 0;
	// 			 padding: 0;
	// 		 }
	// 		 </style>
	// 	 </head>
	// 	 <body>
	// 		 <redoc spec-url='swagger.json'></redoc>
	// 		 <script src="https://cdn.jsdelivr.net/npm/redoc@next/bundles/redoc.standalone.js"> </script>
	// 	 </body>
	// 	 </html>
	// 	 `)
	// 	ctx.Data(200, "text/html; charset=utf-8", doc)
	// case "/swagger.json":
	// 	// Serve the Swagger JSON specification.
	// 	doc, err := swag.ReadDoc()
	// 	if err != nil {
	// 		ctx.String(http.StatusInternalServerError, err.Error())
	// 		return
	// 	}
	// 	ctx.Writer.Write([]byte(doc))
	// }
}

func (s *swagger) json(ctx *gin.Context) {
	doc, err := swag.ReadDoc()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Writer.Write([]byte(doc))
}
