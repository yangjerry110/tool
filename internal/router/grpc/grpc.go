/*
 * @Author: Jerry.Yang
 * @Date: 2024-08-16 17:05:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-16 17:37:01
 * @Description: grpc
 */
package grpc

import (
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router"
)

type Grpc struct{}

/**
 * @description: Init
 * @author: Jerry.Yang
 * @date: 2024-08-16 17:12:35
 * @return {*}
 */
func (g *Grpc) Init() router.RouterInterface {
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
	return nil
}
