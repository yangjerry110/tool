/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-14 16:05:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-24 17:03:57
 * @Description:
 */
package protocgentoolservice

import (
	"fmt"
	"os"

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

	if config.ProtocGenToolConf.IsExtend {
		config.ProjectPathConf.Path = config.ExtendPathConf.Path
	}

	// is exist service path
	if err := s.isExistServicePath(); err != nil {
		return err
	}

	// is exist service file
	if err := s.isExistServiceFile(); err != nil {
		return err
	}

	// if isAppend or if isNew
	// new interface service
	if err := s.newInterfaceService(); err != nil {
		return err
	}

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

// newInterfaceService
//
// new interface service的相关的
// Author Jerry.Yang
// Date 2024-03-05 10:58:46
func (s *Service) newInterfaceService() error {

	// Assemble the parameters related to the template service
	templateNewInterfaceProtobuf := service.NewInterfaceProtobuf{}
	// The serviceName here is protobufFileName
	// so first is fist protobufFileName
	templateNewInterfaceProtobuf.FirstServiceName = config.ProtobufFileConf.FileName[:1]
	// Previously set
	templateNewInterfaceProtobuf.ProjectImportPath = config.ProjectImportPathConf.ImportPath
	// Previously set
	templateNewInterfaceProtobuf.ProjectPath = config.ProjectPathConf.Path
	// The serviceName here is protobufFileName
	templateNewInterfaceProtobuf.ServiceName = config.ProtobufFileConf.FileName
	// The serviceName here is protobufFileName
	// so up is protobufFileName up
	templateNewInterfaceProtobuf.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
	// set time
	templateNewInterfaceProtobuf.Time = template.GetFormatNowTime()

	// templateNewProtoServices
	templateNewInterfaceProtobufServices := []*service.NewInterfaceProtobufService{}
	// set templateNewInterfaceProtobufServices by httpRules
	for _, protocHttpRule := range config.ProtocHttpRules {
		templateNewInterfaceProtobufService := &service.NewInterfaceProtobufService{}
		templateNewInterfaceProtobufService.FirstServiceName = config.ProtobufFileConf.FileName[:1]
		templateNewInterfaceProtobufService.InputReqName = protocHttpRule.InputName
		templateNewInterfaceProtobufService.OutputRespName = protocHttpRule.OutputName
		templateNewInterfaceProtobufService.ServiceFuncUp = template.FirstUpper(protocHttpRule.FuncName)
		templateNewInterfaceProtobufService.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
		templateNewInterfaceProtobufService.Time = template.GetFormatNowTime()
		templateNewInterfaceProtobufServices = append(templateNewInterfaceProtobufServices, templateNewInterfaceProtobufService)
	}

	// set templateNewProtoServices
	templateNewInterfaceProtobuf.Services = templateNewInterfaceProtobufServices

	// Save Template
	if err := template.CreateTemplate(&templateNewInterfaceProtobuf).New(); err != nil {
		return err
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

	// Previously set
	// if isExtend == false, set projectPath = config.ProjectPathConf.Path
	// if isExtend == true, set projectPath = config.ProjectImportPathConf.Path
	templateNewProtobuf.ProjectPath = config.ProjectPathConf.Path

	// The serviceName here is protobufFileName
	// so first is fist protobufFileName
	templateNewProtobuf.FirstServiceName = config.ProtobufFileConf.FileName[:1]
	// Previously set
	templateNewProtobuf.ProjectImportPath = config.ProjectImportPathConf.ImportPath
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
	}

	// set templateNewProtoServices
	templateNewProtobuf.Services = templateNewProtobufServices

	// Save Template
	if err := template.CreateTemplate(&templateNewProtobuf).New(); err != nil {
		return err
	}

	// Append base
	templateAppendBaseService := &service.AppendBaseService{}
	templateAppendBaseService.ServiceName = config.ProtobufFileConf.FileName
	templateAppendBaseService.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
	templateAppendBaseService.Time = template.GetFormatNowTime()
	if err := template.CreateTemplate(templateAppendBaseService).New(); err != nil {
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

	// // Save the contents of the rendered file
	// if err := template.SaveTemplate(filePath, fileName, newFileSetService.FileContent, nil); err != nil {
	// 	fmt.Printf("SaveTemplate Err : %+v", err)
	// 	fmt.Print("\r\n")
	// 	return err
	// }

	// If len AppendFuncs != 0
	// Append func to template
	if len(newFileSetService.AppendFuncs) != 0 {
		for _, appendProtocHttpRule := range newFileSetService.AppendFuncs {

			// Set AppendService
			templateAppendProtobufService := &service.AppendProtobuf{}
			templateAppendProtobufService.ProjectPath = config.ProjectPathConf.Path
			templateAppendProtobufService.ServiceName = config.ProtobufFileConf.FileName
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

// isExistServicePath
//
// is exist service path
// Date 2024-03-04 17:22:25
// Author Jerry.Yang
func (s *Service) isExistServicePath() error {

	// interfaceService path
	interfaceServicePath := fmt.Sprintf("%s/internal/service/interfaceService", config.ProjectPathConf.Path)
	// service path
	servicePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)

	// If exist interfaceService path
	_, err := os.Stat(interfaceServicePath)
	if err != nil {
		// if not exist
		// mkdir all interfaceServicePath 077
		err = os.MkdirAll(interfaceServicePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// if exist service path
	_, err = os.Stat(servicePath)
	if err != nil {
		// if not exist
		// mkdir all servicePath 077
		err = os.MkdirAll(servicePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// isExistServiceFile
//
// is exist service file
// Author Jerry.Yang
// Date 2024-03-04 17:34:40
func (s *Service) isExistServiceFile() error {

	// servicePath
	servicePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)

	// baseService file
	baseServiceFileName := "baseService.go"
	baseServiceFilePath := fmt.Sprintf("%s/%s", servicePath, baseServiceFileName)

	// fileName
	fileName := fmt.Sprintf("%sService.go", config.ProtobufFileConf.FileName)
	filePath := fmt.Sprintf("%s/%s", servicePath, fileName)

	// if exist baseServiceFilePath
	_, err := os.Stat(baseServiceFilePath)
	if err != nil {
		// set newBaseService
		templateNewBaseService := &service.NewBaseService{}
		templateNewBaseService.ProjectImportPath = config.ProjectImportPathConf.ImportPath
		templateNewBaseService.Time = template.GetFormatNowTime()
		// newBaseService
		if err := template.CreateTemplate(templateNewBaseService).New(); err != nil {
			return err
		}
	}

	// if exist filePath
	_, err = os.Stat(filePath)
	if err != nil {
		// set newService
		templateNewService := &service.NewService{}
		templateNewService.ProjectPath = config.ProjectPathConf.Path
		templateNewService.ServiceName = config.ProtobufFileConf.FileName
		templateNewService.ServiceNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
		templateNewService.Time = template.GetFormatNowTime()
		// newService
		if err := template.CreateTemplate(templateNewService).New(); err != nil {
			return err
		}
	}
	return nil
}
