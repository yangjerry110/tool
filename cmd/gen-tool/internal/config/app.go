/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 17:05:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:21:33
 * @Description: app conf
 */
package config

import (
	"fmt"

	"github.com/golib/cli"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/errors"
	"github.com/yangjerry110/tool/conf"
)

type App struct {
	CliContext *cli.Context
	AppName    string
}

/**
 * @description: AppConf
 * @author: Jerry.Yang
 * @date: 2023-12-18 17:06:47
 * @return {*}
 */
var AppConf = &App{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-18 17:06:24
 * @return {*}
 */
func (a *App) SetConfig() error {

	// judge cliContext
	// if nil return err
	if a.CliContext == nil {
		return errors.ErrConfigNoCliContext
	}

	// get first args
	// this first arg is appName
	appName := a.CliContext.Args().First()

	// judge appName
	// if is empty return err
	if appName == "" {
		return errors.ErrConfigNoAppName
	}

	// set ProtobufName
	AppConf.AppName = appName

	// Set ProjectPath
	if err := conf.CreateConf(&ProjectPath{}).SetConfig(); err != nil {
		return err
	}

	// Set ProjectPath And appName
	ProjectPathConf.Path = fmt.Sprintf("%s/%s", ProjectPathConf.Path, appName)
	return nil
}
