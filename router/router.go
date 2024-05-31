/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 16:57:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 17:44:37
 * @Description: router
 */
package router

import "github.com/yangjerry110/tool/internal/router"

// Register
//
// Register
// Date 2024-04-12 17:44:09
// Author Jerry.Yang
func Register(routerName string, routerRegister router.Register) error {
	return routerEnginee().Register(routerName, routerRegister)
}
