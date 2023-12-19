/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 11:43:44
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 16:38:48
 * @Description: gormDb
 */
package db

import (
	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/db"
	gormdb "github.com/yangjerry110/tool/internal/db/gormDb"
)

type GormDb struct{}

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
	return CreateDb(&gormdb.GormDbClient{})
}
