/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 16:57:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 16:26:57
 * @Description:
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router"
	toolgin "github.com/yangjerry110/tool/internal/router/gin"
)

/**
 * @description: CreateGinConf
 * @author: Jerry.Yang
 * @date: 2023-12-19 16:38:45
 * @return {*}
 */
func CreateGinConf() conf.Conf {
	return conf.CreateConf(&toolgin.Gin{})
}

/**
 * @description: CreateGinConfigConf
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:55:04
 * @return {*}
 */
func CreateGinConfigConf() conf.Conf {
	return conf.CreateConf(&toolgin.Config{})
}

/**
 * @description: CreateGinRouter
 * @author: Jerry.Yang
 * @date: 2023-12-19 16:39:48
 * @return {*}
 */
func CreateGinRouter() router.RouterInterface {
	return router.CreateRouter(&toolgin.Gin{})
}

/**
 * @description: GetGinDefaultRouter
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:50:44
 * @return {*}
 */
func GetGinDefaultRouter() *gin.Engine {
	return toolgin.GetGinDefaultRouter()
}
