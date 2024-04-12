/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-12 14:58:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 15:16:14
 * @Description: default
 */
package logrus

import internalLogger "github.com/yangjerry110/tool/internal/logger"

/**
 * @description: defaultLogger
 * @author: Jerry.Yang
 * @date: 2024-04-12 15:15:57
 * @return {*}
 */
var defaultLogger internalLogger.LoggerInterface

// LogrusLogger
//
// Date 2024-04-12 15:15:46
// Author Jerry.Yang
func LogrusLogger() internalLogger.LoggerInterface {
	if defaultLogger == nil {
		createDefaultLogger()
	}
	return defaultLogger
}

// createDefaultLogger
//
// Date 2024-04-12 15:15:07
// Author Jerry.Yang
func createDefaultLogger() error {
	logrus := &Logrus{}
	logger := logrus.Init()
	logger.SetLevel(internalLogger.InfoLevel)
	logger.SetFormatter(internalLogger.JsonFormatter)
	logger.SetReportCaller(true)
	logger.SetEnableHTMLEscape(true)
	defaultLogger = logger
	return nil
}
