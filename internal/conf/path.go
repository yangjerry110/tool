/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:31:43
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-01-30 15:49:06
 * @Description:
 */
package conf

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/internal/errors"
)

/**
 * @description: Path
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:30:41
 * @return {*}
 */
type Path struct {
	ConfigPath string
}

/**
 * @description: configPath
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:30:57
 * @return {*}
 */
type ConfigPath struct{}

/**
 * @description: PathConfig
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:31:30
 * @return {*}
 */
var PathConfig = &Path{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:34:06
 * @return {*}
 */
func (p *Path) SetConfig() error {

	// judge path.ConfigPath is empty
	// if empty => err
	if p.ConfigPath == "" {
		return errors.ErrPathConfigIsNotConfigPath
	}

	// fmt.Sprint
	fmt.Printf("PathConf : %+v", p)
	fmt.Print("\r\n")

	// set
	PathConfig = p
	return nil
}

/**
 * @description: ConfigPath
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:31:26
 * @return {*}
 */
func (c *ConfigPath) SetConfig() error {

	// fmt.Sprint
	fmt.Printf("PathConfig : %+v", PathConfig)
	fmt.Print("\r\n")

	// If ConfigPath != ""
	// Return
	if PathConfig.ConfigPath != "" {
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

		// fmt.Sprint
		fmt.Printf("configPath : %+v", firstArg)
		fmt.Print("\r\n")

		// Set path.ConfigPath
		if err := CreateConf(&Path{ConfigPath: firstArg}).SetConfig(); err != nil {
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

	// fmt.Sprint
	fmt.Printf("configPath : %+v", configPath)
	fmt.Print("\r\n")

	// Set path.ConfigPath
	if err := CreateConf(&Path{ConfigPath: configPath}).SetConfig(); err != nil {
		return err
	}
	return nil
}
