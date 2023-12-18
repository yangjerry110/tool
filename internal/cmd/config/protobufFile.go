/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 15:48:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-12 15:53:51
 * @Description: protobufFile
 */
package config

import (
	"github.com/yangjerry110/tool/internal/errors"
	"google.golang.org/protobuf/compiler/protogen"
)

type ProtobufFile struct {
	File     *protogen.File
	FileName string
}

/**
 * @description: ProtobufFileConf
 * @author: Jerry.Yang
 * @date: 2023-12-12 15:51:47
 * @return {*}
 */
var ProtobufFileConf = &ProtobufFile{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-12 15:51:56
 * @return {*}
 */
func (p *ProtobufFile) SetConfig() error {

	// Judge File
	// If file == nil; return err
	if p.File == nil {
		return errors.ErrProtocGenToolServiceNoFile
	}

	// Get ProtobufFile
	protobufFileName := p.File.Desc.FullName()
	ProtobufFileConf.FileName = string(protobufFileName)
	return nil
}
