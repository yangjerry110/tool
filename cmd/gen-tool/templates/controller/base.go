/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:40:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-22 15:31:29
 * @Description: base
 */
package controller

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
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

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), data)
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
		Time             string
	}

	/**
	 * @step
	 * @渲染参数
	 **/
	data := &Data{ControllerName: controllerName, ControllerNameUp: templates.CreateCommonTemplate().FirstUpper(controllerName), Time: templates.CreateCommonTemplate().GetFormatNowTime()}
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
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: base
	*/
   package controller
   
   /**
	* @description: CreateDemoController
	* @param {...DemoController} DemoControllers
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateDemoController(DemoControllers ...DemoController) DemoController {
	   if len(DemoControllers) == 0 {
		   return &Demo{}
	   }
	   return DemoControllers[0]
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
	* @description: Create{{.ControllerNameUp}}Controller
	* @param {...{{.ControllerNameUp}}Controller} {{.ControllerNameUp}}Controllers
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	 */
	func Create{{.ControllerNameUp}}Controller({{.ControllerName}}Controllers ...{{.ControllerNameUp}}Controller) {{.ControllerNameUp}}Controller {
		if len({{.ControllerName}}Controllers) == 0 {
			return &{{.ControllerNameUp}}{}
		}
		return {{.ControllerName}}Controllers[0]
	}`
}
