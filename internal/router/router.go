/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:30:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-19 11:00:35
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
	Register(routerName string, register Register) error
	Use(useHandler UseHandler) error
	Run(conf conf.Conf) error
}

type Register interface {
	Register(registerEngine RegisterEngine) error
}

type RegisterEngine interface {
	GetGinEngine() *gin.Engine
	GetGrpcEngine() *grpc.Server
}

type Use interface {
	Use(handler UseHandler) error
}

type UseHandler interface {
	GetGinHandlerFunc() gin.HandlerFunc
}
