/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-30 09:59:28
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-09 18:44:49
 * @Description: logger option
 */
package logger

import (
	"io"

	"github.com/yangjerry110/tool/logger"
)

type LoggerOptionPkgInterface interface {
	SetLevel(logLevel Level) logger.LoggerOptionFunc
	SetWithFields(fields map[string]interface{}) logger.LoggerOptionFunc
	SetIsReportcaller(isOpen bool) logger.LoggerOptionFunc
	SetOutput(output io.Writer) logger.LoggerOptionFunc
	SetFormatter(formatter string) logger.LoggerOptionFunc
	SetFormatterDisableTime(isOpen bool) logger.LoggerOptionFunc
	SetFormatterDisableHtmlEscap(isOpen bool) logger.LoggerOptionFunc
	SetCallerDept(dept int) logger.LoggerOptionFunc
}

type LoggerOptionPkg struct {
	LoggerOptionInterface    logger.LoggerOptionInterface
	LoggerOptionPkgInterface LoggerOptionPkgInterface
}

var defaultLoggOption = CreateLoggerOption(&LogrusOptionsPkg{})

/**
 * @description: CreateLoggerOptionPkgInterface
 * @param {LoggerOptionPkgInterface} loggerOptionPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:39:13
 * @return {*}
 */
func CreateLoggerOptionPkgInterface(loggerOptionPkgInterface LoggerOptionPkgInterface) *LoggerOptionPkg {
	return &LoggerOptionPkg{LoggerOptionPkgInterface: loggerOptionPkgInterface}
}

/**
 * @description: CreateLoggerOptionInterface
 * @param {logger.LoggerOptionInterface} loggerOptionInterface
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:39:20
 * @return {*}
 */
func CreateLoggerOptionInterface(loggerOptionInterface logger.LoggerOptionInterface) *LoggerOptionPkg {
	return &LoggerOptionPkg{LoggerOptionInterface: loggerOptionInterface}
}

/**
 * @description: CreateLoggerOption
 * @param {LoggerOptionPkgInterface} loggerOptionPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:40:56
 * @return {*}
 */
func CreateLoggerOption(loggerOptionPkgInterface LoggerOptionPkgInterface) LoggerOptionPkgInterface {
	return loggerOptionPkgInterface
}

/**
 * @description: SetLevel
 * @param {Level} logLevel
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:44:54
 * @return {*}
 */
func SetLevel(logLevel Level) logger.LoggerOptionFunc {
	return defaultLoggOption.SetLevel(logLevel)
}

/**
 * @description: SetWithFields
 * @param {map[string]interface{}} fields
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:45:01
 * @return {*}
 */
func SetWithFields(fields map[string]interface{}) logger.LoggerOptionFunc {
	return defaultLoggOption.SetWithFields(fields)
}

/**
 * @description: SetIsReportcaller
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:45:11
 * @return {*}
 */
func SetIsReportcaller(isOpen bool) logger.LoggerOptionFunc {
	return defaultLoggOption.SetIsReportcaller(isOpen)
}

/**
 * @description: SetOutput
 * @param {io.Writer} output
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:45:19
 * @return {*}
 */
func SetOutput(output io.Writer) logger.LoggerOptionFunc {
	return defaultLoggOption.SetOutput(output)
}

/**
 * @description: SetFormatter
 * @param {string} formatter
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:45:28
 * @return {*}
 */
func SetFormatter(formatter string) logger.LoggerOptionFunc {
	return defaultLoggOption.SetFormatter(formatter)
}

/**
 * @description: SetFormatterDisableTime
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:45:36
 * @return {*}
 */
func SetFormatterDisableTime(isOpen bool) logger.LoggerOptionFunc {
	return defaultLoggOption.SetFormatterDisableTime(isOpen)
}

/**
 * @description: SetFormatterDisableHtmlEscap
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:45:44
 * @return {*}
 */
func SetFormatterDisableHtmlEscap(isOpen bool) logger.LoggerOptionFunc {
	return defaultLoggOption.SetFormatterDisableHtmlEscap(isOpen)
}

/**
 * @description: SetCallDept
 * @param {int} dept
 * @author: Jerry.Yang
 * @date: 2022-10-09 18:44:54
 * @return {*}
 */
func SetCallerDept(dept int) logger.LoggerOptionFunc {
	return defaultLoggOption.SetCallerDept(dept)
}
