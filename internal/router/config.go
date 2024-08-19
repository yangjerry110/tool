/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-12 17:34:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-19 11:13:07
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
 * @description: GrpcConfig
 * @author: Jerry.Yang
 * @date: 2024-08-19 10:46:43
 * @return {*}
 */
type GrpcConfig struct {
	Protocol string `yaml:"protocol"`
	Port     string `yaml:"port"`
	Addr     string `yaml:"addr"`
}

/**
 * @description: GrpcRouterConf
 * @author: Jerry.Yang
 * @date: 2024-08-19 10:46:53
 * @return {*}
 */
var GrpcRouterConf = &GrpcConfig{}

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

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2024-08-19 10:48:01
 * @return {*}
 */
func (g *GrpcConfig) SetConfig() error {

	/**
	 * @step
	 * @返回结果
	 **/
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "grpc_router.yaml", FileType: "yaml", ConfData: &GrpcRouterConf}).SetConfig(); err != nil {
		return err
	}
	return nil
}
