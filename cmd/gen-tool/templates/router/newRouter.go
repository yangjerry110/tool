/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-16 16:35:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-11-30 16:28:01
 * @Description: new router
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
)

type NewRouter interface {
	SaveTemplate(path string, projectImportPath string, routerName string) error
	SaveAppendFuncTemplate(path string, baseRouterName string, routerName string) error
	SaveProtobufTemplate(path string, projectImportPath string, routerName string, protobufRouters []*NewProtobufRouter) error
	GetTemplate() string
	GetAppendFuncTemplate() string
	GetProtobufTemplate() string
}

type New struct{}

/**
 * @description: NewRouters
 * @author: Jerry.Yang
 * @date: 2023-05-23 19:31:09
 * @return {*}
 */
type NewProtobufRouter struct {
	Description     string
	RouterMethod    string
	RouterPath      string
	RouterNameUp    string
	RouterFunc      string
	RouterFuncUp    string
	InputReqName    string
	OutputRespName  string
	FirstRouterName string
	RouterName      string
	Time            string
}

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
 * @description: SaveProtobufTemplate
 * @param {string} path
 * @param {string} projectImportPath
 * @param {string} routerName
 * @param {[]*NewProtobufRouter} protobufRouters
 * @author: Jerry.Yang
 * @date: 2023-05-23 19:37:55
 * @return {*}
 */
func (n *New) SaveProtobufTemplate(path string, projectImportPath string, routerName string, protobufRouters []*NewProtobufRouter) error {
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
		Routers           []*NewProtobufRouter
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
	data := &Data{ProjectImportPath: projectImportPath, FirstRouterName: routerName[:1], RouterNameUp: routerNameUp, RouterName: routerName, Time: templates.CreateCommonTemplate().GetFormatNowTime(), Routers: protobufRouters}
	return templates.CreateCommonTemplate().SaveTemplate(path, fmt.Sprintf("%sRouter.go", routerName), n.GetProtobufTemplate(), data)
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
	   "net/http"
	   "{{.ProjectImportPath}}/logger"
	   "{{.ProjectImportPath}}/service"
	   "{{.ProjectImportPath}}/vo/input"
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
		groupRouter := router.Group("{{.RouterName}}")
   
	   /**
		* @step
		* @定义router
		**/
		groupRouter.GET("/{{.RouterName}}", {{.FirstRouterName}}.{{.RouterNameUp}})
   }

   /**
	* @description: {{.RouterName}}
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
	func ({{.FirstRouterName}} *{{.RouterNameUp}}) {{.RouterNameUp}}(ctx *gin.Context)  {

		/**
		* @step
		* @inputVo
		**/
		inputVo := &input.{{.RouterNameUp}}{}

		/**
		* @step
		* @should bind
		**/
		if err := ctx.ShouldBind(inputVo); err != nil {
			logger.Logger().Errorf("{{.RouterNameUp}} shouldBind Err : %+v", err)
			return 
		}

		/**
		* @step
		* @调用service
		**/
		outputVo, err := service.Create{{.RouterNameUp}}Service().{{.RouterNameUp}}(ctx, inputVo)
		if err != nil {
			logger.Logger().Errorf("{{.RouterNameUp}}Service {{.RouterNameUp}} Err : %+v", err)
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
		groupRouter := router.Group("{{.RouterName}}")
   
	   /**
		* @step
		* @定义router
		**/
		groupRouter.GET("/{{.RouterName}}", {{.FirstBaseRouterName}}.{{.RouterNameUp}})
   }

   /**
	* @description: {{.RouterName}}
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
	func ({{.FirstBaseRouterName}} *{{.BaseRouterNameUp}}) {{.RouterNameUp}}(ctx *gin.Context)  {

		/**
		* @step
		* @inputVo
		**/
		inputVo := &input.{{.RouterNameUp}}{}

		/**
		* @step
		* @should bind
		**/
		if err := ctx.ShouldBind(inputVo); err != nil {
			logger.Logger().Errorf("{{.RouterNameUp}} shouldBind Err : %+v", err)
			return 
		}

		/**
		* @step
		* @调用service
		**/
		outputVo, err := service.Create{{.BaseRouterNameUp}}Service().{{.RouterNameUp}}(ctx, inputVo)
		if err != nil {
			logger.Logger().Errorf("{{.BaseRouterNameUp}}Service {{.RouterNameUp}} Err : %+v", err)
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

/**
 * @description: GetProtobufTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-23 19:27:29
 * @return {*}
 */
func (n *New) GetProtobufTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.RouterName}}
	*/
   package router
   
   import (
	   "net/http"
	   "github.com/gin-gonic/gin"
	   "{{.ProjectImportPath}}/logger"
	   "{{.ProjectImportPath}}/internal/service"
	   "{{.ProjectImportPath}}/vo/protobuf"
   )
   
   type {{.RouterNameUp}}Router interface {
	   CreateRouter()
	   {{- range .Routers}}
	   {{.RouterFuncUp}}(ctx *gin.Context)
	   {{- end}}
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
		{{- range .Routers}}
		
		/** 
		 * @step
		 * @{{.RouterMethod}} {{.RouterPath}}
		 * @{{.Description}}
		 **/
		router.{{.RouterMethod}}("{{.RouterPath}}", {{.FirstRouterName}}.{{.RouterFuncUp}})
		{{- end}}
   }

   {{- range .Routers}}
   /**
   * @description: {{.RouterName}}
   * @param {*gin.Context} ctx
   * @author: Jerry.Yang
   * @date: {{.Time}}
   * @return {*}
   */
   func ({{.FirstRouterName}} *{{.RouterNameUp}}) {{.RouterFuncUp}}(ctx *gin.Context)  {

	   /**
	   * @step
	   * @inputVo
	   **/
	   inputVo := &protobuf.{{.InputReqName}}{}

	   /**
	   * @step
	   * @should bind
	   **/
	   if err := ctx.ShouldBind(inputVo); err != nil {
		   logger.Logger().Errorf("{{.RouterFunc}} shouldBind Err : %+v", err)
		   return 
	   }

	   /**
	   * @step
	   * @调用service
	   **/
	   outputVo, err := service.Create{{.RouterNameUp}}Service().{{.RouterFunc}}(ctx, inputVo)
	   if err != nil {
		   logger.Logger().Errorf("{{.RouterNameUp}}Service {{.RouterFunc}} Err : %+v", err)
		   ctx.JSON(http.StatusOK,&protobuf.{{.OutputRespName}}{RetCode:-1,RetMsg:err.Error()})
		   return
	   }

	   /**
	   * @step
	   * @return
	   **/
	   ctx.JSON(http.StatusOK, outputVo)
   }
   {{- end}}
   `
}
