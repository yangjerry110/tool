/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-12 17:43:11
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 16:36:30
 * @Description: router enginee
 */
package router

import (
	"github.com/yangjerry110/tool/internal/router"
)

var defaultRouterEnginee router.RouterInterface

// SetRouterEnginee
//
// SetRouterEnginee
// Date 2024-04-12 17:39:13
// Author Jerry.Yang
func SetRouterEnginee(routerInterface router.RouterInterface) router.RouterInterface {
	defaultRouterEnginee = routerInterface.Init()
	return defaultRouterEnginee
}

// routerEnginee
//
// routerEnginee
// Date 2024-04-12 17:42:03
// Author Jerry.Yang
func routerEnginee() router.RouterInterface {
	return defaultRouterEnginee
}
