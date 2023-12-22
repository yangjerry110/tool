/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 17:01:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 16:01:57
 * @Description: newApp
 */
package command

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/folder"
	internalfolder "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder"
	internalFolderCache "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/cache"
	internalFolderConfig "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/config"
	internalFolderYamlConfig "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/config/yamlConfig"
	internalFolderDao "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/dao"
	internalFolderErrors "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/errors"
	internalFolderModel "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/model"
	internalFolderQuery "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/query"
	internalFolderService "github.com/yangjerry110/tool/internal/cmd/folder/internalFolder/service"
	folderProto "github.com/yangjerry110/tool/internal/cmd/folder/proto"
	floderRouter "github.com/yangjerry110/tool/internal/cmd/folder/router"
	"github.com/yangjerry110/tool/internal/cmd/folder/vo"
	folderVoProtobuf "github.com/yangjerry110/tool/internal/cmd/folder/vo/protobuf"
	"github.com/yangjerry110/tool/internal/cmd/template"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/cache"
	internalTemplateConfig "github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/config"
	yamlconfig "github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/config/yamlConfig"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/dao"
	internalTemplateErrors "github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/errors"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/model"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/service"
	"github.com/yangjerry110/tool/internal/cmd/template/proto"
	"github.com/yangjerry110/tool/internal/cmd/template/router"
	"github.com/yangjerry110/tool/internal/cmd/template/vo/protobuf"
	"github.com/yangjerry110/tool/internal/errors"
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

	// judge cliContext
	// if nil return err
	if n.CliContext == nil {
		return errors.ErrCmdCommandNoCliContext
	}

	// Set App conf
	if err := config.CreateConfig(&config.App{CliContext: n.CliContext}).SetConfig(); err != nil {
		return err
	}

	// Ask ImportPath
	importPath, err := n.askImportPath()
	if err != nil {
		return err
	}

	// Set ImportPath
	config.ProjectImportPathConf.ImportPath = importPath

	// newFolder
	if err := n.newFloder(); err != nil {
		return err
	}

	// newTemplate
	if err := n.newTemplate(); err != nil {
		return err
	}

	// // Swag Init
	// if err := exec.Command("swag", "init").Run(); err != nil {
	// 	return err
	// }
	return nil
}

/**
 * @description: newFloder
 * @author: Jerry.Yang
 * @date: 2023-12-19 16:03:54
 * @return {*}
 */
func (n *NewApp) newFloder() error {

	// NewAppFloder
	if err := folder.CreateFlod(&folder.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppRouter
	if err := folder.CreateFlod(&floderRouter.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppProto
	if err := folder.CreateFlod(&folderProto.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppInternal
	if err := folder.CreateFlod(&internalfolder.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppCache
	if err := folder.CreateFlod(&internalFolderCache.NewAppCache{}).New(); err != nil {
		return err
	}

	// NewAppConfig
	if err := folder.CreateFlod(&internalFolderConfig.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppConfig YamlConfig
	if err := folder.CreateFlod(&internalFolderYamlConfig.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppQuery
	if err := folder.CreateFlod(&internalFolderQuery.NewAppQuery{}).New(); err != nil {
		return err
	}

	// NewAppDao
	if err := folder.CreateFlod(&internalFolderDao.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppErrors
	if err := folder.CreateFlod(&internalFolderErrors.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppModel
	if err := folder.CreateFlod(&internalFolderModel.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppService
	if err := folder.CreateFlod(&internalFolderService.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppVo
	if err := folder.CreateFlod(&vo.NewApp{}).New(); err != nil {
		return err
	}

	// NewAppVoProtobuf
	if err := folder.CreateFlod(&folderVoProtobuf.NewApp{}).New(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: newTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:50:00
 * @return {*}
 */
func (n *NewApp) newTemplate() error {

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

	// NewAppRedisYamlConfig
	if err := template.CreateTemplate(&yamlconfig.NewAppRedis{}).New(); err != nil {
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

	// NewAppRedisConfig
	if err := template.CreateTemplate(&internalTemplateConfig.NewAppRedis{}).New(); err != nil {
		return err
	}

	// NewAppConfig
	if err := template.CreateTemplate(&internalTemplateConfig.NewAppConfig{}).New(); err != nil {
		return err
	}

	// NewAppCacheRedis
	if err := template.CreateTemplate(&cache.NewAppCacheRedis{}).New(); err != nil {
		return err
	}

	// NewAppBaseDao
	if err := template.CreateTemplate(&dao.NewAppBaseDao{}).New(); err != nil {
		return err
	}

	// NewAppErrors
	if err := template.CreateTemplate(&internalTemplateErrors.NewAppError{}).New(); err != nil {
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

	// NewAppMain
	if err := template.CreateTemplate(&template.NewAppMain{}).New(); err != nil {
		return err
	}

	// NewAppVoProtobuf
	if err := template.CreateTemplate(&protobuf.NewAppDemoProto{}).New(); err != nil {
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
