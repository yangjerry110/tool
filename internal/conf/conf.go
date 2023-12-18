/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:30:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-08 17:41:18
 * @Description: config
 */
package conf

type Conf interface {
	SetConfig() error
}

/**
 * @description: CreateConf
 * @param {Conf} conf
 * @author: Jerry.Yang
 * @date: 2023-12-08 17:40:19
 * @return {*}
 */
func CreateConf(conf Conf) Conf {
	return conf
}
