/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 17:39:35
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 20:37:37
 * @Description: conf
 */
package conf

type Conf interface {
	SetConfig() error
}

/**
 * @description: CreateConf
 * @param {Conf} conf
 * @author: Jerry.Yang
 * @date: 2023-12-22 16:29:03
 * @return {*}
 */
func CreateConf(conf Conf) Conf {
	return conf
}

/**
 * @description: CreatePathConf
 * @param {string} configPath
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:51:21
 * @return {*}
 */
func CreatePathConf(configPath string) Conf {
	return CreateConf(&Path{ConfigPath: configPath})
}

/**
 * @description: CreateConfigPathConf
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:51:30
 * @return {*}
 */
func CreateConfigPathConf() Conf {
	return CreateConf(&ConfigPath{})
}

/**
 * @description: GetPathConf
 * @author: Jerry.Yang
 * @date: 2023-12-26 14:24:47
 * @return {*}
 */
func GetPathConf() *Path {
	return PathConfig
}

/**
 * @description: CreateYamlConf
 * @param {string} fileName
 * @param {interface{}} confData
 * @author: Jerry.Yang
 * @date: 2023-12-22 10:54:45
 * @return {*}
 */
func CreateYamlConf(fileName string, confData interface{}) Conf {
	return CreateConf(&Yaml{FilePath: PathConfig.ConfigPath, FileName: fileName, FileType: "yaml", ConfData: confData})
}
