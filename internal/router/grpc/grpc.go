/*
 * @Author: Jerry.Yang
 * @Date: 2024-08-16 17:05:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 16:42:22
 * @Description: grpc
 */
package grpc

import (
	"fmt"
	"net"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"github.com/yangjerry110/tool/internal/router"
	"google.golang.org/grpc"
)

type Grpc struct {
	syncOnce   sync.Once
	grpcServer *grpc.Server
}

/**
 * @description: Init
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:12:35
 * @return {*}
 */
func (g *Grpc) Init() router.RouterInterface {
	g.syncOnce.Do(func() {
		grpcNewServer := grpc.NewServer()
		g.grpcServer = grpcNewServer
	})
	return g
}

/**
 * @description: Register
 * @param {string} routerName
 * @param {router.Register} register
 * @author: Jerry.Yang
 * @date:
 * @return {*}
 */
func (g *Grpc) Register(routerName string, register router.Register) router.Register {

	// register
	register.RegisterGrpc(g.grpcServer)
	return register
}

/**
 * @description: Use
 * @param {gin.HandlerFunc} ginHandlerFunc
 * @author: Jerry.Yang
 * @date:
 * @return {*}
 */
func (g *Grpc) Use(userName string, useHandler router.Use) error {
	return nil
}

/**
 * @description: Run
 * @param {conf.Conf} conf
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:12:58
 * @return {*}
 */
func (g *Grpc) Run(runConf conf.Conf) error {

	// If you need to add other config later, you can add it before running, and if you don't want to load these, you can replace the others and customize a conf
	// Set RunGin conf
	if err := conf.CreateConf(runConf).SetConfig(); err != nil {
		return err
	}

	/**
	 * @step
	 * @net.Listen
	 **/
	lis, err := net.Listen(router.GrpcRouterConf.Protocol, fmt.Sprintf(":%s", router.GrpcRouterConf.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %+v", err)
	}

	/**
	 * @step
	 * @Serve
	 **/
	if err := g.grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %+v", err)
	}
	return nil
}

/**
 * @description: GetGinEngine
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:30:08
 * @return {*}
 */
func (g *Grpc) GetGinEngine() *gin.Engine {
	panic(errors.ErrGrpcRouterNoGinEngine)
}

/**
 * @description: GetGrpcEngine
 * @author: Jerry.Yang
 * @date: 2024-08-19 11:01:10
 * @return {*}
 */
func (g *Grpc) GetGrpcEngine() *grpc.Server {
	return g.grpcServer
}
