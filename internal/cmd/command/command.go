/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 15:35:36
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2024-10-24 20:51:02
 * @Description: command
 */
package command

import (
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/errors"
)

type Command struct {
	CliApp *cli.App
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-11 15:59:38
 * @return {*}
 */
func (c *Command) New() error {

	// judge cliApp
	// if nil ?
	// if nil return err
	if c.CliApp == nil {
		return errors.ErrCmdCommandNoCliApp
	}

	// set newApp command
	// set newAPp command to cliapp.Commands
	if err := c.setNewAppCommand(); err != nil {
		return err
	}

	// set newApi command
	// set newApi command to cliApp.Commands
	if err := c.setNewApiCommand(); err != nil {
		return err
	}

	// Set newModel command
	// Set newModel command to cliApp.Commands
	if err := c.setNewModelCommand(); err != nil {
		return err
	}

	// Set newDao command
	// Set newDao command to cliApp.Commands
	if err := c.setNewDaoCommand(); err != nil {
		return err
	}

	// Set newProtobuf command
	// Set newProtobuf command to cliApp.Commands
	if err := c.setNewProtobufCommand(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: setNewAppCommand
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:13:25
 * @return {*}
 */
func (c *Command) setNewAppCommand() error {
	// newApp command
	newApiCommand := cli.Command{}
	newApiCommand.Name = "newApp"
	newApiCommand.Aliases = []string{"napp"}
	newApiCommand.Usage = "new an app => new app"
	newApiCommand.Action = func(ctx *cli.Context) error {
		return CreateCommand(&NewApp{CliContext: ctx}).New()
	}

	// set cliApp commands
	c.CliApp.Commands = append(c.CliApp.Commands, newApiCommand)
	return nil
}

/**
 * @description: setNewApiCommand
 * @author: Jerry.Yang
 * @date: 2023-12-11 15:58:41
 * @return {*}
 */
func (c *Command) setNewApiCommand() error {

	// newApi command
	newApiCommand := cli.Command{}
	newApiCommand.Name = "newApi"
	newApiCommand.Aliases = []string{"napi"}
	newApiCommand.Usage = "new an api => new api"
	newApiCommand.Action = func(ctx *cli.Context) error {
		return CreateCommand(&NewApi{CliContext: ctx}).New()
	}

	// set cliApp commands
	c.CliApp.Commands = append(c.CliApp.Commands, newApiCommand)
	return nil
}

/**
 * @description: setNewModelCommand
 * @author: Jerry.Yang
 * @date: 2023-12-21 14:26:26
 * @return {*}
 */
func (c *Command) setNewModelCommand() error {

	// newModel command
	newModelCommand := cli.Command{}
	newModelCommand.Name = "newModel"
	newModelCommand.Aliases = []string{"nmodel"}
	newModelCommand.Usage = "new an model => new model"
	newModelCommand.Action = func(ctx *cli.Context) error {
		return CreateCommand(&NewModel{CliContext: ctx}).New()
	}

	// Set cliApp Commands
	c.CliApp.Commands = append(c.CliApp.Commands, newModelCommand)
	return nil
}

/**
 * @description: setNewDaoCommand
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:48:15
 * @return {*}
 */
func (c *Command) setNewDaoCommand() error {
	// newDao command
	newDaoCommand := cli.Command{}
	newDaoCommand.Name = "newDao"
	newDaoCommand.Aliases = []string{"ndao"}
	newDaoCommand.Usage = "new an dao => new dao"
	newDaoCommand.Action = func(ctx *cli.Context) error {
		return CreateCommand(&NewDao{CliContext: ctx}).New()
	}

	// Set cliApp Commands
	c.CliApp.Commands = append(c.CliApp.Commands, newDaoCommand)
	return nil
}

/**
 * @description: setNewProtobuf
 * @author: Jerry.Yang
 * @date: 2024-10-24 18:32:48
 * @return {*}
 */
func (c *Command) setNewProtobufCommand() error {
	// newProtobuf command
	newProtobufCommand := cli.Command{}
	newProtobufCommand.Name = "newProtobuf"
	newProtobufCommand.Aliases = []string{"nProtobuf"}
	newProtobufCommand.Usage = "new an protobuf => new protobuf"
	newProtobufCommand.Action = func(ctx *cli.Context) error {
		return CreateCommand(&NewProtobuf{CliContext: ctx}).New()
	}

	// Set cliApp Commands
	c.CliApp.Commands = append(c.CliApp.Commands, newProtobufCommand)
	return nil
}
