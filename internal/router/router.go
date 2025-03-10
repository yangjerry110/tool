/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 14:42:11
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-10 18:49:14
 * @Description: router
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
)

type Router interface {
	Init() Router
	Register(registerName string, routerRegister RouterRegister) error
	Use(useName string, use RouterUse) error
	Run(conf conf.Conf) error
}

type RouterRegister interface {
	RegisterGin() RouterRegisterGin
}

type RouterRegisterGin interface {
	Register(ginEnginee *gin.Engine) RouterRegisterGin
	RegisterService(routerService RouterRegisterGinHttpServer) RouterRegisterGin
}

type RouterRegisterGinHttpServer interface {
	mustRegisterServiceHttpServer()
}

type RouterUse interface {
	UseGin() RouterUseGin
}

type RouterUseGin interface {
	Use() gin.HandlerFunc
}
