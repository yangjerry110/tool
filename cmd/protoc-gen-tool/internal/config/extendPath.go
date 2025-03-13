/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-21 17:25:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 18:28:08
 * @Description: protobuf import path
 */
package config

import (
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
	"golang.org/x/tools/go/packages"
)

type ExtendPath struct {
	Path       string
	ImportPath string
}

/**
 * @description: ProtobufImportPathConf
 * @author: Jerry.Yang
 * @date: 2025-02-21 17:26:29
 * @return {*}
 */
var ExtendPathConf = &ExtendPath{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2025-02-21 18:52:49
 * @return {*}
 */
func (e *ExtendPath) SetConfig() error {

	/**
	 * @step
	 * @判断path
	 **/
	if e.Path == "" {
		return nil
	}

	/**
	 * @step
	 * @获取当前mod的module
	 **/
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName,
		Dir:  e.Path,
	})
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @pkgs 判断
	 **/
	if len(pkgs) == 0 {
		return errors.ErrConfigExtendPathNoImportProjectPath
	}
	e.ImportPath = pkgs[0].PkgPath

	/**
	 * @step
	 * @set conf
	 **/
	ExtendPathConf = e
	return nil
}
