/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 10:56:08
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:27:32
 * @Description: The gormdb package provides functionality for managing GORM database configurations.
 * It defines the GormDbConfig struct and provides a method to load configurations from a YAML file.
 */
package gormdb

import (
	"github.com/yangjerry110/tool/conf"
	"gorm.io/gorm/logger"
)

// GormDbConfig defines the configuration for a GORM database connection.
// It includes the DSN (Data Source Name), transaction behavior, and logger level.
type GormDbConfig struct {
	Dsn                    string          `yaml:"dsn"`                      // The Data Source Name (DSN) for the database connection.
	SkipDefaultTransaction bool            `yaml:"skip_default_transaction"` // Whether to skip the default transaction behavior.
	LoggerLevel            logger.LogLevel `yaml:"logger_level"`             // The logging level for the GORM logger.
}

// GormDbConfs is a map that stores GORM database configurations.
// The key is the database name, and the value is a pointer to a GormDbConfig struct.
var GormDbConfs = map[string]*GormDbConfig{}

// SetConfig loads the GORM database configurations from a YAML file.
// It uses the `conf` package to read the configuration from `database.yaml` and populate the GormDbConfs map.
//
// Returns:
//   - error: An error if the configuration file cannot be read or parsed.
func (g *GormDbConfig) SetConfig() error {
	// Load the configurations from the YAML file into the GormDbConfs map.
	err := conf.CreateYamlConf("database.yaml", GormDbConfs).SetConfig()
	if err != nil {
		return err
	}
	return nil
}
