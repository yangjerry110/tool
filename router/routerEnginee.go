/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 15:15:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-10 15:15:45
 * @Description: router enginee
 */
package router

import "github.com/yangjerry110/tool/internal/router"

var defaultRouterEnginee router.Router

func SetRouterEnginee(routerEnginee router.Router) router.Router {
	defaultRouterEnginee = routerEnginee
	return defaultRouterEnginee
}

func routerEnginee() router.Router {
	return defaultRouterEnginee
}
