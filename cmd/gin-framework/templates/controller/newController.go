/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-06 11:17:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:53:14
 * @Description: newController
 */
package controller

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewController interface {
	SaveTemplate(path string, projectName string, ControllerName string, controllerFileName string) error
	GetTemplate() string
	SaveAppendFuncTemplate(path string, controllerName string, baseControllerName string) error
	GetAppendFuncTemplate() string
}

type New struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectName
 * @param {string} controllerName
 * @param {string} controllerFileName
 * @author: Jerry.Yang
 * @date: 2023-05-06 11:29:06
 * @return {*}
 */
func (n *New) SaveTemplate(path string, projectImportPath string, controllerName string, controllerFileName string) error {

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath   string
		ControllerName      string
		FirstControllerName string
		Time                string
	}

	/**
	 * @step
	 * @ControllerName进行大写字母的转换
	 **/
	controllerNameUp := templates.CreateCommonTemplate().FirstUpper(controllerName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{ProjectImportPath: projectImportPath, ControllerName: controllerNameUp, FirstControllerName: controllerName[:1], Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, controllerFileName, n.GetTemplate(), data)
}

/**
 * @description: SaveAppendTemplate
 * @param {string} path
 * @param {string} controllerName
 * @param {string} baseControllerName
 * @author: Jerry.Yang
 * @date: 2023-05-11 17:10:15
 * @return {*}
 */
func (n *New) SaveAppendFuncTemplate(path string, controllerName string, baseControllerName string) error {

	/**
	 * @step
	 * @组合baseController的path
	 **/
	basePath := fmt.Sprintf("%s/%sController.go", path, baseControllerName)

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ControllerName      string
		FirstControllerName string
		BaseControllerName  string
		Time                string
	}

	/**
	 * @step
	 * @ControllerName进行大写字母的转换
	 **/
	controllerNameUp := templates.CreateCommonTemplate().FirstUpper(controllerName)

	/**
	 * @step
	 * @baseControllerName进行大写字母的转换
	 **/
	baseControllerNameUp := templates.CreateCommonTemplate().FirstUpper(baseControllerName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{ControllerName: controllerNameUp, FirstControllerName: controllerName[:1], BaseControllerName: baseControllerNameUp, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, n.GetAppendFuncTemplate(), data)
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
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.ControllerName}}
	*/
	package controller

	import (
		"net/http"
		"{{.ProjectImportPath}}/logger"
		"{{.ProjectImportPath}}/service"
		"{{.ProjectImportPath}}/vo/input"

		"github.com/gin-gonic/gin"
	)

	type {{.ControllerName}}Controller interface {
		{{.ControllerName}}(ctx *gin.Context) error
	}

	type {{.ControllerName}} struct{}

	/**
	* @description: {{.ControllerName}}
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
	func ({{.FirstControllerName}} *{{.ControllerName}}) {{.ControllerName}}(ctx *gin.Context) error {

		/**
		* @step
		* @inputVo
		**/
		inputVo := &input.{{.ControllerName}}{}

		/**
		* @step
		* @should bind
		**/
		if err := ctx.ShouldBind(inputVo); err != nil {
			logger.Logger().Errorf("{{.ControllerName}}Controller {{.ControllerName}} shouldBind Err : %+v", err)
			return err
		}

		/**
		* @step
		* @调用service
		**/
		outputVo, err := service.Create{{.ControllerName}}Service().{{.ControllerName}}(ctx, inputVo)
		if err != nil {
			logger.Logger().Errorf("{{.ControllerName}}Controller {{.ControllerName}}Service {{.ControllerName}} Err : %+v", err)
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

/**
 * @description: GetAppendTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 17:07:09
 * @return {*}
 */
func (n *New) GetAppendFuncTemplate() string {
	return `/**
	* @description: {{.ControllerName}}
	* @param {*gin.Context} ctx
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
	func ({{.FirstControllerName}} *{{.BaseControllerName}}) {{.ControllerName}}(ctx *gin.Context) error {

		/**
		* @step
		* @inputVo
		**/
		inputVo := &input.{{.ControllerName}}{}

		/**
		* @step
		* @should bind
		**/
		if err := ctx.ShouldBind(inputVo); err != nil {
			logger.Logger().Errorf("{{.BaseControllerName}}Controller {{.ControllerName}} shouldBind Err : %+v", err)
			return err
		}

		/**
		* @step
		* @调用service
		**/
		outputVo, err := service.Create{{.BaseControllerName}}Service().{{.ControllerName}}(ctx, inputVo)
		if err != nil {
			logger.Logger().Errorf("{{.ControllerName}}Controller {{.BaseControllerName}}Service {{.ControllerName}} Err : %+v", err)
		}

		/**
		* @step
		* @return
		**/
		ctx.JSON(http.StatusOK, outputVo)
		return nil
	}`
}
