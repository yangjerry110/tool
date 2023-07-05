/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 15:58:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 16:34:00
 * @Description: logger yaml
 */
package config

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type LoggerYamlConfig interface {
	SaveTemplate(configPath string) error
	GetTemplate() string
}

type LoggerYaml struct{}

/**
 * @description: SaveTemplate
 * @param {string} configPath
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:00:49
 * @return {*}
 */
func (l *LoggerYaml) SaveTemplate(configPath string) error {
	return templates.CreateCommonTemplate().SaveTemplate(configPath, "logger.yaml", l.GetTemplate(), nil, "yaml")
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:00:10
 * @return {*}
 */
func (l *LoggerYaml) GetTemplate() string {
	return `level: debug
callerDept: 4`
}
