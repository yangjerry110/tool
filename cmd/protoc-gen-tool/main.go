/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-23 15:00:59
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 22:12:32
 * @Description: main
 */
package main

import (
	"flag"

	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/commands"
	"google.golang.org/protobuf/compiler/protogen"
)

/**
 * @description: main
 * @author: Jerry.Yang
 * @date: 2023-05-23 16:11:17
 * @return {*}
 */
func main() {

	/**
	 * @step
	 * @处理参数
	 **/
	isFirstCreate := flag.Bool("isFirstCreate", false, "isFirstCreate")
	isAppend := flag.Bool("isAppend", false, "isAppend")
	flag.Parse()

	/**
	 * @step
	 * @protogen
	 **/
	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(plugin *protogen.Plugin) error {
		commands.CommandParams.IsFirstCreate = *isFirstCreate
		commands.CommandParams.IsAppend = *isAppend
		return commands.CreateGenerateCommands().Generate(plugin)
	})
}
