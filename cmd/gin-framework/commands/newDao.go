/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:06:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 17:11:12
 * @Description: dao commands
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates/dao"
)

type NewDaoCommands interface {
	CreateDao() error
	CreateWd() error
	CreateFile() error
}

type NewDao struct {
	DaoPath string
}

/**
 * @description: NewDaoParams
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:08:00
 * @return {*}
 */
var NewDaoParams = &NewDao{}

/**
 * @description: CreateDao
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:11:21
 * @return {*}
 */
func (n *NewDao) CreateDao() error {

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
 * @date: 2023-04-24 17:07:24
 * @return {*}
 */
func (n *NewDao) CreateWd() error {
	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParms.ProjectPath, "dao")

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
	NewDaoParams.DaoPath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:10:55
 * @return {*}
 */
func (n *NewDao) CreateFile() error {

	/**
	 * @step
	 * @创建base
	 **/
	err := dao.CreateBaseDao().SaveTemplate(NewDaoParams.DaoPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建common
	 **/
	err = dao.CreateCommonDao().SaveTemplate(NewDaoParams.DaoPath)
	if err != nil {
		return err
	}
	return nil
}
