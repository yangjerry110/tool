/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-28 14:49:02
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-28 14:49:05
 * @Description: router
 */
package router

import (
	"github.com/yangjerry110/tool/internal/router"
	"github.com/yangjerry110/tool/internal/router/routerGin"
)

func CreateGinRouter() router.RegisterRouter {
	return &routerGin.RouterGin{}
}
