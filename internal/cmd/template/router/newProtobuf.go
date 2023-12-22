/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 16:19:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 14:33:43
 * @Description: new protobuf
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/template"
)

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
	SwaggerNotes    string
	Time            string
}

type NewProtobuf struct {
	ProjectPath       string
	ProjectImportPath string
	FirstRouterName   string
	RouterNameUp      string
	RouterName        string
	Time              string
	Routers           []*NewProtobufRouter
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-12 17:03:48
 * @return {*}
 */
func (n *NewProtobuf) New() error {
	if err := template.SaveTemplate(fmt.Sprintf("%s/router", n.ProjectPath), fmt.Sprintf("%sRouter.go", n.RouterName), n.getTemplate(), n); err != nil {
		return err
	}
	return nil
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-12 16:43:05
 * @return {*}
 */
func (n *NewProtobuf) getTemplate() string {
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
	   "git.qutoutiao.net/gopher/qms/pkg/qlog"
	   "github.com/yangjerry110/tool/router"
	   "{{.ProjectImportPath}}/internal/service"
	   "{{.ProjectImportPath}}/vo/protobuf"
   )
   
   type {{.RouterNameUp}}Router interface {
	   Register() error
	   {{- range .Routers}}
	   {{.RouterFuncUp}}(ctx *gin.Context)
	   {{- end}}
   }
   
   type {{.RouterNameUp}} struct{}
   
   /**
	* @description: Register
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstRouterName}} *{{.RouterNameUp}}) Register() error {
   
	   /**
		* @step
		* @获取router
		**/
	   router := router.GetGinDefaultRouter()
		{{- range .Routers}}
		
		/** 
		* @step
		* @{{.RouterMethod}} {{.RouterPath}}
		* @{{.Description}}
		**/
	   router.{{.RouterMethod}}("{{.RouterPath}}", {{.FirstRouterName}}.{{.RouterFuncUp}})
	   {{- end}}
	   return nil
  }

  {{- range .Routers}}
  {{.SwaggerNotes}}
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
		  qlog.Errorf("{{.RouterFunc}} shouldBind Err : %+v", err)
		  return 
	  } 
	  
	  /**
	  * @step
	  * @调用service
	  **/
	  outputVo, err := service.Create{{.RouterNameUp}}Service().{{.RouterFunc}}(ctx, inputVo)
	  if err != nil {
		  qlog.Errorf("{{.RouterNameUp}}Service {{.RouterFunc}} Err : %+v", err)
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