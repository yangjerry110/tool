/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 10:31:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-27 11:20:24
 * @Description: common
 */
package logger

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type CommonLogger interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Common struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:34:19
 * @return {*}
 */
func (c *Common) SaveTemplate(path string, projectPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
	}

	data := &Data{ProjectPath: projectPath}
	return templates.CreateCommonTemplate().SaveTemplate(path, "common.go", c.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:34:03
 * @return {*}
 */
func (c *Common) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-21 15:27:42
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-21 16:44:52
	* @Description: logger
	*/
   package logger
   
   import (
	   "{{.ProjectPath}}/config"
   
	   "github.com/yangjerry110/tool/logger"
	   pkgLogger "github.com/yangjerry110/tool/pkg/logger"
   )
   
   type CommonLogger interface {
	   Logger() pkgLogger.LoggerPkgInterface
   }
   
   type Common struct{}
   
   /**
	* @description: CreateLogger
	* @author: Jerry.Yang
	* @date: 2023-04-21 16:04:38
	* @return {*}
	*/
   func (c *Common) Logger() pkgLogger.LoggerPkgInterface {
	   return pkgLogger.SetOptions([]logger.LoggerOptionFunc{
		   pkgLogger.SetCallerDept(config.LoggerSetConfig.CallerDept),
		   pkgLogger.SetLevel(pkgLogger.Level(config.LoggerSetConfig.Level)),
	   })
   }
   `
}
