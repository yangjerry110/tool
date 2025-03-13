/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 11:36:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:26:30
 * @Description: model config
 */
package config

import (
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/errors"
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
		return errors.ErrConfigNoCliContext
	}

	// get first args
	// this first arg is modelName
	modelName := m.CliContext.Args().First()

	// judge protobufName
	// if is empty return err
	if modelName == "" {
		return errors.ErrConfigNoModelName
	}

	// set modelName
	ModelConf.ModelName = modelName
	return nil
}
