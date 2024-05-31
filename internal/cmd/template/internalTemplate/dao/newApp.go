/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:03:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-26 16:09:58
 * @Description: newApp
 */
package dao

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppBaseDao struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:06:06
 * @return {*}
 */
func (n *NewAppBaseDao) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	filePath := fmt.Sprintf("%s/internal/dao", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "base.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:05:57
 * @return {*}
 */
func (n *NewAppBaseDao) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: base
	*/
   package dao
   
   import (
	   "context"
	   "github.com/yangjerry110/tool/db"
	   "gorm.io/gorm"
   )
   
   /**
	* @description: CreateClient
	* @param {string} dbName
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateClient(ctx context.Context,dbName string) *gorm.DB {
   
	   dbClient, err := db.CreateGormDb().GetClient(ctx,dbName)
	   if err != nil {
		   panic(err)
	   }
	   return dbClient
   }
   `
}
