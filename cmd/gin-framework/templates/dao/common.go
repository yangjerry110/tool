/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:04:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 17:05:48
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
	return templates.CreateCommonTemplate().SaveTemplate(path, "common.go", c.GetTemplate(), nil)
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
	* @Date: 2023-04-23 11:43:53
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 11:43:58
	* @Description: common
	*/
   package dao
   
   type CommonDao interface{}
   
   type Common struct{}
   `
}
