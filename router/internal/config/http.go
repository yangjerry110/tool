/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 14:28:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:34:03
 * @Description: The config package provides functionality for managing HTTP server configuration and related dependencies.
 * It defines the HttpConfig struct and provides methods to load configurations for the HTTP server, database, and cache.
 */
package config

import (
	"github.com/yangjerry110/tool/cache"
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/db"
)

// HttpConfig defines the configuration for the HTTP server.
// It includes the address on which the server will listen.
type HttpConfig struct {
	Addr string `yaml:"addr"` // The address (host:port) on which the HTTP server will listen
}

// HttpConf is a global instance of HttpConfig, used to store the HTTP server configuration.
var HttpConf = &HttpConfig{}

// SetConfig loads the HTTP server configuration from a YAML file.
// It reads the configuration from the specified file and populates the HttpConf instance.
//
// Returns:
//   - error: An error if any issue occurs during the configuration loading process.
func (h *HttpConfig) SetConfig() error {
	return conf.CreateYamlConf("router.yaml", HttpConf).SetConfig()
}

// HttpRouterConfig defines the configuration for the HTTP router and its dependencies.
// It includes methods to set up configurations for the HTTP server, database, and cache.
type HttpRouterConfig struct{}

// SetConfig initializes configurations for the HTTP router and its dependencies.
// It sets up the configuration path, GORM database configuration, Redis cache configuration, and HTTP server configuration.
//
// Returns:
//   - error: An error if any issue occurs during the configuration loading process.
func (h *HttpRouterConfig) SetConfig() error {
	// Set up the configuration path
	err := conf.CreateConfigPathConf().SetConfig()
	if err != nil {
		return err
	}

	// Set up the GORM database configuration
	err = db.SetGormConf().SetConfig()
	if err != nil {
		return err
	}

	// Set up the Redis cache configuration
	err = cache.SetRedisConf().SetConfig()
	if err != nil {
		return err
	}

	// Set up the HTTP server configuration
	err = conf.CreateConf(&HttpConfig{}).SetConfig()
	if err != nil {
		return err
	}

	return nil
}
