/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-20 17:05:11
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-16 17:30:01
 * @Description: Health check
 */
package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/router"
)

type Ping struct{}

/**
 * @description: Register
 * @author: Jerry.Yang
 * @date: 2023-12-20 17:07:11
 * @return {*}
 */
func (p *Ping) Register(registerEngine router.RegisterEngine) error {
	registerEngine.GetGinEngine().GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "success")
	})
	return nil
}
