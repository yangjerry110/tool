/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:53:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 14:55:37
 * @Description: newApp demo router
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppDemoRouter struct{}

/**
 * @description: NewAppDemoRouter
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:55:24
 * @return {*}
 */
func (n *NewAppDemoRouter) New() error {
	filePath := fmt.Sprintf("%s/router", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "demoRouter.go", n.getTemplate(), nil)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:55:14
 * @return {*}
 */
func (n *NewAppDemoRouter) getTemplate() string {
	return `package router`
}
