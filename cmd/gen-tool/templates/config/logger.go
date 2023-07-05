/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 15:23:07
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:55:25
 * @Description: logger
 */
package config

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type LoggerConfig interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Logger struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:37:29
 * @return {*}
 */
func (l *Logger) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "logger.go", l.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:36:33
 * @return {*}
 */
func (l *Logger) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 11:32:20
	* @Description: logger config
	*/
   package config
   
   import (
	   "time"
   
	   "github.com/sirupsen/logrus"
	   "github.com/yangjerry110/tool/pkg/conf"
   )
   
   type LoggerConfig interface {
	   SetConfig() error
   }
   
   type Logger struct {
	   Level      logrus.Level ` + " ` yaml:\"level\"` " + `
	   CallerDept int          ` + " ` yaml:\"callerDept\"` " + `
   }
   
   /**
	* @description: LoggerSetConfig
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   var LoggerSetConfig = &Logger{}
   
   /**
	* @description: SetConfig
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (l *Logger) SetConfig() error {
   
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
		* @渲染配置
		**/
	   err = conf.GetConf(configPath, "logger.yaml", "yaml", 60*time.Second, LoggerSetConfig)
	   if err != nil {
		   return err
	   }
	   return nil
   }
   `
}
