/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-28 14:14:44
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-28 14:51:11
 * @Description: gin router
 */
package routerGin

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/router"
)

// RouterGin 用于管理 Gin 路由
type RouterGin struct {
	ginEngine *gin.Engine
	once      sync.Once
}

// Register 注册路由
func (r *RouterGin) Register(registerRouter router.RegisterRouter) {
	registerRouter.RegisterGinHttpServer(r.getGinEngine())
}

func (r *RouterGin) RegisterGinHttpServer(gin *gin.Engine) {

}

// getGinEngine 获取 gin.Engine 实例（单例模式）
func (r *RouterGin) getGinEngine() *gin.Engine {
	r.once.Do(func() {
		if r.ginEngine == nil {
			r.ginEngine = gin.Default()
		}
	})
	return r.ginEngine
}
