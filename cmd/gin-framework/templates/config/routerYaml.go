/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:01:14
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 16:29:11
 * @Description: router yaml
 */
package config

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type RouterYamlConfig interface {
	SaveTemplate(configPath string) error
	GetTemplate() string
}

type RouterYaml struct{}

/**
 * @description: SaveTemplate
 * @param {string} configPath
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:02:50
 * @return {*}
 */
func (r *RouterYaml) SaveTemplate(configPath string) error {
	return templates.CreateCommonTemplate().SaveTemplate(configPath, "router.yaml", r.GetTemplate(), nil, "yaml")
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:02:19
 * @return {*}
 */
func (r *RouterYaml) GetTemplate() string {
	return `addr: ":12000"`
}
