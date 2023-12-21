/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 11:36:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 14:30:52
 * @Description: model config
 */
package config

import (
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/errors"
)

type Model struct {
	CliContext *cli.Context
	ModelName  string
}

/**
 * @description: ModelConf
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:37:21
 * @return {*}
 */
var ModelConf = &Model{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:38:59
 * @return {*}
 */
func (m *Model) SetConfig() error {

	// If CliContext == nil
	if m.CliContext == nil {
		return errors.ErrCmdCommandNoCliContext
	}

	// get first args
	// this first arg is modelName
	modelName := m.CliContext.Args().First()

	// judge protobufName
	// if is empty return err
	if modelName == "" {
		return errors.ErrCmdConfNoProtobufName
	}

	// set modelName
	ModelConf.ModelName = modelName
	return nil
}
