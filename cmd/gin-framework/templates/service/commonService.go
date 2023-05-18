/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-17 18:25:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 18:28:37
 * @Description: common Service
 */
package service

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type CommonService interface {
	SaveTemplate(path string, projectImportPath string) error
	GetTemplate() string
}

type Common struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectImportPath
 * @author: Jerry.Yang
 * @date: 2023-05-17 18:29:23
 * @return {*}
 */
func (c *Common) SaveTemplate(path string, projectImportPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath string
		Time              string
	}

	data := &Data{ProjectImportPath: projectImportPath, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "commonService.go", c.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-17 18:29:14
 * @return {*}
 */
func (c *Common) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: common service
	*/
   package service
   
   import "{{.ProjectImportPath}}/errors"
   
   type CommonService interface {
	   returnError(err error) interface{}
   }
   
   type Common struct{}
   
   /**
	* @description: returnError
	* @param {error} err
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (c *Common) returnError(err error) interface{} {
   
	   /**
		* @step
		* @定义返回结构
		**/
	   type ReturnErr struct {
		   RetCode int32 ` + " ` json:\"retCode\"` " + `
		   RetMsg string ` + " ` json:\"retMsg\"` " + `
	   }
   
	   /**
		* @step
		* @根据error，获取retCode
		**/
	   retCode, ok := errors.ErrCodes[err]
	   if !ok {
		   returnErr := ReturnErr{RetCode: int32(errors.DefaultErrCode), RetMsg: err.Error()}
		   return returnErr
	   }
   
	   /**
		* @step
		* @赋值
		**/
	   returnErr := ReturnErr{RetCode: retCode, RetMsg: err.Error()}
	   return returnErr
   }
   `
}
