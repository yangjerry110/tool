/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 15:00:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 16:21:04
 * @Description: plugin protoc_gen_tool
 */
package protocgentoolservice

import (
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type Plugin struct {
	Plugin        *protogen.Plugin
	IsFirstCreate bool
	IsAppend      bool
}

/**
 * @description: Generate
 * @author: Jerry.Yang
 * @date: 2023-12-13 15:01:58
 * @return {*}
 */
func (p *Plugin) Generate() error {

	// set projectPath
	if err := conf.CreateConf(&config.ProjectPath{}).SetConfig(); err != nil {
		return err
	}

	// set projectImportPath
	if err := conf.CreateConf(&config.ProjectImportPath{}).SetConfig(); err != nil {
		return err
	}

	// Get Params by flags
	// Set Params to config
	if err := conf.CreateConf(&config.ProtocGenTool{IsFirstCreate: p.IsFirstCreate, IsAppend: p.IsAppend}).SetConfig(); err != nil {
		return err
	}

	// judge plugin
	// if == nil; return err
	if p.Plugin == nil {
		return errors.ErrProtocGenToolServiceNoPlugin
	}

	// judge plugin files
	// if files's num < 1 return err
	if len(p.Plugin.Files) < 1 {
		return errors.ErrProtocGenToolServiceNoFiles
	}

	// With plugin.Fiels, we can get all the input proto files
	// If we need to generate code for this file, then go to the generateFile() logic
	// and pass g and f together
	for _, file := range p.Plugin.Files {
		// Judge Generate
		// If true , run file generate
		if file.Generate {
			if err := CreateProtoGenToolService(&File{File: file}).Generate(); err != nil {
				return err
			}
		}
	}
	return nil
}
