/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-26 15:36:05
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-27 11:33:49
 * @Description:
 */
package protocgentoolservice

import (
	"fmt"
	"strings"

	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type Comment struct {
	ProtocHttpRule *config.ProtocHttpRule
	Comment        string
}

/**
 * @description: Generate
 * @author: Jerry.Yang
 * @date: 2025-02-26 15:36:54
 * @return {*}
 */
func (c *Comment) Generate() error {

	// swaggerNotes
	swaggerNotes := fmt.Sprintf("// %s %s", template.FirstUpper(c.ProtocHttpRule.FuncName), c.ProtocHttpRule.Description)
	swaggerNotes += "\r\n"
	swaggerNotes += fmt.Sprintf("// @ID %s", template.FirstUpper(c.ProtocHttpRule.FuncName))
	swaggerNotes += "\r\n"
	swaggerNotes += fmt.Sprintf("// @Summary %s", c.ProtocHttpRule.Description)
	swaggerNotes += "\r\n"
	swaggerNotes += fmt.Sprintf("// @Tags %s ", template.FirstUpper(config.ProtobufFileConf.FileName))
	swaggerNotes += "\r\n"

	// get req swagger
	reqSwagger := c.getMessageSwaggerReq(c.ProtocHttpRule.InputMessage)
	swaggerNotes += reqSwagger

	// Add Output
	swaggerNotes += fmt.Sprintf("// @Success 200 {object} protobuf.%s", c.ProtocHttpRule.OutputName)
	swaggerNotes += "\r\n"

	// Add Router
	swaggerNotes += fmt.Sprintf("// @Router %s [%s]", c.ProtocHttpRule.Url, c.ProtocHttpRule.Method)

	// set Comment
	c.Comment = swaggerNotes
	return nil
}

/**
 * @description: getMessageSwaggerReq
 * @param {*protogen.Message} message
 * @author: Jerry.Yang
 * @date: 2025-02-26 18:58:56
 * @return {*}
 */
func (c *Comment) getMessageSwaggerReq(message *protogen.Message) string {

	// define result
	swaggerNotes := ""

	// 假如是post
	if c.ProtocHttpRule.Method == "POST" {
		// get req comment
		messageComment := c.getComment(&message.Comments)
		swaggerNotes += fmt.Sprintf("// @Param input body protobuf.%s false \"%s\"", c.ProtocHttpRule.InputName, messageComment)
		swaggerNotes += "\r\n"
		return swaggerNotes
	}

	// for messages
	for _, message := range message.Messages {
		c.getMessageSwaggerReq(message)
	}

	// for message.fields
	for _, field := range message.Fields {

		// get fieldComment
		fieldComment := c.getComment(&field.Comments)

		// judge method
		if c.ProtocHttpRule.Method == "GET" {
			swaggerNotes += fmt.Sprintf("// @Param %s query %s false \"%s\"", field.Desc.Name(), field.Desc.Kind(), fieldComment)
			swaggerNotes += "\r\n"
			continue
		}
	}
	return swaggerNotes
}

/**
 * @description: getFieldComment
 * @param {*protogen.Field} field
 * @author: Jerry.Yang
 * @date: 2025-02-26 15:47:02
 * @return {*}
 */
func (c *Comment) getComment(protogenComment *protogen.CommentSet) string {

	// define comment
	comment := ""

	// get field comment
	fieldLeadingComment := protogenComment.Leading.String()
	fieldLeadingComment = strings.ReplaceAll(fieldLeadingComment, "//", "")
	comment = strings.TrimSpace(fieldLeadingComment)

	// judge fieldComment
	if comment == "" {
		fieldTrailingComment := protogenComment.Trailing.String()
		fieldTrailingComment = strings.ReplaceAll(fieldTrailingComment, "//", "")
		comment = strings.TrimSpace(fieldTrailingComment)
	}

	// judge comment
	if comment != "" {
		return comment
	}
	return "-"
}
