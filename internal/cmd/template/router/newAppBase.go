/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:51:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 14:53:38
 * @Description: newApp base router
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppBaseRouter struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:53:24
 * @return {*}
 */
func (n *NewAppBaseRouter) New() error {
	filePath := fmt.Sprintf("%s/router", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "base.go", n.getTemplate(), nil)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:52:38
 * @return {*}
 */
func (n *NewAppBaseRouter) getTemplate() string {
	return `package router`
}
