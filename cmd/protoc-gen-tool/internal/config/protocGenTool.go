/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-12 11:22:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-05-16 14:12:07
 * @Description: protocGenTool
 */
package config

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/conf"
	"gopkg.in/yaml.v2"
)

type ProtocGenTool struct {
	IsFirstCreate bool
	IsAppend      bool
	IsExtend      bool
	ExtendPath    string
}

/**
 * @description: ProtocGenToolConf
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:23:52
 * @return {*}
 */
var ProtocGenToolConf = &ProtocGenTool{}

var protoGenToolConfPath = "/opt/protobuf/tool/"
var protoGenToolConfName = "protoGenToolConf.yaml"

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-12 11:26:39
 * @return {*}
 */
func (p *ProtocGenTool) SetConfig() error {
	if err := conf.CreatePathConf(protoGenToolConfPath).SetConfig(); err != nil {
		return err
	}

	return conf.CreateYamlConf(protoGenToolConfName, ProtocGenToolConf).SetConfig()
}

func (p *ProtocGenTool) SetConf() error {

	// 确保文件夹存在
	err := os.MkdirAll(protoGenToolConfPath, 0755) // 0755 表示权限设置，表示所有者有读写执行权限，组和其他用户有读执行权限
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	// 打开文件
	file, err := os.Create(fmt.Sprintf("%s/%s", protoGenToolConfPath, protoGenToolConfName))
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建编码器并保存数据
	encoder := yaml.NewEncoder(file)
	err = encoder.Encode(p)
	if err != nil {
		return err
	}
	return nil
}
