/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 17:39:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 10:54:54
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
 * @param {string} filePath
 * @param {string} fileName
 * @param {string} fileType
 * @param {interface{}} confData
 * @author: Jerry.Yang
 * @date: 2023-12-22 10:54:45
 * @return {*}
 */
func CreateYamlConf(filePath string, fileName string, fileType string, confData interface{}) error {
	return CreateConf(&conf.Yaml{FilePath: filePath, FileName: fileName, FileType: fileType, ConfData: confData}).SetConfig()
}
