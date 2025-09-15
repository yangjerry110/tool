/*
 * @Author: Jerry.Yang
 * @Date: 2025-09-15 16:53:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-09-15 16:55:23
 * @Description: metrics 放开
 */
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type metrics struct{}

func (m *metrics) RegisterHTTP(ginEngine gin.IRouter) {
	ginEngine.GET("/metrics", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "success")
	})
}

func (m *metrics) RouterName() string {
	return "metrics"
}

func (m metrics) RegisterHTTPService(service RouterHTTPService) {

}
