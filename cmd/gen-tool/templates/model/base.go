/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 11:34:13
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 11:37:07
 * @Description: base
 */
package model

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type BaseModel interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Base struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-05-18 11:37:15
 * @return {*}
 */
func (b *Base) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-18 11:37:06
 * @return {*}
 */
func (b *Base) GetTemplate() string {
	return `package model
	
	/**
	* @description: delete标识
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/	
	var Is_Deleted = 1
	var No_Deleted = 0
	`
}
