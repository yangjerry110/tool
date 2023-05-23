/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 15:40:36
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-22 15:50:44
 * @Description: new protobuf
 */
package commands

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates/protobuf"
)

type NewProtobufCommands interface {
	CreateProtobuf() error
	CreateWd() error
	CreateFile() error
}

type NewProtobuf struct {
	ProtobufPath string
}

/**
 * @description: NewProtobufParams
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:43:18
 * @return {*}
 */
var NewProtobufParams = &NewProtobuf{}

/**
 * @description: CreateProtobuf
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:46:02
 * @return {*}
 */
func (n *NewProtobuf) CreateProtobuf() error {

	/**
	 * @step
	 * @创建wd
	 **/
	err := n.CreateWd()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建file
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
 * @date: 2023-04-24 16:47:15
 * @return {*}
 */
func (n *NewProtobuf) CreateWd() error {

	/**
	 * @step
	 * @获取config的path
	 **/
	path := fmt.Sprintf("%s/%s", InitParams.ProjectPath, "protobuf")

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
	NewProtobufParams.ProtobufPath = path
	return nil
}

/**
 * @description: CreateFile
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:44:48
 * @return {*}
 */
func (n *NewProtobuf) CreateFile() error {

	/**
	 * @step
	 * @创建demo
	 **/
	err := protobuf.CreateDemoProtobuf().SaveTemplate(NewProtobufParams.ProtobufPath, InitParams.ProjectImportPath)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建http
	 **/
	err = protobuf.CreateHttpProtobuf().SaveTemplate(NewProtobufParams.ProtobufPath)
	if err != nil {
		return err
	}
	return nil
}
