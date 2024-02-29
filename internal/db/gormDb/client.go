/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:56:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-02-29 17:00:35
 * @Description: gorm db client
 */
package gormdb

import (
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
var GormDbClients = map[string]*gorm.DB{}
var GormDbTransactionClients = map[string]*gorm.DB{}

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
	GormDbClients[dbName] = db
	return nil
}

/**
 * @description: GetClient
 * @param {string} dbName
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:22:53
 * @return {*}
 */
func (g *GormDbClient) GetClient(dbName string) (*gorm.DB, error) {

	// if GormDbTransactionClients exist
	gormDbTransactionClient, isExistGormDbTransactionClient := GormDbTransactionClients[dbName]
	if isExistGormDbTransactionClient && gormDbTransactionClient != nil {
		return gormDbTransactionClient, nil
	}

	// get gormClient
	// judge client is exist ?
	gormDbClient, isExist := GormDbClients[dbName]
	if !isExist {
		return nil, errors.ErrGormDbClientIsNotExist
	}

	// return client
	return gormDbClient, nil
}

// begin
//
// transaction begin
// Author yangjie04@qutoutiao.net
// Date 2024-01-30 11:26:07
func (g *GormDbClient) TransactionBegin(dbName string) error {

	// get gormDb conf
	// judge conf is exist
	gormDbConf, isExist := GormDbConfs[dbName]
	if !isExist {
		return errors.ErrGormDbConfIsNotExist
	}

	// Set dbConfig
	config := &gorm.Config{}
	config.SkipDefaultTransaction = true
	config.Logger = logger.Default.LogMode(gormDbConf.LoggerLevel)

	// init client
	// init conf
	db, err := gorm.Open(mysql.Open(gormDbConf.Dsn), config)
	if err != nil {
		return err
	}

	// set begin db to clients
	GormDbTransactionClients[dbName] = db.Begin()
	return nil
}

// commit
//
// transaction commit
// Author yangjie04@qutoutiao.net
// Date 2024-01-30 11:25:37
func (g *GormDbClient) TransactionCommit(dbName string) error {

	// Get db client by GormDbCliens
	gormDbClient, isExist := GormDbTransactionClients[dbName]
	if !isExist {
		return errors.ErrGormDbClientIsNotExist
	}

	// Set rollback
	gormDbClient.Commit()

	// Reset db client
	GormDbTransactionClients[dbName] = nil
	return nil
}

// rollback
//
// transaction rollback
// Author yangjie04@qutoutiao.net
// Date 2024-01-30 11:25:37
func (g *GormDbClient) TransactionRollback(dbName string) error {

	// Get db client by GormDbCliens
	gormDbClient, isExist := GormDbTransactionClients[dbName]
	if !isExist {
		return errors.ErrGormDbClientIsNotExist
	}

	// Set rollback
	gormDbClient.Rollback()

	// Reset db client
	GormDbTransactionClients[dbName] = nil
	return nil
}
