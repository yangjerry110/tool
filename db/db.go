/*
 * @Author: Jerry.Yang
 * @Date: 2023-02-09 14:48:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:29:38
 * @Description: The db package provides utility functions for managing database configurations and clients.
 * It includes functions for setting GORM configurations and creating GORM database clients.
 */
package db

import (
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/db/internal/db"
	gormdb "github.com/yangjerry110/tool/db/internal/db/gormDb"
)

// SetGormConf creates and returns a configuration object for GORM database settings.
// It initializes a configuration instance using the GormDbConfig struct.
//
// Returns:
//   - conf.Conf: A configuration object for GORM database settings.
func SetGormConf() conf.Conf {
	return conf.CreateConf(&gormdb.GormDbConfig{})
}

// CreateGormDb creates and returns an instance of a GORM database client.
// It uses the factory function `CreateDb` to initialize a GORM database client.
//
// Returns:
//   - db.DbInterface: An instance of the GORM database client implementing the DbInterface.
func CreateGormDb() db.DbInterface {
	return db.CreateDb(&gormdb.GormDbClient{})
}
