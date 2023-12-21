/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-14 16:05:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-20 18:25:34
 * @Description:
 */
package protocgentoolservice

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	newfilesetservice "github.com/yangjerry110/tool/internal/cmd/service/newFileSetService"
	"github.com/yangjerry110/tool/internal/cmd/template"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/service"
)

type Service struct{}

/**
 * @description: Generate
 * @author: Jerry.Yang
 * @date: 2023-12-14 16:15:12
 * @return {*}
 */
func (s *Service) Generate() error {

	// Judge isFirstCreate isAppend
	// If IsFirstCreate and is not isAppend
	// Create New service
	if config.ProtocGenToolConf.IsFirstCreate && !config.ProtocGenToolConf.IsAppend {
		if err := s.newService(); err != nil {
			return err
		}
	}

	// If is not IsFirstCreate and isAppend
	// Create Append service
	if !config.ProtocGenToolConf.IsFirstCreate && config.ProtocGenToolConf.IsAppend {
		if err := s.appendService(); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @description: newService
 * @author: Jerry.Yang
 * @date: 2023-12-14 16:25:28
 * @return {*}
 */
func (s *Service) newService() error {

	// Assemble the parameters related to the template service
	templateNewProtobuf := service.NewProtobuf{}
	// The serviceName here is protobufFileName
	// so first is fist protobufFileName
	templateNewProtobuf.FirstServiceName = config.ProtobufFileConf.FileName[:1]
	// Previously set
	templateNewProtobuf.ProjectImportPath = config.ProjectImportPathConf.ImportPath
	// Previously set
	templateNewProtobuf.ProjectPath = config.ProjectPathConf.Path
	// The serviceName here is protobufFileName
	templateNewProtobuf.ServiceName = config.ProtobufFileConf.FileName
	// The serviceName here is protobufFileName
	// so up is protobufFileName up
	templateNewProtobuf.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
	// set time
	templateNewProtobuf.Time = template.GetFormatNowTime()

	// templateNewProtoServices
	templateNewProtobufServices := []*service.NewProtobufService{}
	// set templateNewProtobufServices by httpRules
	for _, protocHttpRule := range config.ProtocHttpRules {
		templateNewProtobufService := &service.NewProtobufService{}
		templateNewProtobufService.FirstServiceName = config.ProtobufFileConf.FileName[:1]
		templateNewProtobufService.InputReqName = protocHttpRule.InputName
		templateNewProtobufService.OutputRespName = protocHttpRule.OutputName
		templateNewProtobufService.ServiceFuncUp = template.FirstUpper(protocHttpRule.FuncName)
		templateNewProtobufService.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
		templateNewProtobufService.Time = template.GetFormatNowTime()
		templateNewProtobufServices = append(templateNewProtobufServices, templateNewProtobufService)

		// Append base
		templateAppendBase := &service.AppendBase{}
		templateAppendBase.ServiceName = protocHttpRule.FuncName
		templateAppendBase.ServiceNameUp = template.FirstUpper(protocHttpRule.FuncName)
		templateAppendBase.Time = template.GetFormatNowTime()
		if err := template.CreateTemplate(templateAppendBase).New(); err != nil {
			return err
		}
	}

	// set templateNewProtoServices
	templateNewProtobuf.Services = templateNewProtobufServices

	// Save Template
	if err := template.CreateTemplate(&templateNewProtobuf).New(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: appendService
 * @author: Jerry.Yang
 * @date: 2023-12-18 16:46:25
 * @return {*}
 */
func (s *Service) appendService() error {

	// Define
	filePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)
	fileName := fmt.Sprintf("%sService.go", config.ProtobufFileConf.FileName)
	newFileSetService := &newfilesetservice.Service{}

	// First, re-render the interface portion of the service based on the protocHttpRules we've obtained
	if err := newfilesetservice.CreateNewFileSetService(newFileSetService).NewFileSet(filePath, fileName); err != nil {
		fmt.Printf("newFileSet Err : %+v", err)
		fmt.Print("\r\n")
		return err
	}

	// Save the contents of the rendered file
	if err := template.SaveTemplate(filePath, fileName, newFileSetService.FileContent, nil); err != nil {
		fmt.Printf("SaveTemplate Err : %+v", err)
		fmt.Print("\r\n")
		return err
	}

	// If len AppendFuncs != 0
	// Append func to template
	if len(newFileSetService.AppendFuncs) != 0 {
		for _, appendProtocHttpRule := range newFileSetService.AppendFuncs {

			// Set AppendService
			templateAppendProtobufService := &service.AppendProtobuf{}
			templateAppendProtobufService.FirstServiceName = config.ProtobufFileConf.FileName[:1]
			templateAppendProtobufService.InputReqName = appendProtocHttpRule.InputName
			templateAppendProtobufService.OutputRespName = appendProtocHttpRule.OutputName
			templateAppendProtobufService.ServiceFuncUp = template.FirstUpper(appendProtocHttpRule.FuncName)
			templateAppendProtobufService.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
			templateAppendProtobufService.Time = template.GetFormatNowTime()

			// Append template
			if err := template.CreateTemplate(templateAppendProtobufService).New(); err != nil {
				return err
			}
		}
	}
	return nil
}
