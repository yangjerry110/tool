/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-24 11:40:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-26 11:25:40
 * @Description: service
 */
package commands

import (
	"fmt"
	"os"

	genToolCommands "github.com/yangjerry110/tool/cmd/gen-tool/commands"
	genToolService "github.com/yangjerry110/tool/cmd/gen-tool/service"
	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
	"github.com/yangjerry110/tool/cmd/gen-tool/templates/service"
	"google.golang.org/protobuf/compiler/protogen"
)

type ServiceCommands interface {
	GenService(service *protogen.Service) ([]*HttpRule, error)
	CreateService(serviceName string, httpRules []*HttpRule, newProtobufServices []*service.NewProtobufService) error
	CreateNewService(newProtobufServices []*service.NewProtobufService) error
	SetServiceName(serviceName string) error
	IsServiceFileExist() (bool, error)
	GetServiceProtobufs(fileName string, httpRules []*HttpRule) ([]*service.NewProtobufService, error)
	GetReviseInterfaceHttpRules(httpRules []*HttpRule) (map[string]*genToolService.HttpRule, error)
}

type Service struct{}

/**
 * @description: GenService
 * @param {*protogen.Service} service
 * @author: Jerry.Yang
 * @date: 2023-05-24 14:18:32
 * @return {*}
 */
func (s *Service) GenService(service *protogen.Service) ([]*HttpRule, error) {

	/**
	 * @step
	 * @返回数据
	 **/
	httpRules := []*HttpRule{}

	/**
	 * @step
	 * @service 内部有很多 rpc 关键字的方法
	 **/
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}

		/**
		 * @step
		 * @由于我们自定义的是就是MethodOptions，所以就来到了这里来进行判断
		 **/
		httpRule, err := CreateHttpCommands().GetHttpRule(method)
		if err != nil {
			return nil, err
		}

		/**
		 * @step
		 * @复制Name
		 **/
		httpRule.FuncName = method.GoName
		httpRule.InputName = string(method.Desc.Input().Name())
		httpRule.OutputName = string(method.Desc.Output().Name())
		httpRules = append(httpRules, httpRule)
	}
	return httpRules, nil
}

/**
 * @description: CreateService
 * @param {string} serviceName
 * @param {[]*HttpRule} httpRules
 * @param {[]*service.NewProtobufService} newProtobufServices
 * @author: Jerry.Yang
 * @date: 2023-05-25 10:58:49
 * @return {*}
 */
