/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:55:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 15:21:50
 * @Description: db
 */
package db

import (
	"gorm.io/gorm"
)

type DbInterface interface {
	CreateAllClient() error
	CreateClient(dbName string) error
	GetClient(dbName string) (*gorm.DB, error)
	TransactionBegin(dbName string) error
	TransactionCommit(dbName string) error
	TransactionRollback(dbName string) error
}

/**
 * @description: CreateDb
 * @param {DbInterface} DbInterface
 * @author: Jerry.Yang
 * @date: 2023-12-21 14:01:30
 * @return {*}
 */
func CreateDb(DbInterface DbInterface) DbInterface {
	return DbInterface
}
