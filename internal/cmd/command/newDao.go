/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 14:58:22
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 16:52:27
 * @Description: newDao
 */
package command

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/golib/cli"
	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
	"github.com/yangjerry110/tool/internal/cmd/template/internalTemplate/dao"
	"github.com/yangjerry110/tool/internal/conf"
)

type NewDao struct {
	CliContext *cli.Context
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:31:51
 * @return {*}
 */
func (n *NewDao) New() error {

	// Set ProjectPath
	if err := conf.CreateConf(&config.ProjectPath{}).SetConfig(); err != nil {
		return err
	}

	// Set ImportProjectPath
	if err := conf.CreateConf(&config.ProjectImportPath{}).SetConfig(); err != nil {
		return err
	}

	// Set DaoName
	if err := conf.CreateConf(&config.Dao{CliContext: n.CliContext}).SetConfig(); err != nil {
		return err
	}

	// Ask DB Name
	dbName, err := n.askDbName()
	if err != nil {
		return err
	}

	// Ask Model Name
	modelName, err := n.askModelName()
	if err != nil {
		return err
	}

	// Set NewDao
	templateNewDao := &dao.NewDao{}
	templateNewDao.DaoName = config.DaoConf.DaoName
	templateNewDao.DaoNameUp = template.FirstUpper(config.DaoConf.DaoName)
	templateNewDao.DbName = dbName
	templateNewDao.FirstDaoName = config.DaoConf.DaoName[:1]
	templateNewDao.ProjectImportPath = config.ProjectImportPathConf.ImportPath
	templateNewDao.ModelName = modelName
	templateNewDao.Time = template.GetFormatNowTime()

	// Action NewDao
	if err := template.CreateTemplate(templateNewDao).New(); err != nil {
		return err
	}

	// Append Base
	templateNewDaoAppendBase := &dao.NewDaoAppendBase{}
	templateNewDaoAppendBase.DaoName = config.DaoConf.DaoName
	templateNewDaoAppendBase.DaoNameUp = template.FirstUpper(config.DaoConf.DaoName)
	templateNewDaoAppendBase.Time = template.GetFormatNowTime()

	// Action Append Base
	if err := template.CreateTemplate(templateNewDaoAppendBase).New(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: askDbName
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:35:21
 * @return {*}
 */
func (n *NewDao) askDbName() (string, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "dbName",
			Prompt: &survey.Input{
				Message: "please input dbName ? ",
				Default: "",
			},
		},
	}

	// set answer
	type Answer struct {
		DbName string `survey:"dbName"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return "", err
	}
	return answer.DbName, nil
}

/**
 * @description: askModelName
 * @author: Jerry.Yang
 * @date: 2023-12-21 16:18:07
 * @return {*}
 */
func (n *NewDao) askModelName() (string, error) {

	// define question
	questions := []*survey.Question{
		{
			Name: "modelName",
			Prompt: &survey.Input{
				Message: "please input modelName ? ",
				Default: "",
			},
		},
	}

	// set answer
	type Answer struct {
		ModelName string `survey:"modelName"`
	}

	// action ask
	// Set the answer data to Answer
	answer := &Answer{}
	err := survey.Ask(questions, answer)
	if err != nil {
		return "", err
	}
	return answer.ModelName, nil
}
