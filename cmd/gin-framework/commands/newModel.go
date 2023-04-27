/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 10:45:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 15:00:32
 * @Description: new model
 */
package commands

import (
	"fmt"
	"os"
)

type NewModelCommands interface {
	CreateModel() error
	CreateWd() error
}

type NewModel struct {
	ModelPath string
}

/**
 * @description: NewModelParams
 * @author: Jerry.Yang
 * @date: 2023-04-25 14:14:13
 * @return {*}
 */
var NewModelParams = &NewModel{}

/**
 * @description: CreateModel
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:00:39
 * @return {*}
 */
func (n *NewModel) CreateModel() error {

	/**
	 * @step
	 * @创建目录
	 **/
	err := n.CreateWd()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateWd
 * @author: Jerry.Yang
 * @date: 2023-04-25 14:15:53
 * @return {*}
 */
func (n *NewModel) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParms.ProjectPath, "model")

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
	NewModelParams.ModelPath = path
	return nil
}
