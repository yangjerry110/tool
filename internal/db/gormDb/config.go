/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:56:08
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 11:06:17
 * @Description: gorm config
 */
package gormdb

import "github.com/yangjerry110/tool/internal/conf"

type GormDbConfig struct {
	Dsn string `yaml:"dsn"`
}

/**
 * @description: GormDbConf
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:00:49
 * @return {*}
 */
var GormDbConfs = map[string]*GormDbConfig{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:05:48
 * @return {*}
 */
func (g *GormDbConfig) SetConfig() error {

	/**
	 * @step
	 * @返回结果
	 **/
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "redis.yaml", FileType: "yaml", ConfData: GormDbConfs}).SetConfig(); err != nil {
		return err
	}
	return nil
}
