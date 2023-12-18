/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:48:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 11:29:49
 * @Description: db
 */
package db

import (
	"github.com/yangjerry110/tool/internal/db"
	gormdb "github.com/yangjerry110/tool/internal/db/gormDb"
)

/**
 * @description: defaultDB
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:28:37
 * @return {*}
 */
var DefaultDb = &gormdb.GormDbClient{}

/**
 * @description: CreateDb
 * @param {...db.DbInterface} dbs
 * @author: Jerry.Yang
 * @date: 2023-12-11 11:28:27
 * @return {*}
 */
func CreateDb(dbs ...db.DbInterface) db.DbInterface {
	if len(dbs) == 0 {
		return DefaultDb
	}
	return dbs[0]
}
