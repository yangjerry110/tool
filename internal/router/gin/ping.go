/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-20 17:05:11
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 17:28:47
 * @Description: Health check
 */
package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct{}

/**
 * @description: Register
 * @author: Jerry.Yang
 * @date: 2023-12-20 17:07:11
 * @return {*}
 */
func (p *Ping) Register(ginEngine *gin.Engine) error {
	ginEngine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "success")
	})
	return nil
}
