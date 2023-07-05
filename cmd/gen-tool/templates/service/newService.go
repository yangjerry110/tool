/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-08 11:33:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-26 11:44:33
 * @Description: new service
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
)

type NewService interface {
	SaveTemplate(path string, projectPath string, serviceName string, serviceFileName string) error
	SaveAppendFuncTemplate(path string, serviceName string, baseServiceName string) error
	SaveProtobufTemplate(path string, projectImportPath string, serviceName string, serviceFileName string, protobufServices []*NewProtobufService) error
	SaveAppendProtobufTemplate(path string, baseServiceName string, serviceName string, newProtobufService *NewProtobufService) error
	GetTemplate() string
	GetAppendFuncTemplate() string
	GetAppendProtobufTemplate() string
}

type New struct{}

/**
 * @description: NewProtobufService
 * @author: Jerry.Yang
 * @date: 2023-05-25 10:44:02
 * @return {*}
 */
type NewProtobufService struct {
	FirstServiceName string
	ServiceNameUp    string
	ServiceFuncUp    string
	InputReqName     string
	OutputRespName   string
	Time             string
}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @param {string} serviceName
 * @author: Jerry.Yang
 * @date: 2023-05-08 11:37:16
 * @return {*}
 */
func (n *New) SaveTemplate(path string, projectImportPath string, serviceName string, serviceFileName string) error {

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath string
		ServiceNameUp     string
		FirstServiceName  string
		Time              string
	}

	/**
	 * @step
	 * @ServiceNameUp进行大写字母的转换
	 **/
	serviceNameUp := templates.CreateCommonTemplate().FirstUpper(serviceName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{ProjectImportPath: projectImportPath, ServiceNameUp: serviceNameUp, FirstServiceName: serviceName[:1], Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, serviceFileName, n.GetTemplate(), data)
}

/**
 * @description: SaveAppendFuncTemplate
 * @param {string} path
 * @param {string} baseServiceName
 * @param {string} serviceName
 * @author: Jerry.Yang
 * @date: 2023-05-11 18:26:29
 * @return {*}
 */
func (n *New) SaveAppendFuncTemplate(path string, serviceName string, baseServiceName string) error {

	/**
	 * @step
	 * @要追加的文件地址
	 **/
	basePath := fmt.Sprintf("%s/%sService.go", path, baseServiceName)

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		BaseServiceNameUp    string
		ServiceNameUp        string
		FirstBaseServiceName string
		Time                 string
	}

	/**
	 * @step
	 * @baseServiceNameUpUp进行大写字母的转换
	 **/
	baseServiceNameUp := templates.CreateCommonTemplate().FirstUpper(baseServiceName)

	/**
	 * @step
	 * @ServiceNameUp进行大写字母的转换
	 **/
	serviceNameUp := templates.CreateCommonTemplate().FirstUpper(serviceName)

	/**
	 * @step
	 * @执行添加
	 **/
	data := &Data{BaseServiceNameUp: baseServiceNameUp, ServiceNameUp: serviceNameUp, FirstBaseServiceName: baseServiceName[:1], Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, n.GetAppendFuncTemplate(), data)
}

/**
 * @description: SaveProtobufTemplate
 * @param {string} path
 * @param {string} projectImportPath
 * @param {string} serviceName
 * @param {string} serviceFileName
 * @param {[]*NewProtobufService} protobufServices
 * @author: Jerry.Yang
 * @date: 2023-05-25 10:50:55
 * @return {*}
 */
func (n *New) SaveProtobufTemplate(path string, projectImportPath string, serviceName string, serviceFileName string, protobufServices []*NewProtobufService) error {

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath string
		ServiceNameUp     string
		FirstServiceName  string
		Time              string
		Services          []*NewProtobufService
	}

	/**
	 * @step
	 * @ServiceNameUp进行大写字母的转换
	 **/
	serviceNameUp := templates.CreateCommonTemplate().FirstUpper(serviceName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{ProjectImportPath: projectImportPath, ServiceNameUp: serviceNameUp, FirstServiceName: serviceName[:1], Time: templates.CreateCommonTemplate().GetFormatNowTime(), Services: protobufServices}
	return templates.CreateCommonTemplate().SaveTemplate(path, serviceFileName, n.GetProtobufTemplate(), data)
}

