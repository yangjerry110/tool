/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:30:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 16:40:22
 * @Description: router
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"google.golang.org/grpc"
)

type RouterInterface interface {
	Init() RouterInterface
	Register(routerName string, register Register) Register
	Use(userName string, use Use) error
	Run(conf conf.Conf) error
}

type Register interface {
	Register(ginEngine *gin.Engine) error
	RegisterService(service interface{}) error
	RegisterGrpc(grpc *grpc.Server) error
}

type Use interface {
	Use() gin.HandlerFunc
}
