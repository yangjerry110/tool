/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:56:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-03 15:17:58
 * @Description: gorm db client
 */
package gormdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/yangjerry110/tool/internal/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDbClient struct{}

/**
 * @description: GormDbClients
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:20:03
 * @return {*}
 */
type transactionContextKey string

var gormDbClients = sync.Map{}
var transactionKey transactionContextKey = "transaction"

/**
 * @description: CreateAllClient
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:16:59
 * @return {*}
 */
func (g *GormDbClient) CreateAllClient() error {

	// If len GormDbConf == 0
	// return
	if len(GormDbConfs) == 0 {
		return nil
	}

	// For GormDbConf
	for dbName := range GormDbConfs {
		if err := g.CreateClient(dbName); err != nil {
			return err
		}
	}
	return nil
}

/**
 * @description: CreateClient
 * @param {string} dbName
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:19:56
 * @return {*}
 */
func (g *GormDbClient) CreateClient(dbName string) error {

	// get gormDb conf
	// judge conf is exist
	gormDbConf, isExist := GormDbConfs[dbName]
	if !isExist {
		return errors.ErrGormDbConfIsNotExist
	}

	// Set dbConfig
	config := &gorm.Config{}
	config.SkipDefaultTransaction = gormDbConf.SkipDefaultTransaction
	config.Logger = logger.Default.LogMode(gormDbConf.LoggerLevel)

	// init client
	// init conf
	db, err := gorm.Open(mysql.Open(gormDbConf.Dsn), config)
	if err != nil {
		return err
	}

	// set db to clients
	gormDbClients.Store(dbName, db)
	return nil
}

/**
 * @description: GetClient
 * @param {string} dbName
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:22:53
 * @return {*}
 */
func (g *GormDbClient) GetClient(ctx context.Context, dbName string) (*gorm.DB, error) {

	// get context transaction db client
	transactionDbClient, isExisttransactionDbClient := ctx.Value(g.transactionKey(dbName)).(*gorm.DB)
	if isExisttransactionDbClient && transactionDbClient != nil {
		return transactionDbClient, nil
	}

	// get gormClient
	// judge client is exist ?
	gormDbClient, isExist := gormDbClients.Load(dbName)
	if !isExist {
		return nil, errors.ErrGormDbClientIsNotExist
	}

	// return client
	return gormDbClient.(*gorm.DB), nil
}

// begin
//
// transaction begin
// Author yangjie04@qutoutiao.net
// Date 2024-04-02 16:32:18
func (g *GormDbClient) TransactionBegin(ctx context.Context, dbName string) (context.Context, error) {

	// get db client by dbName
	dbClient, dbClientIsExist := gormDbClients.Load(dbName)
	if !dbClientIsExist {
		return ctx, errors.ErrGormDbClientIsNotExist
	}

	// set begin db to ctx
	transactionDbClient := dbClient.(*gorm.DB).Begin()

	// judge transactionDbClient err
	// if err != nil; return err
	if transactionDbClient.Error != nil {
		return ctx, transactionDbClient.Error
	}

	// set context transactionBegion client
	return context.WithValue(ctx, g.transactionKey(dbName), transactionDbClient), nil
}

// commit
//
// transaction commit
// Author yangjie04@qutoutiao.net
// Date 2024-01-30 11:25:37
func (g *GormDbClient) TransactionCommit(ctx context.Context, dbName string) error {

	// get db client by ctx
	transactionDbClient, isExisttransactionDbClient := ctx.Value(g.transactionKey(dbName)).(*gorm.DB)
	if !isExisttransactionDbClient {
		return errors.ErrGormDbClientIsNotExist
	}

	// Set rollback
	if err := transactionDbClient.Commit().Error; err != nil {
		return err
	}
	return nil
}

// rollback
//
// transaction rollback
// Author yangjie04@qutoutiao.net
// Date 2024-01-30 11:25:37
func (g *GormDbClient) TransactionRollback(ctx context.Context, dbName string) error {

	// get db client by ctx
	transactionDbClient, isExisttransactionDbClient := ctx.Value(g.transactionKey(dbName)).(*gorm.DB)
	if !isExisttransactionDbClient {
		return errors.ErrGormDbClientIsNotExist
	}

	// Set rollback
	if err := transactionDbClient.Rollback().Error; err != nil {
		return err
	}
	return nil
}

/**
 * @description: transactionKey
 * @param {string} dbName
 * @author: Jerry.Yang
 * @date: 2024-04-03 10:52:12
 * @return {*}
 */
func (g *GormDbClient) transactionKey(dbName string) transactionContextKey {
	return transactionContextKey(fmt.Sprintf("%s-%s", dbName, transactionKey))
}
