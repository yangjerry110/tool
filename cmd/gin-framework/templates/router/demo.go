/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:35:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:56:57
 * @Description: Demo router
 */
package router

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type DemoRouter interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Demo struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:49:04
 * @return {*}
 */
func (d *Demo) SaveTemplate(path string, projectPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
		Time        string
	}

	data := &Data{ProjectPath: projectPath, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "demoRouter.go", d.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:49:19
 * @return {*}
 */
func (d *Demo) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: Demo
	*/
   package router
   
   import (
	   "{{.ProjectPath}}/controller"
   
	   "github.com/gin-gonic/gin"
   )
   
   type DemoRouter interface {
	   CreateRouter()
   }
   
   type Demo struct{}
   
   /**
	* @description: CreateRouter
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (t *Demo) CreateRouter() {
   
	   /**
		* @step
		* @获取router
		**/
	   router := CreateCommonRouter().GetRouter()
   
	   /**
		* @step
		* @定义group
		**/
	   groupRouter := router.Group("demo")
   
	   /**
		* @step
		* @定义router
		**/
	   groupRouter.GET("/demo", func(ctx *gin.Context) { controller.CreateDemoController().Demo(ctx) })
   }
   `
}
