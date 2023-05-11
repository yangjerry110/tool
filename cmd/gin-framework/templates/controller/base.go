/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:40:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 15:23:27
 * @Description: base
 */
package controller

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type BaseController interface {
	SaveTemplate(path string) error
	GetTemplate() string
	AppendFuncTemplate(path string, ControllerName string) error
	GetAppendFuncTemplate() string
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
 * @description: AppendFuncTemplate
 * @param {string} path
 * @param {string} controllerName
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:13:22
 * @return {*}
 */
func (b *Base) AppendFuncTemplate(path string, controllerName string) error {

	/**
	 * @step
	 * @获取base的文件地址
	 **/
	basePath := fmt.Sprintf("%s/%s", path, "base.go")

	/**
	 * @step
	 * @定义需要渲染的数据结构
	 **/
	type Data struct {
		ControllerName   string
		ControllerNameUp string
	}

	/**
	 * @step
	 * @渲染参数
	 **/
	data := &Data{ControllerName: controllerName, ControllerNameUp: templates.CreateCommonTemplate().FirstUpper(controllerName)}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, b.GetAppendFuncTemplate(), data)
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

/**
 * @description: GetAppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:22:26
 * @return {*}
 */
func (b *Base) GetAppendFuncTemplate() string {
	return `/**
	* @description: Create{{.ControllerNameUp}}Dao
	* @param {...{{.ControllerNameUp}}Dao} {{.ControllerName}}Daos
	* @author: Jerry.Yang
	* @date: 2023-04-24 17:02:59
	* @return {*}
	*/
   func Create{{.ControllerNameUp}}Dao({{.ControllerName}}Daos ...{{.ControllerNameUp}}Dao) {{.ControllerNameUp}}Dao {
	   if len({{.ControllerName}}Daos) == 0 {
		   return &{{.ControllerNameUp}}{}
	   }
	   return {{.ControllerName}}Daos[0]
   }`
}
