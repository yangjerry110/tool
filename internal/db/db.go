/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:55:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 11:25:40
 * @Description: db
 */
package db

import (
	"context"

	"gorm.io/gorm"
)

type DbInterface interface {
	CreateClient(ctx context.Context, dbName string) error
	GetClient(ctx context.Context, dbName string) (*gorm.DB, error)
}
