/*
 * @Author: Jerry.Yang
 * @Date: 2024-03-05 10:45:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-03-05 14:39:43
 * @Description: NewInterfaceProtobuf
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewInterfaceProtobufService struct {
	FirstServiceName string
	ServiceNameUp    string
	ServiceFuncUp    string
	InputReqName     string
	OutputRespName   string
	Time             string
}

type NewInterfaceProtobuf struct {
	ProjectPath       string
	ProjectImportPath string
	ServiceName       string
	ServiceNameUp     string
	FirstServiceName  string
	Time              string
	Services          []*NewInterfaceProtobufService
}

// New
//
// new
// Author Jerry.Yang
// Date 2024-03-05 10:49:00
func (i *NewInterfaceProtobuf) New() error {
	if err := template.SaveTemplate(fmt.Sprintf("%s/internal/service/interfaceService", i.ProjectPath), fmt.Sprintf("%sInterfaceService.go", i.ServiceName), i.getTemplate(), i); err != nil {
		return err
	}
	return nil
}

// getTemplate
//
// 获取模板
// Author Jerry.Yang
// Date 2024-03-05 10:49:24
func (i *NewInterfaceProtobuf) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.ServiceNameUp}} service
	*/
   package interfaceService
   
   import (
	   "context"
	   "{{.ProjectImportPath}}/vo/protobuf"
   )

   type {{.ServiceNameUp}}Service interface {
	{{- range .Services}}
	{{.ServiceFuncUp}}(ctx context.Context, inputVo *protobuf.{{.InputReqName}}) (*protobuf.{{.OutputRespName}}, error)
	{{- end}}
	}`
}
