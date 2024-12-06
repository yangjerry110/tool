/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 17:28:34
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2024-10-24 20:58:07
 * @Description: project path
 */
package config

import (
	"os"
)

type ProjectPath struct {
	Path string
}

/**
 * @description: ProjectPathConf
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:31:22
 * @return {*}
 */
var ProjectPathConf = &ProjectPath{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:31:14
 * @return {*}
 */
func (p *ProjectPath) SetConfig() error {

	/**
	 * @step
	 * @获取当前目录
	 **/
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	// set projectPathConf
	ProjectPathConf.Path = path
	return nil
}
