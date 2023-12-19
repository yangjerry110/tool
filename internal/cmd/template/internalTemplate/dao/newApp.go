/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:03:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 14:06:02
 * @Description: newApp
 */
package dao

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppBaseDao struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:06:06
 * @return {*}
 */
func (n *NewAppBaseDao) New() error {
	filePath := fmt.Sprintf("%s/internal/dao", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "base.go", n.getTemplate(), nil)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:05:57
 * @return {*}
 */
func (n *NewAppBaseDao) getTemplate() string {
	return `package dao`
}
