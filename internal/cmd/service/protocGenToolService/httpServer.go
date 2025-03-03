/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-28 18:26:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 14:12:02
 * @Description: httpServer
 */
package protocgentoolservice

import (
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
	"github.com/yangjerry110/tool/internal/cmd/template/vo/protobuf"
)

type HttpServer struct{}

/**
 * @description: Generate
 * @author: Jerry.Yang
 * @date: 2025-03-03 14:12:33
 * @return {*}
 */
func (h *HttpServer) Generate() error {

	// Assemble the parameters related to the template service
	templateHttpProtobuf := protobuf.HttpProtobuf{}
	// The serviceName here is protobufFileName
	// so first is fist protobufFileName
	templateHttpProtobuf.FirstServiceName = config.ProtobufFileConf.FileName[:1]
	// Previously set
	templateHttpProtobuf.ProjectImportPath = config.ProjectImportPathConf.ImportPath
	// Previously set
	templateHttpProtobuf.ProjectPath = config.ProjectPathConf.Path
	// The serviceName here is protobufFileName
	templateHttpProtobuf.ServiceName = config.ProtobufFileConf.FileName
	// The serviceName here is protobufFileName
	// so up is protobufFileName up
	templateHttpProtobuf.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
	// set time
	templateHttpProtobuf.Time = template.GetFormatNowTime()

	// templateNewProtoServices
	templateHttpProtobufServices := []*protobuf.HttpProtobufService{}
	// set templateHttpProtobufServices by httpRules
	for _, protocHttpRule := range config.ProtocHttpRules {
		templateHttpProtobufService := &protobuf.HttpProtobufService{}
		templateHttpProtobufService.FirstServiceName = config.ProtobufFileConf.FileName[:1]
		templateHttpProtobufService.InputReqName = protocHttpRule.InputName
		templateHttpProtobufService.OutputRespName = protocHttpRule.OutputName
		templateHttpProtobufService.ServiceFuncUp = template.FirstUpper(protocHttpRule.FuncName)
		templateHttpProtobufService.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
		templateHttpProtobufService.Time = template.GetFormatNowTime()
		templateHttpProtobufServices = append(templateHttpProtobufServices, templateHttpProtobufService)
	}

	// set templateNewProtoServices
	templateHttpProtobuf.Services = templateHttpProtobufServices

	// Save Template
	if err := template.CreateTemplate(&templateHttpProtobuf).New(); err != nil {
		return err
	}
	return nil
}
