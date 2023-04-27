/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:04:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 15:32:09
 * @Description: base
 */
package router

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type BaseRouter interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Base struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:32:17
 * @return {*}
 */
func (b *Base) SaveTemplate(path string) error {
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), nil)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:31:33
 * @return {*}
 */
func (b *Base) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-21 14:46:07
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:29:47
	* @Description: base
	*/
   package router
   
   /**
	* @description: CreateCommonRouter
	* @param {...CommonRouter} CommonRouters
	* @author: Jerry.Yang
	* @date: 2023-04-21 14:59:43
	* @return {*}
	*/
   func CreateCommonRouter(CommonRouters ...CommonRouter) CommonRouter {
	   if len(CommonRouters) == 0 {
		   return &Common{}
	   }
	   return CommonRouters[0]
   }
   
   /**
	* @description: CreateTestRouter
	* @param {...TestRouter} TestRouters
	* @author: Jerry.Yang
	* @date: 2023-04-23 14:29:52
	* @return {*}
	*/
   func CreateTestRouter(TestRouters ...TestRouter) TestRouter {
	   if len(TestRouters) == 0 {
		   return &Test{}
	   }
	   return TestRouters[0]
   }
   `
}
