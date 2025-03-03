/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-20 17:05:11
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 16:32:45
 * @Description: Health check
 */
package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
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

/**
 * @description: RegisterService
 * @param {interface{}} service
 * @author: Jerry.Yang
 * @date: 2025-03-03 15:41:33
 * @return {*}
 */
func (p *Ping) RegisterService(service interface{}) error {
	return nil
}

/* @description: RegisterGrpc
 * @param {*grpc.Server} grpcEngine
 * @author: Jerry.Yang
 * @date: 2024-08-19 14:57:02
 * @return {*}
 */
func (p *Ping) RegisterGrpc(grpcEngine *grpc.Server) error {
	return nil
}
