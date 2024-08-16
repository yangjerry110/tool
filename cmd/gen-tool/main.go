/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 16:57:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-16 16:25:20
 * @Description: gen-tool main
 */
package main

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/command"
	"github.com/yangjerry110/tool/internal/toolErrors"
)

func main() {

	// set cli app
	if err := command.CreateCommand(&command.CliNewApp{}).New(); err != nil {
		fmt.Printf("gen-tool Err : %+v", toolErrors.NewError(err))
		fmt.Print("\r\n")
		panic(err)
	}
}
