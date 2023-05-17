/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:04:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:26:47
 * @Description: common
 */
package dao

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type CommonDao interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Common struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:05:55
 * @return {*}
 */
func (c *Common) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "commonDao.go", c.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:05:17
 * @return {*}
 */
func (c *Common) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: common
	*/
   package dao
   
   type CommonDao interface{}
   
   type Common struct{}
   `
}
