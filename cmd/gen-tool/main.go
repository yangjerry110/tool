/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 16:57:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:08:55
 * @Description: gen-tool main
 */
package main

import (
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/command"
	"github.com/yangjerry110/tool/toolerrors"
)

func main() {

	// set cli app
	if err := command.CreateCommand(&command.CliNewApp{}).New(); err != nil {
		toolerrors.NewError(err)
		panic(err)
	}
}
