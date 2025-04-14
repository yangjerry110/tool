/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 17:42:56
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-28 15:29:25
 * @Description: newApp
 */
package config

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

/**
 * @description: NewAppDatabase
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:31:33
 * @return {*}
 */
type NewAppDatabase struct{}

/**
 * @description: NewAppLogger
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:31:33
 * @return {*}
 */
type NewAppLogger struct{}

/**
 * @description: NewAppRouter
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:31:49
 * @return {*}
 */
type NewAppRouter struct{}

/**
 * @description: NewAppRedis
 * @author: Jerry.Yang
 * @date: 2023-12-21 10:59:34
 * @return {*}
 */
type NewAppRedis struct{}

/**
 * @description: NewAppConfig
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:59:32
 * @return {*}
 */
type NewAppConfig struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:22:16
 * @return {*}
 */
func (n *NewAppDatabase) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	// return
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config")
	return template.SaveTemplate(filePath, "database.go", n.getTemplate(), data)
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:22:26
 * @return {*}
 */
func (n *NewAppLogger) New() error {
	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	// return
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config")
	return template.SaveTemplate(filePath, "logger.go", n.getTemplate(), data)
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:30:24
 * @return {*}
 */
func (n *NewAppRouter) New() error {
	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	// return
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config")
	return template.SaveTemplate(filePath, "router.go", n.getTemplate(), data)
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:00:11
 * @return {*}
 */
func (n *NewAppRedis) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	// return
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config")
	return template.SaveTemplate(filePath, "redis.go", n.getTemplate(), data)
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 16:00:53
 * @return {*}
 */
func (n *NewAppConfig) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	// return
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config")
	return template.SaveTemplate(filePath, "config.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:22:02
 * @return {*}
 */
func (n *NewAppDatabase) getTemplate() string {
	return `
/*
 * @Author: Jerry.Yang
 * @Date: {{.Time}}
 * @LastEditors: Jerry.Yang
 * @LastEditTime: {{.Time}}
 * @Description: database config
 */
package config

import (
	"github.com/yangjerry110/tool/db"
)

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
	* @setDatabaseConfig
	**/
	return db.SetGormConf().SetConfig()
}

`
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:21:53
 * @return {*}
 */
func (n *NewAppLogger) getTemplate() string {
	return "package config"
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:27:33
 * @return {*}
 */
func (n *NewAppRouter) getTemplate() string {
	return `
/*
 * @Author: Jerry.Yang
 * @Date: {{.Time}}
 * @LastEditors: Jerry.Yang
 * @LastEditTime: {{.Time}}
 * @Description: router config
 */
package config

import "github.com/yangjerry110/tool/router"

type Router struct{}

/**
* @description: SetConfig
* @author: Jerry.Yang
* @date: {{.Time}}
* @return {*}
 */
func (r *Router) SetConfig() error {

	/**
	* @step
	* @setRouterConfig
	**/
	return router.SetHTTPRouterConfig().SetConfig()
}
`
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:01:13
 * @return {*}
 */
func (n *NewAppRedis) getTemplate() string {
	return `
/*
 * @Author: Jerry.Yang
 * @Date: {{.Time}}
 * @LastEditors: Jerry.Yang
 * @LastEditTime: {{.Time}}
 * @Description: Redis config
 */
	package config

	import "github.com/yangjerry110/tool/cache"
	
	type Redis struct{}
	
	/**
	 * @description: SetConfig
	 * @author: Jerry.Yang
	 * @date: {{.Time}}
	 * @return {*}
	 */
	func (r *Redis) SetConfig() error {
		return cache.SetRedisConf().SetConfig()
	}
	`
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-21 16:00:35
 * @return {*}
 */
func (n *NewAppConfig) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: Config
	*/
   package config
   
   import (
	   "github.com/yangjerry110/tool/cache"
	   "github.com/yangjerry110/tool/router"
	   "github.com/yangjerry110/tool/db"
   )
   
   type Config struct{}
   
   /**
	* @description: SetConfig
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (c *Config) SetConfig() error {
   
	   // set gin conf
	   if err := router.SetHTTPRouterConfig().SetConfig(); err != nil {
			return err
		}
   
	   // Get All Db Clients
	   if err := db.CreateGormDb().CreateAllClient(); err != nil {
		   return err
	   }
   
	   // Get All Redis Clients
	   if err := cache.CreateRedisCache().CreateAllClient(); err != nil {
		   return err
	   }
	   return nil
   }
   `
}
