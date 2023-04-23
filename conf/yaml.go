/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 17:15:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-23 15:24:19
 * @Description: yaml conf
 */
package conf

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

/**
 * @description: YamlConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:58
 * @return {*}
 */
type YamlConf struct {
	FilePath  string
	FileName  string
	FileType  string
	Intervals time.Duration
	Conf      interface{}
}

/**
 * @description: Init
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:29:02
 * @return {*}
 */
func (y *YamlConf) Init(conf interface{}) ConfInterface {
	return y
}

/**
 * @description: GetParseConf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:29:10
 * @return {*}
 */
func (y *YamlConf) GetParseConf() interface{} {
	return y.Conf
}

/**
 * @description: GetHotUpdateConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:17:57
 * @return {*}
 */
func (y *YamlConf) GetNewConf() error {
	conf := &Conf{
		FilePath:  y.FilePath,
		FileName:  y.FileName,
		FileType:  "yaml",
		Intervals: y.Intervals,
		Data:      y.Conf,
	}
	return conf.GetNewConf()
}

/**
 * @description: Parse()
 * @author: Jerry.Yang
 * @date: 2022-09-22 17:19:26
 * @return {*}
 */
func (y *YamlConf) GetConf(conf interface{}) error {
	/**
	 * @step
	 * @获取yaml文件的数据
	 **/
	yamlData, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", y.FilePath, y.FileName))
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @渲染到结构体
	 **/
	decoder := yaml.NewDecoder(bytes.NewReader(yamlData))
	err = decoder.Decode(y.Conf)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: GetErr
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:59:23
 * @return {*}
 */
func (y *YamlConf) GetErr() error {
	return nil
}
