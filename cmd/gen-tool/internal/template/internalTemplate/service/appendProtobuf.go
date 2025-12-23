/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-14 19:04:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-12-01 14:52:35
 * @Description: append protobuf
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type AppendProtobuf struct {
	ProjectPath      string
	ServiceName      string
	FirstServiceName string
	ServiceNameUp    string
	ServiceFuncUp    string
	InputReqName     string
	OutputRespName   string
	Time             string
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-14 19:07:57
 * @return {*}
 */
func (a *AppendProtobuf) New() error {
	if err := template.AppendTemplate(fmt.Sprintf("%s/internal/service/%sService.go", a.ProjectPath, a.ServiceName), a.getTemplate(), a); err != nil {
		return err
	}
	return nil
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-14 19:08:09
 * @return {*}
 */
func (a *AppendProtobuf) getTemplate() string {
	return `
	/**
	* @description: {{.ServiceFuncUp}}
	* @param {context.Context} ctx
	* @param {*protobuf.{{.InputReqName}}} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func ({{.FirstServiceName}} *{{.ServiceNameUp}}) {{.ServiceFuncUp}}(ctx *gin.Context, inputVo *protobuf.{{.InputReqName}}) (*protobuf.{{.OutputRespName}}, error) {
   
	   /**
		* @step
		* @result
		**/
	   resp := &protobuf.{{.OutputRespName}}{}
	   return resp, nil
   }`
}
