/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 11:28:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 14:59:59
 * @Description: template config
 */
package config

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type BaseConfig interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Base struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:01:21
 * @return {*}
 */
func (b *Base) SaveTemplate(path string) error {
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), nil)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 14:41:57
 * @return {*}
 */
func (b *Base) GetTemplate() string {

	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-21 15:15:48
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-21 17:28:31
	* @Description: base
	*/
   package config
   
   /**
	* @description: CreatePathConfig
	* @param {...PathConfig} PathConfigs
	* @author: Jerry.Yang
	* @date: 2023-04-21 16:28:37
	* @return {*}
	*/
   func CreatePathConfig(PathConfigs ...PathConfig) PathConfig {
	   if len(PathConfigs) == 0 {
		   return &Path{}
	   }
	   return PathConfigs[0]
   }
   
   /**
	* @description: CreateLoggerConfig
	* @param {...LoggerConfig} LoggerConfigs
	* @author: Jerry.Yang
	* @date: 2023-04-21 16:33:04
	* @return {*}
	*/
   func CreateLoggerConfig(LoggerConfigs ...LoggerConfig) LoggerConfig {
	   if len(LoggerConfigs) == 0 {
		   return &Logger{}
	   }
	   return LoggerConfigs[0]
   }
   
   /**
	* @description: CreateRouterConfig
	* @param {...RouterConfig} RouterConfigs
	* @author: Jerry.Yang
	* @date: 2023-04-21 17:28:48
	* @return {*}
	*/
   func CreateRouterConfig(RouterConfigs ...RouterConfig) RouterConfig {
	   if len(RouterConfigs) == 0 {
		   return &Router{}
	   }
	   return RouterConfigs[0]
   }
   `
}