/**
 * @description: SaveAppendProtobufTemplate
 * @param {string} path
 * @param {string} baseServiceName
 * @param {string} serviceName
 * @param {*NewProtobufService} newProtobufService
 * @author: Jerry.Yang
 * @date: 2023-05-25 21:14:20
 * @return {*}
 */
func (n *New) SaveAppendProtobufTemplate(path string, baseServiceName string, serviceName string, newProtobufService *NewProtobufService) error {

	/**
	 * @step
	 * @要追加的文件地址
	 **/
	basePath := fmt.Sprintf("%s/%sService.go", path, baseServiceName)

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		FirstBaseServiceName string
		BaseServiceName      string
		BaseServiceNameUp    string
		ServiceName          string
		ServiceNameUp        string
		ServiceFuncUp        string
		InputReqName         string
		OutputRespName       string
		Time                 string
	}

	/**
	 * @step
	 * @baseServiceNameUpUp进行大写字母的转换
	 **/
	baseServiceNameUp := templates.CreateCommonTemplate().FirstUpper(baseServiceName)

	/**
	 * @step
	 * @ServiceNameUp进行大写字母的转换
	 **/
	serviceNameUp := templates.CreateCommonTemplate().FirstUpper(serviceName)

	/**
	 * @step
	 * @执行添加
	 **/
	data := &Data{BaseServiceName: baseServiceName, ServiceName: serviceName, BaseServiceNameUp: baseServiceNameUp, ServiceNameUp: serviceNameUp, ServiceFuncUp: newProtobufService.ServiceFuncUp, FirstBaseServiceName: baseServiceName[:1], InputReqName: newProtobufService.InputReqName, OutputRespName: newProtobufService.OutputRespName, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, n.GetAppendProtobufTemplate(), data)
}

/**
 * @description: GetAppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 18:19:27
 * @return {*}
 */
func (n *New) GetAppendFuncTemplate() string {
	return `/**
	* @description: {{.ServiceNameUp}}
	* @param {context.Context} ctx
	* @param {*input.{{.ServiceNameUp}}} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstBaseServiceName}} *{{.BaseServiceNameUp}}) {{.ServiceNameUp}}(ctx context.Context, inputVo *input.{{.ServiceNameUp}}) (*output.{{.ServiceNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
		result := &output.{{.ServiceNameUp}}{}
	   return result, nil
   }`
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-08 11:37:05
 * @return {*}
 */
func (n *New) GetTemplate() string {
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
	   "{{.ProjectImportPath}}/vo/input"
	   "{{.ProjectImportPath}}/vo/output"
   )
   
   type {{.ServiceNameUp}}Service interface {
		{{.ServiceNameUp}}(ctx context.Context, inputVo *input.{{.ServiceNameUp}}) (*output.{{.ServiceNameUp}}, error)
   }
   
   type {{.ServiceNameUp}} struct{}
   
   /**
	* @description: {{.ServiceNameUp}}
	* @param {context.Context} ctx
	* @param {*input.{{.ServiceNameUp}}} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstServiceName}} *{{.ServiceNameUp}}) {{.ServiceNameUp}}(ctx context.Context, inputVo *input.{{.ServiceNameUp}}) (*output.{{.ServiceNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := &output.{{.ServiceNameUp}}{}
	   return result, nil
   }
   `
}

func (n *New) GetProtobufTemplate() string {
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

/**
 * @description: GetAppendProtobufTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-25 21:07:56
 * @return {*}
 */
func (n *New) GetAppendProtobufTemplate() string {
	return `
	/**
	* @description: {{.ServiceFuncUp}}
	* @param {context.Context} ctx
	* @param {*protobuf.{{.InputReqName}}} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstBaseServiceName}} *{{.BaseServiceNameUp}}) {{.ServiceFuncUp}}(ctx context.Context, inputVo *protobuf.{{.InputReqName}}) (*protobuf.{{.OutputRespName}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   result := &protobuf.{{.OutputRespName}}{}
	   return result, nil
   }
	`
}
