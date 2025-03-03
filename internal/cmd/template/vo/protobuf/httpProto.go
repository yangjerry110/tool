/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-03 14:03:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 14:37:32
 * @Description: httpProtobuf
 */
package protobuf

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/template"
)

type HttpProtobufService struct {
	FirstServiceName string
	ServiceNameUp    string
	ServiceFuncUp    string
	InputReqName     string
	OutputRespName   string
	Time             string
}

type HttpProtobuf struct {
	ProjectPath       string
	ProjectImportPath string
	ServiceName       string
	ServiceNameUp     string
	FirstServiceName  string
	Time              string
	Services          []*HttpProtobufService
}

// New
//
// new
// Author Jerry.Yang
// Date 2024-03-05 10:49:00
func (i *HttpProtobuf) New() error {
	if err := template.SaveTemplate(fmt.Sprintf("%s/vo/protobuf", i.ProjectPath), fmt.Sprintf("%s_http.go", i.ServiceName), i.getTemplate(), i); err != nil {
		return err
	}
	return nil
}

// getTemplate
//
// 获取模板
// Author Jerry.Yang
// Date 2024-03-05 10:49:24
func (i *HttpProtobuf) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.ServiceNameUp}} httpServer
	*/
   package protobuf
   
   import (
	   "context"
   )

   type {{.ServiceNameUp}}HttpServer interface {
	{{- range .Services}}
	{{.ServiceFuncUp}}(ctx context.Context, inputVo *{{.InputReqName}}) (*{{.OutputRespName}}, error)
	{{- end}}
	}`
}
