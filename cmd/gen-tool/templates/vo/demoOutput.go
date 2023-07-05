/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 17:18:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:59:13
 * @Description: Demo output
 */
package vo

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type DemoOutputVo interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type DemoOutput struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:22:22
 * @return {*}
 */
func (d *DemoOutput) SaveTemplate(path string) error {
	/**
	 * @step
	 * @定义数据结构
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "demoOutputVo.go", d.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:21:49
 * @return {*}
 */
func (d *DemoOutput) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:17:14
	* @Description: {{.Time}}
	*/
   package output
   
   /**
    * @description: Demo
    * @author: Jerry.Yang
    * @date: {{.Time}}
    * @return {*}
    */   
   type Demo struct {
	   RetCode   int32  ` + " ` json:\"retCode\"` " + `
	   RetMsg    string ` + " ` json:\"retMsg\"` " + `
	   RetResult bool   ` + " ` json:\"retResult\"` " + `
   }
   `
}
