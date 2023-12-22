/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 17:39:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 10:57:43
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
