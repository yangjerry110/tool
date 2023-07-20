/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 15:36:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-10 19:51:15
 * @Description: newapi
 */
package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golib/cli"
)

type NewApiCommands interface {
	NewApi(ctx *cli.Context) error
	CreateNewApi() error
	AskIsFirstCreate() error
}

/**
 * @description: NewApi
 * @author: Jerry.Yang
 * @date: 2023-05-24 17:09:27
 * @return {*}
 */
type NewApi struct {
	IsFirstCreate bool
	IsAppend      bool
}

/**
 * @description: NewApiParams
 * @author: Jerry.Yang
 * @date: 2023-05-24 17:09:17
 * @return {*}
 */
var NewApiParams = &NewApi{}

/**
 * @description: NewApi
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:37:56
 * @return {*}
 */
func (n *NewApi) NewApi(ctx *cli.Context) error {

	/**
	 * @step
	 * @setProject
	 **/
	if err := CreateInitCommands().SetImportProjectPath(); err != nil {
		return err
	}

	/**
	 * @step
	 * @setProtobufName
	 **/
	if err := CreateInitCommands().SetProtobufName(ctx); err != nil {
		return err
	}

	/**
	 * @step
	 * @执行询问，是否第一次创建
	 **/
	if err := n.AskIsFirstCreate(); err != nil {
		return err
	}

	/**
	 * @step
	 * @执行询问，是否追加
	 **/
	if err := n.AskIsAppend(); err != nil {
		return err
	}

	/**
	 * @step
	 * @createNewApi
	 **/
	if err := n.CreateNewApi(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateNewApi
 * @author: Jerry.Yang
 * @date: 2023-05-24 17:23:28
 * @return {*}
 */
func (n *NewApi) CreateNewApi() error {

	/**
	 * @step
	 * @执行命令行
	 **/
	cmd := exec.Command("protoc", "-I", "protobuf", fmt.Sprintf("--go_opt=module=%s/vo/protobuf", InitParams.ProjectImportPath), "--go_out=plugins=grpc:vo/protobuf", fmt.Sprintf("--tool_out=isFirstCreate=%t,isAppend=%t:vo/protobuf", NewApiParams.IsFirstCreate, NewApiParams.IsAppend), InitParams.ProtobufName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

/**
 * @description: AskIsFirstCreate
 * @author: Jerry.Yang
 * @date: 2023-05-24 17:06:59
 * @return {*}
 */
func (n *NewApi) AskIsFirstCreate() error {
	/**
	 * @step
	 * @定义问题列表
	 **/
	questions := []*survey.Question{
		{
			Name: "isFirstCreate",
			Prompt: &survey.Confirm{
				Message: "please select is first create ? ",
				Default: false,
			},
		},
	}

	/**
	 * @step
	 * @定义答案列表
	 **/
	type Answer struct {
		IsFirstCreate bool `survey:"isFirstCreate"`
	}

	/**
	 * @step
	 * @执行问答
	 **/
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @获取答案
	 **/
	NewApiParams.IsFirstCreate = answer.IsFirstCreate
	return nil
}

/**
 * @description: AskIsAppend
 * @author: Jerry.Yang
 * @date: 2023-05-25 20:48:22
 * @return {*}
 */
func (n *NewApi) AskIsAppend() error {
	/**
	 * @step
	 * @定义问题列表
	 **/
	questions := []*survey.Question{
		{
			Name: "isAppend",
			Prompt: &survey.Confirm{
				Message: "please select is append ? ",
				Default: false,
			},
		},
	}

	/**
	 * @step
	 * @定义答案
	 **/
	type Answer struct {
		IsAppend bool `survey:"isAppend"`
	}

	/**
	 * @step
	 * @执行问答
	 **/
	answer := &Answer{}
	survey.Ask(questions, answer)

	/**
	 * @step
	 * @赋值答案
	 **/
	NewApiParams.IsAppend = answer.IsAppend
	return nil
}
