/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:52:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 17:21:04
 * @Description: test controller
 */
package controller

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type TestController interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Test struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:55:10
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
	return templates.CreateCommonTemplate().SaveTemplate(path, "testController.go", t.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:55:19
 * @return {*}
 */
func (t *Test) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-23 11:44:49
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:30:54
	* @Description: test
	*/
   package controller
   
   import (
	   "net/http"
	   "{{.ProjectPath}}/logger"
	   "{{.ProjectPath}}/service"
	   "{{.ProjectPath}}/vo/input"
   
	   "github.com/gin-gonic/gin"
   )
   
   type TestController interface {
	   Test(ctx *gin.Context) error
   }
   
   type Test struct{}
   
   /**
	* @description: Test
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: 2023-04-23 11:45:45
	* @return {*}
	*/
   func (t *Test) Test(ctx *gin.Context) error {
   
	   /**
		* @step
		* @inputVo
		**/
	   inputVo := &input.Test{}
   
	   /**
		* @step
		* @should bind
		**/
	   if err := ctx.ShouldBind(inputVo); err != nil {
		   logger.Logger().Errorf("TestController Test shouldBind Err : %+v", err)
		   return err
	   }
   
	   /**
		* @step
		* @调用service
		**/
	   outputVo, err := service.CreateTestService().Test(ctx, inputVo)
	   if err != nil {
		   logger.Logger().Errorf("TestController TestService Test Err : %+v", err)
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
