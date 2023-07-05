/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:48:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-05 20:17:39
 * @Description: db
 */
package db

import "gorm.io/gorm"

type DbInterface interface {
	RenderDbConfig(dbConfigPath string, dbConfigFileName string) error
	Client(dbname string, gormConfig *gorm.Config) *gorm.DB
}

type Db struct{}
