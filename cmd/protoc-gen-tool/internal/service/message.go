/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-25 10:34:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 18:59:37
 * @Description: message
 */
package service

import (
	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/template"
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/template/vo/protobuf"
)

type Message struct {
	Messages []*protogen.Message
}

/**
 * @description: message Generate
 * @author: Jerry.Yang
 * @date: 2025-02-26 14:36:42
 * @return {*}
 */
func (m *Message) Generate() error {

	// set protobufMessage
	templateProtobufMessage := &protobuf.ProtobufMessage{}
	templateProtobufMessage.ProjectPath = config.ProjectPathConf.Path
	templateProtobufMessage.FileName = config.ProtobufFileConf.FileName

	// define protobufMessageFields
	templateProtobufMessageFields := []*protobuf.ProtobufMessageField{}

	// set ProtobufMessageField
	for _, message := range m.Messages {
		// 为每个 optional 字段生成 Has<Field> 方法并写入文件
		for _, field := range message.Fields {
			if field.Desc.HasOptionalKeyword() {
				templateProtobufMessageField := &protobuf.ProtobufMessageField{}
				templateProtobufMessageField.MessageName = message.GoIdent.GoName
				templateProtobufMessageField.FieldName = template.FirstUpper(field.GoName)
				templateProtobufMessageField.FieldGap = ""
				templateProtobufMessageFields = append(templateProtobufMessageFields, templateProtobufMessageField)
			}
		}
	}

	// set templateProtobufMessageFields
	templateProtobufMessage.Fields = templateProtobufMessageFields

	// Save Template
	if err := template.CreateTemplate(templateProtobufMessage).New(); err != nil {
		return err
	}
	return nil
}
