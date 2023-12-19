/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:57:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 14:57:28
 * @Description:
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppDemoService struct{}

/**
 * @description: NewDemo
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:28:41
 * @return {*}
 */
func (n *NewAppDemoService) New() error {
	filePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "demoService.go", n.getTemplate(), nil)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:28:33
 * @return {*}
 */
func (n *NewAppDemoService) getTemplate() string {
	return `package service`
}
