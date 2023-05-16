/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 14:44:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 14:14:30
 * @Description: conf
 */
package conf

import (
	"time"

	toolConf "github.com/yangjerry110/tool/conf"
)

type ConfInterface interface {
	CreateConfInterface(confInterface toolConf.ConfInterface) toolConf.ConfInterface
	GetYamlConf(yamlConfPath string, config interface{}) error
	GetConf(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) error
	Init(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) toolConf.ConfInterface
	GetParseConf(filepath string, fileName string) interface{}
}

type Conf struct {
	ConfInterface toolConf.ConfInterface
}

/**
 * @description: CreateConfInterface
 * @param {conf.ConfInterface} confInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:39:40
 * @return {*}
 */
func CreateConfInterface(confInterface toolConf.ConfInterface) *Conf {
	return &Conf{ConfInterface: confInterface}
}

/**
 * @description: GetConf
 * @param {string} filepath
 * @param {string} fileName
 * @param {string} fileType
 * @param {time.Duration} intervals
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:20:33
 * @return {*}
 */
func GetConf(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) error {
	toolConf := &toolConf.Conf{
		FilePath:  filepath,
		FileName:  fileName,
		FileType:  fileType,
		Intervals: intervals,
		Data:      conf,
	}
	return CreateConfInterface(toolConf).ConfInterface.GetNewConf()
}

/**
 * @description: Init
 * @param {string} filepath
 * @param {string} fileName
 * @param {string} fileType
 * @param {time.Duration} intervals
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:45:24
 * @return {*}
 */
func Init(filepath string, fileName string, fileType string, intervals time.Duration, conf interface{}) toolConf.ConfInterface {
	mytoolInitConf := &toolConf.Conf{
		FilePath:  filepath,
		FileName:  fileName,
		FileType:  fileType,
		Intervals: intervals,
	}
	return CreateConfInterface(mytoolInitConf).ConfInterface.Init(conf)
}

/**
 * @description: GetParseConf
 * @param {string} filepath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-15 19:00:04
 * @return {*}
 */
func GetParseConf(filepath string, fileName string) interface{} {
	mytoolInitConf := &toolConf.Conf{
		FilePath: filepath,
		FileName: fileName,
	}
	return CreateConfInterface(mytoolInitConf).ConfInterface.GetParseConf()
}
