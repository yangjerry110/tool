/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 14:42:36
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 16:05:45
 * @Description: base config
 */
package config

/**
 * @description: CreateBaseConfig
 * @param {...BaseConfig} BaseConfigs
 * @author: Jerry.Yang
 * @date: 2023-04-24 14:44:01
 * @return {*}
 */
func CreateBaseConfig(BaseConfigs ...BaseConfig) BaseConfig {
	if len(BaseConfigs) == 0 {
		return &Base{}
	}
	return BaseConfigs[0]
}

/**
 * @description: CreateLoggerConfig
 * @param {...LoggerConfig} LoggerConfigs
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:04:55
 * @return {*}
 */
func CreateLoggerConfig(LoggerConfigs ...LoggerConfig) LoggerConfig {
	if len(LoggerConfigs) == 0 {
		return &Logger{}
	}
	return LoggerConfigs[0]
}

/**
 * @description: CreatePathConfig
 * @param {...PathConfig} PathConfigs
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:05:04
 * @return {*}
 */
func CreatePathConfig(PathConfigs ...PathConfig) PathConfig {
	if len(PathConfigs) == 0 {
		return &Path{}
	}
	return PathConfigs[0]
}

/**
 * @description: CreateRouterConfig
 * @param {...RouterConfig} RouterConfigs
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:05:13
 * @return {*}
 */
func CreateRouterConfig(RouterConfigs ...RouterConfig) RouterConfig {
	if len(RouterConfigs) == 0 {
		return &Router{}
	}
	return RouterConfigs[0]
}

/**
 * @description: CreateLoggerYamlConfig
 * @param {...LoggerYamlConfig} LoggerYamlConfigs
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:05:52
 * @return {*}
 */
func CreateLoggerYamlConfig(LoggerYamlConfigs ...LoggerYamlConfig) LoggerYamlConfig {
	if len(LoggerYamlConfigs) == 0 {
		return &LoggerYaml{}
	}
	return LoggerYamlConfigs[0]
}

/**
 * @description: CreateRouterYamlConfig
 * @param {...RouterYamlConfig} RouterYamlConfigs
 * @author: Jerry.Yang
 * @date: 2023-04-24 16:06:35
 * @return {*}
 */
func CreateRouterYamlConfig(RouterYamlConfigs ...RouterYamlConfig) RouterYamlConfig {
	if len(RouterYamlConfigs) == 0 {
		return &RouterYaml{}
	}
	return RouterYamlConfigs[0]
}
