/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-11 14:28:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:15:04
 * @Description: The config package provides functionality for managing HTTP server configuration.
 * It defines the HttpConfig struct and provides a method to load the configuration from a YAML file.
 */
package config

import "github.com/yangjerry110/tool/conf"

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
