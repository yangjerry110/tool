/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 15:50:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 16:12:53
 * @Description: new
 */
package commands

import (
	"github.com/golib/cli"
)

type NewCommands interface {
	Commands(app *cli.App) error
}

type New struct{}

/**
 * @description: Commands
 * @author: Jerry.Yang
 * @date: 2023-04-23 16:54:42
 * @return {*}
 */
func (n *New) Commands(app *cli.App) error {

	/**
	 * @step
	 * @定义命令
	 **/
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "new a project => new project",
			Action: func(c *cli.Context) error {
				return n.New(c)
			},
		},
		{
			Name:    "newApp",
			Aliases: []string{"na"},
			Usage:   "new a app => new app",
			Action: func(c *cli.Context) error {
				return CreateNewAppCommands().NewApp(c)
			},
		},
		{
			Name:    "newApi",
			Aliases: []string{"napi"},
			Usage:   "new a api => new api",
			Action: func(c *cli.Context) error {
				return CreateNewApiCommands().NewApi(c)
			},
		},
		{
			Name:    "newController",
			Aliases: []string{"nc"},
			Usage:   "new a controller => new controller",
			Action: func(c *cli.Context) error {
				return CreateNewControllerCommands().NewController(c)
			},
		},
		{
			Name:    "newRouter",
			Aliases: []string{"nr"},
			Usage:   "new a router => new router",
			Action: func(c *cli.Context) error {
				return CreateNewRouterCommands().NewRouter(c)
			},
		},
		{
			Name:    "newService",
			Aliases: []string{"ns"},
			Usage:   "new a service => new service",
			Action: func(c *cli.Context) error {
				return CreateNewServiceCommands().NewService(c)
			},
		},
		{
			Name:    "newVo",
			Aliases: []string{"nv"},
			Usage:   "new a vo => new vo",
			Action: func(c *cli.Context) error {
				return CreateNewVoCommands().NewVo(c)
			},
		},
		{
			Name:    "newModel",
			Aliases: []string{"nm"},
			Usage:   "new a model => new model",
			Action: func(c *cli.Context) error {
				return CreateNewModelCommands().NewModel(c)
			},
		},
		{
			Name:    "newDao",
			Aliases: []string{"nd"},
			Usage:   "new a dao => new dao",
			Action: func(c *cli.Context) error {
				return CreateNewDaoCommands().NewDao(c)
			},
		},
	}
	return nil
}

/**
 * @description: New
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-04-23 16:51:53
 * @return {*}
 */
func (n *New) New(ctx *cli.Context) error {

	/**
	 * @step
	 * @设置projectName
	 **/
	err := CreateInitCommands().SetProjectName(ctx)
	if err != nil {
		return nil
	}

	/**
	 * @step
	 * @问新创建的项目名称
	 **/
	err = CreateInitCommands().AskInit()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @设置projectPath
	 **/
	err = CreateInitCommands().SetProjectPath()
	if err != nil {
		return err
	}
	/**
	 * @step
	 * @创建config
	 **/
	err = CreateNewConfigCommands().CreateConfig()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建controller
	 **/
	err = CreateNewControllerCommands().CreateController()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建dao
	 **/
	err = CreateNewDaoCommands().CreateDao()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建error
	 **/
	err = CreateNewErrorCommands().CreateError()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建logger
	 **/
	err = CreateNewLoggerCommands().CreateLogger()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建model
	 **/
	err = CreateNewModelCommands().CreateModel()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建router
	 **/
	err = CreateNewRouterCommands().CreateRouter()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建service
	 **/
	err = CreateNewServiceCommands().CreateService()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建vo
	 **/
	err = CreateNewVoCommands().CreateVo()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建protobuf
	 **/
	err = CreateNewProtobufCommands().CreateProtobuf()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建template
	 **/
	err = CreateNewTemplateCommands().CreateTemplate()
	if err != nil {
		return err
	}
	return nil
}
