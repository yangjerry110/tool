/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-25 14:47:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-13 10:49:21
 * @Description: new service
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/template"
)

type NewService struct {
	ProjectPath   string
	Time          string
	ServiceName   string
	ServiceNameUp string
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2024-04-25 14:50:05
 * @return {*}
 */
func (n *NewService) New() error {
	if err := template.SaveTemplate(fmt.Sprintf("%s/internal/service", n.ProjectPath), fmt.Sprintf("%sService.go", n.ServiceName), n.getTemplate(), n); err != nil {
		return err
	}
	return nil
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2024-04-25 14:49:56
 * @return {*}
 */
func (n *NewService) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: {{.ServiceNameUp}} service
	*/
   package service
   `
}
