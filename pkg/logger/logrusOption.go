/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-30 10:05:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-09 18:45:33
 * @Description: logrus
 */
package logger

import (
	"io"

	"github.com/yangjerry110/tool/logger"
)

type LogrusOptionsPkg struct{}

/**
 * @description: SetLevel
 * @param {Level} level
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:02
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetLevel(level Level) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetLevel(logger.Level(level))
}

/**
 * @description: SetWithFields
 * @param {map[string]interface{}} fields
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:15
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetWithFields(fields map[string]interface{}) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetWithFields(fields)
}

/**
 * @description: SetIsReportcaller
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:23
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetIsReportcaller(isOpen bool) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetIsReportcaller(isOpen)
}

/**
 * @description: SetOutput
 * @param {io.Writer} output
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:31
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetOutput(output io.Writer) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetOutput(output)
}

/**
 * @description: SetFormatter
 * @param {string} format
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:39
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetFormatter(format string) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetFormatter(format)
}

/**
 * @description: SetFormatterDisableHtmlEscap
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:47
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetFormatterDisableHtmlEscap(isOpen bool) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetFormatterDisableHtmlEscap(isOpen)
}

/**
 * @description: SetFormatterDisableTime
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:08:54
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetFormatterDisableTime(isOpen bool) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetFormatterDisableTime(isOpen)
}

/**
 * @description: SetCallDept
 * @param {int} dept
 * @author: Jerry.Yang
 * @date: 2022-10-09 18:44:17
 * @return {*}
 */
func (l *LogrusOptionsPkg) SetCallerDept(dept int) logger.LoggerOptionFunc {
	return CreateLoggerOptionInterface(&logger.LogrusOption{}).LoggerOptionInterface.SetCallDept(dept)
}
