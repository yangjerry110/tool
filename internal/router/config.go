/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-12 17:34:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 16:31:37
 * @Description: router Config
 */
package router

import "github.com/yangjerry110/tool/internal/conf"

/**
 * @description: config
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:49:13
 * @return {*}
 */
type Config struct {
	Addr string `yaml:"addr"`
}

/**
 * @description: GinConf
 * @author: Jerry.Yang
 * @date: 2023-12-13 18:40:07
 * @return {*}
 */
var RouteConf = &Config{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-13 18:41:43
 * @return {*}
 */
func (c *Config) SetConfig() error {

	/**
	 * @step
	 * @返回结果
	 **/
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "router.yaml", FileType: "yaml", ConfData: &RouteConf}).SetConfig(); err != nil {
		return err
	}
	return nil
}
