/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-17 16:24:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 16:27:44
 * @Description: database config
 */
package config

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type DataBaseConfig interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type DataBase struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-05-17 16:26:27
 * @return {*}
 */
func (d *DataBase) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "database.go", d.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-17 16:25:55
 * @return {*}
 */
func (d *DataBase) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: database config
	*/
   package config
   
   import "github.com/yangjerry110/tool/pkg/db"
   
   type DataBaseConfig interface {
	   SetConfig() error
   }
   
   type DataBase struct{}
   
   /**
	* @description: SetConfig
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (d *DataBase) SetConfig() error {
   
	   /**
	   * @step
	   * @获取configPath
	   **/
	   configPath, err := CreatePathConfig().GetConfigPath()
	   if err != nil {
		   return err
	   }
   
	   /**
		* @step
		* @setDatabaseConfig
		**/
	   return db.RenderDbConfig(configPath, "database.yaml")
   }
   `
}
