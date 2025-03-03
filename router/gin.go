/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-03 16:28:45
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 16:36:09
 * @Description: gin
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

/**
 * @description: InitGinRouter
 * @author: Jerry.Yang
 * @date: 2024-08-19 15:25:19
 * @return {*}
 */
func InitGinRouter() error {
	SetRouterEnginee(&toolgin.Gin{})
	return nil
}
