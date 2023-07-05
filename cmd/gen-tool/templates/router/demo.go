/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:35:21
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 16:27:16
 * @Description: Demo router
 */
package router

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type DemoRouter interface {
	SaveTemplate(path string, projectImportPath string) error
	GetTemplate() string
}

type Demo struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectImportPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:49:04
 * @return {*}
 */
func (d *Demo) SaveTemplate(path string, projectImportPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath string
		Time              string
	}

	data := &Data{ProjectImportPath: projectImportPath, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
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
	    "net/http"
	    "{{.ProjectImportPath}}/logger"
	    "{{.ProjectImportPath}}/service"
	    "{{.ProjectImportPath}}/vo/input"
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
	   groupRouter.GET("/demo", t.Demo)
   }

   /**
	* @description: Demo
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (t *Demo) Demo(ctx *gin.Context) {
   
		/**
		* @step
		* @inputVo
		**/
		inputVo := &input.Demo{}

		/**
		* @step
		* @should bind
		**/
		if err := ctx.ShouldBind(inputVo); err != nil {
			logger.Logger().Errorf("DemoController Demo shouldBind Err : %+v", err)
			return 
		}

		/**
		* @step
		* @调用service
		**/
		outputVo, err := service.CreateDemoService().Demo(ctx, inputVo)
		if err != nil {
			logger.Logger().Errorf("DemoController DemoService Demo Err : %+v", err)
		}

		/**
		* @step
		* @return
		**/
		ctx.JSON(http.StatusOK, outputVo)
		return 
	}
   `
}
