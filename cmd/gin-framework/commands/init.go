/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 16:36:03
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 16:36:42
 * @Description: init
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/yangjerry110/tool/cmd/gin-framework/errors"
)

type InitCommands interface {
	AskInit() error
	AskProjectName() error
	SetProjectPath() error
}

type Init struct {
	ProjectImportPath string
	ProjectPath       string
	ProjectName       string
}

/**
 * @description: initParam
 * @author: Jerry.Yang
 * @date: 2023-04-23 16:37:53
 * @return {*}
 */
var InitParms = &Init{}

/**
 * @description: askInit
 * @author: Jerry.Yang
 * @date: 2023-04-23 16:43:06
 * @return {*}
 */
func (i *Init) AskInit() error {

	/**
	 * @step
	 * @ask projectName
	 **/
	err := i.AskProjectName()
	if err != nil {
		return err
	}

	return nil
}

/**
 * @description: AskProjectName
 * @author: Jerry.Yang
 * @date: 2023-04-23 16:40:07
 * @return {*}
 */
func (i *Init) AskProjectName() error {

	/**
	 * @step
	 * @初始化reader
	 **/
	reader := bufio.NewReader(os.Stdin)

	/**
	 * @step
	 * @定义输入的提示
	 **/
	fmt.Println("请输入你的项目目录，按回车结束")
	fmt.Print("\r\n")
	fmt.Print("类似: github.com/test/test    ")
	fmt.Print("=> ")

	/**
	 * @step
	 * @获取输入的内容
	 **/
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("发生错误 : %+v \r\n", err)
		return err
	}

	/**
	 * @step
	 * @假如输入内容为空，报错直接
	 **/
	if len(text) == 1 {
		return errors.ErrProjectPathIsEmpty
	}

	/**
	 * @step
	 * @删除换行
	 **/
	text = strings.ReplaceAll(text, "\n", "")

	/**
	 * @step
	 * @赋值
	 **/
	InitParms.ProjectImportPath = text
	return nil
}

/**
 * @description: SetProjectPath
 * @author: Jerry.Yang
 * @date: 2023-04-24 10:49:32
 * @return {*}
 */
func (i *Init) SetProjectPath() error {

	/**
	 * @step
	 * @判断当前的projectPath是否存在值
	 **/
	if InitParms.ProjectPath != "" {
		return nil
	}

	/**
	 * @step
	 * @获取当前目录
	 **/
	thisPath, err := os.Getwd()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @进行赋值
	 **/
	InitParms.ProjectPath = fmt.Sprintf("%s/%s", thisPath, InitParms.ProjectName)
	return nil
}
