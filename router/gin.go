/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 16:37:28
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 15:55:14
 * @Description: gin router
 */
package router

import (
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/internal/router"
	"github.com/yangjerry110/tool/internal/router/gin"
)

/**
 * @description: CreateGinConf
 * @author: Jerry.Yang
 * @date: 2023-12-19 16:38:45
 * @return {*}
 */
func CreateGinConf() error {
	return conf.CreateConf(&gin.Gin{}).SetConfig()
}

/**
 * @description: CreateGinConfigConf
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:55:04
 * @return {*}
 */
func CreateGinConfigConf() error {
	return conf.CreateConf(&gin.Config{}).SetConfig()
}

/**
 * @description: CreateGinRouter
 * @author: Jerry.Yang
 * @date: 2023-12-19 16:39:48
 * @return {*}
 */
func CreateGinRouter() router.RouterInterface {
	return router.CreateRouter(&gin.Gin{})
}
