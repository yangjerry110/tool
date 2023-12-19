/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:16:43
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 17:40:23
 * @Description: newApp service
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
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
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	filePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "base.go", n.getTemplate(), data)
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
   
   /**
	* @description: CreateDemoService
	* @param {...DemoService} DemoServices
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateDemoService(DemoServices ...DemoService) DemoService {
	   if len(DemoServices) == 0 {
		   return &Demo{}
	   }
	   return DemoServices[0]
   }
   `
}
