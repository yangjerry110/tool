package router

import "github.com/gin-gonic/gin"

type router interface {
	register(registerName string, routerRegister routerRegister) error
	run() error
}

type routerRegister interface {
	routerRegisterHttp() RouterRegisterHttp
}

type RouterRegisterHttp interface {
	Register(ginEnginee *gin.Engine) RouterRegisterHttp
}
