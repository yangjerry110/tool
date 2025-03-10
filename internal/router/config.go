/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 14:42:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-10 14:46:37
 * @Description: router config
 */
package router

import "github.com/yangjerry110/tool/internal/conf"

type GinRouter struct {
	Addr string `yaml:"addr"`
}

type GrcpRouter struct {
	Protocol string `yaml:"protocol"`
	Port     string `yaml:"port"`
	Addr     string `yaml:"addr"`
}

var GinRouterConf = &GinRouter{}
var GrpcRouterConf = &GrcpRouter{}

func (g *GinRouter) SetConfig() error {
	/**
	 * @step
	 * @返回结果
	 **/
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "router.yaml", FileType: "yaml", ConfData: &GinRouterConf}).SetConfig(); err != nil {
		return err
	}
	return nil
}

func (g *GrcpRouter) SetConfig() error {
	/**
	 * @step
	 * @返回结果
	 **/
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "grpc_router.yaml", FileType: "yaml", ConfData: &GrpcRouterConf}).SetConfig(); err != nil {
		return err
	}
	return nil
}
