/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:30:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 15:48:56
 * @Description: router
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
)

type RouterInterface interface {
	Init() RouterInterface
	Register(routerName string, register Register) Register
	Use(ginHandlerFunc gin.HandlerFunc) error
	Run(conf conf.Conf) error
}

type Register interface {
	Register(ginEngine *gin.Engine) error
	RegisterService(service interface{}) error
}
