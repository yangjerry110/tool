/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:44:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 16:57:00
 * @Description: controller
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates/controller"
)

type NewControllerCommands interface {
	CreateController() error
	CreateWd() error
	CreateFile() error
}

type NewController struct {
	ControllerPath string
}

/**
 * @description: NewControllerParams
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:47:07
 * @return {*}
 */
var NewControllerParams = &NewController{}

/**
 * @description: CreateConfig
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:06:35
 * @return {*}
 */
func (n *NewController) CreateController() error {

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
 * @date: 2023-04-24 16:47:15
 * @return {*}
 */
func (n *NewController) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParms.ProjectPath, "controller")

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
	NewControllerParams.ControllerPath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:48:30
 * @return {*}
 */
func (n *NewController) CreateFile() error {

	/**
	 * @step
	 * @创建base
	 **/
	err := controller.CreateBaseController().SaveTemplate(NewControllerParams.ControllerPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建test
	 **/
	err = controller.CreateTestController().SaveTemplate(NewControllerParams.ControllerPath, InitParms.ProjectImportPath)
	if err != nil {
		return err
	}
	return nil
}
