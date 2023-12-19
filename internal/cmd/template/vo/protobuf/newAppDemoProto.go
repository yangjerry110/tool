/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 17:03:45
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 21:30:55
 * @Description:
 */
package protobuf

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppDemoProto struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 17:06:10
 * @return {*}
 */
func (n *NewAppDemoProto) New() error {
	filePath := fmt.Sprintf("%s/vo/protobuf", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "demo.proto.go", n.getTemplate(), nil, "proto.go")
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 17:06:17
 * @return {*}
 */
func (n *NewAppDemoProto) getTemplate() string {
	return `// demo.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.4
// source: demo.proto

package protobuf

import (
	context "context"
	protoreflect "github.com/yangjerry110/protoc-gen-go/reflect/protoreflect"
	protoimpl "github.com/yangjerry110/protoc-gen-go/runtime/protoimpl"
	_ "github.com/yangjerry110/tool/pkg/protocol/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 空的结构体
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

// 增加Demo请求结构体
type AddDemoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// swagger:parameters AddDemo
	App     string` + " `protobuf:\"bytes,1,opt,name=app,proto3\" json:\"app\" form:\"app\" uri:\"app\" `" + `              // 应用名称
	Version string` + " ` protobuf:\"bytes,2,opt,name=version,proto3\" json:\"version\" form:\"version\" uri:\"version\"`" + ` // 应用版本
}

func (x *AddDemoReq) Reset() {
	*x = AddDemoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDemoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDemoReq) ProtoMessage() {}

func (x *AddDemoReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDemoReq.ProtoReflect.Descriptor instead.
func (*AddDemoReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *AddDemoReq) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *AddDemoReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// 删除Demo请求结构体
type DeleteDemoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// swagger:parameters DeleteDemo
	Id int32 ` + "` protobuf:\"varint,1,opt,name=id,proto3\" json:\"id\" form:\"id\" uri:\"id\"`" + ` // 应用ID
}

func (x *DeleteDemoReq) Reset() {
	*x = DeleteDemoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDemoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDemoReq) ProtoMessage() {}

func (x *DeleteDemoReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDemoReq.ProtoReflect.Descriptor instead.
func (*DeleteDemoReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteDemoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 更新Demo请求结构体
type UpdateDemoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// swagger:parameters UpdateDemo
	Id  int32 ` + " `protobuf:\"varint,1,opt,name=id,proto3\" json:\"id\" form:\"id\" uri:\"id\"`" + `   // 应用ID
	Obj *UpdateMessage ` + " `protobuf:\"bytes,2,opt,name=obj,proto3\" json:\"obj\" form:\"obj\" uri:\"obj\"`" + ` // 更新信息
}

func (x *UpdateDemoReq) Reset() {
	*x = UpdateDemoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDemoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDemoReq) ProtoMessage() {}

func (x *UpdateDemoReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDemoReq.ProtoReflect.Descriptor instead.
func (*UpdateDemoReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDemoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateDemoReq) GetObj() *UpdateMessage {
	if x != nil {
		return x.Obj
	}
	return nil
}

// 更新信息结构体
type UpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App     string ` + " `protobuf:\"bytes,2,opt,name=app,proto3\" json:\"app\" form:\"app\" uri:\"app\"`" + `              // 应用名称
	Version string ` + " `protobuf:\"bytes,3,opt,name=version,proto3\" json:\"version\" form:\"version\" uri:\"version\"`" + ` // 应用版本
}

func (x *UpdateMessage) Reset() {
	*x = UpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessage) ProtoMessage() {}

func (x *UpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessage.ProtoReflect.Descriptor instead.
func (*UpdateMessage) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateMessage) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *UpdateMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// 查找Demo请求结构体
type GetDemoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// swagger:parameters GetDemo
	Id int32 ` + " `protobuf:\"varint,1,opt,name=id,proto3\" json:\"id\" form:\"id\" uri:\"id\"`" + ` // 应用ID
}

func (x *GetDemoReq) Reset() {
	*x = GetDemoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoReq) ProtoMessage() {}

func (x *GetDemoReq) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoReq.ProtoReflect.Descriptor instead.
func (*GetDemoReq) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{5}
}

func (x *GetDemoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 查找Demo响应结构体
type GetDemoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 ` + " `protobuf:\"varint,1,opt,name=id,proto3\" json:\"id\" form:\"id\" uri:\"id\"`" + `                   // 应用ID
	App     string ` + " `protobuf:\"bytes,2,opt,name=app,proto3\" json:\"app\" form:\"app\" uri:\"app\"`" + `             // 应用名称
	Version string ` + " `protobuf:\"bytes,3,opt,name=version,proto3\" json:\"version\" form:\"version\" uri:\"version\"`" + ` // 应用版本
}

func (x *GetDemoResp) Reset() {
	*x = GetDemoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDemoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDemoResp) ProtoMessage() {}

func (x *GetDemoResp) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDemoResp.ProtoReflect.Descriptor instead.
func (*GetDemoResp) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{6}
}

func (x *GetDemoResp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetDemoResp) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *GetDemoResp) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x65,
	0x6d, 0x6f, 0x1a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x2a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x71, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09,
	0xca, 0xb0, 0x87, 0x4d, 0x04, 0x70, 0x61, 0x74, 0x68, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x19,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xca, 0xb0, 0x87, 0x4d,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x6f, 0x62, 0x6a,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6f, 0x62, 0x6a,
	0x22, 0x3b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xb0, 0x87, 0x4d, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x32, 0x87, 0x03, 0x0a, 0x07, 0x44, 0x65, 0x6d, 0x6f, 0x41, 0x70, 0x69, 0x12, 0x58,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x10, 0x2e, 0x64, 0x65, 0x6d, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0xaa, 0xaa, 0x87, 0x4d, 0x29, 0x22,
	0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x19, 0xe5,
	0xa2, 0x9e, 0xe5, 0x8a, 0xa0, 0x44, 0x65, 0x6d, 0x6f, 0xe7, 0x9a, 0x84, 0xe8, 0xaf, 0xa6, 0xe7,
	0xbb, 0x86, 0xe6, 0x8f, 0x8f, 0xe8, 0xbf, 0xb0, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x13, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2f, 0xaa, 0xaa, 0x87, 0x4d, 0x2a, 0x2a,
	0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x3a, 0x69, 0x64, 0x42, 0x19,
	0xe5, 0x88, 0xa0, 0xe9, 0x99, 0xa4, 0x44, 0x65, 0x6d, 0x6f, 0xe7, 0x9a, 0x84, 0xe8, 0xaf, 0xa6,
	0xe7, 0xbb, 0x86, 0xe6, 0x8f, 0x8f, 0xe8, 0xbf, 0xb0, 0x12, 0x64, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x13, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x34, 0xaa, 0xaa, 0x87, 0x4d, 0x2f,
	0x32, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x3a, 0x69, 0x64, 0x3a,
	0x03, 0x6f, 0x62, 0x6a, 0x42, 0x19, 0xe6, 0x9b, 0xb4, 0xe6, 0x96, 0xb0, 0x44, 0x65, 0x6d, 0x6f,
	0xe7, 0x9a, 0x84, 0xe8, 0xaf, 0xa6, 0xe7, 0xbb, 0x86, 0xe6, 0x8f, 0x8f, 0xe8, 0xbf, 0xb0, 0x12,
	0x5b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x10, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x2b, 0xaa, 0xaa, 0x87, 0x4d, 0x26, 0x12, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d,
	0x6f, 0x42, 0x19, 0xe6, 0x9f, 0xa5, 0xe6, 0x89, 0xbe, 0x44, 0x65, 0x6d, 0x6f, 0xe7, 0x9a, 0x84,
	0xe8, 0xaf, 0xa6, 0xe7, 0xbb, 0x86, 0xe6, 0x8f, 0x8f, 0xe8, 0xbf, 0xb0, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x2e, 0x71, 0x75, 0x74, 0x6f, 0x75, 0x74, 0x69, 0x61, 0x6f, 0x2e, 0x6e, 0x65,
	0x74, 0x2f, 0x65, 0x65, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData = file_demo_proto_rawDesc
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_proto_rawDescData)
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_demo_proto_goTypes = []interface{}{
	(*Empty)(nil),         // 0: demo.Empty
	(*AddDemoReq)(nil),    // 1: demo.AddDemoReq
	(*DeleteDemoReq)(nil), // 2: demo.DeleteDemoReq
	(*UpdateDemoReq)(nil), // 3: demo.UpdateDemoReq
	(*UpdateMessage)(nil), // 4: demo.UpdateMessage
	(*GetDemoReq)(nil),    // 5: demo.GetDemoReq
	(*GetDemoResp)(nil),   // 6: demo.GetDemoResp
}
var file_demo_proto_depIdxs = []int32{
	4, // 0: demo.UpdateDemoReq.obj:type_name -> demo.UpdateMessage
	1, // 1: demo.DemoApi.AddDemo:input_type -> demo.AddDemoReq
	2, // 2: demo.DemoApi.DeleteDemo:input_type -> demo.DeleteDemoReq
	3, // 3: demo.DemoApi.UpdateDemo:input_type -> demo.UpdateDemoReq
	5, // 4: demo.DemoApi.GetDemo:input_type -> demo.GetDemoReq
	0, // 5: demo.DemoApi.AddDemo:output_type -> demo.Empty
	0, // 6: demo.DemoApi.DeleteDemo:output_type -> demo.Empty
	0, // 7: demo.DemoApi.UpdateDemo:output_type -> demo.Empty
	6, // 8: demo.DemoApi.GetDemo:output_type -> demo.GetDemoResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDemoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDemoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDemoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_demo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDemoResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_rawDesc = nil
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DemoApiClient is the client API for DemoApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoApiClient interface {
	// swagger:route POST /api/demo demo AddDemo
	// 增加Demo的详细描述
	AddDemo(ctx context.Context, in *AddDemoReq, opts ...grpc.CallOption) (*Empty, error)
	// swagger:route DELETE /api/demo/{id} demo DeleteDemo
	// 删除Demo的详细描述
	DeleteDemo(ctx context.Context, in *DeleteDemoReq, opts ...grpc.CallOption) (*Empty, error)
	// swagger:route PATCH /api/demo/{id} demo UpdateDemo
	// 更新Demo的详细描述
	UpdateDemo(ctx context.Context, in *UpdateDemoReq, opts ...grpc.CallOption) (*Empty, error)
	// swagger:route GET /api/demo demo GetDemo
	// 查找Demo的详细描述
	GetDemo(ctx context.Context, in *GetDemoReq, opts ...grpc.CallOption) (*GetDemoResp, error)
}

type demoApiClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoApiClient(cc grpc.ClientConnInterface) DemoApiClient {
	return &demoApiClient{cc}
}

func (c *demoApiClient) AddDemo(ctx context.Context, in *AddDemoReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/demo.DemoApi/AddDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoApiClient) DeleteDemo(ctx context.Context, in *DeleteDemoReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/demo.DemoApi/DeleteDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoApiClient) UpdateDemo(ctx context.Context, in *UpdateDemoReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/demo.DemoApi/UpdateDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoApiClient) GetDemo(ctx context.Context, in *GetDemoReq, opts ...grpc.CallOption) (*GetDemoResp, error) {
	out := new(GetDemoResp)
	err := c.cc.Invoke(ctx, "/demo.DemoApi/GetDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoApiServer is the server API for DemoApi service.
type DemoApiServer interface {
	// swagger:route POST /api/demo demo AddDemo
	// 增加Demo的详细描述
	AddDemo(context.Context, *AddDemoReq) (*Empty, error)
	// swagger:route DELETE /api/demo/{id} demo DeleteDemo
	// 删除Demo的详细描述
	DeleteDemo(context.Context, *DeleteDemoReq) (*Empty, error)
	// swagger:route PATCH /api/demo/{id} demo UpdateDemo
	// 更新Demo的详细描述
	UpdateDemo(context.Context, *UpdateDemoReq) (*Empty, error)
	// swagger:route GET /api/demo demo GetDemo
	// 查找Demo的详细描述
	GetDemo(context.Context, *GetDemoReq) (*GetDemoResp, error)
}

// UnimplementedDemoApiServer can be embedded to have forward compatible implementations.
type UnimplementedDemoApiServer struct {
}

func (*UnimplementedDemoApiServer) AddDemo(context.Context, *AddDemoReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDemo not implemented")
}
func (*UnimplementedDemoApiServer) DeleteDemo(context.Context, *DeleteDemoReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDemo not implemented")
}
func (*UnimplementedDemoApiServer) UpdateDemo(context.Context, *UpdateDemoReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDemo not implemented")
}
func (*UnimplementedDemoApiServer) GetDemo(context.Context, *GetDemoReq) (*GetDemoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDemo not implemented")
}

func RegisterDemoApiServer(s *grpc.Server, srv DemoApiServer) {
	s.RegisterService(&_DemoApi_serviceDesc, srv)
}

func _DemoApi_AddDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDemoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoApiServer).AddDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoApi/AddDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoApiServer).AddDemo(ctx, req.(*AddDemoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoApi_DeleteDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDemoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoApiServer).DeleteDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoApi/DeleteDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoApiServer).DeleteDemo(ctx, req.(*DeleteDemoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoApi_UpdateDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDemoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoApiServer).UpdateDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoApi/UpdateDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoApiServer).UpdateDemo(ctx, req.(*UpdateDemoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoApi_GetDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDemoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoApiServer).GetDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoApi/GetDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoApiServer).GetDemo(ctx, req.(*GetDemoReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.DemoApi",
	HandlerType: (*DemoApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDemo",
			Handler:    _DemoApi_AddDemo_Handler,
		},
		{
			MethodName: "DeleteDemo",
			Handler:    _DemoApi_DeleteDemo_Handler,
		},
		{
			MethodName: "UpdateDemo",
			Handler:    _DemoApi_UpdateDemo_Handler,
		},
		{
			MethodName: "GetDemo",
			Handler:    _DemoApi_GetDemo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo.proto",
}
`
}
