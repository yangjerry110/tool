/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:31:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 15:49:03
 * @Description: gin router
 */
package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router"
)

type Gin struct {
	*gin.Engine
}

// Gin Init
//
// Init
// Date 2024-04-12 17:21:38
// Author Jerry.Yang
func (g *Gin) Init() router.RouterInterface {

	// create g.Engine
	g.Engine = gin.Default()
	return g
}

/**
 * @description: Register
 * @author: Jerry.Yang
 * @date: 2023-12-13 17:36:43
 * @return {*}
 */
func (g *Gin) Register(routerName string, registerRouter router.Register) router.Register {

	// register router
	registerRouter.Register(g.Engine)
	return registerRouter
}

/**
 * @description: Use
 * @param {gin.HandlerFunc} ginHandlerFunc
 * @author: Jerry.Yang
 * @date: 2024-08-07 15:47:17
 * @return {*}
 */
func (g *Gin) Use(ginHandlerFunc gin.HandlerFunc) error {

	// use gin.HandlerFunc
	g.Engine.Use(ginHandlerFunc)
	return nil
}

// Run gin
//
// Param runConf conf.conf
// Return error
// date 2023-12-14 11:29:26
// Author Jerry.Yang
func (g *Gin) Run(runConf conf.Conf) error {

	// If you need to add other config later, you can add it before running, and if you don't want to load these, you can replace the others and customize a conf
	// Set RunGin conf
	if err := conf.CreateConf(runConf).SetConfig(); err != nil {
		return err
	}

	// Register Swagger
	g.Register("swagger", &Swagger{})

	// Register ping
	g.Register("ping", &Ping{})

	// Run
	return g.Engine.Run(router.RouteConf.Addr)
}
