/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-13 10:55:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-13 14:35:15
 * @Description: ping
 */
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ping struct{}

func (p *ping) RegisterHTTP(ginEngine gin.IRouter) {
	ginEngine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "success")
	})
}

func (p *ping) RouterName() string {
	return "ping"
}

func (p ping) RegisterHTTPService(service RouterHTTPService) {

}
