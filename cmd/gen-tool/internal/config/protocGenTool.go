/**
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 11:22:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-24 16:55:50
 * @Description: protocGenTool
 * The `protocGenTool` package provides configuration management for the protoc generation tool.
 */
package config

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/conf"
	"gopkg.in/yaml.v2"
)

// ProtocGenTool represents the configuration structure for the protoc generation tool.
type ProtocGenTool struct {
	IsFirstCreate bool   // Indicates whether this is the first time the configuration is being created.
	IsAppend      bool   // Indicates whether the tool should append to existing files.
	IsExtend      bool   // Indicates whether the tool should extend the existing configuration.
	ExtendPath    string // The path where the extension configuration is located.
}

/**
 * @description: ProtocGenToolConf
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:23:52
 * @return {*}
 * ProtocGenToolConf is a global instance of the ProtocGenTool configuration.
 */
var ProtocGenToolConf = &ProtocGenTool{}

// Path and file name for the protoc generation tool configuration.
var protoGenToolConfPath = "/data/protobuf/tool/"
var protoGenToolConfName = "protoGenToolConf.yaml"

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:26:39
 * @return {*}
 * SetConfig initializes the configuration path and sets the configuration for the protoc generation tool.
 */
func (p *ProtocGenTool) SetConfig() error {
	// Create the configuration path if it doesn't exist.
	if err := conf.CreatePathConf(protoGenToolConfPath).SetConfig(); err != nil {
		return err
	}

	// Set the YAML configuration for the protoc generation tool.
	return conf.CreateYamlConf(protoGenToolConfName, ProtocGenToolConf).SetConfig()
}

/**
 * @description: SetConf
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:26:39
 * @return {*}
 * SetConf ensures the directory exists, creates the configuration file, and encodes the configuration data into YAML format.
 */
func (p *ProtocGenTool) SetConf() error {

	// Ensure the directory exists with the specified permissions.
	err := os.MkdirAll(protoGenToolConfPath, 0755) // 0755 permissions: owner has read/write/execute, group and others have read/execute.
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	// Create the configuration file.
	file, err := os.Create(fmt.Sprintf("%s/%s", protoGenToolConfPath, protoGenToolConfName))
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the configuration data into YAML format and write it to the file.
	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(p)
	if err != nil {
		return err
	}
	return nil
}
