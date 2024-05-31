/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 16:37:28
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 17:37:17
 * @Description: gin router
 */
package router

import (
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
	return conf.CreateConf(&router.Config{})
}
