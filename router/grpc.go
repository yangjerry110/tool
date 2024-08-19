/*
 * @Author: Jerry.Yang
 * @Date: 2024-08-19 10:49:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-19 15:26:12
 * @Description: grpc
 */
package router

import (
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/internal/router"
	routerGrpc "github.com/yangjerry110/tool/internal/router/grpc"
)

/**
 * @description: CreateGrpcConf
 * @author: Jerry.Yang
 * @date: 2024-08-19 11:15:46
 * @return {*}
 */
func CreateGrpcConf() error {
	return conf.CreateConf(&routerGrpc.GrpcConfig{}).SetConfig()
}

/**
 * @description: CreateGrpcRouterConf
 * @author: Jerry.Yang
 * @date: 2024-08-19 11:15:57
 * @return {*}
 */
func CreateGrpcRouterConf() error {
	return conf.CreateConf(&router.GrpcConfig{}).SetConfig()
}

/**
 * @description: InitGrpcRouter
 * @author: Jerry.Yang
 * @date: 2024-08-19 15:25:53
 * @return {*}
 */
func InitGrpcRouter() error {
	SetRouterEnginee(&routerGrpc.Grpc{})
	return nil
}
