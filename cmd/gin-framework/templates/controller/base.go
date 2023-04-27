/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:40:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 16:43:54
 * @Description: base
 */
package controller

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type BaseController interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Base struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:42:49
 * @return {*}
 */
func (b *Base) SaveTemplate(path string) error {
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), nil)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:42:18
 * @return {*}
 */
func (b *Base) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-21 19:40:24
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:26:07
	* @Description: base
	*/
   package controller
   
   /**
	* @description: CreateTestController
	* @param {...TestController} TestControllers
	* @author: Jerry.Yang
	* @date: 2023-04-23 14:26:15
	* @return {*}
	*/
   func CreateTestController(TestControllers ...TestController) TestController {
	   if len(TestControllers) == 0 {
		   return &Test{}
	   }
	   return TestControllers[0]
   }
   `
}
