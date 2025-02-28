/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 16:57:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-28 15:13:07
 * @Description: router
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router"
)

// Register
//
// Register
// Date 2024-04-12 17:44:09
// Author Jerry.Yang
func Register(routerName string, routerRegister router.Register) error {
	return routerEnginee().Register(routerName, routerRegister)
}

// Use
//
// Use
// Data 2024-08-07 15:49:38
// Author Jerry.Yang
func Use(ginHandlerFunc gin.HandlerFunc) error {
	return routerEnginee().Use(ginHandlerFunc)
}

// Run
//
// Run
// Date 2024-05-31 11:31:45
// Author Jerry.Yang
func Run(conf conf.Conf) error {
	return routerEnginee().Run(conf)
}
