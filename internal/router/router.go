/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:30:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 17:34:21
 * @Description: router
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
)

type RouterInterface interface {
	Init() RouterInterface
	Register(routerName string, register Register) error
	Run(conf conf.Conf) error
}

type Register interface {
	Register(ginEngine *gin.Engine) error
}
