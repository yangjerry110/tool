/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:35:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-27 11:11:20
 * @Description: test router
 */
package router

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type TestRouter interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Test struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:49:04
 * @return {*}
 */
func (t *Test) SaveTemplate(path string, projectPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
	}

	data := &Data{ProjectPath: projectPath}
	return templates.CreateCommonTemplate().SaveTemplate(path, "test.go", t.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:49:19
 * @return {*}
 */
func (t *Test) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-23 14:18:10
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:39:34
	* @Description: test
	*/
   package router
   
   import (
	   "{{.ProjectPath}}/controller"
   
	   "github.com/gin-gonic/gin"
   )
   
   type TestRouter interface {
	   CreateRouter()
   }
   
   type Test struct{}
   
   /**
	* @description: CreateRouter
	* @author: Jerry.Yang
	* @date: 2023-04-23 14:31:32
	* @return {*}
	*/
   func (t *Test) CreateRouter() {
   
	   /**
		* @step
		* @获取router
		**/
	   router := CreateCommonRouter().GetRouter()
   
	   /**
		* @step
		* @定义group
		**/
	   gourpRouter := router.Group("test")
   
	   /**
		* @step
		* @定义router
		**/
	   gourpRouter.GET("/test", func(ctx *gin.Context) { controller.CreateTestController().Test(ctx) })
   }
   `
}
