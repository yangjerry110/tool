/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 15:50:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-27 10:29:40
 * @Description: new
 */
package commands

import (
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/errors"
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
	 * @获取项目名称
	 **/
	projectName := ctx.Args().First()
	if projectName == "" {
		return errors.ErrProjectNameIsEmpty
	}
	InitParms.ProjectName = projectName

	/**
	 * @step
	 * @问新创建的项目名称
	 **/
	err := CreateInitCommands().AskInit()
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
	 * @创建template
	 **/
	err = CreateNewTemplateCommands().CreateTemplate()
	if err != nil {
		return err
	}
	return nil
}
