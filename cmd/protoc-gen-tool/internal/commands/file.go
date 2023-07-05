/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-23 17:03:26
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 21:49:29
 * @Description: genfile 相关的操作
 */
package commands

import (
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type FileCommands interface {
	GetProtobufFile(file *protogen.File) (string, error)
	GenerateFile(file *protogen.File) error
}

type File struct{}

/**
 * @description: GetProtobufFile
 * @param {*protogen.File} file
 * @author: Jerry.Yang
 * @date: 2023-05-23 17:31:01
 * @return {*}
 */
func (f *File) GetProtobufFile(file *protogen.File) (string, error) {
	return string(file.Desc.FullName()), nil
}

/**
 * @description: GenerateFile
 * @param {*protogen.File} file
 * @author: Jerry.Yang
 * @date: 2023-05-24 14:24:36
 * @return {*}
 */
func (f *File) GenerateFile(file *protogen.File) error {

	/**
	 * @step
	 * @定义httpRules
	 **/
	httpRules := []*HttpRule{}

	/**
	 * @step
	 * @这一段代码仅仅只是为了忽略包含proto文件中包含了streamClient和streamServer的代码
	 **/
	isGenerated := false
	for _, srv := range file.Services {
		for _, method := range srv.Methods {
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}
			isGenerated = true
		}
	}

	/**
	 * @step
	 * @判断
	 **/
	if !isGenerated {
		return errors.Err_Gen_File_isGenerated_False
	}

	/**
	 * @step
	 * @file 的下一层级就是 services 层级
	 **/
	for _, service := range file.Services {
		methodHttpRules, err := CreateServiceCommands().GenService(service)
		if err != nil {
			return err
		}

		httpRules = append(httpRules, methodHttpRules...)
	}

	/**
	 * @step
	 * @判断httpRules
	 **/
	if len(httpRules) == 0 {
		return errors.Err_Http_Rules_Is_Empty
	}

	/**
	 * @step
	 * @获取当前的proto的文件名称
	 **/
	fileName, err := f.GetProtobufFile(file)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @GetNewRouterProtobufs
	 **/
	newRouterProtobufs, err := CreateRouterCommands().GetNewRouterProtobufs(fileName, httpRules)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @执行创建router
	 **/
	if err := CreateRouterCommands().CreateRouter(fileName, newRouterProtobufs); err != nil {
		return err
	}

	/**
	 * @step
	 * @getNewProtobufServices
	 **/
	newProtobufServies, err := CreateServiceCommands().GetServiceProtobufs(fileName, httpRules)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @创建service
	 **/
	if err := CreateServiceCommands().CreateService(fileName, httpRules, newProtobufServies); err != nil {
		return err
	}
	return nil
}
