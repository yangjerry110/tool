/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-24 14:26:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 16:30:16
 * @Description: http
 */
package commands

import (
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
	httpProto "github.com/yangjerry110/tool/pkg/protocol/api"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type HttpCommands interface {
	GetHttpRule(method *protogen.Method) (*HttpRule, error)
	FormatHttpRule(httpProtoRule *httpProto.HttpRule) (*HttpRule, error)
}

type Http struct{}

/**
 * @description: GetHttpRule
 * @param {*protogen.Method} method
 * @author: Jerry.Yang
 * @date: 2023-05-24 14:31:26
 * @return {*}
 */
func (h *Http) GetHttpRule(method *protogen.Method) (*HttpRule, error) {

	/**
	 * @step
	 * @因为我们通过method.Desc.Options() 拿到的数据类型是`interface{}` 类型
	 * @所以这里我们需要对Options，明确指定转换为 *descriptorpb.MethodOptions 类型
	 * @这样子就能拿到我们的MethodOption对象
	 **/
	options, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, errors.Err_Http_Rules_Options_Is_Empty
	}

	/**
	 * @step
	 * @PS：重点
	 * @这里我们看到我们借助了一个非protogen下的包的内容
	 * @原因就是，protobuf编译器会把自定义的Option全部指定为Extension，由于并非内置的属性和值
	 * @protobuf官方是没办法拿到和你对应的可读的内容的，只能通过拿到经过序列化之后的数据。
	 * @因此，我们这里通过 proto.GetExtension的方法，把刚才unknow.proto单独编译好的 unknow.pb.proto 文件下的 pb. E_HTTP 加载进来，指定了我需要在自定义扩展的MethodOptions中，拿到该Http下里面的value
	 * @也因此，我们可以再经过一次类型转换，就可以拿到了具体的httpRulehttpRule, ok := proto.GetExtension(options, pb.E_Http).(*pb.HttpRule)
	 **/
	httpProtoRule, ok := proto.GetExtension(options, httpProto.E_Http).(*httpProto.HttpRule)
	if !ok {
		return nil, errors.Err_Http_Rules_Extensions_Is_Empty
	}

	/**
	 * @step
	 * @根据不同类型的pattern，获取url,method
	 **/
	httpRule, err := h.FormatHttpRule(httpProtoRule)
	if err != nil {
		return nil, err
	}
	return httpRule, nil
}

/**
 * @description: FormatHttpRule
 * @param {*httpProto.HttpRule} httpProtoRule
 * @author: Jerry.Yang
 * @date: 2023-05-24 14:31:15
 * @return {*}
 */
func (h *Http) FormatHttpRule(httpProtoRule *httpProto.HttpRule) (*HttpRule, error) {

	/**
	 * @step
	 * @返回结构体
	 **/
	httpRule := &HttpRule{}

	/**
	 * @step
	 * @赋值description
	 **/
	httpRule.Description = httpProtoRule.GetDescription()

	/**
	 * @step
	 * @接下来，我们就可以通过GetXxx的方式，来获取我们设置在其Message内部filed
	 * @获取定义的是什么请求类型，是get，post，put，delete，path
	 **/
	httpRulePattern := httpProtoRule.GetPattern()

	/**
	 * @step
	 * @判断pattern的类型
	 * @get
	 **/
	_, httpGetOk := httpRulePattern.(*httpProto.HttpRule_Get)
	if httpGetOk {
		httpRule.Method = "GET"
		httpRule.Url = httpProtoRule.GetGet()
		return httpRule, nil
	}

	/**
	 * @step
	 * @判断pattern的类型
	 * @post
	 **/
	_, httpPostOk := httpRulePattern.(*httpProto.HttpRule_Post)
	if httpPostOk {
		httpRule.Method = "POST"
		httpRule.Url = httpProtoRule.GetPost()
		return httpRule, nil
	}

	/**
	 * @step
	 * @判断pattern的类型
	 * @put
	 **/
	_, httpPutOk := httpRulePattern.(*httpProto.HttpRule_Put)
	if httpPutOk {
		httpRule.Method = "PUT"
		httpRule.Url = httpProtoRule.GetPut()
		return httpRule, nil
	}

	/**
	 * @step
	 * @判断pattern的类型
	 * @delete
	 **/
	_, httpDeleteOk := httpRulePattern.(*httpProto.HttpRule_Delete)
	if httpDeleteOk {
		httpRule.Method = "DELETE"
		httpRule.Url = httpProtoRule.GetDelete()
		return httpRule, nil
	}

	/**
	 * @step
	 * @判断pattern的类型
	 * @patch
	 **/
	_, httpPatchOk := httpRulePattern.(*httpProto.HttpRule_Patch)
	if httpPatchOk {
		httpRule.Method = "PATCH"
		httpRule.Url = httpProtoRule.GetPatch()
		return httpRule, nil
	}

	/**
	 * @step
	 * @return
	 **/
	return nil, errors.Err_Http_Rules_Methods_No_Match
}
