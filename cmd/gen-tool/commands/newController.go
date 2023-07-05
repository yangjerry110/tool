/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 16:44:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 15:20:02
 * @Description: controller
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gen-tool/errors"
	"github.com/yangjerry110/tool/cmd/gen-tool/templates/controller"
)

type NewControllerCommands interface {
	NewController(ctx *cli.Context) error
	CreateNewController() error
	AppendFuncBaseDao() error
	AppendFuncController() error
	CreateController() error
	CreateWd() error
	CreateFile() error
	AskIsAppend() (bool, error)
	AskAppendFileName() error
}

type NewController struct {
	ControllerPath           string
	AppendBaseControllerName string
}

/**
 * @description: NewControllerParams
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:47:07
 * @return {*}
 */
var NewControllerParams = &NewController{}

/**
 * @description: NewController
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-09 17:02:15
 * @return {*}
 */
func (n *NewController) NewController(ctx *cli.Context) error {

	/**
	 * @step
	 * @设置projectPath
	 **/
	err := CreateInitCommands().SetProjectPath()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @设置projectImportPath
	 **/
	err = CreateInitCommands().SetImportProjectPath()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @设置appName
	 **/
	err = CreateInitCommands().SetControllerName(ctx)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @提问是否是追加
	 **/
	isAppend, err := n.AskIsAppend()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @假如是追加，追问追加的文件路径
	 **/
	if isAppend {
		err = n.AskAppendFileName()
		if err != nil {
			return err
		}

		/**
		 * @step
		 * @追加文件
		 **/
		err = n.AppendFuncController()
		if err != nil {
			return err
		}
		return nil
	}

	/**
	 * @step
	 * @创建controller
	 **/
	err = n.CreateNewController()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateConfig
 * @author: Jerry.Yang
 * @date: 2023-04-24 15:06:35
 * @return {*}
 */
func (n *NewController) CreateController() error {

	/**
	 * @step
	 * @创建config的文件夹
	 **/
	err := n.CreateWd()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建文件
	 **/
	err = n.CreateFile()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateNewController
 * @author: Jerry.Yang
 * @date: 2023-05-10 10:48:19
 * @return {*}
 */
func (n *NewController) CreateNewController() error {

	/**
	 * @step
	 * @创建controller
	 **/
	NewAppParams.AppControllerFileName = fmt.Sprintf("%sController.go", InitParams.ControllerName)
	err := controller.CreateNewController().SaveTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "controller"), InitParams.ProjectImportPath, InitParams.ControllerName, NewAppParams.AppControllerFileName)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @appendBase
	 **/
	err = n.AppendFuncBaseDao()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: AppendController
 * @author: Jerry.Yang
 * @date: 2023-05-11 17:13:30
 * @return {*}
 */
func (n *NewController) AppendFuncController() error {

	/**
	 * @step
	 * @append controller
	 **/
	err := controller.CreateNewController().SaveAppendFuncTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "controller"), InitParams.ControllerName, NewControllerParams.AppendBaseControllerName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: AppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:16:22
 * @return {*}
 */
func (n *NewController) AppendFuncBaseDao() error {
	err := controller.CreateBaseController().AppendFuncTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "controller"), InitParams.ControllerName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateWd
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:47:15
 * @return {*}
 */
func (n *NewController) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "controller")

	/**
	 * @step
	 * @创建configPath
	 **/
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @赋值
	 **/
	NewControllerParams.ControllerPath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:48:30
 * @return {*}
 */
func (n *NewController) CreateFile() error {

	/**
	 * @step
	 * @创建base
	 **/
	err := controller.CreateBaseController().SaveTemplate(NewControllerParams.ControllerPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建test
	 **/
	err = controller.CreateDemoController().SaveTemplate(NewControllerParams.ControllerPath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: AskIsAppend
 * @author: Jerry.Yang
 * @date: 2023-05-11 16:47:06
 * @return {*}
 */
func (n *NewController) AskIsAppend() (bool, error) {
	/**
	 * @step
	 * @初始化reader
	 **/
	reader := bufio.NewReader(os.Stdin)

	/**
	 * @step
	 * @定义输入的提示
	 **/
	fmt.Print("\r\n")
	fmt.Println("是否在当前已有的controller文件追加？")
	fmt.Print("\r\n")
	fmt.Print("回答: yes or no   ")
	fmt.Print("=> ")

	/**
	 * @step
	 * @获取输入的内容
	 **/
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("发生错误 : %+v \r\n", err)
		return false, err
	}

	/**
	 * @step
	 * @假如输入内容为空，报错直接
	 **/
	if len(text) == 1 {
		return false, nil
	}

	/**
	 * @step
	 * @删除换行
	 **/
	text = strings.ReplaceAll(text, "\n", "")
	if text == "yes" {
		return true, nil
	}
	return false, nil
}

/**
 * @description: AskAppendFileName
 * @author: Jerry.Yang
 * @date: 2023-05-11 16:53:47
 * @return {*}
 */
func (n *NewController) AskAppendFileName() error {

	/**
	 * @step
	 * @初始化reader
	 **/
	reader := bufio.NewReader(os.Stdin)

	/**
	 * @step
	 * @定义输入的提示
	 **/
	fmt.Print("\r\n")
	fmt.Println("追加的文件名称？")
	fmt.Print("\r\n")
	fmt.Print("回答: 只需要controller名称; ps: testControoler.go => test   ")
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
		return errors.ErrAppendControllerPathIsEmpyty
	}

	/**
	 * @step
	 * @删除换行
	 **/
	text = strings.ReplaceAll(text, "\n", "")
	NewControllerParams.AppendBaseControllerName = text
	return nil
}
