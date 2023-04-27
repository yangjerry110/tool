/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:14:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 17:15:45
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
	return templates.CreateCommonTemplate().SaveTemplate(path, "errors.go", r.GetTemplate(), nil)
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
	* @Date: 2023-04-21 16:08:08
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-21 17:06:39
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
