/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-17 16:26:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 16:30:55
 * @Description: databaseYaml config
 */
package config

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type DataBaseYamlConfig interface {
	SaveTemplate(path string) error
}

type DataBaseYaml struct{}

/**
 * @description: SaveTemplate
 * @param {string} configPath
 * @author: Jerry.Yang
 * @date: 2023-05-17 16:29:16
 * @return {*}
 */
func (d *DataBaseYaml) SaveTemplate(configPath string) error {
	return templates.CreateCommonTemplate().SaveTemplate(configPath, "database.yaml", d.GetTemplate(), nil)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-17 16:28:39
 * @return {*}
 */
func (d *DataBaseYaml) GetTemplate() string {
	return `master: 
	dsn: "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	`
}
