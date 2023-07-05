/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:04:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-22 15:35:30
 * @Description: base
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
)

type BaseRouter interface {
	SaveTemplate(path string) error
	SaveAppendFuncTemplate(path string, routerName string) error
	GetTemplate() string
	GetAppendFuncTemplate() string
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
 * @description: SaveAppendFuncTemplate
 * @param {string} path
 * @param {string} routerName
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:25:49
 * @return {*}
 */
func (b *Base) SaveAppendFuncTemplate(path string, routerName string) error {

	/**
	 * @step
	 * @basePath
	 **/
	basePath := fmt.Sprintf("%s/base.go", path)

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		RouterNameUp string
		RouterName   string
		Time         string
	}

	/**
	 * @step
	 * @大写
	 **/
	routerNameUp := templates.CreateCommonTemplate().FirstUpper(routerName)

	/**
	 * @step
	 * @append
	 **/
	data := &Data{RouterNameUp: routerNameUp, RouterName: routerName, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, b.GetAppendFuncTemplate(), data)
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
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: base
	*/
   package router
   
   /**
	* @description: CreateCommonRouter
	* @param {...CommonRouter} CommonRouters
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateCommonRouter(CommonRouters ...CommonRouter) CommonRouter {
	   if len(CommonRouters) == 0 {
		   return &Common{}
	   }
	   return CommonRouters[0]
   }
   
   /**
	* @description: CreateDemoRouter
	* @param {...DemoRouter} DemoRouters
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateDemoRouter(DemoRouters ...DemoRouter) DemoRouter {
	   if len(DemoRouters) == 0 {
		   return &Demo{}
	   }
	   return DemoRouters[0]
   }
   `
}

/**
 * @description: GetAppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:22:29
 * @return {*}
 */
func (b *Base) GetAppendFuncTemplate() string {
	return `
	/**
	* @description: Create{{.RouterNameUp}}Router
	* @param {...{{.RouterNameUp}}Router} {{.RouterNameUp}}Routers
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func Create{{.RouterNameUp}}Router({{.RouterName}}Routers ...{{.RouterNameUp}}Router) {{.RouterNameUp}}Router {
	   if len({{.RouterName}}Routers) == 0 {
		   return &{{.RouterNameUp}}{}
	   }
	   return {{.RouterName}}Routers[0]
   }`
}
