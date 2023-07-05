/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-26 10:42:10
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-26 10:46:54
 * @Description: newTemplates
 */
package commands

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type NewTemplateCommands interface {
	CreateTemplate() error
	CreateWd() error
	CreateFile() error
}

type NewTemplate struct{}

/**
 * @description: CreateTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:47:09
 * @return {*}
 */
func (n *NewTemplate) CreateTemplate() error {
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
 * @date: 2023-04-26 10:44:18
 * @return {*}
 */
func (n *NewTemplate) CreateWd() error {
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:46:43
 * @return {*}
 */
func (n *NewTemplate) CreateFile() error {

	/**
	 * @step
	 * @创建main
	 **/
	err := templates.CreateMainTemplate().SaveTemplate(InitParams.ProjectPath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建mod
	 **/
	err = templates.CreateModTemplate().SaveTemplate(InitParams.ProjectPath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}
	return nil
}
