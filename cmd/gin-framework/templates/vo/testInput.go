/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:50:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 17:31:04
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
	return templates.CreateCommonTemplate().SaveTemplate(path, "test.go", t.GetTemplate(), nil)
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
	* @Date: 2023-04-23 14:16:24
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:16:32
	* @Description: test
	*/
   package input
   
   type Test struct {
	   Id int32 ` + " ` json:\"id\"` " + `
   }
   `
}
