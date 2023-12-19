/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 17:42:56
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 11:31:47
 * @Description: newApp
 */
package config

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
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
	"github.com/yangjerry110/tool/pkg/db/gormdb"
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
	return gormdb.CreateDbConf()
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
	return ""
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
* @Description: database config
*/
package config

import (
	"github.com/yangjerry110/tool/pkg/router/gin"
)

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
	* @setDatabaseConfig
	**/
	return gin.CreateRouterConf()
`
}
