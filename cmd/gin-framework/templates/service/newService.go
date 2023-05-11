/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-08 11:33:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-09 16:54:58
 * @Description: new service
 */
package service

import (
	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewService interface {
	SaveTemplate(path string, projectPath string, appName string, serviceName string) error
	GetTemplate() string
}

type New struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @param {string} appName
 * @param {string} serviceName
 * @author: Jerry.Yang
 * @date: 2023-05-08 11:37:16
 * @return {*}
 */
func (n *New) SaveTemplate(path string, projectImportPath string, appName string, serviceName string) error {

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath string
		AppName           string
	}

	/**
	 * @step
	 * @appName进行大写字母的转换
	 **/
	appName = templates.CreateCommonTemplate().FirstUpper(appName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{ProjectImportPath: projectImportPath, AppName: appName}
	return templates.CreateCommonTemplate().SaveTemplate(path, serviceName, n.GetTemplate(), data)
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
	* @Date: 2023-04-23 14:21:15
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:37:13
	* @Description: {{.AppName}} service
	*/
   package service
   
   import (
	   "context"
	   "{{.ProjectImportPath}}/vo/input"
	   "{{.ProjectImportPath}}/vo/output"
   )
   
   type {{.AppName}}Service interface {
	   Test(ctx context.Context, inputVo *input.{{.AppName}}) (*output.{{.AppName}}, error)
   }
   
   type {{.AppName}} struct{}
   
   /**
	* @description: {{.AppName}}
	* @param {context.Context} ctx
	* @param {*input.Test} inputVo
	* @author: Jerry.Yang
	* @date: 2023-04-23 14:23:07
	* @return {*}
	*/
   func (t *{{.AppName}}) {{.AppName}}(ctx context.Context, inputVo *input.{{.AppName}}) (*output.{{.AppName}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   output := &output.{{.AppName}}{}
	   output.RetCode = 0
	   output.RetMsg = ""
	   output.RetResult = true
	   return output, nil
   }
   `
}
