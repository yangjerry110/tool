/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 15:31:06
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 16:30:45
 * @Description: path
 */
package config

import (
	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
)

type PathConfig interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Path struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectName
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:54:19
 * @return {*}
 */
func (p *Path) SaveTemplate(path string, projectPath string) error {

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
	}

	data := &Data{ProjectPath: projectPath}
	return templates.CreateCommonTemplate().SaveTemplate(path, "path.go", p.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:54:34
 * @return {*}
 */
func (p *Path) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-21 15:16:29
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 10:35:12
	* @Description: common
	*/
   package config
   
   import (
	   "fmt"
	   "os"
	   "{{.ProjectPath}}/errors"
   )
   
   type PathConfig interface {
	   SetConfigPath(path string) error
	   GetConfigPath() (string, error)
   }
   
   type Path struct{}
   
   /**
	* @description: configPath
	* @author: Jerry.Yang
	* @date: 2023-04-21 15:25:45
	* @return {*}
	*/
   var configPath = ""
   
   /**
	* @description: SetConfigPath
	* @param {string} configPath
	* @author: Jerry.Yang
	* @date: 2023-04-21 15:22:17
	* @return {*}
	*/
   func (p *Path) SetConfigPath(path string) error {
   
	   /**
		* @step
		* @判断是否为空
		**/
	   if path == "" {
		   return errors.Err_Config_SetConfigPath_PathIsEmpty
	   }
	   configPath = path
	   return nil
   }
   
   /**
	* @description: GetConfigPath
	* @author: Jerry.Yang
	* @date: 2023-04-21 15:25:00
	* @return {*}
	*/
   func (p *Path) GetConfigPath() (string, error) {
   
	   /**
		* @step
		* @判断configPath是否有值
		**/
	   if configPath != "" {
		   return configPath, nil
	   }
   
	   /**
		* @step
		* @获取当前目录+config
		**/
	   localPath, err := p.GetLocalPath()
	   if err != nil {
		   return "", err
	   }
   
	   /**
		* @step
		* @设置configPath
		**/
	   setConfigPath := fmt.Sprintf("%s/config/yamlConfig", localPath)
	   err = p.SetConfigPath(setConfigPath)
	   if err != nil {
		   return "", err
	   }
   
	   /**
		* @step
		* @定义默认的config的目录
		**/
	   return configPath, nil
   }
   
   /**
	* @description: GetLocalPath
	* @author: Jerry.Yang
	* @date: 2023-04-21 16:18:30
	* @return {*}
	*/
   func (p *Path) GetLocalPath() (string, error) {
   
	   /**
		* @step
		* @获取当前目录
		**/
	   thisPath, err := os.Getwd()
	   if err != nil {
		   return "", err
	   }
	   return thisPath, nil
   }
   `
}
