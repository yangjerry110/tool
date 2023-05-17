/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 10:25:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:24:59
 * @Description: base
 */
package logger

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type BaseLogger interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Base struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:29:09
 * @return {*}
 */
func (b *Base) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:28:58
 * @return {*}
 */
func (b *Base) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-21 16:16:29
	* @Description: base
	*/
   package logger
   
   import "github.com/yangjerry110/tool/pkg/logger"
   
   /**
	* @description: Logger
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func Logger() logger.LoggerPkgInterface {
	   return CreateLogger().Logger()
   }
   
   /**
	* @description: CreateLogger
	* @param {...CommonLogger} CommonLoggers
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateLogger(CommonLoggers ...CommonLogger) CommonLogger {
	   if len(CommonLoggers) == 0 {
		   return &Common{}
	   }
	   return CommonLoggers[0]
   }
   `
}
