/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 16:36:03
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 17:26:40
 * @Description: init
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gen-tool/errors"
	"golang.org/x/tools/go/packages"
)

type InitCommands interface {
	AskInit() error
	AskProjectName() error
	SetProjectPath() error
	SetProjectName(ctx *cli.Context) error
	SetImportProjectPath() error
	SetBaseProjectPath() error
	SetAppName(ctx *cli.Context) error
	SetControllerName(ctx *cli.Context) error
	SetRouterName(ctx *cli.Context) error
	SetServiceName(ctx *cli.Context) error
	SetDaoName(ctx *cli.Context) error
	SetModelName(ctx *cli.Context) error
	SetVoName(ctx *cli.Context) error
	SetProtobufName(ctx *cli.Context) error
}

type Init struct {
	ProjectImportPath string
	ProjectBasePath   string
	ProjectPath       string
	ModelConfigPath   string
	ProjectName       string
	AppName           string
	ControllerName    string
	RouterName        string
	ServiceName       string
	ModelName         string
	DaoName           string
	VoName            string
	ProtobufName      string
}

/**
 * @description: initParam
 * @author: Jerry.Yang
 * @date: 2023-04-23 16:37:53
 * @return {*}
 */
var InitParams = &Init{}

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
	InitParams.ProjectImportPath = text
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
	if InitParams.ProjectPath != "" {
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
	InitParams.ProjectPath = fmt.Sprintf("%s/%s", thisPath, InitParams.ProjectName)
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
		Dir:  InitParams.ProjectPath,
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
	InitParams.ProjectImportPath = pkgs[0].PkgPath
	return nil
}

/**
 * @description: GetBaseProjectPath
 * @author: Jerry.Yang
 * @date: 2023-05-11 17:27:58
 * @return {*}
 */
func (i *Init) SetBaseProjectPath() error {

	/**
	 * @step
	 * @获取项目跟目录
	 **/
	binary, err := os.Executable()
	if err != nil {
		return err
	}

	rootDir := filepath.Dir(filepath.Dir(binary))
	InitParams.ProjectBasePath = rootDir
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
	InitParams.ProjectName = projectName
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
	if InitParams.AppName != "" {
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
	InitParams.AppName = appName
	InitParams.ControllerName = appName
	InitParams.RouterName = appName
	InitParams.ServiceName = appName
	InitParams.VoName = appName
	return nil
}

/**
 * @description: SetControllerName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-11 16:30:16
 * @return {*}
 */
func (i *Init) SetControllerName(ctx *cli.Context) error {

	/**
	 * @step
	 * @判断ControllerName是否有值
	 **/
	if InitParams.ControllerName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数的名称
	 **/
	ControllerName := ctx.Args().First()
	if ControllerName == "" {
		return errors.ErrControllerNameIsEmpty
	}
	InitParams.ControllerName = ControllerName
	return nil
}

/**
 * @description: SetRouterName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-16 16:32:53
 * @return {*}
 */
func (i *Init) SetRouterName(ctx *cli.Context) error {

	/**
	 * @step
	 * @判断routerName是否有值
	 **/
	if InitParams.RouterName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数
	 **/
	routerName := ctx.Args().First()
	if routerName == "" {
		return errors.ErrRouterNameIsEmpty
	}
	InitParams.RouterName = routerName
	return nil
}

/**
 * @description: SetServiceName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-11 16:31:31
 * @return {*}
 */
func (i *Init) SetServiceName(ctx *cli.Context) error {
	/**
	 * @step
	 * @判断ServiceName是否有值
	 **/
	if InitParams.ServiceName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数的名称
	 **/
	ServiceName := ctx.Args().First()
	if ServiceName == "" {
		return errors.ErrServiceNameIsEmpty
	}
	InitParams.ServiceName = ServiceName
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
	if InitParams.DaoName != "" {
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
	InitParams.DaoName = daoName
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
	if InitParams.ModelName != "" {
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
	InitParams.ModelName = modelName
	return nil
}

/**
 * @description: SetVoName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-11 16:32:43
 * @return {*}
 */
func (i *Init) SetVoName(ctx *cli.Context) error {
	/**
	 * @step
	 * @判断VoName是否有值
	 **/
	if InitParams.VoName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数的名称
	 **/
	VoName := ctx.Args().First()
	if VoName == "" {
		return errors.ErrVoNameIsEmpty
	}
	InitParams.VoName = VoName
	return nil
}

/**
 * @description: SetProtobufName
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-24 17:26:28
 * @return {*}
 */
func (i *Init) SetProtobufName(ctx *cli.Context) error {

	/**
	 * @step
	 * @判断ProtobufName是否有值
	 **/
	if InitParams.ProtobufName != "" {
		return nil
	}

	/**
	 * @step
	 * @获取第一个参数的名称
	 **/
	ProtobufName := ctx.Args().First()
	if ProtobufName == "" {
		return errors.ErrProtobufNameIsEmpty
	}
	InitParams.ProtobufName = ProtobufName
	return nil
}
