/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 15:59:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:24:53
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

type NewHttp struct {
	CliContext *cli.Context
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:42:58
 * @return {*}
 */
func (n *NewHttp) New() error {

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

	// ask protobufImportPath
	protobufImportPath, err := n.askProtobufImportPath()
	if err != nil {
		return err
	}

	// Exec Command
	// Exec protoc
	cmd := exec.Command("protoc", "-I", "protobuf", fmt.Sprintf("--tool_out=isFirstCreate=%t,isAppend=%t,protobufImportPath=%s:vo/protobuf", isFirstCreate, isAppend, protobufImportPath))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd Run Err : %+v", err)
		fmt.Print("\r\n")
		return toolerrors.NewError(err)
	}

	// Exec Command
	// Exec swag init
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
 * @description: askProtobufImportPath
 * @author: Jerry.Yang
 * @date: 2025-02-21 17:23:43
 * @return {*}
 */
func (n *NewHttp) askProtobufImportPath() (string, error) {
	// define question
	questions := []*survey.Question{
		{
			Name: "protobufImportPath",
			Prompt: &survey.Input{
				Message: "please input protobufImportPath ? ",
				Default: "",
			},
		},
	}

	// set answer
	type Answer struct {
		ProtobufImportPath string `survey:"protobufImportPath"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return "", err
	}
	return answer.ProtobufImportPath, nil
}

/**
 * @description: askIsFirstCreate
 * @author: Jerry.Yang
 * @date: 2023-12-11 16:15:07
 * @return {*}
 */
func (n *NewHttp) askIsFirstCreate() (bool, error) {

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
func (n *NewHttp) askIsAppend() (bool, error) {

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
