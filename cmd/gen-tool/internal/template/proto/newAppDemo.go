/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:43:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-26 16:20:23
 * @Description:
 */
package proto

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type NewAppDemoProto struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:45:59
 * @return {*}
 */
func (n *NewAppDemoProto) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time              string
		ImportProjectPath string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()
	data.ImportProjectPath = config.ProjectImportPathConf.ImportPath

	filePath := fmt.Sprintf("%s/protobuf", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "demo.proto", n.getTemplate(), data, "proto")
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:45:51
 * @return {*}
 */
func (n *NewAppDemoProto) getTemplate() string {
	return `// demo.proto
syntax = "proto3";
package demo;
option go_package = "{{.ImportProjectPath}}/vo/protobuf";
import "http.proto";

// 服务相关的注释
service DemoApi {

	// swagger:route POST /api/demo demo AddDemo
	// 增加Demo的详细描述
	rpc AddDemo (AddDemoReq) returns (Empty) {
		option (api.http) = {
			post: "/api/demo"
			body: "*"
			description: "增加Demo的详细描述"
		};
	};

	// swagger:route DELETE /api/demo/{id} demo DeleteDemo
	// 删除Demo的详细描述
	rpc DeleteDemo (DeleteDemoReq) returns (Empty) {
		option (api.http) = {
			delete: "/api/demo/:id"
			description: "删除Demo的详细描述"
		};
	};

	// swagger:route PATCH /api/demo/{id} demo UpdateDemo
	// 更新Demo的详细描述
	rpc UpdateDemo (UpdateDemoReq) returns (Empty) {
		option (api.http) = {
			patch: "/api/demo/:id"
			body: "obj"
			description: "更新Demo的详细描述"
		};
	};

	// swagger:route GET /api/demo demo GetDemo
	// 查找Demo的详细描述
	rpc GetDemo (GetDemoReq) returns (GetDemoResp) {
		option (api.http) = {
			get: "/api/demo"
			description: "查找Demo的详细描述"
		};
	};
}

// 空的结构体
message Empty {}

// 增加Demo请求结构体
message AddDemoReq {
	// swagger:parameters AddDemo
	string app = 1;		// 应用名称
	string version = 2;	// 应用版本
}

// 删除Demo请求结构体
message DeleteDemoReq {
	// swagger:parameters DeleteDemo
	int32 id = 1 [(api.in)="path"];	// 应用ID
}

// 更新Demo请求结构体
message UpdateDemoReq {
	// swagger:parameters UpdateDemo
	int32 id = 1 [(api.in)="path"]; // 应用ID
	UpdateMessage obj = 2;			// 更新信息
}

// 更新信息结构体
message UpdateMessage {
	string app = 2;		// 应用名称
	string version = 3;	// 应用版本
}

// 查找Demo请求结构体
message GetDemoReq {
	// swagger:parameters GetDemo
	int32 id = 1 [(api.in)="query"];	// 应用ID
}

// 查找Demo响应结构体
message GetDemoResp {
	int32 id = 1;		// 应用ID
	string app = 2;		// 应用名称
	string version = 3;	// 应用版本
}
`
}
