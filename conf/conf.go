/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 17:39:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 15:51:29
 * @Description: conf
 */
package conf

import "github.com/yangjerry110/tool/internal/conf"

/**
 * @description: CreateConf
 * @param {conf.Conf} conf
 * @author: Jerry.Yang
 * @date: 2023-12-11 10:50:44
 * @return {*}
 */
func CreateConf(conf conf.Conf) conf.Conf {
	return conf
}

/**
 * @description: CreatePathConf
 * @param {string} configPath
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:51:21
 * @return {*}
 */
func CreatePathConf(configPath string) error {
	return CreateConf(&conf.Path{ConfigPath: configPath}).SetConfig()
}

/**
 * @description: CreateConfigPathConf
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:51:30
 * @return {*}
 */
func CreateConfigPathConf() error {
	return CreateConf(&conf.ConfigPath{}).SetConfig()
}

/**
 * @description: CreateYamlConf
 * @param {string} fileName
 * @param {interface{}} confData
 * @author: Jerry.Yang
 * @date: 2023-12-22 10:54:45
 * @return {*}
 */
func CreateYamlConf(fileName string, confData interface{}) error {
	return CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: fileName, FileType: "yaml", ConfData: confData}).SetConfig()
}
