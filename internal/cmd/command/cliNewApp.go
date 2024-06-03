/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 15:46:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-31 15:39:58
 * @Description:
 */
package command

import (
	"os"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/toolErrors"
)

type CliNewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-11 15:49:42
 * @return {*}
 */
func (c *CliNewApp) New() error {

	// define app
	app := cli.NewApp()

	// set app name
	// set app usage
	app.Name = "newApp"
	app.Usage = "create a new app "

	// set commands
	if err := CreateCommand(&Command{CliApp: app}).New(); err != nil {
		return toolErrors.NewError(err)
	}

	// run app
	app.Run(os.Args)
	return nil
}
