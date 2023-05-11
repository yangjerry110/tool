/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 17:25:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 14:44:28
 * @Description: vo commands
 */
package commands

import (
	"fmt"
	"os"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/templates/vo"
)

type NewVoCommands interface {
	NewVo(ctx *cli.Context) error
	CreateNewVo() error
	CreateVo() error
	CreateWd() error
	CreateFile() error
}

type NewVo struct {
	VoInputPath  string
	VoOutputPath string
}

/**
 * @description: NewVoParams
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:27:52
 * @return {*}
 */
var NewVoParams = &NewVo{}

/**
 * @description: NewVo
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-09 17:05:19
 * @return {*}
 */
func (n *NewVo) NewVo(ctx *cli.Context) error {

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
	err = n.CreateNewVo()
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateNewVo
 * @author: Jerry.Yang
 * @date: 2023-05-10 10:49:53
 * @return {*}
 */
func (n *NewVo) CreateNewVo() error {
	/**
	 * @step
	 * @创建vo
	 **/
	NewAppParams.AppVoInputFileName = fmt.Sprintf("%sInputVo.go", InitParms.AppName)
	NewAppParams.AppVoOutputFileName = fmt.Sprintf("%sOutputVo.go", InitParms.AppName)
	err := vo.CreateNewVo().SaveTemplate(fmt.Sprintf("%s%s/%s", InitParms.ProjectPath, "vo", "input"), fmt.Sprintf("%s%s/%s", InitParms.ProjectPath, "vo", "output"), NewAppParams.AppVoInputFileName, NewAppParams.AppVoOutputFileName, InitParms.AppName)
	if err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateVo
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:32:51
 * @return {*}
 */
func (n *NewVo) CreateVo() error {

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
 * @date: 2023-04-25 17:30:12
 * @return {*}
 */
func (n *NewVo) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	inputPath := fmt.Sprintf("%s/%s/%s", InitParms.ProjectPath, "vo", "input")

	/**
	 * @step
	 * @创建configPath
	 **/
	err := os.MkdirAll(inputPath, 0777)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @获取config的path
	 **/
	outputPath := fmt.Sprintf("%s/%s/%s", InitParms.ProjectPath, "vo", "output")

	/**
	 * @step
	 * @创建configPath
	 **/
	err = os.MkdirAll(outputPath, 0777)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @赋值
	 **/
	NewVoParams.VoInputPath = inputPath
	NewVoParams.VoOutputPath = outputPath
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:32:25
 * @return {*}
 */
func (n *NewVo) CreateFile() error {

	/**
	 * @step
	 * @创建input test
	 **/
	err := vo.CreateTestInputVo().SaveTemplate(NewVoParams.VoInputPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建output test
	 **/
	err = vo.CreateTestOutputVo().SaveTemplate(NewVoParams.VoOutputPath)
	if err != nil {
		return err
	}
	return nil
}
