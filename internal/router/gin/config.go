/*
* @Author: Jerry.Yang
* @Date: 2023-12-13 18:39:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-14 11:04:23
* @Description: gin conf
*/
package gin

import (
	"github.com/yangjerry110/tool/internal/conf"
	gormdb "github.com/yangjerry110/tool/internal/db/gormDb"
)

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
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "router.yaml", FileType: "yaml", ConfData: RouteConf}).SetConfig(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:47:37
 * @return {*}
 */
func (g *Gin) SetConfig() error {

	// Set config path
	if err := conf.CreateConf(&conf.ConfigPath{}).SetConfig(); err != nil {
		return err
	}

	// Set db config
	if err := conf.CreateConf(&gormdb.GormDbConfig{}).SetConfig(); err != nil {
		return err
	}

	// Set gin config
	if err := conf.CreateConf(&Config{}).SetConfig(); err != nil {
		return err
	}
	return nil
}
