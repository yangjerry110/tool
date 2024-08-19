/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:31:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-19 15:21:54
 * @Description: gin router
 */
package gin

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"github.com/yangjerry110/tool/internal/router"
	"google.golang.org/grpc"
)

type Gin struct {
	syncOnce sync.Once
	*gin.Engine
}

// Gin Init
//
// Init
// Date 2024-04-12 17:21:38
// Author Jerry.Yang
func (g *Gin) Init() router.RouterInterface {

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
func (g *Gin) Register(routerName string, registerRouter router.Register) error {

	// register router
	if err := registerRouter.Register(g.Engine); err != nil {
		return err
	}
	return nil
}

/**
 * @description: Use
 * @param {router.UseHandler} useHandler
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:35:21
 * @return {*}
 */
func (g *Gin) Use(useName string, use router.Use) error {

	// use gin.HandlerFunc
	g.Engine.Use(use.Use())
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

/**
 * @description: GetGinEngine
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:30:08
 * @return {*}
 */
func (g *Gin) GetGinEngine() *gin.Engine {
	return g.Engine
}

/**
 * @description: GetGrpcEngine
 * @author: Jerry.Yang
 * @date: 2024-08-19 11:01:10
 * @return {*}
 */
func (g *Gin) GetGrpcEngine() *grpc.Server {
	panic(errors.ErrGinRouterNoGrpcEngine)
}
