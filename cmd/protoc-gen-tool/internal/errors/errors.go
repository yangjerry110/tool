/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-12 17:35:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 19:03:45
 * @Description:
 */
package errors

import "github.com/yangjerry110/tool/toolerrors"

var (
	ErrConfigTemplateNoName                = toolerrors.New("protoc-gen-tool Err : config template no name")
	ErrConfigExtendPathNoImportProjectPath = toolerrors.New("protoc-gen-tool Err : config extendPath no import project path")
	ErrConfigPluginNoPlugin                = toolerrors.New("protoc-gen-tool Err : config plugin no plugin")
	ErrConfigProjectFileNoFiles            = toolerrors.New("protoc-gen-tool Err : projectFile no file")
	ErrConfigProjectImportPathNoPkgs       = toolerrors.New("protoc-gen-tool Err : config project import path no pkgs")
	ErrConfigProtocHttpRuleNoMethods       = toolerrors.New("protoc-gen-tool Err : config protoHttpRule no methods")
	ErrConfigProtocHttpRuleNoHttpRules     = toolerrors.New("protoc-gen-tool Err : config protoHttpRule no httpRules")
	ErrConfigProtocHttpRuleNoHttpRuleType  = toolerrors.New("proto-gen-tool Err : config protoHttpRule no httpRuleType")
	ErrConfigProtocServiceNoServices       = toolerrors.New("proto-gen-tool Err : config protocService no services")

	ErrTemplateNoPath = toolerrors.New("protoc-gen-tool Err : template no path")

	ErrServiceNoPlugin    = toolerrors.New("protoc-gen-tool Err : service no plugin")
	ErrServiceNoFile      = toolerrors.New("protoc-gen-tool Err : service no file")
	ErrServiceNoMethod    = toolerrors.New("protoc-gen-tool Err : service no method")
	ErrServiceNoMethods   = toolerrors.New("protoc-gen-tool Err : service no methods")
	ErrServiceNoMessages  = toolerrors.New("protoc-gen-tool Err : service no messages")
	ErrServiceNoNoOptions = toolerrors.New("protoc-gen-tool Err : service no options")
	ErrServiceNoExtension = toolerrors.New("protoc-gen-tool Err : service no extension")
)
