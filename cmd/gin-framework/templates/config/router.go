/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 15:55:13
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:56:53
 * @Description: router
 */
package config

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type RouterConfig interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Router struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:56:55
 * @return {*}
 */
func (r *Router) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义渲染数据
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "router.go", r.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:56:47
 * @return {*}
 */
func (r *Router) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: router
	*/
   package config
   
   import (
	   "time"
   
	   "github.com/yangjerry110/tool/pkg/conf"
   )
   
   type RouterConfig interface {
	   SetConfig() error
   }
   
   type Router struct {
	   Addr string ` + " ` yaml:\"addr\"` " + `
   }
   
   /**
	* @description: RouterSetConfig
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   var RouterSetConfig = &Router{}
   
   /**
	* @description: GetConfig
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (r *Router) SetConfig() error {
   
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
	   err = conf.GetConf(configPath, "router.yaml", "yaml", 60*time.Second, RouterSetConfig)
	   if err != nil {
		   return err
	   }
	   return nil
   }
   `
}
