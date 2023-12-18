/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 11:16:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-13 15:16:35
 * @Description: main
 */
package main

import (
	"flag"

	protocgentoolservice "github.com/yangjerry110/tool/internal/cmd/service/protocGenToolService"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {

	// Get isFirstCreate by flag
	isFirstCreate := flag.Bool("isFirstCreate", false, "isFirstCreate")

	// Get isAppend by flag
	isAppend := flag.Bool("isAppend", false, "isAppend")

	// Flag parse
	flag.Parse()

	// Define proto option
	protogenOptions := protogen.Options{}

	// Add flag.command.Set to paramsFunc
	protogenOptions.ParamFunc = flag.CommandLine.Set

	// To run by protogenOptions
	protogenOptions.Run(func(plugin *protogen.Plugin) error {
		// Plugin Generate
		return protocgentoolservice.CreateProtoGenToolService(&protocgentoolservice.Plugin{IsFirstCreate: *isFirstCreate, IsAppend: *isAppend, Plugin: plugin}).Generate()
	})

}
