/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 15:47:11
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-22 17:30:28
 * @Description: test
 */
package protobuf

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type DemoProtobuf interface {
	SaveTemplate(path string, projectImportPath string) error
	GetTemplate() string
}

type Demo struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectImportPath
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:53:36
 * @return {*}
 */
func (d *Demo) SaveTemplate(path string, projectImportPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectImportPath string
	}

	data := &Data{ProjectImportPath: projectImportPath}
	return templates.CreateCommonTemplate().SaveTemplate(path, "demo.proto", d.GetTemplate(), data, "proto")
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:52:36
 * @return {*}
 */
func (d *Demo) GetTemplate() string {
	return `syntax = "proto3";

package demo;

option go_package = "{{.ProjectImportPath}}/vo/protobuf";

import "http.proto";

service DemoApi {

	// 增加
	rpc AddDemo (AddDemoReq) returns (Empty) {
		option (api.http) = {
			post: "/api/demo"
			body: "*"
			description: "增加Demo的详细描述"
		};
	};

	// 删除
	rpc DeleteDemo (DeleteDemoReq) returns (Empty) {
		option (api.http) = {
			delete: "/api/demo/:id"
			description: "删除Demo的详细描述"
		};
	};

	// 更新
	rpc UpdateDemo (UpdateDemoReq) returns (Empty) {
		option (api.http) = {
			patch: "/api/demo/:id"
			body: "obj"	// 指定的body 必须是个message类型
			description: "更新Demo的详细描述"
		};
	};

	// 查找
	rpc GetDemo (GetDemoReq) returns (GetDemoResp) {
		option (api.http) = {
			get: "/api/demo"
			description: "查找Demo的详细描述"
		};
	};
}

// 空的结构体
message Empty {}

message AddDemoReq {
	string app = 1;		// 应用名称
	string version = 2;	// 应用版本
}

message DeleteDemoReq {
	int32 id = 1 [(api.in)="path"];	// 应用ID
}

message UpdateDemoReq {
	int32 id = 1 [(api.in)="path"]; // 应用ID
	UpdateMessage obj = 2;			// 更新信息
}

message UpdateMessage {
	string app = 2;		// 应用名称
	string version = 3;	// 应用版本
}

message GetDemoReq {
	int32 id = 1 [(api.in)="query"];	// 应用ID
}

message GetDemoResp {
	int32 id = 1;		// 应用ID
	string app = 2;		// 应用名称
	string version = 3;	// 应用版本
}
`
}
