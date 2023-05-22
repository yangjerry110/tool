/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 17:25:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 16:06:46
 * @Description: vo commands
 */
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gin-framework/errors"
	"github.com/yangjerry110/tool/cmd/gin-framework/templates/vo"
)

type NewVoCommands interface {
	NewVo(ctx *cli.Context) error
	CreateNewVo() error
	CreateVo() error
	CreateWd() error
	CreateFile() error
	AppendVo() error
	AskIsAppend() (bool, error)
	AskAppendFileName() error
}

type NewVo struct {
	AppendBaseVoName string
	VoInputPath      string
	VoOutputPath     string
	VoProtobufPath   string
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
	err = CreateInitCommands().SetVoName(ctx)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @追问是否是追加
	 **/
	isAppend, err := n.AskIsAppend()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @判断追问是否添加之后，再执行添加
	 **/
	if isAppend {

		/**
		 * @step
		 * @追问添加文件的名称
		 **/
		err := n.AskAppendFileName()
		if err != nil {
			return err
		}

		/**
		 * @step
		 * @添加内容到追加的文件
		 **/
		err = n.AppendVo()
		if err != nil {
			return err
		}
		return nil
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
 * @description: AppendVo
 * @author: Jerry.Yang
 * @date: 2023-05-16 11:03:01
 * @return {*}
 */
func (n *NewVo) AppendVo() error {

	/**
	 * @step
	 * @append vo
	 **/
	err := vo.CreateNewVo().SaveAppendFuncInputTemplate(fmt.Sprintf("%s%s/%s", InitParams.ProjectPath, "vo", "input"), InitParams.VoName, NewVoParams.AppendBaseVoName)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @add output
	 **/
	err = vo.CreateNewVo().SaveAppendFuncOutputTemplate(fmt.Sprintf("%s%s/%s", InitParams.ProjectPath, "vo", "output"), InitParams.VoName, NewVoParams.AppendBaseVoName)
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
	NewAppParams.AppVoInputFileName = fmt.Sprintf("%sInputVo.go", InitParams.VoName)
	NewAppParams.AppVoOutputFileName = fmt.Sprintf("%sOutputVo.go", InitParams.VoName)
	err := vo.CreateNewVo().SaveTemplate(fmt.Sprintf("%s%s/%s", InitParams.ProjectPath, "vo", "input"), fmt.Sprintf("%s%s/%s", InitParams.ProjectPath, "vo", "output"), NewAppParams.AppVoInputFileName, NewAppParams.AppVoOutputFileName, InitParams.VoName)
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
	inputPath := fmt.Sprintf("%s/%s/%s", InitParams.ProjectPath, "vo", "input")

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
	outputPath := fmt.Sprintf("%s/%s/%s", InitParams.ProjectPath, "vo", "output")

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
	 * @创建protobufPath
	 **/
	protobufPath := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "protobufVo")
	err = os.MkdirAll(protobufPath, 0777)
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
	err := vo.CreateDemoInputVo().SaveTemplate(NewVoParams.VoInputPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建output test
	 **/
	err = vo.CreateDemoOutputVo().SaveTemplate(NewVoParams.VoOutputPath)
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
func (n *NewVo) AskIsAppend() (bool, error) {
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
	fmt.Println("是否在当前已有的vo文件追加？")
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
func (n *NewVo) AskAppendFileName() error {

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
	fmt.Print("回答: 只需要vo名称; ps: testInputVo.go => test   ")
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
	NewVoParams.AppendBaseVoName = text
	return nil
}
