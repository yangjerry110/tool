/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-06 11:17:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-09 16:54:45
 * @Description: newController
 */
package controller

import (
	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewController interface {
	SaveTemplate(path string, projectName string, appName string, controllerFileName string) error
	GetTemplate() string
}

type New struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectName
 * @param {string} appName
 * @param {string} controllerFileName
 * @author: Jerry.Yang
 * @date: 2023-05-06 11:29:06
 * @return {*}
 */
func (n *New) SaveTemplate(path string, projectImportPath string, appName string, controllerFileName string) error {

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
	return templates.CreateCommonTemplate().SaveTemplate(path, controllerFileName, n.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-06 11:25:32
 * @return {*}
 */
func (n *New) GetTemplate() string {

	return `
	/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-23 11:44:49
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:30:54
	* @Description: {{.AppName}}
	*/
	package controller

	import (
		"net/http"
		"{{.ProjectImportPath}}/logger"
		"{{.ProjectImportPath}}/service"
		"{{.ProjectImportPath}}/vo/input"

		"github.com/gin-gonic/gin"
	)

	type {{.AppName}}Controller interface {
		{{.AppName}}(ctx *gin.Context) error
	}

	type {{.AppName}} struct{}

	/**
	* @description: {{.AppName}}
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: 2023-04-23 11:45:45
	* @return {*}
	*/
	func (t *{{.AppName}}) {{.AppName}}(ctx *gin.Context) error {

		/**
		* @step
		* @inputVo
		**/
		inputVo := &input.{{.AppName}}{}

		/**
		* @step
		* @should bind
		**/
		if err := ctx.ShouldBind(inputVo); err != nil {
			logger.Logger().Errorf("{{.AppName}}Controller {{.AppName}} shouldBind Err : %+v", err)
			return err
		}

		/**
		* @step
		* @调用service
		**/
		outputVo, err := service.Create{{.AppName}}Service().{{.AppName}}(ctx, inputVo)
		if err != nil {
			logger.Logger().Errorf("{{.AppName}}Controller {{.AppName}}Service Test Err : %+v", err)
		}

		/**
		* @step
		* @return
		**/
		ctx.JSON(http.StatusOK, outputVo)
		return nil
	}
	`
}
