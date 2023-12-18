/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:30:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-14 15:41:32
 * @Description: router
 */
package router

import (
	"sync"

	"github.com/yangjerry110/tool/internal/conf"
)

type RouterInterface interface {
	SetDefaultRouter() error
	Register(routerName string, register Register) error
	Run(conf conf.Conf) error
}

type Register interface {
	Register() error
}

/**
 * @description: RegisterRouters
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:41:47
 * @return {*}
 */
var RegisterRouters = sync.Map{}

/**
 * @description: CreateRouter
 * @param {RouterInterface} RouterInterface
 * @author: Jerry.Yang
 * @date: 2023-12-13 17:38:02
 * @return {*}
 */
func CreateRouter(RouterInterface RouterInterface) RouterInterface {
	return RouterInterface
}

// RegisterRouter
//
// Param Register
// Date 2023-12-14 15:41:19
// Author Jerry.Yang
func RegisterRouter(Register Register) Register {
	return Register
}
