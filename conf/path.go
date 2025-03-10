/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:31:43
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 20:44:23
 * @Description:
 */
package conf

import (
	"fmt"
	"os"
)

/**
 * @description: Path
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:30:41
 * @return {*}
 */
type path struct {
	configPath string
}

/**
 * @description: configPath
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:30:57
 * @return {*}
 */
type configPath struct{}

/**
 * @description: PathConfig
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:31:30
 * @return {*}
 */
var pathConfig = &path{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:34:06
 * @return {*}
 */
func (p *path) SetConfig() error {

	// judge path.ConfigPath is empty
	// if empty => err
	if p.configPath == "" {
		return errors.ErrPathConfigIsNotConfigPath
	}

	// set
	pathConfig = p
	return nil
}

/**
 * @description: ConfigPath
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:31:26
 * @return {*}
 */
func (c *configPath) SetConfig() error {

	// If ConfigPath != ""
	// Return
	if pathConfig.configPath != "" {
		return nil
	}

	// Judge args
	// If len args != 0;
	// Set secend arg => path.ConfigPath
	// Else set currenPath => path.ConfigPath
	args := os.Args

	// Judge len args
	// If != 0
	if len(args) >= 2 {

		// Secend arg
		// For first is other
		firstArg := args[1]

		// Set path.ConfigPath
		if err := CreateConf(&path{configPath: firstArg}).SetConfig(); err != nil {
			return err
		}
		return nil
	}

	// If len args < 1
	// Get current Path + "internal/config/yamlConfig" => path.ConfigPath
	// Get currentPath
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	// Splicing configPath
	configPath := fmt.Sprintf("%s/internal/config/yamlConfig", currentPath)

	// Set path.ConfigPath
	if err := CreateConf(&path{configPath: configPath}).SetConfig(); err != nil {
		return err
	}
	return nil
}
