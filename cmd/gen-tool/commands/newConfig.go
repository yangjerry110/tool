/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 10:57:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-23 19:07:51
 * @Description: new config
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates/config"
)

type NewConfigCommands interface {
	CreateConfig() error
	CreateWd() error
	CreateFile() error
}

type NewConfig struct {
	ConfigPath     string
	YamlConfigPath string
}

/**
 * @description: NewConfigParams
 * @author: Jerry.Yang
 * @date: 2023-04-24 11:24:38
 * @return {*}
 */
var NewConfigParams = &NewConfig{}

/**
 * @description: CreateConfig
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:06:35
 * @return {*}
 */
func (n *NewConfig) CreateConfig() error {

	/**
	 * @step
	 * @创建config的文件夹
	 **/
	err := n.CreateWd()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建文件
	 **/
	err = n.CreateFile()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateWd
 * @author: Jerry.Yang
 * @date: 2023-04-24 11:25:19
 * @return {*}
 */
func (n *NewConfig) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "config")

	/**
	 * @step
	 * @创建configPath
	 **/
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建yamlConfig
	 **/
	yamlConfigPath := fmt.Sprintf("%s/%s", path, "yamlConfig")

	/**
	 * @step
	 * @创建yamlconfig
	 **/
	err = os.MkdirAll(yamlConfigPath, 0777)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @赋值
	 **/
	NewConfigParams.ConfigPath = path
	NewConfigParams.YamlConfigPath = yamlConfigPath
	return nil
}

/**
 * @description: CreateConfigFile
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:03:08
 * @return {*}
 */
func (n *NewConfig) CreateFile() error {

	/**
	 * @step
	 * @创建config下base文件
	 **/
	err := config.CreateBaseConfig().SaveTemplate(NewConfigParams.ConfigPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建logger
	 **/
	err = config.CreateLoggerConfig().SaveTemplate(NewConfigParams.ConfigPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建path
	 **/
	err = config.CreatePathConfig().SaveTemplate(NewConfigParams.ConfigPath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建router
	 **/
	err = config.CreateRouterConfig().SaveTemplate(NewConfigParams.ConfigPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建database
	 **/
	err = config.CreateDatabaseConfig().SaveTemplate(NewConfigParams.ConfigPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建loggerYaml
	 **/
	err = config.CreateLoggerYamlConfig().SaveTemplate(NewConfigParams.YamlConfigPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建routerYaml
	 **/
	err = config.CreateRouterYamlConfig().SaveTemplate(NewConfigParams.YamlConfigPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建databaseYaml
	 **/
	err = config.CreateDatabaseYamlConfig().SaveTemplate(NewConfigParams.YamlConfigPath)
	if err != nil {
		return err
	}
	return nil
}
