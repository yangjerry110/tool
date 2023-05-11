/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:14:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 15:29:26
 * @Description: service
 */
package commands

import (
	"fmt"
	"os"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/templates/service"
)

type NewServiceCommands interface {
	NewService(ctx *cli.Context) error
	CreateNewService() error
	AppendFuncBaseDao() error
	CreateService() error
	CreateWd() error
	CreateFile() error
}

type NewService struct {
	ServicePath string
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
	err = CreateInitCommands().SetAppName(ctx)
	if err != nil {
		return err
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
	NewAppParams.AppServiceFileName = fmt.Sprintf("%sService.go", InitParms.AppName)
	err := service.CreateNewService().SaveTemplate(fmt.Sprintf("%s%s", InitParms.ProjectPath, "service"), InitParms.ProjectImportPath, InitParms.AppName, NewAppParams.AppServiceFileName)
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
	err := service.CreateBaseService().AppendFuncTemplate(fmt.Sprintf("%s%s", InitParms.ProjectPath, "service"), InitParms.AppName)
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
	path := fmt.Sprintf("%s/%s", InitParms.ProjectPath, "service")

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
	err = service.CreateBeforStartService().SaveTemplate(NewServiceParams.ServicePath, InitParms.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建test
	 **/
	err = service.CreateTestService().SaveTemplate(NewServiceParams.ServicePath, InitParms.ProjectImportPath)
	if err != nil {
		return err
	}
	return nil
}
