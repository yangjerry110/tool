/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:52:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:56:15
 * @Description: Demo controller
 */
package controller

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type DemoController interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Demo struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:55:10
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
	return templates.CreateCommonTemplate().SaveTemplate(path, "demoController.go", d.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:55:19
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
   package controller
   
   import (
	   "net/http"
	   "{{.ProjectPath}}/logger"
	   "{{.ProjectPath}}/service"
	   "{{.ProjectPath}}/vo/input"
   
	   "github.com/gin-gonic/gin"
   )
   
   type DemoController interface {
	   Demo(ctx *gin.Context) error
   }
   
   type Demo struct{}
   
   /**
	* @description: Demo
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (t *Demo) Demo(ctx *gin.Context) error {
   
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
		   return err
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
	   return nil
   }
   `
}
