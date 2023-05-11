/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-08 15:05:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-09 16:55:04
 * @Description: new vo
 */
package vo

import (
	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewVo interface {
	SaveTemplate(inputPath string, outputPath string, inputVoName string, outputVoName string, appName string) error
}

type New struct{}

/**
 * @description: SaveTemplate
 * @param {string} inputPath
 * @param {string} outputPath
 * @param {string} inputVoName
 * @param {string} outputVoName
 * @param {string} appName
 * @author: Jerry.Yang
 * @date: 2023-05-08 15:09:48
 * @return {*}
 */
func (n *New) SaveTemplate(inputPath string, outputPath string, inputVoName string, outputVoName string, appName string) error {

	/**
	 * @step
	 * @需要渲染的数据
	 **/
	type Data struct {
		AppName string
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
	data := &Data{AppName: appName}

	/**
	 * @step
	 * @input
	 **/
	err := templates.CreateCommonTemplate().SaveTemplate(inputPath, inputVoName, n.GetInputTemplate(), data)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @output
	 **/
	err = templates.CreateCommonTemplate().SaveTemplate(outputPath, outputVoName, n.GetOutputTemplate(), data)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: GetInputTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-09 16:22:50
 * @return {*}
 */
func (n *New) GetInputTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-23 14:16:24
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:16:32
	* @Description: {{.AppName}} inputVo
	*/
   package input
   
   type {{.AppName}} struct {
	   
   }`
}

/**
 * @description: GetOutputTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-09 16:22:41
 * @return {*}
 */
func (n *New) GetOutputTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-23 14:16:24
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:16:32
	* @Description: {{.AppName}} outputVo
	*/
   package output
   
   type {{.AppName}} struct {
	   
   }`
}
