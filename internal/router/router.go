/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:30:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-19 15:08:36
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
	Use(useName string, use Use) error
	Run(conf conf.Conf) error
}

type Register interface {
	Register(gin *gin.Engine) error
	RegisterGrpc(grpc *grpc.Server) error
}

type Use interface {
	Use() gin.HandlerFunc
}
