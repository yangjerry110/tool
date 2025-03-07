/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 15:06:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-07 16:39:51
 * @Description: method
 */
package protocgentoolservice

import (
	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
	"github.com/yangjerry110/protoc-gen-go/proto"
	"github.com/yangjerry110/protoc-gen-go/types/descriptorpb"
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	httpProto "github.com/yangjerry110/tool/pkg/protocol/api"
)

type Method struct {
	Method *protogen.Method
}

/**
 * @description: Generate
 * @author: Jerry.Yang
 * @date: 2023-12-12 15:32:22
 * @return {*}
 */
func (m *Method) Generate() error {

	// Judge Method
	// if Method == nil; return err
	if m.Method == nil {
		return errors.ErrProtocGenToolServiceNoMethod
	}

	/**
	 * @step
	 * @因为我们通过method.Desc.Options() 拿到的数据类型是`interface{}` 类型
	 * @所以这里我们需要对Options，明确指定转换为 *descriptorpb.MethodOptions 类型
	 * @这样子就能拿到我们的MethodOption对象
	 **/
	options, ok := m.Method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return errors.ErrProtocGenToolServiceNoOptions
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
		return errors.ErrProtocGenToolServiceNoExtensions
	}

	// Set HttpProtoRules
	if err := conf.CreateConf(&config.ProtocHttpRule{HttpProtoMethod: m.Method, HttpProtoRule: httpProtoRule}).SetConfig(); err != nil {
		return err
	}
	return nil
}
