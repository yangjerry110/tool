/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:22:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:23:20
 * @Description: newApp Main
 */
package template

import "github.com/yangjerry110/tool/internal/cmd/config"

type NewAppMain struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:23:46
 * @return {*}
 */
func (n *NewAppMain) New() error {
	return SaveTemplate(config.ProjectPathConf.Path, "main.go", n.getTemplate(), nil)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:23:06
 * @return {*}
 */
func (n *NewAppMain) getTemplate() string {
	return `
	package main
	
	func main() {

	}`
}
