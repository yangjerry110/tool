/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:50:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:13:57
 * @Description: test input
 */
package vo

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type TestInputVo interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type TestInput struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:18:34
 * @return {*}
 */
func (t *TestInput) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义数据结构
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "testInputVo.go", t.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:17:58
 * @return {*}
 */
func (t *TestInput) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:16:32
	* @Description:  {{.Time}}
	*/
   package input
   
   /**
    * @description: Test
    * @author: Jerry.Yang
    * @date: {{.Time}}
    * @return {*}
    */   
   type Test struct {
	   Id int32 ` + " ` json:\"id\"` " + `
   }
   `
}
