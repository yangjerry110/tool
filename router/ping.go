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

func (p ping) RegisterHTTPService(service RouterHTTPService) {

}
