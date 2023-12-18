/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 17:35:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 17:40:00
 * @Description: protobuf config
 */
package config

import (
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/errors"
)

type Protobuf struct {
	CliContext   *cli.Context
	ProtobufName string
}

/**
 * @description: ProtobufConf
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:36:18
 * @return {*}
 */
var ProtobufConf = &Protobuf{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:40:36
 * @return {*}
 */
func (p *Protobuf) SetConfig() error {

	// get first args
	// this first arg is protobufName
	ProtobufName := p.CliContext.Args().First()

	// judge protobufName
	// if is empty return err
	if ProtobufName == "" {
		return errors.ErrCmdConfNoProtobufName
	}

	// set ProtobufName
	ProtobufConf.ProtobufName = ProtobufName
	return nil
}
