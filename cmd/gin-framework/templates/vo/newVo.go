/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-08 15:05:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:10:09
 * @Description: new vo
 */
package vo

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type NewVo interface {
	SaveTemplate(inputPath string, outputPath string, inputVoName string, outputVoName string, VoName string) error
	SaveAppendFuncInputTemplate(path string, voName string, baseVoName string) error
	SaveAppendFuncOutputTemplate(path string, voName string, baseVoName string) error
	GetAppendFuncInputTemplate() string
	GetAppendFuncOutputTemplate() string
	GetInputTemplate() string
	GetOutputTemplate() string
}

type New struct{}

/**
 * @description: SaveTemplate
 * @param {string} inputPath
 * @param {string} outputPath
 * @param {string} inputVoName
 * @param {string} outputVoName
 * @param {string} voName
 * @author: Jerry.Yang
 * @date: 2023-05-08 15:09:48
 * @return {*}
 */
func (n *New) SaveTemplate(inputPath string, outputPath string, inputVoName string, outputVoName string, voName string) error {

	/**
	 * @step
	 * @需要渲染的数据
	 **/
	type Data struct {
		VoName   string
		VoNameUp string
		Time     string
	}

	/**
	 * @step
	 * @VoName进行大写字母的转换
	 **/
	voNameUp := templates.CreateCommonTemplate().FirstUpper(voName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{VoName: voName, VoNameUp: voNameUp, Time: templates.CreateCommonTemplate().GetFormatNowTime()}

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
 * @description: SaveAppendFuncInputTemplate
 * @param {string} path
 * @param {string} voName
 * @param {string} baseVoName
 * @author: Jerry.Yang
 * @date: 2023-05-16 10:45:08
 * @return {*}
 */
func (n *New) SaveAppendFuncInputTemplate(path string, voName string, baseVoName string) error {

	/**
	 * @step
	 * @inputVoPath
	 **/
	baseInputVoPath := fmt.Sprintf("%s/%sInputVo.go", path, baseVoName)

	/**
	 * @step
	 * @需要渲染的数据
	 **/
	type Data struct {
		VoName   string
		VoNameUp string
		Time     string
	}

	/**
	 * @step
	 * @VoName进行大写字母的转换
	 **/
	voNameUp := templates.CreateCommonTemplate().FirstUpper(voName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{VoName: voName, VoNameUp: voNameUp, Time: templates.CreateCommonTemplate().GetFormatNowTime()}

	/**
	 * @step
	 * @执行追加
	 **/
	return templates.CreateCommonTemplate().AppendTemplate(baseInputVoPath, n.GetAppendFuncInputTemplate(), data)
}

/**
 * @description: SaveAppendFuncOutputTemplate
 * @param {string} path
 * @param {string} voName
 * @param {string} baseVoName
 * @author: Jerry.Yang
 * @date: 2023-05-16 10:48:00
 * @return {*}
 */
func (n *New) SaveAppendFuncOutputTemplate(path string, voName string, baseVoName string) error {

	/**
	 * @step
	 * @baseOutputVoPath
	 **/
	baseOutputVoPath := fmt.Sprintf("%s/%sOutputVo.go", path, baseVoName)

	/**
	 * @step
	 * @需要渲染的数据
	 **/
	type Data struct {
		VoName   string
		VoNameUp string
		Time     string
	}

	/**
	 * @step
	 * @VoName进行大写字母的转换
	 **/
	voNameUp := templates.CreateCommonTemplate().FirstUpper(voName)

	/**
	 * @step
	 * @进行赋值
	 **/
	data := &Data{VoName: voName, VoNameUp: voNameUp, Time: templates.CreateCommonTemplate().GetFormatNowTime()}

	/**
	 * @step
	 * @执行追加
	 **/
	return templates.CreateCommonTemplate().AppendTemplate(baseOutputVoPath, n.GetAppendFuncOutputTemplate(), data)
}

/**
 * @description: GetAppendFuncInputTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-12 11:00:48
 * @return {*}
 */
func (n *New) GetAppendFuncInputTemplate() string {
	return `/**
	* @description: {{.VoNameUp}}
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/	
	type {{.VoNameUp}} struct {}`
}

/**
 * @description: GetAppendFuncOutputTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-12 11:16:08
 * @return {*}
 */
func (n *New) GetAppendFuncOutputTemplate() string {
	return `type {{.VoNameUp}} struct {}`
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
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.VoName}} inputVo
	*/
   package input
   
   type {{.VoNameUp}} struct {
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
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.VoName}} outputVo
	*/
   package output
   
   type {{.VoNameUp}} struct {
	   
   }`
}
