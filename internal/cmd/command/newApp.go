/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 17:01:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:12:29
 * @Description: newApp
 */
package command

import (
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
	internalTemplateConfig "github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/config"
	yamlconfig "github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/config/yamlConfig"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/dao"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/errors"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/model"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/service"
	"github.com/yangjerry110/tool/internal/cmd/template/proto"
	"github.com/yangjerry110/tool/internal/cmd/template/router"
)

type NewApp struct {
	CliContext *cli.Context
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:12:29
 * @return {*}
 */
func (n *NewApp) New() error {

	// Set App conf
	if err := config.CreateConfig(&config.App{}).SetConfig(); err != nil {
		return err
	}

	// Ask ImportPath
	importPath, err := n.askImportPath()
	if err != nil {
		return err
	}

	// Set ImportPath
	config.ProjectImportPathConf.ImportPath = importPath

	// NewAppDatabaseYamlConfig
	if err := template.CreateTemplate(&yamlconfig.NewAppDatabase{}).New(); err != nil {
		return err
	}

	// NewAppLoggerYamlConfig
	if err := template.CreateTemplate(&yamlconfig.NewAppLogger{}).New(); err != nil {
		return err
	}

	// NewAppRouterYamlConfig
	if err := template.CreateTemplate(&yamlconfig.NewAppRouter{}).New(); err != nil {
		return err
	}

	// NewAppDatabaseConfig
	if err := template.CreateTemplate(&internalTemplateConfig.NewAppDatabase{}).New(); err != nil {
		return err
	}

	// NewAppLoggerConfig
	if err := template.CreateTemplate(&internalTemplateConfig.NewAppLogger{}).New(); err != nil {
		return err
	}

	// NewAppRouterConfig
	if err := template.CreateTemplate(&internalTemplateConfig.NewAppRouter{}).New(); err != nil {
		return err
	}

	// NewAppBaseDao
	if err := template.CreateTemplate(&dao.NewAppBaseDao{}).New(); err != nil {
		return err
	}

	// NewAppErrors
	if err := template.CreateTemplate(&errors.NewAppError{}).New(); err != nil {
		return err
	}

	// NewAppBaseModel
	if err := template.CreateTemplate(&model.NewAppBaseModel{}).New(); err != nil {
		return err
	}

	// NewAppBaseService
	if err := template.CreateTemplate(&service.NewAppBaseService{}).New(); err != nil {
		return err
	}

	// NewAppDemoService
	if err := template.CreateTemplate(&service.NewAppDemoService{}).New(); err != nil {
		return err
	}

	// NewAppHttpProto
	if err := template.CreateTemplate(&proto.NewAppHttpProto{}).New(); err != nil {
		return err
	}

	// NewAppDemoProto
	if err := template.CreateTemplate(&proto.NewAppDemoProto{}).New(); err != nil {
		return err
	}

	// NewAppBaseRouter
	if err := template.CreateTemplate(&router.NewAppBaseRouter{}).New(); err != nil {
		return err
	}

	// NewAppDemoRouter
	if err := template.CreateTemplate(&router.NewAppDemoRouter{}).New(); err != nil {
		return err
	}

	// NewAppGoMod
	if err := template.CreateTemplate(&template.NewAppGoMod{}).New(); err != nil {
		return err
	}

	// Action go mod tidy
	if err := exec.Command("go", "mod", "tidy").Run(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: askImportPath
 * @author: Jerry.Yang
 * @date: 2023-12-18 17:04:03
 * @return {*}
 */
func (n *NewApp) askImportPath() (string, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "importPath",
			Prompt: &survey.Input{
				Message: "please input importPath ? ",
				Default: "",
			},
		},
	}

	// set answer
	type Answer struct {
		ImportPath string `survey:"importPath"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return "", err
	}
	return answer.ImportPath, nil
}
