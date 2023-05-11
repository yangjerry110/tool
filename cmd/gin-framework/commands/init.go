/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 16:36:03
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 11:35:34
 * @Description: init
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/errors"
	"golang.org/x/tools/go/packages"
)

type InitCommands interface {
	AskInit() error
	AskProjectName() error
	SetProjectPath() error
	SetProjectName(ctx *cli.Context) error
	SetAppName(ctx *cli.Context) error
	SetDaoName(ctx *cli.Context) error
	SetModelName(ctx *cli.Context) error
	SetImportProjectPath() error
}

type Init struct {
	ProjectImportPath string
	ProjectPath       string
	ProjectName       string
	AppName           string
	ModelName         string
	DaoName           string
	ModelConfigPath   string
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

/**
 * @description: SetProjectName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-10 16:49:02
 * @return {*}
 */
func (i *Init) SetProjectName(ctx *cli.Context) error {
	/**
	 * @step
	 * @获取项目名称
	 **/
	projectName := ctx.Args().First()
	if projectName == "" {
		return errors.ErrProjectNameIsEmpty
	}
	InitParms.ProjectName = projectName
	return nil
}

/**
 * @description: SetAppName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-10 16:48:20
 * @return {*}
 */
func (i *Init) SetAppName(ctx *cli.Context) error {

	/**
	 * @step
	 * @判断appName是否有值
	 **/
	if InitParms.AppName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数的名称
	 **/
	appName := ctx.Args().First()
	if appName == "" {
		return errors.ErrAppNameIsEmpty
	}
	InitParms.AppName = appName
	return nil
}

/**
 * @description: SetDaoName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-11 11:29:45
 * @return {*}
 */
func (i *Init) SetDaoName(ctx *cli.Context) error {

	/**
	 * @step
	 * @判断daoName是否有值
	 **/
	if InitParms.DaoName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数
	 **/
	daoName := ctx.Args().First()
	if daoName == "" {
		return errors.ErrDaoNameIsEmpty
	}

	/**
	 * @step
	 * @进行赋值
	 **/
	InitParms.DaoName = daoName
	return nil
}

/**
 * @description: SetModelName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-11 11:35:33
 * @return {*}
 */
func (i *Init) SetModelName(ctx *cli.Context) error {

	/**
	 * @step
	 * @判断modelName是否有值
	 **/
	if InitParms.ModelName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数的名称
	 **/
	modelName := ctx.Args().First()
	if modelName == "" {
		return errors.ErrModelNameIsEmpty
	}

	/**
	 * @step
	 * @进行赋值
	 **/
	InitParms.ModelName = modelName
	return nil
}

/**
 * @description: SetImportProjectPath
 * @author: Jerry.Yang
 * @date: 2023-05-10 16:45:47
 * @return {*}
 */
func (i *Init) SetImportProjectPath() error {

	/**
	 * @step
	 * @获取当前mod的module
	 **/
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName,
		Dir:  InitParms.ProjectPath,
	})
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @pkgs 判断
	 **/
	if len(pkgs) == 0 {
		return errors.ErrImportProjectPathIsEmpty
	}
	InitParms.ProjectImportPath = pkgs[0].PkgPath
	return nil
}
