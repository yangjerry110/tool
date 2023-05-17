/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-16 16:35:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 14:42:17
 * @Description: new router
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewRouter interface {
	SaveTemplate(path string, projectImportPath string, routerName string) error
	SaveAppendFuncTemplate(path string, baseRouterName string, routerName string) error
	GetTemplate() string
	GetAppendFuncTemplate() string
}

type New struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectImportPath
 * @param {string} routerName
 * @author: Jerry.Yang
 * @date: 2023-05-16 16:42:38
 * @return {*}
 */
func (n *New) SaveTemplate(path string, projectImportPath string, routerName string) error {

	/**
	 * @step
	 * @需要渲染的数据结构
	 **/
	type Data struct {
		ProjectImportPath string
		FirstRouterName   string
		RouterNameUp      string
		RouterName        string
		Time              string
	}

	/**
	 * @step
	 * @大写
	 **/
	routerNameUp := templates.CreateCommonTemplate().FirstUpper(routerName)

	/**
	 * @step
	 * @data
	 **/
	data := &Data{ProjectImportPath: projectImportPath, FirstRouterName: routerName[:1], RouterNameUp: routerNameUp, RouterName: routerName, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, fmt.Sprintf("%sRouter.go", routerName), n.GetTemplate(), data)
}

/**
 * @description: SaveAppendFuncTemplate
 * @param {string} path
 * @param {string} baseRouterName
 * @param {string} routerName
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:17:19
 * @return {*}
 */
func (n *New) SaveAppendFuncTemplate(path string, baseRouterName string, routerName string) error {

	/**
	 * @step
	 * @渲染需要append的文件地址
	 **/
	basePath := fmt.Sprintf("%s/%sRouter.go", path, baseRouterName)

	/**
	 * @step
	 * @需要渲染的数据结构
	 **/
	type Data struct {
		FirstBaseRouterName string
		BaseRouterNameUp    string
		RouterNameUp        string
		RouterName          string
		Time                string
	}

	/**
	 * @step
	 * @大写
	 **/
	baseRouterNameUp := templates.CreateCommonTemplate().FirstUpper(baseRouterName)
	routerNameUp := templates.CreateCommonTemplate().FirstUpper(routerName)

	/**
	 * @step
	 * @append
	 **/
	data := &Data{FirstBaseRouterName: baseRouterName[:1], BaseRouterNameUp: baseRouterNameUp, RouterNameUp: routerNameUp, RouterName: routerName, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, n.GetAppendFuncTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-16 16:38:15
 * @return {*}
 */
func (n *New) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.RouterName}}
	*/
   package router
   
   import (
	   "{{.ProjectImportPath}}/controller"
   
	   "github.com/gin-gonic/gin"
   )
   
   type {{.RouterNameUp}}Router interface {
	   CreateRouter()
   }
   
   type {{.RouterNameUp}} struct{}
   
   /**
	* @description: CreateRouter
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstRouterName}} *{{.RouterNameUp}}) CreateRouter() {
   
	   /**
		* @step
		* @获取router
		**/
	   router := CreateCommonRouter().GetRouter()
   
	   /**
		* @step
		* @定义group
		**/
	   gourpRouter := router.Group("{{.RouterName}}")
   
	   /**
		* @step
		* @定义router
		**/
	   gourpRouter.GET("/{{.RouterName}}", func(ctx *gin.Context) { controller.Create{{.RouterNameUp}}Controller().{{.RouterNameUp}}(ctx) })
   }
   `
}

/**
 * @description: GetAppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:10:14
 * @return {*}
 */
func (n *New) GetAppendFuncTemplate() string {
	return `
	/**
	* @description: {{.RouterName}}
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstBaseRouterName}} *{{.BaseRouterNameUp}}) Create{{.RouterNameUp}}Router() {
   
	   /**
		* @step
		* @获取router
		**/
	   router := CreateCommonRouter().GetRouter()
   
	   /**
		* @step
		* @定义group
		**/
	   gourpRouter := router.Group("{{.RouterName}}")
   
	   /**
		* @step
		* @定义router
		**/
	   gourpRouter.GET("/{{.RouterName}}", func(ctx *gin.Context) { controller.Create{{.BaseRouterNameUp}}Controller().{{.RouterNameUp}}(ctx) })
   }
	`
}
