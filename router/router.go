package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/internal/router"
	toolgin "github.com/yangjerry110/tool/internal/router/gin"
)

/**
 * @description: defaultRouter
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:51:12
 * @return {*}
 */
var defaultRouter = &toolgin.Gin{}

/**
 * @description: CreateRouter
 * @param {...router.RouterInterface} routers
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:52:26
 * @return {*}
 */
func CreateRouter(routers ...router.RouterInterface) router.RouterInterface {
	if len(routers) == 0 {
		return defaultRouter
	}
	return routers[0]
}

/**
 * @description: GetGinDefaultRouter
 * @author: Jerry.Yang
 * @date: 2023-12-14 15:50:44
 * @return {*}
 */
func GetGinDefaultRouter() *gin.Engine {
	return toolgin.GetGinDefaultRouter()
}
