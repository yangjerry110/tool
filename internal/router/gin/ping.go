/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 15:31:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-10 15:36:11
 * @Description:
 */
package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/router"
)

type Ping struct{}

func (p *Ping) Register(ginEngine *gin.Engine) router.RouterRegisterGin {
	ginEngine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "success")
	})
	return p
}

func (p *Ping) RegisterGin() router.RouterRegisterGin {
	return p
}

func (p *Ping) RegisterService(routerService router.RouterRegisterGinHttpServer) router.RouterRegisterGin {
	return p
}
