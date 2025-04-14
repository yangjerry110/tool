/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 15:59:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:14:08
 * @Description: newApi Command
 */
package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golib/cli"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/errors"
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/toolerrors"
)

type NewApi struct {
	CliContext *cli.Context
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:42:58
 * @return {*}
 */
func (n *NewApi) New() error {

	// judge cliContext
	// if nil return err
	if n.CliContext == nil {
		return errors.ErrCommandNoCliContext
	}

	// set projectPath
	if err := conf.CreateConf(&config.ProjectPath{}).SetConfig(); err != nil {
		panic(err)
	}

	// set projectImportPath
	if err := conf.CreateConf(&config.ProjectImportPath{}).SetConfig(); err != nil {
		panic(err)
	}

	// Get protobufName
	// Set protobufName
	if err := conf.CreateConf(&config.Protobuf{CliContext: n.CliContext}).SetConfig(); err != nil {
		return err
	}

	// get is firstCreate
	isFirstCreate, err := n.askIsFirstCreate()
	if err != nil {
		return err
	}

	// get is append
	isAppend, err := n.askIsAppend()
	if err != nil {
		return err
	}

	// get isExtend
	isExtend, err := n.askIsExtend()
	if err != nil {
		return err
	}

	// define extendPath
	extendPath := ""

	// judge isExtend
	// if == true; ask protobufImportPath
	if isExtend {
		extendPath, err = n.askExtendPath()
		if err != nil {
			return err
		}
	}

	// set protoGenToolConf
	protoGenToolConfig := &config.ProtocGenTool{}
	protoGenToolConfig.IsFirstCreate = isFirstCreate
	protoGenToolConfig.IsAppend = isAppend
	protoGenToolConfig.IsExtend = isExtend
	protoGenToolConfig.ExtendPath = extendPath
	if err := protoGenToolConfig.SetConf(); err != nil {
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
		"--tool_out=vo/protobuf",
		config.ProtobufConf.ProtobufName,
	)

	// cmd := exec.Command(
	// 	"protoc",
	// 	"--proto_path=protobuf", // 指定 proto 路径
	// 	"--go_out=.",            // 指定 Go 输出目录（当前目录）
	// 	"--go-grpc_out=.",       // 指定 gRPC Go 输出目录（当前目录）
	// 	"--tool_out=.",
	// 	config.ProtobufConf.ProtobufName, // 指定 Proto 文件名
	// )

	// 捕获命令的标准输出和标准错误
	output, err := cmd.CombinedOutput()
	if err != nil {
		// 如果执行出错，打印错误信息
		fmt.Printf("执行 protoc 命令时出错: %v\n", err)
		fmt.Printf("错误输出: %s\n", output)
		return err
	}

	// Exec Command
	// Exec swag init
	// swagInitCmd := exec.Command("swag", "init", "--parseDependency")
	swagInitCmd := exec.Command("swag", "init")
	swagInitCmd.Stdout = os.Stdout
	swagInitCmd.Stderr = os.Stderr
	if err := swagInitCmd.Run(); err != nil {
		fmt.Printf("swagInitCmd Run Err : %+v", err)
		fmt.Print("\r\n")
		return toolerrors.NewError(err)
	}
	return nil
}

/**
 * @description: askIsFirstCreate
 * @author: Jerry.Yang
 * @date: 2023-12-11 16:15:07
 * @return {*}
 */
func (n *NewApi) askIsFirstCreate() (bool, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "isFirstCreate",
			Prompt: &survey.Confirm{
				Message: "please select is first create ? ",
				Default: false,
			},
		},
	}

	// set answer
	type Answer struct {
		IsFirstCreate bool `survey:"isFirstCreate"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return false, err
	}
	return answer.IsFirstCreate, nil
}

/**
 * @description: askIsAppend
 * @author: Jerry.Yang
 * @date: 2023-12-11 16:16:04
 * @return {*}
 */
func (n *NewApi) askIsAppend() (bool, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "isAppend",
			Prompt: &survey.Confirm{
				Message: "please select is append ? ",
				Default: false,
			},
		},
	}

	// set answer
	type Answer struct {
		IsAppend bool `survey:"isAppend"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return false, err
	}
	return answer.IsAppend, nil
}

/**
 * @description: askIsExtend
 * @author: Jerry.Yang
 * @date: 2025年02月21日18:38:24
 * @return {*}
 */
func (n *NewApi) askIsExtend() (bool, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "isExtend",
			Prompt: &survey.Confirm{
				Message: "please select is extend ? ",
				Default: false,
			},
		},
	}

	// set answer
	type Answer struct {
		IsExtend bool `survey:"isExtend"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return false, err
	}
	return answer.IsExtend, nil
}

/**
 * @description: askProtobufImportPath
 * @author: Jerry.Yang
 * @date: 2025-02-21 17:23:43
 * @return {*}
 */
func (n *NewApi) askExtendPath() (string, error) {
	// define question
	questions := []*survey.Question{
		{
			Name: "extendPath",
			Prompt: &survey.Input{
				Message: "please input extendPath ? ",
				Default: "",
			},
		},
	}

	// set answer
	type Answer struct {
		ExtendPath string `survey:"extendPath"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return "", err
	}
	return answer.ExtendPath, nil
}
