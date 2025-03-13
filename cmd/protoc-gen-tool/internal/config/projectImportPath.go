package config

import (
	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
	"golang.org/x/tools/go/packages"
)

type ProjectImportPath struct {
	ImportPath string
}

/**
 * @description: ProjectImportPathConf
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:33:36
 * @return {*}
 */
var ProjectImportPathConf = &ProjectImportPath{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-11 17:33:28
 * @return {*}
 */
func (p *ProjectImportPath) SetConfig() error {

	/**
	 * @step
	 * @获取当前mod的module
	 **/
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName,
		Dir:  ProjectPathConf.Path,
	})
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @pkgs 判断
	 **/
	if len(pkgs) == 0 {
		return errors.ErrConfigProjectImportPathNoPkgs
	}
	ProjectImportPathConf.ImportPath = pkgs[0].PkgPath
	return nil
}
