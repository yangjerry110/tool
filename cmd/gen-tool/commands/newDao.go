/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:06:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 15:19:44
 * @Description: dao commands
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gen-tool/errors"
	"github.com/yangjerry110/tool/cmd/gen-tool/templates/dao"
)

type NewDaoCommands interface {
	NewDao(ctx *cli.Context) error
	CreateNewDao() error
	AppendFuncBaseDao() error
	CreateDao() error
	CreateWd() error
	CreateFile() error
	AskIsAppend() (bool, error)
	AskAppendFileName() error
}

type NewDao struct {
	AppendBaseDaoPath string
	DaoPath           string
}

/**
 * @description: NewDaoParams
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:08:00
 * @return {*}
 */
var NewDaoParams = &NewDao{}

/**
 * @description: NewDao
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-11 11:36:30
 * @return {*}
 */
func (n *NewDao) NewDao(ctx *cli.Context) error {

	/**
	 * @step
	 * @设置daoName
	 **/
	err := CreateInitCommands().SetDaoName(ctx)
	if err != nil {
		return nil
	}

	/**
	 * @step
	 * @设置projectPath
	 **/
	err = CreateInitCommands().SetProjectPath()
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
	 * @判断是否是追加
	 **/
	isAppend, err := n.AskIsAppend()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @假如是追加，则追问追加的文件
	 **/
	if isAppend {

		/**
		 * @step
		 * @追问需要追加的文件
		 **/
		err := n.AskAppendFileName()
		if err != nil {
			return err
		}

		/**
		 * @step
		 * @添加追加的数据
		 **/
		err = n.AppendFuncDao()
		if err != nil {
			return err
		}
		return nil
	}

	/**
	 * @step
	 * @创建newDao
	 **/
	err = n.CreateNewDao()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateNewDao
 * @author: Jerry.Yang
 * @date: 2023-05-11 11:24:28
 * @return {*}
 */
func (n *NewDao) CreateNewDao() error {
	NewAppParams.AppDaoFileName = fmt.Sprintf("%sDao.go", InitParams.DaoName)
	err := dao.CreateNewDao().SaveTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "dao"), InitParams.ProjectImportPath, InitParams.DaoName, NewAppParams.AppDaoFileName)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @appendFuncBaseGo
	 **/
	err = n.AppendFuncBaseDao()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: AppendFuncDao
 * @author: Jerry.Yang
 * @date: 2023-05-16 11:29:28
 * @return {*}
 */
func (n *NewDao) AppendFuncDao() error {

	/**
	 * @step
	 * @appendFunDao
	 **/
	err := dao.CreateNewDao().AppendFuncTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "dao"), InitParams.DaoName, NewDaoParams.AppendBaseDaoPath)
	if err != nil {
		return err
	}
	return err
}

/**
 * @description: AppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:16:22
 * @return {*}
 */
func (n *NewDao) AppendFuncBaseDao() error {
	err := dao.CreateBaseDao().AppendFuncTemplate(fmt.Sprintf("%s%s", InitParams.ProjectPath, "dao"), InitParams.DaoName)
	if err != nil {
		return err
	}
	return nil
}

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
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "dao")

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

/**
 * @description: AskIsAppend
 * @author: Jerry.Yang
 * @date: 2023-05-11 16:47:06
 * @return {*}
 */
func (n *NewDao) AskIsAppend() (bool, error) {
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
	fmt.Println("是否在当前已有的dao文件追加？")
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
func (n *NewDao) AskAppendFileName() error {

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
	fmt.Print("回答: 只需要dao名称; ps: testDao.go => test   ")
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
		return errors.ErrAppendVoPathIsEmprty
	}

	/**
	 * @step
	 * @删除换行
	 **/
	text = strings.ReplaceAll(text, "\n", "")
	NewDaoParams.AppendBaseDaoPath = text
	return nil
}
