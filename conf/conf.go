/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 17:39:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-26 14:25:02
 * @Description: conf
 */
package conf

import "github.com/yangjerry110/tool/internal/conf"

/**
 * @description: CreateConf
 * @param {conf.Conf} conf
 * @author: Jerry.Yang
 * @date: 2023-12-22 16:29:03
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
func CreatePathConf(configPath string) conf.Conf {
	return conf.CreateConf(&conf.Path{ConfigPath: configPath})
}

/**
 * @description: CreateConfigPathConf
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:51:30
 * @return {*}
 */
func CreateConfigPathConf() conf.Conf {
	return conf.CreateConf(&conf.ConfigPath{})
}

/**
 * @description: GetPathConf
 * @author: Jerry.Yang
 * @date: 2023-12-26 14:24:47
 * @return {*}
 */
func GetPathConf() *conf.Path {
	return conf.PathConfig
}

/**
 * @description: CreateYamlConf
 * @param {string} fileName
 * @param {interface{}} confData
 * @author: Jerry.Yang
 * @date: 2023-12-22 10:54:45
 * @return {*}
 */
func CreateYamlConf(fileName string, confData interface{}) conf.Conf {
	return conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: fileName, FileType: "yaml", ConfData: confData})
}
