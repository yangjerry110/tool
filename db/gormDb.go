/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 11:43:44
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 14:00:05
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
	if err := conf.CreateConf(&gormdb.GormDbConfig{}).SetConfig(); err != nil {
		return err
	}
	return nil
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
