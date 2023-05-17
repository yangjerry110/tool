/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:05:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:18:12
 * @Description: beforStart
 */
package service

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type BeforStartService interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type BeforStart struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:07:45
 * @return {*}
 */
func (b *BeforStart) SaveTemplate(path string, projectPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
		Time        string
	}

	data := &Data{ProjectPath: projectPath, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "beforStart.go", b.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:07:00
 * @return {*}
 */
func (b *BeforStart) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime:  {{.Time}}
	* @Description: before start
	*/
   package service
   
   import (
	   "os"
	   "{{.ProjectPath}}/config"
	   "{{.ProjectPath}}/logger"
   )
   
   type BeforeStartService interface {
	   Preparing() error
   }
   
   type BeforeStart struct{}
   
   /**
	* @description: Preparing
	* @author: Jerry.Yang
	* @date:  {{.Time}}
	* @return {*}
	*/
   func (b *BeforeStart) Preparing() error {
   
	   /**
		* @step
		* @定义参数
		**/
	   configPath := ""
   
	   /**
		* @step
		* @获取命令行参数
		**/
	   args := os.Args
   
	   /**
		* @step
		* @判断args是否大于1，configPath是否有设置
		**/
	   if len(args) > 1 {
		   configPath = args[1]
	   }
   
	   /**
		* @step
		* @设置configPath
		**/
	   if configPath != "" {
		   err := config.CreatePathConfig().SetConfigPath(configPath)
		   if err != nil {
			   logger.Logger().Errorf("beforeStartService SetConfigPath Err : %+v", err)
			   return err
		   }
	   }
   
	   /**
		* @step
		* @设置loggerConfig
		**/
	   err := config.CreateLoggerConfig().SetConfig()
	   if err != nil {
		   logger.Logger().Errorf("beforeStartService SetLoggerConfig Err : %+v", err)
		   return err
	   }
   
	   /**
		* @step
		* @设置routerConfig
		**/
	   err = config.CreateRouterConfig().SetConfig()
	   if err != nil {
		   logger.Logger().Errorf("beforeStartService SetRouterConfig Err : %+v", err)
		   return err
	   }
	   return nil
   }
   `
}
