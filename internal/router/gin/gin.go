/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 17:31:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-14 15:46:20
 * @Description: gin router
 */
package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"github.com/yangjerry110/tool/internal/router"
)

type Gin struct{}

/**
 * @description: defaultGinRouter
 * @author: Jerry.Yang
 * @date: 2023-12-13 17:39:12
 * @return {*}
 */
var defaultGinRouter *gin.Engine

/**
 * @description: GetGinDefaultRouter
 * @author: Jerry.Yang
 * @date: 2023-12-13 17:39:04
 * @return {*}
 */
func GetGinDefaultRouter() *gin.Engine {

	// First Resgister
	// if Err != nil; return err
	if err := router.CreateRouter(&Gin{}).SetDefaultRouter(); err != nil {
		panic(err)
	}

	// Judge defaultRouter
	// If == nil; return err
	if defaultGinRouter == nil {
		panic(errors.ErrGinRouterIsNoDefault)
	}

	// Return DefaultRouter
	return defaultGinRouter
}

/**
 * @description: Register
 * @author: Jerry.Yang
 * @date: 2023-12-13 17:36:43
 * @return {*}
 */
func (g *Gin) Register(routerName string, registerRouter router.Register) error {

	// Set router
	router.RegisterRouters.Store(routerName, registerRouter)
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
	g.Register("swagger", &SwaggerGinRouter{})

	// Circular ginRouters, executing registered ginrouters
	router.RegisterRouters.Range(func(routerName, registerRouter interface{}) bool {
		ginRouter := registerRouter.(*router.Register)
		if err := router.Register(*ginRouter).Register(); err != nil {
			panic(err)
		}
		return true
	})

	// Run
	return defaultGinRouter.Run(RouteConf.Addr)
}

// setDefaultRouter
//
// Return error
// Date 2023-12-14 11:35:30
// Author Jerry.Yang
func (g *Gin) SetDefaultRouter() error {

	// Judge defaultGinRouter
	// if not nil; return
	if defaultGinRouter != nil {
		return nil
	}

	// Register gin router
	// Register Default
	defaultGinRouter = gin.Default()
	return nil
}
