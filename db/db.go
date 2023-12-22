/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:48:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 16:21:25
 * @Description: db
 */
package db

import (
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/db"
	gormdb "github.com/yangjerry110/tool/internal/db/gormDb"
)

/**
 * @description: CreateGormDbConf
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:44:51
 * @return {*}
 */
func CreateGormDbConf() error {
	return conf.CreateConf(&gormdb.GormDbConfig{}).SetConfig()
}

/**
 * @description: CreateGormDb
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:00:11
 * @return {*}
 */
func CreateGormDb() db.DbInterface {
	return db.CreateDb(&gormdb.GormDbClient{})
}
