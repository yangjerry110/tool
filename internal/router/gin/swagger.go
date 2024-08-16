/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-14 11:17:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-16 17:28:54
 * @Description: swagger
 */
package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"github.com/yangjerry110/tool/internal/router"
)

type Swagger struct{}

/**
 * @description: Register
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:45:31
 * @return {*}
 */
func (s *Swagger) Register(registerEngine router.RegisterEngine) error {

	// Get Default Router
	// Define apidoc router
	registerEngine.GetGinEngine().GET("/api/apidoc/*any", s.apidoc)
	return nil
}

/**
 * @description: apidoc
 * @param {*gin.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:45:02
 * @return {*}
 */
func (s *Swagger) apidoc(ctx *gin.Context) {

	any := ctx.Param("any")

	switch any {
	case "/index.html":
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
		doc, err := swag.ReadDoc()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.Writer.Write([]byte(doc))
	}
}
