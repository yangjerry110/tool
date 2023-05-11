/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:01:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 15:24:36
 * @Description: base
 */
package dao

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type BaseDao interface {
	SaveTemplate(path string) error
	GetTemplate() string
	AppendFuncTemplate(path string, daoName string) error
	GetAppendFuncTemplate() string
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
 * @description: AppendFunc
 * @param {string} path
 * @param {string} daoName
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:13:22
 * @return {*}
 */
func (b *Base) AppendFuncTemplate(path string, daoName string) error {

	/**
	 * @step
	 * @获取base的文件地址
	 **/
	basePath := fmt.Sprintf("%s/%s", path, "base.go")

	/**
	 * @step
	 * @定义需要渲染的数据结构
	 **/
	type Data struct {
		DaoName   string
		DaoNameUp string
	}

	/**
	 * @step
	 * @渲染参数
	 **/
	data := &Data{DaoName: daoName, DaoNameUp: templates.CreateCommonTemplate().FirstUpper(daoName)}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, b.GetAppendFuncTemplate(), data)
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

/**
 * @description: GetAppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:24:35
 * @return {*}
 */
func (b *Base) GetAppendFuncTemplate() string {
	return `/**
	* @description: Create{{.DaoNameUp}}Dao
	* @param {...{{.DaoNameUp}}Dao} {{.DaoName}}Daos
	* @author: Jerry.Yang
	* @date: 2023-04-24 17:02:59
	* @return {*}
	*/
   func Create{{.DaoNameUp}}Dao({{.DaoName}}Daos ...{{.DaoNameUp}}Dao) {{.DaoNameUp}}Dao {
	   if len({{.DaoName}}Daos) == 0 {
		   return &{{.DaoNameUp}}{}
	   }
	   return {{.DaoName}}Daos[0]
   }`
}
