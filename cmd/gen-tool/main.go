/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 16:57:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:25:56
 * @Description: gen-tool main
 */
package main

import (
	"github.com/yangjerry110/tool/internal/cmd/command"
)

func main() {

	// set cli app
	if err := command.CreateCommand(&command.CliNewApp{}).New(); err != nil {
		panic(err)
	}
}
