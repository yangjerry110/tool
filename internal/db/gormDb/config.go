/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:56:08
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-01-26 16:32:44
 * @Description: gorm config
 */
package gormdb

import (
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/errors"
	"gorm.io/gorm/logger"
)

type GormDbConfig struct {
	Dsn                    string          `yaml:"dsn"`
	SkipDefaultTransaction bool            `yaml:"skip_default_transaction"`
	LoggerLevel            logger.LogLevel `yaml:"logger_level"`
}

/**
 * @description: GormDbConf
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:00:49
 * @return {*}
 */
var GormDbConfs = map[string]*GormDbConfig{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:05:48
 * @return {*}
 */
func (g *GormDbConfig) SetConfig() error {

	/**
	 * @step
	 * @返回结果
	 **/
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "database.yaml", FileType: "yaml", ConfData: GormDbConfs}).SetConfig(); err != nil {
		return err
	}
	return nil
}

/**
 * @description: ResetConfigSkipDefaultTransaction
 * @param {string} clientName
 * @param {bool} skipDefaultTransaction
 * @author: Jerry.Yang
 * @date: 2024-01-26 16:32:42
 * @return {*}
 */
func ResetConfigSkipDefaultTransaction(clientName string, skipDefaultTransaction bool) error {

	// get gormDbConf by clientName
	gormDbConf, isExistGormDbConf := GormDbConfs[clientName]
	if !isExistGormDbConf {
		return errors.ErrGormDbConfIsNotExist
	}

	// Set GormDbConf
	gormDbConf.SkipDefaultTransaction = skipDefaultTransaction

	// Set GormDbConfs
	GormDbConfs[clientName] = gormDbConf

	// Return
	return nil
}
