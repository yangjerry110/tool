/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 15:11:58
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 18:37:36
 * @Description: protocHttpRules
 */
package config

import (
	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
	httpProto "github.com/yangjerry110/tool/pkg/protocol/api"
)

type ProtocHttpRule struct {
	HttpProtoMethod *protogen.Method
	HttpProtoRule   *httpProto.HttpRule
	InputFields     []*protogen.Field
	Method          string
	Url             string
	Description     string
	FuncName        string
	InputName       string
	OutputName      string
	InputMessage    *protogen.Message
	OutputMessage   *protogen.Message
}

/**
 * @description: ProtocHttpRules
 * @author: Jerry.Yang
 * @date: 2023-12-12 15:13:59
 * @return {*}
 */
var ProtocHttpRules = []*ProtocHttpRule{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-12 15:29:17
 * @return {*}
 */
func (p *ProtocHttpRule) SetConfig() error {

	// Judge p.HttpProtoMethod
	// If == nil ; return err
	if p.HttpProtoMethod == nil {
		return errors.ErrConfigProtocHttpRuleNoMethods
	}

	// Judge p.HttpProtoRule
	// If == nil; return err
	if p.HttpProtoRule == nil {
		return errors.ErrConfigProtocHttpRuleNoHttpRules
	}

	// Set Description
	// Set FuncName
	// Set InputName
	// Set OutputName
	p.Description = p.HttpProtoRule.GetDescription()
	p.FuncName = p.HttpProtoMethod.GoName
	p.InputFields = p.HttpProtoMethod.Input.Fields
	p.InputName = string(p.HttpProtoMethod.Desc.Input().Name())
	p.OutputName = string(p.HttpProtoMethod.Desc.Output().Name())
	// set inputMessage
	// set outputMessage
	p.InputMessage = p.HttpProtoMethod.Input
	p.OutputMessage = p.HttpProtoMethod.Output

	// Next, we can use GetXxx to get the internal file we set in its Message
	// get what request type is defined, is it get, post, put, delete, path
	httpRulePattern := p.HttpProtoRule.GetPattern()

	// Judge httpRulePattern type
	// switch
	switch httpRulePattern.(type) {
	case *httpProto.HttpRule_Get:
		p.Method = "GET"
		p.Url = p.HttpProtoRule.GetGet()
	case *httpProto.HttpRule_Post:
		p.Method = "POST"
		p.Url = p.HttpProtoRule.GetPost()
	case *httpProto.HttpRule_Put:
		p.Method = "PUT"
		p.Url = p.HttpProtoRule.GetPut()
	case *httpProto.HttpRule_Patch:
		p.Method = "PATCH"
		p.Url = p.HttpProtoRule.GetPatch()
	case *httpProto.HttpRule_Delete:
		p.Method = "DELETE"
		p.Url = p.HttpProtoRule.GetDelete()
	default:
		return errors.ErrConfigProtocHttpRuleNoHttpRuleType
	}

	// set p to ProtocHttpRules
	ProtocHttpRules = append(ProtocHttpRules, p)
	return nil
}