func (s *Service) CreateService(serviceName string, httpRules []*HttpRule, newProtobufServices []*service.NewProtobufService) error {

	/**
	 * @step
	 * @设置projectPath
	 **/
	if err := genToolCommands.CreateInitCommands().SetProjectPath(); err != nil {
		return err
	}

	/**
	 * @step
	 * @设置projectImportPath
	 **/
	if err := genToolCommands.CreateInitCommands().SetImportProjectPath(); err != nil {
		return err
	}

	/**
	 * @step
	 * @setRouteeName
	 **/
	if err := s.SetServiceName(serviceName); err != nil {
		return err
	}

	/**
	 * @step
	 * @判断serviceFile是否存在
	 **/
	isServiceFileExist, err := s.IsServiceFileExist()
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @假如存在，则判断是否是追加还是覆盖
	 **/
	if isServiceFileExist && CommandParams.IsAppend {
		if err := s.AppendNewSevice(serviceName, httpRules); err != nil {
			return err
		}
		return nil
	}

	/**
	 * @step
	 * @create new service
	 **/
	if err := s.CreateNewService(newProtobufServices); err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateNewService
 * @param {[]*service.NewProtobufService} newProtobufServices
 * @author: Jerry.Yang
 * @date: 2023-05-25 11:06:22
 * @return {*}
 */
func (s *Service) CreateNewService(newProtobufServices []*service.NewProtobufService) error {

	/**
	 * @step
	 * @saveProtobufTemplate
	 **/
	if err := service.CreateNewService().SaveProtobufTemplate(
		fmt.Sprintf("%s%s", genToolCommands.InitParams.ProjectPath, "service"),
		genToolCommands.InitParams.ProjectImportPath,
		genToolCommands.InitParams.ServiceName,
		fmt.Sprintf("%sService.go", genToolCommands.InitParams.ServiceName),
		newProtobufServices,
	); err != nil {
		return err
	}

	/**
	 * @step
	 * @判断是否是首次创建
	 **/
	if CommandParams.IsFirstCreate {
		if err := genToolCommands.CreateNewServiceCommands().AppendFuncBaseService(); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @description: AppendNewSevice
 * @param {string} serviceName
 * @param {[]*HttpRule} httpRules
 * @author: Jerry.Yang
 * @date: 2023-05-25 21:06:21
 * @return {*}
 */
func (s *Service) AppendNewSevice(serviceName string, httpRules []*HttpRule) error {

	/**
	 * @step
	 * @获取reviseInterfaceHttpRules
	 **/
	reviseInterfaceHttpRules, err := s.GetReviseInterfaceHttpRules(httpRules)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @拼接filePath
	 **/
	filePath := fmt.Sprintf("%s%s/%sService.go", genToolCommands.InitParams.ProjectPath, "service", serviceName)

	/**
	 * @step
	 * @获取渲染完interface的文件内容
	 **/
	serviceContent, err := genToolService.CreateReviseInterfaceService().GetReviseInterface(filePath, reviseInterfaceHttpRules)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @获取servicePath
	 **/
	servicePath := fmt.Sprintf("%s%s", genToolCommands.InitParams.ProjectPath, "service")

	/**
	 * @step
	 * @先保存当前渲染完interface的文件
	 **/
	if err := templates.CreateCommonTemplate().SaveTemplate(servicePath, fmt.Sprintf("%sService.go", serviceName), serviceContent, nil); err != nil {
		return err
	}

	/**
	 * @step
	 * @获取不存在的funcNames
	 **/
	noExistFuncNames, err := genToolService.CreateParseService().GetNoExistFuncByParseFile(filePath, reviseInterfaceHttpRules)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @判断不存在funcNames是否为空
	 **/
	if len(noExistFuncNames) == 0 {
		return nil
	}

	/**
	 * @step
	 * @appendFuncName
	 **/
	for _, noExistFuncName := range noExistFuncNames {

		/**
		 * @step
		 * @获取当前funcName的httpRule
		 **/
		noExistFuncHttpRule, isOk := reviseInterfaceHttpRules[noExistFuncName]
		if !isOk {
			continue
		}

		/**
		 * @step
		 * @append protobuf
		 **/
		if err := service.CreateNewService().SaveAppendProtobufTemplate(
			servicePath,
			serviceName,
			noExistFuncName,
			&service.NewProtobufService{
				FirstServiceName: noExistFuncName[:1],
				ServiceNameUp:    templates.CreateCommonTemplate().FirstUpper(serviceName),
				ServiceFuncUp:    templates.CreateCommonTemplate().FirstUpper(noExistFuncName),
				InputReqName:     noExistFuncHttpRule.InputName,
				OutputRespName:   noExistFuncHttpRule.OutputName,
				Time:             templates.CreateCommonTemplate().GetFormatNowTime(),
			},
		); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @description: SetServiceName
 * @param {string} serviceName
 * @author: Jerry.Yang
 * @date: 2023-05-25 10:59:44
 * @return {*}
 */
func (s *Service) SetServiceName(serviceName string) error {
	genToolCommands.InitParams.ServiceName = serviceName
	return nil
}

/**
 * @description: IsServiceFileExist
 * @author: Jerry.Yang
 * @date: 2023-05-25 11:31:26
 * @return {*}
 */
func (s *Service) IsServiceFileExist() (bool, error) {

	/**
	 * @step
	 * @判断serviceFile是否存在
	 **/
	serviceFilePath := fmt.Sprintf("%s%s/%sService.go", genToolCommands.InitParams.ProjectPath, "service", genToolCommands.InitParams.ServiceName)

	/**
	 * @step
	 * @判断是否存在
	 **/
	_, err := os.Stat(serviceFilePath)
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 * @description: GetServiceProtobufs
 * @param {string} fileName
 * @param {[]*HttpRule} httpRules
 * @author: Jerry.Yang
 * @date: 2023-05-25 10:57:21
 * @return {*}
 */
func (s *Service) GetServiceProtobufs(fileName string, httpRules []*HttpRule) ([]*service.NewProtobufService, error) {

	/**
	 * @step
	 * @定义返回
	 **/
	newProtobufServices := []*service.NewProtobufService{}

	/**
	 * @step
	 * @渲染newRouterProtobuf
	 **/
	for _, httpRule := range httpRules {
		newProtobufServices = append(newProtobufServices, &service.NewProtobufService{
			FirstServiceName: fileName[:1],
			ServiceNameUp:    templates.CreateCommonTemplate().FirstUpper(fileName),
			ServiceFuncUp:    templates.CreateCommonTemplate().FirstUpper(httpRule.FuncName),
			InputReqName:     httpRule.InputName,
			OutputRespName:   httpRule.OutputName,
			Time:             templates.CreateCommonTemplate().GetFormatNowTime(),
		})
	}
	return newProtobufServices, nil
}

/**
 * @description: GetReviseInterfaceHttpRules
 * @param {[]*HttpRule} httpRules
 * @author: Jerry.Yang
 * @date: 2023-05-25 20:57:51
 * @return {*}
 */
func (s *Service) GetReviseInterfaceHttpRules(httpRules []*HttpRule) (map[string]*genToolService.HttpRule, error) {

	/**
	 * @step
	 * @定义返回
	 **/
	genToolServiceHttpRules := map[string]*genToolService.HttpRule{}

	/**
	 * @step
	 * @渲染gentoolServiceHttpRules
	 **/
	for _, httpRule := range httpRules {
		genToolServiceHttpRules[httpRule.FuncName] = &genToolService.HttpRule{
			Description: httpRule.Description,
			FuncName:    httpRule.FuncName,
			Method:      httpRule.Method,
			Url:         httpRule.Url,
			InputName:   httpRule.InputName,
			OutputName:  httpRule.OutputName,
		}
	}
	return genToolServiceHttpRules, nil
}
