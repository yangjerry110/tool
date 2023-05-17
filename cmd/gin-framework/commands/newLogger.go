/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 10:36:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 10:41:19
 * @Description: new Logger
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates/logger"
)

type NewLoggerCommands interface {
	CreateLogger() error
	CreateWd() error
	CreateFile() error
}

type NewLogger struct {
	LoggerPath string
}

/**
 * @description: NewLoggerParams
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:38:07
 * @return {*}
 */
var NewLoggerParams = &NewLogger{}

/**
 * @description: CreateLogger
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:22:55
 * @return {*}
 */
func (n *NewLogger) CreateLogger() error {
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
 * @date: 2023-04-25 10:38:46
 * @return {*}
 */
func (n *NewLogger) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "logger")

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
	 * @赋值
	 **/
	NewLoggerParams.LoggerPath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:41:18
 * @return {*}
 */
func (n *NewLogger) CreateFile() error {

	/**
	 * @step
	 * @创建base
	 **/
	err := logger.CreateBaseLogger().SaveTemplate(NewLoggerParams.LoggerPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建common
	 **/
	err = logger.CreateCommonLogger().SaveTemplate(NewLoggerParams.LoggerPath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}
	return nil
}
