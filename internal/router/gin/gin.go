/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 14:47:57
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-10 15:38:10
 * @Description:
 */
package gin

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router"
)

type Gin struct {
	syncOnce sync.Once
	*gin.Engine
}

func (g *Gin) Init() router.Router {
	// sync once
	g.syncOnce.Do(func() {
		// create g.Engine
		g.Engine = gin.Default()
	})
	return g
}

/**
 * @description: Register
 * @author: Jerry.Yang
 * @date: 2023-12-13 17:36:43
 * @return {*}
 */
func (g *Gin) Register(routerName string, routerRegister router.RouterRegister) error {

	// register router
	routerRegister.RegisterGin().Register(g.Engine)
	return nil
}

/**
 * @description: Use
 * @param {gin.HandlerFunc} ginHandlerFunc
 * @author: Jerry.Yang
 * @date: 2024-08-07 15:47:17
 * @return {*}
 */
func (g *Gin) Use(useName string, routerUse router.RouterUse) error {

	// use gin.HandlerFunc
	g.Engine.Use(routerUse.UseGin().Use())
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

	// register ping
	g.Register("ping", &Ping{})

	// register swagger
	g.Register("swagger", &Swagger{})

	// Run
	return g.Engine.Run(router.GinRouterConf.Addr)
}
