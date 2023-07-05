/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-23 15:30:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 22:12:15
 * @Description: generate
 */
package commands

import (
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type GenerateCommands interface {
	Generate(plugin *protogen.Plugin) error
}

type Generate struct{}

/**
 * @description: Generate
 * @param {*protogen.Plugin} plugin
 * @author: Jerry.Yang
 * @date: 2023-05-25 22:12:50
 * @return {*}
 */
func (g *Generate) Generate(plugin *protogen.Plugin) error {

	/**
	 * @step
	 * @判断pluginfiles
	 **/
	if len(plugin.Files) < 1 {
		return errors.Err_Plugins_Files_Is_Empty
	}

	/**
	 * @step
	 * @通过plugin.Fiels，我们可以拿到所有的输入的proto文件
	 * @如果我们需要对这个文件生成代码的话，那么就进入到generateFile()逻辑
	 * @并且把g和f一起传递过去
	 **/
	for _, file := range plugin.Files {
		if file.Generate {
			err := CreateFileCommands().GenerateFile(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
