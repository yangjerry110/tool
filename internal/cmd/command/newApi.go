/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 15:59:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:26:22
 * @Description: newApi Command
 */
package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
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

	// Exec Command
	// Exec protoc
	cmd := exec.Command("protoc", "-I", "protobuf", fmt.Sprintf("--go_opt=module=%s/vo/protobuf", config.ProjectImportPathConf.ImportPath), "--go_out=plugins=grpc:vo/protobuf", fmt.Sprintf("--tool_out=isFirstCreate=%t,isAppend=%t:vo/protobuf", isFirstCreate, isAppend), config.ProtobufConf.ProtobufName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// Exec Command
	// Exec swag init
	if err := exec.Command("swag", "init").Run(); err != nil {
		return err
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
