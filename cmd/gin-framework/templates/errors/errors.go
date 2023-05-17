/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:14:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:25:26
 * @Description: errors
 */
package errors

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type ErrorsInterface interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Errors struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:16:14
 * @return {*}
 */
func (r *Errors) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "errors.go", r.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:15:44
 * @return {*}
 */
func (r *Errors) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: error
	*/
   package errors
   
   import "errors"
   
   // config err
   var Err_Config_SetConfigPath_PathIsEmpty = errors.New("config setConfigPath Err : path is empty")
   
   // service before start
   var Err_Service_BeforeStart_Args = errors.New("service beforeStartService Err : args is err")
   `
}
