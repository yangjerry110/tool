/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 16:19:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 15:47:04
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

type NewProtobufService struct {
	ServiceName  string
	RouterNameUp string
}

type NewProtobuf struct {
	ExtendImportPath  string
	ProjectPath       string
	ProjectImportPath string
	FirstRouterName   string
	RouterNameUp      string
	RouterName        string
	Time              string
	Services          []*NewProtobufService
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
	   "google.golang.org/grpc"
	   "{{.ProjectImportPath}}/vo/protobuf"
	   "github.com/yangjerry110/protoc-gen-go/proto"
   )
   
   type {{.RouterNameUp}} struct{
	HttpServer protobuf.{{.RouterNameUp}}HttpServer
   }
   
   /**
	* @description: Register
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstRouterName}} *{{.RouterNameUp}}) Register(router *gin.Engine) error {
   
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

  /**
	* @description: RegisterGrpc
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
	func ({{.FirstRouterName}} *{{.RouterNameUp}}) RegisterGrpc(grpc *grpc.Server) error {
		{{- range .Services}}
		// register {{.ServiceName}}
		protobuf.Register{{.ServiceName}}Server(grpc, &protobuf.Unimplemented{{.RouterNameUp}}Server{})
		{{- end}}
		return nil
	}

	/**
	* @description: RegisterService
	* @param {interface{}} service
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
	func ({{.FirstRouterName}} *{{.RouterNameUp}}) RegisterService(service interface{}) error {
		r.HttpServer = service.(protobuf.{{.RouterNameUp}}HttpServer)
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
	  outputVo, err := {{.FirstRouterName}}.HttpServer.{{.RouterFunc}}(ctx, inputVo)
	  if err != nil {
		  qlog.Errorf("{{.RouterNameUp}}Service {{.RouterFunc}} Err : %+v", err)
		  ctx.JSON(http.StatusBadRequest,&protobuf.{{.OutputRespName}}{RetCode:proto.Int32(-1),RetMsg: proto.String(err.Error())})
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
