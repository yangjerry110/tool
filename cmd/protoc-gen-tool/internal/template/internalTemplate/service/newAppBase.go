/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:16:43
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-13 10:49:03
 * @Description: newApp service
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/template"
)

type NewAppBaseService struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:18:18
 * @return {*}
 */
func (n *NewAppBaseService) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time              string
		ImportProjectPath string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()
	data.ImportProjectPath = config.ProjectImportPathConf.ImportPath

	filePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "baseService.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:18:06
 * @return {*}
 */
func (n *NewAppBaseService) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: base
	*/
   package service

   import "{{.ImportProjectPath}}/internal/service/interfaceService"
   
   /**
	* @description: CreateDemoService
	* @param {...DemoService} DemoServices
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateDemoService(DemoServices ...interfaceService.DemoService) interfaceService.DemoService {
	   if len(DemoServices) == 0 {
		   return &Demo{}
	   }
	   return DemoServices[0]
   }
   `
}
