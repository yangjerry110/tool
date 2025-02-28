/*
 * @Author: Jerry.Yang
 * @Date: 2024-10-24 18:27:47
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-02-27 22:56:36
 * @Description: new protobuf
 */
package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"github.com/yangjerry110/tool/internal/toolErrors"
)

type NewProtobuf struct {
	CliContext *cli.Context
}

/**
 * @description: newProtobuf new
 * @author: Jerry.Yang
 * @date: 2024-10-24 18:29:03
 * @return {*}
 */
func (n *NewProtobuf) New() error {

	// judge cliContext
	// if nil return err
	if n.CliContext == nil {
		return errors.ErrCmdCommandNoCliContext
	}

	// set projectPath
	if err := config.CreateConfig(&config.ProjectPath{}).SetConfig(); err != nil {
		panic(err)
	}

	// set projectImportPath
	if err := config.CreateConfig(&config.ProjectImportPath{}).SetConfig(); err != nil {
		panic(err)
	}

	// Get protobufName
	// Set protobufName
	if err := conf.CreateConf(&config.Protobuf{CliContext: n.CliContext}).SetConfig(); err != nil {
		return err
	}

	// Exec Command
	// Exec protoc
	cmd := exec.Command(
		"protoc",
		"-I",
		"protobuf",
		fmt.Sprintf("--go_opt=module=%s/vo/protobuf", config.ProjectImportPathConf.ImportPath),
		// "--go_out=plugins=grpc:vo/protobuf",
		"--go_out=vo/protobuf",
		fmt.Sprintf("--go-grpc_out=module=%s/vo/protobuf:vo/protobuf", config.ProjectImportPathConf.ImportPath),
		"grpc-gateway_out=.",
		config.ProtobufConf.ProtobufName,
	)
	// cmd := exec.Command(
	// 	"protoc",
	// 	"--proto_path=protobuf",
	// 	"--go_out=.",
	// 	"--go-grpc_out=.",
	// 	config.ProtobufConf.ProtobufName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd Run Err : %+v", err)
		fmt.Print("\r\n")
		return toolErrors.NewError(err)
	}
	return nil
}
