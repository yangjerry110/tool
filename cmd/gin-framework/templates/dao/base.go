/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:01:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 17:03:10
 * @Description: base
 */
package dao

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type BaseDao interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Base struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:03:49
 * @return {*}
 */
func (b *Base) SaveTemplate(path string) error {
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), nil)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:03:41
 * @return {*}
 */
func (b *Base) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-23 11:40:11
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 11:40:18
	* @Description: base
	*/
   package dao
   
   /**
	* @description: CreateCommonDao
	* @param {...CommonDao} commonDaos
	* @author: Jerry.Yang
	* @date: 2023-04-24 17:02:59
	* @return {*}
	*/
   func CreateCommonDao(commonDaos ...CommonDao) CommonDao {
	   if len(commonDaos) == 0 {
		   return &Common{}
	   }
	   return commonDaos[0]
   }
   `
}
