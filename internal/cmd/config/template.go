package config

import (
	"text/template"

	"github.com/yangjerry110/tool/internal/errors"
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
		return errors.ErrCmdTemplateNoName
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
