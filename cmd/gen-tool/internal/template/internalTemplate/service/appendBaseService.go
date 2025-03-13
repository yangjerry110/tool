/*
 * @Author: Jerry.Yang
 * @Date: 2024-03-05 14:09:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-03-05 14:41:40
 * @Description: append base service
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type AppendBaseService struct {
	ServiceName   string
	ServiceNameUp string
	Time          string
}

// new
//
// new
// Author Jerry.Yang
// Date 2024-03-05 14:10:34
func (a *AppendBaseService) New() error {
	return template.AppendTemplate(fmt.Sprintf("%s/internal/service/baseService.go", config.ProjectPathConf.Path), a.getTemplate(), a)
}

// getTemplate
//
// getTemplate
// Author Jerry.Yang
// Date 2024-03-05 14:10:51
func (a *AppendBaseService) getTemplate() string {
	return `/**
	* @description: Create{{.ServiceNameUp}}Service
	* @param {...{{.ServiceNameUp}}Service} {{.ServiceName}}Services
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func Create{{.ServiceNameUp}}Service({{.ServiceName}}Services ...interfaceService.{{.ServiceNameUp}}Service) interfaceService.{{.ServiceNameUp}}Service {
	   if len({{.ServiceName}}Services) == 0 {
		   return &{{.ServiceNameUp}}{}
	   }
	   return {{.ServiceName}}Services[0]
   }`
}
