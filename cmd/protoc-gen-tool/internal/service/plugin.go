/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-13 15:00:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 19:03:10
 * @Description: plugin protoc_gen_tool
 */
package service

import (
	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
	"github.com/yangjerry110/tool/conf"
)

type Plugin struct {
	Plugin *protogen.Plugin
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

	// set config
	if err := conf.CreateConf(&config.ProtocGenTool{}).SetConfig(); err != nil {
		return err
	}

	// set extendImportPath
	if err := conf.CreateConf(&config.ExtendPath{Path: config.ProtocGenToolConf.ExtendPath}).SetConfig(); err != nil {
		return err
	}

	// Get Params by flags
	// Set Params to config

	// judge plugin
	// if == nil; return err
	if p.Plugin == nil {
		return errors.ErrServiceNoExtension
	}

	// judge plugin files
	// if files's num < 1 return err
	if len(p.Plugin.Files) < 1 {
		return errors.ErrConfigPluginNoPlugin
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
