/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:17:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 17:24:58
 * @Description: error commands
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates/errors"
)

type NewErrorCommands interface {
	CreateError() error
	CreateWd() error
	CreateFile() error
}

type NewError struct {
	ErrorPath string
}

/**
 * @description: NewErrorParams
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:18:41
 * @return {*}
 */
var NewErrorParams = &NewError{}

/**
 * @description: CreateError
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:22:55
 * @return {*}
 */
func (n *NewError) CreateError() error {
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
 * @date: 2023-04-24 17:19:53
 * @return {*}
 */
func (n *NewError) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "errors")

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
	NewErrorParams.ErrorPath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:20:17
 * @return {*}
 */
func (n *NewError) CreateFile() error {

	/**
	 * @step
	 * @创建error
	 **/
	err := errors.CreateError().SaveTemplate(NewErrorParams.ErrorPath)
	if err != nil {
		return err
	}
	return nil
}
