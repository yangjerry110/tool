/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:31:43
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 22:27:18
 * @Description: 配置文件路径管理模块，提供路径配置的加载和设置功能。
 * 支持从命令行参数或默认路径加载配置路径，并确保路径的有效性。
 */
package conf

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/toolerrors"
)

/**
 * @description: path 结构体，用于存储配置路径
 * @field configPath string 配置文件的路径
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:30:41
 */
type path struct {
	configPath string
}

/**
 * @description: configPath 结构体，用于默认路径配置
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:30:57
 */
type configPath struct{}

/**
 * @description: pathConfig 全局变量，存储当前的路径配置
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:31:30
 */
var pathConfig = &path{}

/**
 * @description: SetConfig 实现 Conf 接口，用于设置路径配置
 * @receiver p *path 路径配置对象
 * @return error 如果路径为空，返回错误；否则返回 nil
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:34:06
 */
func (p *path) SetConfig() error {
	// 检查路径是否为空
	if p.configPath == "" {
		return toolerrors.New("conf Err : path config is not configPath")
	}

	// 设置全局路径配置
	pathConfig = p
	return nil
}

/**
 * @description: SetConfig 实现 Conf 接口，用于设置默认路径配置
 * @receiver c *configPath 默认路径配置对象
 * @return error 如果路径已设置，直接返回；否则从命令行参数或默认路径加载配置路径
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:31:26
 */
func (c *configPath) SetConfig() error {
	// 如果路径已设置，直接返回
	if pathConfig.configPath != "" {
		return nil
	}

	// 获取命令行参数
	args := os.Args

	// 如果命令行参数数量大于等于 2，使用第二个参数作为配置路径
	if len(args) >= 2 {
		firstArg := args[1] // 第一个参数通常是程序名称，第二个参数是配置路径

		// 设置路径配置
		if err := CreateConf(&path{configPath: firstArg}).SetConfig(); err != nil {
			return err
		}
		return nil
	}

	// 如果命令行参数不足，使用默认路径
	currentPath, err := os.Getwd() // 获取当前工作目录
	if err != nil {
		return err
	}

	// 拼接默认配置路径
	configPath := fmt.Sprintf("%s/internal/config/yamlConfig", currentPath)

	// 设置路径配置
	if err := CreateConf(&path{configPath: configPath}).SetConfig(); err != nil {
		return err
	}
	return nil
}
