/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:50:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-05 20:06:41
 * @Description: gorm db
 */
package db

import (
	"fmt"

	"github.com/yangjerry110/tool/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDbInterface interface{}

type GormDb struct{}

/**
 * @description: Client
 * @param {string} dbname
 * @param {*gorm.Config} gormConfig
 * @author: Jerry.Yang
 * @date: 2023-07-05 20:06:48
 * @return {*}
 */
func (g *GormDb) Client(dbname string, gormConfig *gorm.Config) *gorm.DB {

	/**
	 * @step
	 * 获取当前db的配置信息
	 **/
	dbConfig, isExist := DbConfig[dbname]
	if !isExist {
		panic(fmt.Sprintf("%s is not match config", dbname))
	}

	/**
	 * @step
	 * @创建db连接
	 **/
	db, err := gorm.Open(mysql.Open(dbConfig.Dsn), gormConfig)
	if err != nil {
		panic(fmt.Sprintf("getDb Err : %v", err))
	}
	return db
}

/**
 * @description: renderDbConfig
 * @param {string} dbConfigPath
 * @param {string} dbConfigFileName
 * @author: Jerry.Yang
 * @date: 2023-02-09 14:59:42
 * @return {*}
 */
func (g *GormDb) RenderDbConfig(dbConfigPath string, dbConfigFileName string) error {

	/**
	 * @step
	 * @返回结果
	 **/
	yamlConf := conf.YamlConf{FilePath: dbConfigPath, FileName: dbConfigFileName, Conf: &DbConfig}
	return yamlConf.GetConf(DbConfig)
}
