/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:14:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 18:30:56
 * @Description: service
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/errors"
	"github.com/yangjerry110/tool/cmd/gin-framework/templates/service"
)

type NewServiceCommands interface {
	NewService(ctx *cli.Context) error
	CreateNewService() error
	AppendFuncService() error
	AppendFuncBaseDao() error
	CreateService() error
	CreateWd() error
	CreateFile() error
	AskIsAppend() (bool, error)
	AskAppendFileName() error
}

type NewService struct {
	ServicePath           string
	AppendBaseServiceName string
}

/**
 * @description: ServiceParams
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:15:10
 * @return {*}
 */
var NewServiceParams = &NewService{}

/**
 * @description: NewService
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-09 17:03:47
 * @return {*}
 */
func (n *NewService) NewService(ctx *cli.Context) error {

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
	err = CreateInitCommands().SetServiceName(ctx)
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
		err = n.AppendFuncService()
		if err != nil {
			return err
		}
		return nil
	}

	/**
	 * @step
	 * @创建controller
	 **/
	err = n.CreateNewService()
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
 * @description: CreateNewService
 * @author: Jerry.Yang
 * @date: 2023-05-10 10:49:10
 * @return {*}
 */
func (n *NewService) CreateNewService() error {
	/**
	 * @step
	 * @创建service
	 **/
	NewAppParams.AppServiceFileName = fmt.Sprintf("%sService.go", InitParams.ServiceName)
	err := service.CreateNewService().SaveTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "service"), InitParams.ProjectImportPath, InitParams.ServiceName, NewAppParams.AppServiceFileName)
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
 * @description: AppendService
 * @author: Jerry.Yang
 * @date: 2023-05-11 18:29:23
 * @return {*}
 */
func (n *NewService) AppendFuncService() error {

	/**
	 * @step
	 * @append controller
	 **/
	err := service.CreateNewService().SaveAppendFuncTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "service"), InitParams.ServiceName, NewServiceParams.AppendBaseServiceName)
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
func (n *NewService) AppendFuncBaseDao() error {
	err := service.CreateBaseService().AppendFuncTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "service"), InitParams.ServiceName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateLogger
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:22:55
 * @return {*}
 */
func (n *NewService) CreateService() error {

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
func (n *NewService) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "service")

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
	NewServiceParams.ServicePath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:59:32
 * @return {*}
 */
func (n *NewService) CreateFile() error {

	/**
	 * @step
	 * @创建base
	 **/
	err := service.CreateBaseService().SaveTemplate(NewServiceParams.ServicePath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建beforStart
	 **/
	err = service.CreateBeforStartService().SaveTemplate(NewServiceParams.ServicePath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建test
	 **/
	err = service.CreateTestService().SaveTemplate(NewServiceParams.ServicePath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建common
	 **/
	err = service.CreateCommonService().SaveTemplate(NewServiceParams.ServicePath, InitParams.ProjectImportPath)
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
func (n *NewService) AskIsAppend() (bool, error) {
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
	fmt.Println("是否在当前已有的service文件追加？")
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
func (n *NewService) AskAppendFileName() error {

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
	fmt.Print("回答: 只需要service名称; ps: testService.go => test   ")
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
		return errors.ErrAppendServicePathIsEmpty
	}

	/**
	 * @step
	 * @删除换行
	 **/
	text = strings.ReplaceAll(text, "\n", "")
	NewServiceParams.AppendBaseServiceName = text
	return nil
}
