/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-10 14:27:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 10:27:14
 * @Description: logrus config
 */
package logrus

import (
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/logger/internal/logger"
)

/**
 * @description: LogrusOption
 * @author: Jerry.Yang
 * @date: 2024-04-10 15:20:35
 * @return {*}
 */
type LogrusOption struct {
	Level            logger.Level     `yaml:"level"`
	Formatter        logger.Formatter `yaml:"formatter"`
	EnableHTMLEscape bool             `yaml:"enable_html_escape"`
	ReportCaller     bool             `yaml:"report_caller"`
}

/**
 * @description: LogrusOptionConf
 * @author: Jerry.Yang
 * @date: 2024-04-10 14:28:58
 * @return {*}
 */
var LogrusOptionConf = &LogrusOption{}

// set logrus config
//
// SetConfig
// Date 2024-04-10 14:31:03
// Author Jerry.Yang
func (l *LogrusOption) SetConfig() error {
	return conf.CreateYamlConf("logger.yaml", &LogrusOptionConf).SetConfig()
}
