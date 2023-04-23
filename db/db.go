/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:48:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-02-09 15:24:08
 * @Description: db
 */
package db

import "gorm.io/gorm"

type DbInterface interface {
	RenderDbConfig(dbConfigPath string, dbConfigFileName string)
	Client(dbname string) *gorm.DB
}

type Db struct{}
