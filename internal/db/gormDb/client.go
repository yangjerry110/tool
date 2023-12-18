/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:56:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 11:20:59
 * @Description: gorm db client
 */
package gormdb

import (
	"context"

	"github.com/yangjerry110/tool/internal/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDbClient struct {
	Config *gorm.Config
}

/**
 * @description: GormDbClients
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:20:03
 * @return {*}
 */
var GormDbClients = map[string]*gorm.DB{}

/**
 * @description: CreateClient
 * @param {context.Context} ctx
 * @param {string} dbName
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:19:56
 * @return {*}
 */
func (g *GormDbClient) CreateClient(ctx context.Context, dbName string) error {

	// get gormDb conf
	// judge conf is exist
	gormDbConf, isExist := GormDbConfs[dbName]
	if !isExist {
		return errors.ErrGormDbConfIsNotExist
	}

	// init client
	// init conf
	db, err := gorm.Open(mysql.Open(gormDbConf.Dsn), g.Config)
	if err != nil {
		return err
	}

	// set db to clients
	GormDbClients[dbName] = db
	return nil
}

/**
 * @description: GetClient
 * @param {context.Context} ctx
 * @param {string} dbName
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:22:53
 * @return {*}
 */
func (g *GormDbClient) GetClient(ctx context.Context, dbName string) (*gorm.DB, error) {

	// get gormClient
	// judge client is exist ?
	gormDbClient, isExist := GormDbClients[dbName]
	if !isExist {
		return nil, errors.ErrGormDbClientIsNotExist
	}

	// return client
	return gormDbClient, nil
}
