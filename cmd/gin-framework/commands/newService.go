/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:14:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 17:32:36
 * @Description: service
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates/service"
)

type NewServiceCommands interface {
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
	err = service.CreateBeforStartService().SaveTemplate(NewRouterParams.RouterPath, InitParms.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建test
	 **/
	err = service.CreateTestService().SaveTemplate(NewRouterParams.RouterPath, InitParms.ProjectImportPath)
	if err != nil {
		return err
	}
	return nil
}
