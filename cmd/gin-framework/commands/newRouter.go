/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:53:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 16:01:16
 * @Description: new router
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/errors"
	"github.com/yangjerry110/tool/cmd/gin-framework/templates/router"
)

type NewRouterCommands interface {
	NewRouter(ctx *cli.Context) error
	CreateNewRouter() error
	AppendFuncRotuer() error
	AppendBaseFuncRouter() error
	CreateRouter() error
	CreateWd() error
	CreateFile() error
	AskIsAppend() (bool, error)
	AskAppendFileName() error
}

type NewRouter struct {
	RouterPath           string
	AppendBaseRouterName string
}

/**
 * @description: NewRouterParams
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:54:25
 * @return {*}
 */
var NewRouterParams = &NewRouter{}

/**
 * @description: NewRouter
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:38:21
 * @return {*}
 */
func (n *NewRouter) NewRouter(ctx *cli.Context) error {

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
	 * @设置routerName
	 **/
	err = CreateInitCommands().SetRouterName(ctx)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @追问是否追加
	 **/
	isAppend, err := n.AskIsAppend()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @假如追问，则直接进入到append环节
	 **/
	if isAppend {

		/**
		 * @step
		 * @追问追加的文件地址
		 **/
		err := n.AskAppendFileName()
		if err != nil {
			return err
		}

		/**
		 * @step
		 * @append
		 **/
		err = n.AppendFuncRotuer()
		if err != nil {
			return err
		}
		return nil
	}

	/**
	 * @step
	 * @创建new
	 **/
	err = n.CreateNewRouter()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @追加base
	 **/
	err = n.AppendBaseFuncRouter()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateNewRouter
 * @author: Jerry.Yang
 * @date: 2023-05-16 16:35:09
 * @return {*}
 */
func (n *NewRouter) CreateNewRouter() error {
	NewAppParams.AppRouterFileName = fmt.Sprintf("%sRouter.go", InitParams.RouterName)
	err := router.CreateNewRouter().SaveTemplate(fmt.Sprintf("%srouter", InitParams.ProjectPath), InitParams.ProjectImportPath, InitParams.RouterName)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @添加func到base
	 **/
	err = router.CreateBaseRouter().SaveAppendFuncTemplate(fmt.Sprintf("%srouter", InitParams.ProjectPath), InitParams.RouterName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: AppendFuncRotuer
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:35:46
 * @return {*}
 */
func (n *NewRouter) AppendFuncRotuer() error {
	err := router.CreateNewRouter().SaveAppendFuncTemplate(fmt.Sprintf("%srouter", InitParams.ProjectPath), NewRouterParams.AppendBaseRouterName, InitParams.RouterName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: AppendBaseFuncRouter
 * @author: Jerry.Yang
 * @date: 2023-05-16 17:38:00
 * @return {*}
 */
func (n *NewRouter) AppendBaseFuncRouter() error {
	err := router.CreateBaseRouter().SaveAppendFuncTemplate(fmt.Sprintf("%srouter", InitParams.ProjectPath), InitParams.RouterName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateRouter
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:22:55
 * @return {*}
 */
func (n *NewRouter) CreateRouter() error {
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
 * @description: CreateWd
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:55:04
 * @return {*}
 */
func (n *NewRouter) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "router")

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
	NewRouterParams.RouterPath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:59:32
 * @return {*}
 */
func (n *NewRouter) CreateFile() error {

	/**
	 * @step
	 * @创建base
	 **/
	err := router.CreateBaseRouter().SaveTemplate(NewRouterParams.RouterPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建common
	 **/
	err = router.CreateCommonRouter().SaveTemplate(NewRouterParams.RouterPath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建test
	 **/
	err = router.CreateDemoRouter().SaveTemplate(NewRouterParams.RouterPath, InitParams.ProjectImportPath)
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
func (n *NewRouter) AskIsAppend() (bool, error) {
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
	fmt.Println("是否在当前已有的Router文件追加？")
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
func (n *NewRouter) AskAppendFileName() error {

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
	fmt.Print("回答: 只需要router名称; ps: testRouter.go => test   ")
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
		return errors.ErrAppendRouterPathIsEmpty
	}

	/**
	 * @step
	 * @删除换行
	 **/
	text = strings.ReplaceAll(text, "\n", "")
	NewRouterParams.AppendBaseRouterName = text
	return nil
}
