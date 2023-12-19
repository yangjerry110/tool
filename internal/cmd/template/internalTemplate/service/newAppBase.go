/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:16:43
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 14:57:34
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
	filePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "base.go", n.getTemplate(), nil)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:18:06
 * @return {*}
 */
func (n *NewAppBaseService) getTemplate() string {
	return `package service`
}
