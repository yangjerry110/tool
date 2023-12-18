/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 17:25:55
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 17:42:27
 * @Description:
 */
package config

type ConfigInterface interface {
	SetConfig() error
}

type Config struct{}

/**
 * @description: CreateConfig
 * @param {ConfigInterface} ConfigInterface
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:28:18
 * @return {*}
 */
func CreateConfig(ConfigInterface ConfigInterface) ConfigInterface {
	return ConfigInterface
}
