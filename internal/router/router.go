/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-28 14:13:57
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-28 14:50:18
 * @Description: router
 */
package router

import "github.com/gin-gonic/gin"

type Router interface {
	Register(RegisterRouter RegisterRouter)
}

type RegisterRouter interface {
	RegisterGinHttpServer(gin *gin.Engine)
}
