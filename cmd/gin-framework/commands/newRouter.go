/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 15:53:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 16:00:17
 * @Description: new router
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates/router"
)

type NewRouterCommands interface {
	CreateRouter() error
	CreateWd() error
	CreateFile() error
}

type NewRouter struct {
	RouterPath string
}

/**
 * @description: NewRouterParams
 * @author: Jerry.Yang
 * @date: 2023-04-25 15:54:25
 * @return {*}
 */
var NewRouterParams = &NewRouter{}

/**
 * @description: CreateLogger
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
	path := fmt.Sprintf("%s/%s", InitParms.ProjectPath, "router")

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
	err = router.CreateCommonRouter().SaveTemplate(NewRouterParams.RouterPath, InitParms.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建test
	 **/
	err = router.CreateTestRouter().SaveTemplate(NewRouterParams.RouterPath, InitParms.ProjectImportPath)
	if err != nil {
		return err
	}
	return nil
}
