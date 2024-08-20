/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 11:41:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-08-19 16:33:01
 * @Description: file service
 */
package protocgentoolservice

import (
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type File struct {
	File *protogen.File
}

/**
 * @description: Generate
 * @author: Jerry.Yang
 * @date: 2023-12-12 15:32:38
 * @return {*}
 */
func (f *File) Generate() error {

	// Set ProtobufFileConf
	if err := conf.CreateConf(&config.ProtobufFile{File: f.File}).SetConfig(); err != nil {
		return err
	}

	// Define
	protoServices := []*protogen.Service{}
	protoMethods := []*protogen.Method{}

	// Judge f.File
	// If == nil; return err
	if f.File == nil {
		return errors.ErrProtocGenToolServiceNoFile
	}

	// This code simply ignores the code that contains streamClient and streamServer in the proto file
	// isGenerate := false
	for _, service := range f.File.Services {
		// for service.Methods
		for _, method := range service.Methods {
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}

			// append protoMethods
			protoMethods = append(protoMethods, method)
		}

		// append protoServices
		protoServices = append(protoServices, service)
	}

	// Judge protoServices
	// if len == 0; return err
	if len(protoServices) == 0 {
		return errors.ErrProtocGenToolServiceNoServices
	}

	// Judge protoMethods
	// if len == 0 ; return err
	if len(protoMethods) == 0 {
		return errors.ErrProtocGenToolServiceNoMethods
	}

	// for protoServices
	// set ProtocSerivce
	for _, protocService := range protoServices {
		// set conf
		if err := conf.CreateConf(&config.ProtocService{ProtoService: protocService}).SetConfig(); err != nil {
			return err
		}
	}

	// The next level of file is the services level
	// Generate services
	for _, protoMethod := range protoMethods {
		// method generate
		if err := CreateProtoGenToolService(&Method{Method: protoMethod}).Generate(); err != nil {
			return err
		}
	}

	// generate router
	if err := CreateProtoGenToolService(&Router{}).Generate(); err != nil {
		return err
	}

	// generate service
	if err := CreateProtoGenToolService(&Service{}).Generate(); err != nil {
		return err
	}
	return nil
}
