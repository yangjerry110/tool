/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:22:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-20 17:01:08
 * @Description: newApp Main
 */
package template

import "github.com/yangjerry110/tool/cmd/gen-tool/internal/config"

type NewAppMain struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:23:46
 * @return {*}
 */
func (n *NewAppMain) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time       string
		ImportPath string
	}

	// Set Data
	data := &Data{}
	data.ImportPath = config.ProjectImportPathConf.ImportPath
	data.Time = GetFormatNowTime()

	return SaveTemplate(config.ProjectPathConf.Path, "main.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:23:06
 * @return {*}
 */
func (n *NewAppMain) getTemplate() string {
	return `
/*
 * @Author: Jerry.Yang
 * @Date: {{.Time}}
 * @LastEditors: Jerry.Yang
 * @LastEditTime: {{.Time}}
 * @Description: main
 */
package main

import (
	"{{.ImportPath}}/router"
	_ "{{.ImportPath}}/docs"
)

/**
 * @description: main
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
func main() {

	// Run Router
	router.RunRouter()
}
`
}
