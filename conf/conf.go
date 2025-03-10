/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 21:48:52
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 22:26:31
 * @FilePath: /tool/conf/conf.go
 * @Description: 配置文件管理模块，提供统一的配置加载和管理接口。
 * 支持多种配置类型（如路径配置、YAML配置等），通过工厂模式创建配置对象。
 */
package conf

// Conf 定义配置接口，所有配置类型必须实现该接口
type Conf interface {
	// SetConfig 加载并设置配置
	SetConfig() error
}

/**
 * @description: CreateConf 工厂函数，用于创建配置对象
 * @param {Conf} conf 实现了 Conf 接口的配置对象
 * @return {Conf} 返回传入的配置对象
 * @author: Jerry.Yang
 * @date: 2023-12-22 16:29:03
 */
func CreateConf(conf Conf) Conf {
	return conf
}

/**
 * @description: CreatePathConf 创建路径配置对象
 * @param {string} configPath 配置文件的路径
 * @return {Conf} 返回路径配置对象
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:51:21
 */
func CreatePathConf(configPath string) Conf {
	return CreateConf(&path{configPath: configPath})
}

/**
 * @description: CreateConfigPathConf 创建默认路径配置对象
 * @return {Conf} 返回默认路径配置对象
 * @author: Jerry.Yang
 * @date: 2023-12-22 15:51:30
 */
func CreateConfigPathConf() Conf {
	return CreateConf(&configPath{})
}

/**
 * @description: GetPathConf 获取全局路径配置对象
 * @return {*path} 返回路径配置对象的指针
 * @author: Jerry.Yang
 * @date: 2023-12-26 14:24:47
 */
func GetPathConf() *path {
	return pathConfig
}

/**
 * @description: CreateYamlConf 创建 YAML 配置对象
 * @param {string} fileName YAML 配置文件的名称
 * @param {interface{}} confData 配置数据的结构体指针
 * @return {Conf} 返回 YAML 配置对象
 * @author: Jerry.Yang
 * @date: 2023-12-22 10:54:45
 */
func CreateYamlConf(fileName string, confData interface{}) Conf {
	return CreateConf(&yamlConf{
		filePath: pathConfig.configPath,
		fileName: fileName,
		fileType: "yaml",
		confData: confData,
	})
}
