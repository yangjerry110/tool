/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-08 11:33:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 15:20:50
 * @Description: new service
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewService interface {
	SaveTemplate(path string, projectPath string, serviceName string, serviceFileName string) error
	SaveAppendFuncTemplate(path string, serviceName string, baseServiceName string) error
	GetTemplate() string
	GetAppendFuncTemplate() string
}

type New struct{}

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
 * @description: GetAppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 18:19:27
 * @return {*}
 */
func (n *New) GetAppendFuncTemplate() string {
	return `/**
	* @description: {{.ServiceNameUp}}
	* @param {context.Context} ctx
	* @param {*input.Test} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstBaseServiceNameUp}} *{{.BaseServiceNameUp}}) {{.ServiceNameUp}}(ctx context.Context, inputVo *input.{{.ServiceNameUp}}) (*output.{{.ServiceNameUp}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   output := &output.{{.ServiceNameUp}}{}
	   return output, nil
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
	   output := &output.{{.ServiceNameUp}}{}
	   return output, nil
   }
   `
}
