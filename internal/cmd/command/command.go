/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 15:35:36
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:14:14
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
