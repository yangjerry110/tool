/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-20 18:18:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-03-05 14:11:16
 * @Description: appendBase
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type AppendBase struct {
	ServiceName   string
	ServiceNameUp string
	Time          string
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-20 18:20:59
 * @return {*}
 */
func (a *AppendBase) New() error {
	return template.AppendTemplate(fmt.Sprintf("%s/internal/service/base.go", config.ProjectPathConf.Path), a.getTemplate(), a)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-20 18:21:08
 * @return {*}
 */
func (a *AppendBase) getTemplate() string {
	return `/**
	* @description: Create{{.ServiceNameUp}}Service
	* @param {...{{.ServiceNameUp}}Service} {{.ServiceName}}Services
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func Create{{.ServiceNameUp}}Service({{.ServiceName}}Services ...{{.ServiceNameUp}}Service) {{.ServiceNameUp}}Service {
	   if len({{.ServiceName}}Services) == 0 {
		   return &{{.ServiceNameUp}}{}
	   }
	   return {{.ServiceName}}Services[0]
   }`
}
