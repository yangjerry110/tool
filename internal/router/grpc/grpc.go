/*
 * @Author: Jerry.Yang
 * @Date: 2024-08-16 17:05:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-19 11:14:42
 * @Description: grpc
 */
package grpc

import (
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"github.com/yangjerry110/tool/internal/router"
	"google.golang.org/grpc"
)

type Grpc struct {
	grpcServer *grpc.Server
}

/**
 * @description: Init
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:12:35
 * @return {*}
 */
func (g *Grpc) Init() router.RouterInterface {
	grpcNewServer := grpc.NewServer()
	g.grpcServer = grpcNewServer
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
func (g *Grpc) Register(routerName string, register router.Register) error {

	// register
	if err := register.Register(g); err != nil {
		return err
	}
	return nil
}

/**
 * @description: Use
 * @param {gin.HandlerFunc} ginHandlerFunc
 * @author: Jerry.Yang
 * @date:
 * @return {*}
 */
func (g *Grpc) Use(useHandler router.UseHandler) error {
	return nil
}

/**
 * @description: Run
 * @param {conf.Conf} conf
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:12:58
 * @return {*}
 */
func (g *Grpc) Run(conf conf.Conf) error {

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
