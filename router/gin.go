/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 15:14:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-10 15:30:17
 * @Description: gin router
 */
package router

import (
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router"
	"github.com/yangjerry110/tool/internal/router/gin"
)

func InitGinRouter() router.Router {
	return SetRouterEnginee(&gin.Gin{})
}

func RegisterHttpServer(registerName string, routerRegister router.RouterRegister) router.RouterRegisterGin {
	routerEnginee().Register(registerName, routerRegister)
	return routerRegister.RegisterGin()
}

func UseHttpServer(useName string, routerUser router.RouterUse) router.RouterUseGin {
	routerEnginee().Use(useName, routerUser)
	return routerUser.UseGin()
}

func RunHttpServer(conf conf.Conf) error {
	return routerEnginee().Run(conf)
}
