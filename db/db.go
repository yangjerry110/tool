/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:48:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-17 16:19:42
 * @Description: db
 */
package db

import "gorm.io/gorm"

type DbInterface interface {
	RenderDbConfig(dbConfigPath string, dbConfigFileName string) error
	Client(dbname string) *gorm.DB
}

type Db struct{}
