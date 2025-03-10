/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 16:24:23
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 22:29:18
 * @Description: YAML 配置文件管理模块，提供 YAML 配置文件的加载和热更新功能。
 * 支持通过 `yaml.v3` 解析 YAML 文件，并使用 `watch` 模块监听文件变化，实现配置的热更新。
 */
package conf

import (
	"fmt"
	"io/ioutil"

	"github.com/yangjerry110/tool/toolerrors"
	"gopkg.in/yaml.v3"
)

// yamlConf 结构体，用于管理 YAML 配置文件
type yamlConf struct {
	filePath string      // 文件路径
	fileName string      // 文件名称
	fileType string      // 文件类型（固定为 "yaml"）
	confData interface{} // 配置文件的数据结构
}

/**
 * @description: SetConfig 实现 Conf 接口，用于加载 YAML 配置文件并启动监听
 * @receiver y *yamlConf YAML 配置对象
 * @return error 如果文件路径、文件名或配置数据结构为空，返回错误；否则加载配置并启动监听
 * @author: Jerry.Yang
 * @date: 2023-12-08 17:35:25
 */
func (y *yamlConf) SetConfig() error {
	// 检查文件路径是否为空
	if y.filePath == "" {
		return toolerrors.New("conf Err : yaml conf is not filePath")
	}

	// 检查文件名是否为空
	if y.fileName == "" {
		return toolerrors.New("conf Err : yaml conf is not FileName")
	}

	// 检查配置数据结构是否为空
	if y.confData == nil {
		return toolerrors.New("conf Err : yaml conf is not ConfData")
	}

	// 热更新 YAML 配置的三种实现思路：
	// (1) 使用信号（Signal）触发更新，但信号仅适用于用户连接和行为，不确定是否可行。
	// (2) 使用文件的最后修改时间，这种情况下需要为每个配置文件启动一个协程，成本较高。
	// (3) 使用 viper，viper 的实现对协程数量和性能有较好的控制，但当前不采用 viper。
	// 最终选择使用 fsnotify 监听文件变化。

	// 读取文件内容
	fileContent, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", y.filePath, y.fileName))
	if err != nil {
		return err
	}

	// 解析 YAML 文件内容到配置数据结构
	if err := yaml.Unmarshal(fileContent, y.confData); err != nil {
		return err
	}

	// 启动文件监听，实现热更新
	if err := CreateConf(&watch{
		watchfile: &watchFile{
			filePath: y.filePath,
			fileName: y.fileName,
			fileType: y.fileType,
			confData: y.confData,
		},
	}).SetConfig(); err != nil {
		return err
	}
	return nil
}
