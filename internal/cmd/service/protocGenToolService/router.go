/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 15:46:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-12-06 16:42:39
 * @Description: router service
 */
package protocgentoolservice

import (
	"fmt"

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
	// Previously set
	templateNewProtobuf.ProjectImportPath = config.ProjectImportPathConf.ImportPath
	// Previously set
	templateNewProtobuf.ProjectPath = config.ProjectPathConf.Path
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
		newProtobufRouters = append(newProtobufRouters, newProtobufRouter)
		// Format SwaggerNotes
		newProtobufRouter.SwaggerNotes = r.formatSwaggerNotes(protocHttpRule)
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

/**
 * @description: formatSwaggerNotes
 * @param {*config.ProtocHttpRule} protocHttpRule
 * @author: Jerry.Yang
 * @date: 2023-12-13 16:37:04
 * @return {*}
 */
func (r *Router) formatSwaggerNotes(protocHttpRule *config.ProtocHttpRule) string {

	// swaggerNotes
	swaggerNotes := fmt.Sprintf("// %s %s", template.FirstUpper(protocHttpRule.FuncName), protocHttpRule.Description)
	swaggerNotes += "\r\n"
	swaggerNotes += fmt.Sprintf("// @ID %s", template.FirstUpper(protocHttpRule.FuncName))
	swaggerNotes += "\r\n"
	swaggerNotes += fmt.Sprintf("// @Summary %s", protocHttpRule.Description)
	swaggerNotes += "\r\n"
	swaggerNotes += fmt.Sprintf("// @Tags %s ", template.FirstUpper(config.ProtobufFileConf.FileName))
	swaggerNotes += "\r\n"

	// Judge Func'Method is Get Or Post
	// If method == Get, add query string
	if protocHttpRule.Method == "GET" {
		// Get Fields
		if len(protocHttpRule.InputFields) != 0 {
			for _, inputField := range protocHttpRule.InputFields {
				swaggerNotes += fmt.Sprintf("// @Param %s query %s false \"-\"", inputField.Desc.Name(), inputField.Desc.Kind())
				swaggerNotes += "\r\n"
			}
		}
	}

	// If method == Post, add input body
	if protocHttpRule.Method == "POST" {
		swaggerNotes += fmt.Sprintf("// @Param input body protobuf.%s false \" - \"", protocHttpRule.InputName)
		swaggerNotes += "\r\n"
	}

	// Add Output
	swaggerNotes += fmt.Sprintf("// @Success 200 {object} protobuf.%s", protocHttpRule.OutputName)
	swaggerNotes += "\r\n"

	// Add Router
	swaggerNotes += fmt.Sprintf("// @Router %s [%s]", protocHttpRule.Url, protocHttpRule.Method)
	return swaggerNotes
}
