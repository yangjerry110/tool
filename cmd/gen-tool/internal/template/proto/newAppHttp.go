package proto

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type NewAppHttpProto struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:49:41
 * @return {*}
 */
func (n *NewAppHttpProto) New() error {
	filePath := fmt.Sprintf("%s/protobuf", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "http.proto", n.getTemplate(), nil, "proto")
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:49:16
 * @return {*}
 */
func (n *NewAppHttpProto) getTemplate() string {
	return `syntax = "proto3";

package api;

option go_package = "github.com/yangjerry110/tool/pkg/protocol/api";

import "google/protobuf/descriptor.proto";

extend google.protobuf.MethodOptions {
	HttpRule http = 20200101;
}

extend google.protobuf.FieldOptions {
	string in = 20200201;   // path/header/query/body
}

// HTTP API定义
message HttpRule {
	string selector = 1;    // 和Google保持一致
	oneof pattern {
		string get = 2;
		string put = 3;
		string post = 4;
		string delete = 5;
		string patch = 6;
	}

	string body = 7;
	string description = 8;		// API描述 给API文档使用
	string response_body = 12;  // 暂时没有用到
	repeated HttpRule additional_bindings = 11; // 和Google保持一致
}

// HTTP
message Http {
	repeated HttpRule rules = 1;
}
	`
}
