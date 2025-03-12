/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:11:10
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 16:21:50
 * @Description: newApp Model
 */
package model

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type NewAppBaseModel struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:12:52
 * @return {*}
 */
func (n *NewAppBaseModel) New() error {
	filePath := fmt.Sprintf("%s/internal/model", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "base.go", n.getTemplate(), nil)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:11:59
 * @return {*}
 */
func (n *NewAppBaseModel) getTemplate() string {
	return `package model
	
	// Delete status
	var Is_Deleted = 1
	var No_Deleted = 0`
}
