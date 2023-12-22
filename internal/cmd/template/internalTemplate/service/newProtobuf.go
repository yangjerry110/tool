/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 16:18:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-14 16:04:34
 * @Description: new protobuf
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewProtobufService struct {
	FirstServiceName string
	ServiceNameUp    string
	ServiceFuncUp    string
	InputReqName     string
	OutputRespName   string
	Time             string
}

type NewProtobuf struct {
	ProjectPath       string
	ProjectImportPath string
	ServiceName       string
	ServiceNameUp     string
	FirstServiceName  string
	Time              string
	Services          []*NewProtobufService
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-14 16:04:39
 * @return {*}
 */
func (n *NewProtobuf) New() error {
	if err := template.SaveTemplate(fmt.Sprintf("%s/internal/service", n.ProjectPath), fmt.Sprintf("%sService.go", n.ServiceName), n.getTemplate(), n); err != nil {
		return err
	}
	return nil
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-14 16:04:49
 * @return {*}
 */
func (n *NewProtobuf) getTemplate() string {

	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.ServiceNameUp}} service
	*/
   package service
   
   import (
	   "context"
	   "{{.ProjectImportPath}}/vo/protobuf"
   )

   type {{.ServiceNameUp}}Service interface {
	{{- range .Services}}
	{{.ServiceFuncUp}}(ctx context.Context, inputVo *protobuf.{{.InputReqName}}) (*protobuf.{{.OutputRespName}}, error)
	{{- end}}
	}

	type {{.ServiceNameUp}} struct{}

	{{- range .Services}}
	/**
	* @description: {{.ServiceFuncUp}}
	* @param {context.Context} ctx
	* @param {*protobuf.{{.InputReqName}}} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstServiceName}} *{{.ServiceNameUp}}) {{.ServiceFuncUp}}(ctx context.Context, inputVo *protobuf.{{.InputReqName}}) (*protobuf.{{.OutputRespName}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := &protobuf.{{.OutputRespName}}{}
	   return result, nil
   }
   {{- end}}
   `
}