/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-23 19:03:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 11:04:33
 * @Description: router
 */
package commands

import (
	"fmt"

	genToolCommands "github.com/yangjerry110/tool/cmd/gen-tool/commands"
	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
	"github.com/yangjerry110/tool/cmd/gen-tool/templates/router"
)

type RouterCommands interface {
	CreateRouter(routerName string, protobufRouters []*router.NewProtobufRouter) error
	CreateNewProtobufRouter(protobufRouters []*router.NewProtobufRouter) error
	SetRouterName(routerName string) error
	GetNewRouterProtobufs(fileName string, httpRules []*HttpRule) ([]*router.NewProtobufRouter, error)
}

type Router struct{}

/**
 * @description: RouterParams
 * @author: Jerry.Yang
 * @date: 2023-05-23 19:04:48
 * @return {*}
 */
var RouterParams = &Router{}

/**
 * @description: CreateRouter
 * @param {string} routerName
 * @param {[]*router.NewProtobufRouter} protobufRouters
 * @author: Jerry.Yang
 * @date: 2023-05-24 11:35:17
 * @return {*}
 */
func (r *Router) CreateRouter(routerName string, protobufRouters []*router.NewProtobufRouter) error {

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
	if err := r.SetRouterName(routerName); err != nil {
		return err
	}

	/**
	 * @step
	 * @执行new
	 **/
	if err := r.CreateNewProtobufRouter(protobufRouters); err != nil {
		return err
	}
	return nil
}

/**
 * @description: CreateNewProtobufRouter
 * @author: Jerry.Yang
 * @date: 2023-05-24 11:14:33
 * @return {*}
 */
func (r *Router) CreateNewProtobufRouter(protobufRouters []*router.NewProtobufRouter) error {
	genToolCommands.NewAppParams.AppRouterFileName = fmt.Sprintf("%sRouter.go", genToolCommands.InitParams.RouterName)
	if err := router.CreateNewRouter().SaveProtobufTemplate(fmt.Sprintf("%srouter", genToolCommands.InitParams.ProjectPath), genToolCommands.InitParams.ProjectImportPath, genToolCommands.InitParams.RouterName, protobufRouters); err != nil {
		return err
	}

	/**
	 * @step
	 * @追加到base,判断是否第一次创建
	 **/
	if CommandParams.IsFirstCreate {
		if err := genToolCommands.CreateNewRouterCommands().AppendBaseFuncRouter(); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @description: SetRouterName
 * @param {string} routerName
 * @author: Jerry.Yang
 * @date: 2023-05-24 11:07:27
 * @return {*}
 */
func (r *Router) SetRouterName(routerName string) error {
	genToolCommands.InitParams.RouterName = routerName
	return nil
}

/**
 * @description: GetNewRouterProtobufs
 * @param {string} fileName
 * @param {[]*HttpRule} httpRules
 * @author: Jerry.Yang
 * @date: 2023-05-25 10:52:35
 * @return {*}
 */
func (r *Router) GetNewRouterProtobufs(fileName string, httpRules []*HttpRule) ([]*router.NewProtobufRouter, error) {

	/**
	 * @step
	 * @定义返回
	 **/
	newRouterProtobufs := []*router.NewProtobufRouter{}

	/**
	 * @step
	 * @渲染newRouterProtobuf
	 **/
	for _, httpRule := range httpRules {
		newRouterProtobufs = append(newRouterProtobufs, &router.NewProtobufRouter{
			Description:     httpRule.Description,
			RouterMethod:    httpRule.Method,
			RouterPath:      httpRule.Url,
			RouterNameUp:    templates.CreateCommonTemplate().FirstUpper(fileName),
			RouterFunc:      httpRule.FuncName,
			RouterFuncUp:    templates.CreateCommonTemplate().FirstUpper(httpRule.FuncName),
			InputReqName:    httpRule.InputName,
			FirstRouterName: fileName[:1],
			RouterName:      fileName,
			Time:            templates.CreateCommonTemplate().GetFormatNowTime(),
		})
	}
	return newRouterProtobufs, nil
}
