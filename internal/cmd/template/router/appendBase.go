/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 17:27:13
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-12 17:42:36
 * @Description:
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type AppendBase struct {
	RouterNameUp string
	RouterName   string
	Time         string
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-12 17:30:31
 * @return {*}
 */
func (a *AppendBase) New() error {
	if err := template.AppendTemplate(fmt.Sprintf("%s/router/base.go", config.ProjectPathConf.Path), a.getTemplate(), a); err != nil {
		return err
	}
	return nil
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-12 17:30:21
 * @return {*}
 */
func (a *AppendBase) getTemplate() string {
	return `
	/**
	* @description: Create{{.RouterNameUp}}Router
	* @param {...{{.RouterNameUp}}Router} {{.RouterNameUp}}Routers
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func Create{{.RouterNameUp}}Router({{.RouterName}}Routers ...{{.RouterNameUp}}Router) {{.RouterNameUp}}Router {
	   if len({{.RouterName}}Routers) == 0 {
		   return &{{.RouterNameUp}}{}
	   }
	   return {{.RouterName}}Routers[0]
   }`
}
