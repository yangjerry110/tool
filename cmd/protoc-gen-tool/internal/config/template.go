/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-12 18:41:36
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 18:41:44
 * @Description:
 */
package config

import (
	"text/template"

	"github.com/yangjerry110/tool/cmd/protoc-gen-tool/internal/errors"
)

type Template struct {
	Name string
}

/**
 * @description: TemplateConfs
 * @author: Jerry.Yang
 * @date: 2023-12-12 16:56:51
 * @return {*}
 */
var TemplateConfs = map[string]*template.Template{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-12 16:58:50
 * @return {*}
 */
func (t *Template) SetConfig() error {

	// Judge name
	// if name == ""; return err
	if t.Name == "" {
		return errors.ErrConfigTemplateNoName
	}

	// Judge this name's conf is exist
	_, isExist := TemplateConfs[t.Name]
	if isExist {
		return nil
	}

	// not exist;
	// set template
	TemplateConfs[t.Name] = template.New(t.Name)
	return nil
}
