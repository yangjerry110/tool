/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 15:46:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-27 11:33:55
 * @Description: router service
 */
package protocgentoolservice

import (
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
	"github.com/yangjerry110/tool/internal/cmd/template/router"
)

type Router struct{}

/**
 * @description: Generate
 * @author: Jerry.Yang
 * @date: 2023-12-12 17:37:40
 * @return {*}
 */
func (r *Router) Generate() error {
	// Assemble the parameters related to the template router
	templateNewProtobuf := &router.NewProtobuf{}
	// The routerName here is protobufFileName
	// so first is fist protobufFileName
	templateNewProtobuf.FirstRouterName = config.ProtobufFileConf.FileName[:1]
	// set ExtendImportPath
	// is isExtend == false, set config.ProjectImportPathConf.ImportPath
	// is isExtend == true, set config.ExtendImportPathConf.Path
	templateNewProtobuf.ExtendImportPath = config.ProjectImportPathConf.ImportPath
	if config.ProtocGenToolConf.IsExtend {
		templateNewProtobuf.ExtendImportPath = config.ExtendPathConf.ImportPath
	}

	// Previously set
	// if isExtend == false, set projectPath = config.ProjectPathConf.Path
	// if isExtend == true, set projectPath = config.ProjectImportPathConf.Path
	templateNewProtobuf.ProjectPath = config.ProjectPathConf.Path
	if config.ProtocGenToolConf.IsExtend {
		templateNewProtobuf.ProjectPath = config.ExtendPathConf.Path
	}

	// Previously set
	templateNewProtobuf.ProjectImportPath = config.ProjectImportPathConf.ImportPath
	// The routerName here is protobufFileName
	templateNewProtobuf.RouterName = config.ProtobufFileConf.FileName
	// The routerName here is protobufFileName
	// so up is protobufFileName up
	templateNewProtobuf.RouterNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
	// set time
	templateNewProtobuf.Time = template.GetFormatNowTime()

	// set newProtobufServices
	newProtobufServices := []*router.NewProtobufService{}
	// set newProtobufServices by ProtocServiceConfs
	for _, protocServiceConf := range config.ProtocServiceConfs {
		newProtobufService := &router.NewProtobufService{}
		newProtobufService.ServiceName = protocServiceConf.ProtocService.GoName
		newProtobufService.RouterNameUp = templateNewProtobuf.RouterNameUp
		newProtobufServices = append(newProtobufServices, newProtobufService)
	}

	// set newProtobufRouters
	newProtobufRouters := []*router.NewProtobufRouter{}
	// set newProtobufRouters by httpRules
	for _, protocHttpRule := range config.ProtocHttpRules {
		newProtobufRouter := &router.NewProtobufRouter{}
		newProtobufRouter.Description = protocHttpRule.Description
		newProtobufRouter.FirstRouterName = config.ProtobufFileConf.FileName[:1]
		newProtobufRouter.InputReqName = protocHttpRule.InputName
		newProtobufRouter.OutputRespName = protocHttpRule.OutputName
		newProtobufRouter.RouterFunc = protocHttpRule.FuncName
		newProtobufRouter.RouterFuncUp = template.FirstUpper(protocHttpRule.FuncName)
		newProtobufRouter.RouterMethod = protocHttpRule.Method
		newProtobufRouter.RouterName = config.ProtobufFileConf.FileName
		newProtobufRouter.RouterNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
		newProtobufRouter.RouterPath = protocHttpRule.Url
		newProtobufRouter.Time = template.GetFormatNowTime()
		// get comment
		commonProtoGenToolService := &Comment{ProtocHttpRule: protocHttpRule}
		if err := commonProtoGenToolService.Generate(); err != nil {
			return err
		}
		newProtobufRouters = append(newProtobufRouters, newProtobufRouter)
		// Format SwaggerNotes
		newProtobufRouter.SwaggerNotes = commonProtoGenToolService.Comment
	}

	// set NewProtobufServices
	templateNewProtobuf.Services = newProtobufServices
	// set newProtobufRouters
	templateNewProtobuf.Routers = newProtobufRouters

	// Create template by newProtobufRouter
	if err := template.CreateTemplate(templateNewProtobuf).New(); err != nil {
		return err
	}

	// // Judge is fistCreate
	// // Judge is append
	// // append func to base
	// if config.ProtocGenToolConf.IsFirstCreate && !config.ProtocGenToolConf.IsAppend {
	// 	appendBaseRouter := router.AppendBase{}
	// 	appendBaseRouter.RouterName = config.ProtobufFileConf.FileName
	// 	appendBaseRouter.RouterNameUp = template.FirstUpper(config.ProtobufFileConf.FileName)
	// 	appendBaseRouter.Time = template.GetFormatNowTime()
	// 	if err := template.CreateTemplate(&appendBaseRouter).New(); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}
