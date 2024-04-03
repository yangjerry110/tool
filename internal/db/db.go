/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:55:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-03 15:17:31
 * @Description: db
 */
package db

import (
	"context"

	"gorm.io/gorm"
)

type DbInterface interface {
	CreateAllClient() error
	CreateClient(dbName string) error
	GetClient(ctx context.Context, dbName string) (*gorm.DB, error)
	TransactionBegin(ctx *context.Context, dbName string) error
	TransactionCommit(ctx context.Context, dbName string) error
	TransactionRollback(ctx context.Context, dbName string) error
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
