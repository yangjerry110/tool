/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-06 11:09:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 15:06:04
 * @Description: new app
 */
package commands

import (
	"github.com/golib/cli"
)

type NewAppCommands interface {
	NewApp(ctx *cli.Context) error
	CreateApp() error
}

type NewApp struct {
	AppControllerFileName string
	AppRouterFileName     string
	AppServiceFileName    string
	AppVoInputFileName    string
	AppVoOutputFileName   string
	AppModelFileName      string
	AppDaoFileName        string
}

/**
 * @description: NewAppParams
 * @author: Jerry.Yang
 * @date: 2023-05-06 11:15:55
 * @return {*}
 */
var NewAppParams = &NewApp{}

/**
 * @description: NewApp
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-09 16:58:24
 * @return {*}
 */
func (n *NewApp) NewApp(ctx *cli.Context) error {

	/**
	 * @step
	 * @设置projectPath
	 **/
	err := CreateInitCommands().SetProjectPath()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @设置projectImportPath
	 **/
	err = CreateInitCommands().SetImportProjectPath()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @设置appName
	 **/
	err = CreateInitCommands().SetAppName(ctx)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建app
	 **/
	err = n.CreateApp()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateApp
 * @author: Jerry.Yang
 * @date: 2023-05-06 11:16:03
 * @return {*}
 */
func (n *NewApp) CreateApp() error {

	// /**
	//  * @step
	//  * @创建controller
	//  **/
	// err := CreateNewControllerCommands().CreateNewController()
	// if err != nil {
	// 	return err
	// }

	/**
	 * @step
	 * @创建router
	 **/
	err := CreateNewRouterCommands().CreateNewRouter()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建service
	 **/
	err = CreateNewServiceCommands().CreateNewService()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建Vo
	 **/
	err = CreateNewVoCommands().CreateNewVo()
	if err != nil {
		return err
	}
	return nil
}
