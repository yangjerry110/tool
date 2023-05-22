/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 16:18:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 16:25:47
 * @Description: http protobuf
 */
package protobuf

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type HttpProtobuf interface {
	SaveTemplate(path string) error
	GetTemplate() string
}

type Http struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-05-18 16:21:54
 * @return {*}
 */
func (h *Http) SaveTemplate(path string) error {
	return templates.CreateCommonTemplate().SaveTemplate(path, "http.proto", h.GetTemplate(), nil, "proto")
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-18 16:20:35
 * @return {*}
 */
func (h *Http) GetTemplate() string {
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
