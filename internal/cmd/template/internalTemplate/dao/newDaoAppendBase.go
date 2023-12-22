/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 16:48:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 16:50:34
 * @Description:
 */
package dao

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewDaoAppendBase struct {
	DaoName   string
	DaoNameUp string
	Time      string
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 16:50:10
 * @return {*}
 */
func (n *NewDaoAppendBase) New() error {
	// filePath
	filePath := fmt.Sprintf("%s/internal/dao/base.go", config.ProjectPathConf.Path)
	return template.AppendTemplate(filePath, n.getTemplate(), n)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-21 16:50:19
 * @return {*}
 */
func (n *NewDaoAppendBase) getTemplate() string {
	return `/**
	* @description: Create{{.DaoNameUp}}Dao
	* @param {...{{.DaoNameUp}}Dao} {{.DaoName}}Daos
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func Create{{.DaoNameUp}}Dao({{.DaoName}}Daos ...{{.DaoNameUp}}Dao) {{.DaoNameUp}}Dao {
	   if len({{.DaoName}}Daos) == 0 {
		   return &{{.DaoNameUp}}{}
	   }
	   return {{.DaoName}}Daos[0]
   }`
}
