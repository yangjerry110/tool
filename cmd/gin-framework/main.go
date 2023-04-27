/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 15:37:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-23 17:42:21
 * @Description: main
 */
package main

import (
	"os"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/commands"
)

func main() {

	/**
	 * @创建app
	 **/
	app := cli.NewApp()

	/**
	 * @step
	 * @定义app相关的参数
	 **/
	app.Name = "newApp"
	app.Usage = "create a new app "

	/**
	 * @step
	 * @创建commands
	 **/
	commands.CreateNewCommands().Commands(app)

	/**
	 * @step
	 * @run
	 **/
	app.Run(os.Args)
}
