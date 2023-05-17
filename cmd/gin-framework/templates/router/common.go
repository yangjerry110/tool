/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:32:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 14:31:12
 * @Description: common router
 */
package router

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type CommonRouter interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Common struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:34:57
 * @return {*}
 */
func (c *Common) SaveTemplate(path string, projectPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
		Time        string
	}

	data := &Data{ProjectPath: projectPath, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "commonRouter.go", c.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:34:01
 * @return {*}
 */
func (c *Common) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: common
	*/
   package router
   
   import (
	   "{{.ProjectPath}}/config"
   
	   "github.com/gin-gonic/gin"
   )
   
   type CommonRouter interface {
	   GetRouter() *gin.Engine
	   CreateRouter()
   }
   
   type Common struct{}
   
   /**
	* @description: default router
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   var defaultRouter *gin.Engine
   
   /**
	* @description: CreateRouter
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (c *Common) GetRouter() *gin.Engine {
   
	   /**
		* @step
		* @假如默认有值，则直接返回默认
		**/
	   if defaultRouter != nil {
		   return defaultRouter
	   }
   
	   /**
		* @step
		* @假如没有默认的则创建
		**/
	   defaultRouter = gin.Default()
	   return defaultRouter
   }
   
   /**
	* @description: CreateRouter
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (c *Common) CreateRouter() {
   
	   /**
		* @step
		* @test router
		**/
	   CreateTestRouter().CreateRouter()
   
	   /**
		* @step
		* @run
		**/
	   defaultRouter.Run(config.RouterSetConfig.Addr)
   }
   `
}
