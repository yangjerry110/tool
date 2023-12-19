/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:59:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 16:09:48
 * @Description: newApp go mod
 */
package template

import (
	"github.com/yangjerry110/tool/internal/cmd/config"
)

type NewAppGoMod struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:02:55
 * @return {*}
 */
func (n *NewAppGoMod) New() error {

	// Data struct
	type Data struct {
		ImportPath string
	}

	// Set Data
	data := &Data{}
	data.ImportPath = config.ProjectImportPathConf.ImportPath

	// Return
	return SaveTemplate(config.ProjectPathConf.Path, "go.mod", n.getTemplate(), data, "mod")
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:02:36
 * @return {*}
 */
func (n *NewAppGoMod) getTemplate() string {
	return `module {{.ImportPath}}

go 1.20`
}
