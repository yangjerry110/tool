/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 15:08:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-05 20:18:02
 * @Description: db
 */
package db

import (
	"github.com/yangjerry110/tool/db"
	"gorm.io/gorm"
)

type DbPkgInterface interface{}

type DbPkg struct{}

/**
 * @description: defaultClient
 * @author: Jerry.Yang
 * @date: 2023-02-09 15:14:48
 * @return {*}
 */
var defaultClient = CreateClient(&db.GormDb{})

/**
 * @description: CreateClient
 * @param {db.BaseDbInterface} dbInterface
 * @author: Jerry.Yang
 * @date: 2023-02-09 15:17:11
 * @return {*}
 */
func CreateClient(dbInterface db.DbInterface) db.DbInterface {
	return dbInterface
}

/**
 * @description: SetClient
 * @param {db.DbInterface} dbInterface
 * @author: Jerry.Yang
 * @date: 2023-02-09 15:30:00
 * @return {*}
 */
func SetClient(dbInterface db.DbInterface) db.DbInterface {
	defaultClient = CreateClient(dbInterface)
	return defaultClient
}

/**
 * @description: Client
 * @param {string} dbname
 * @author: Jerry.Yang
 * @date: 2023-02-09 15:27:06
 * @return {*}
 */
func Client(dbname string, gormConfig *gorm.Config) *gorm.DB {
	return defaultClient.Client(dbname, gormConfig)
}

/**
 * @description: RenderDbConfig
 * @param {string} filePath
 * @param {string} filename
 * @author: Jerry.Yang
 * @date: 2023-02-09 15:28:27
 * @return {*}
 */
func RenderDbConfig(filePath string, filename string) error {
	return defaultClient.RenderDbConfig(filePath, filename)
}
