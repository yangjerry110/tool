/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 15:05:10
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-09 18:41:02
 * @Description: logrus log option
 */
package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type LoggerOptionInterface interface {
	SetOptions(options []LoggerOptionFunc) LoggerOptionInterface
	SetLevel(logLevel Level) LoggerOptionFunc
	SetWithFields(fields map[string]interface{}) LoggerOptionFunc
	SetIsReportcaller(isOpen bool) LoggerOptionFunc
	SetOutput(output io.Writer) LoggerOptionFunc
	SetFormatter(formatter string) LoggerOptionFunc
	SetFormatterDisableTime(isOpen bool) LoggerOptionFunc
	SetFormatterDisableHtmlEscap(isOpen bool) LoggerOptionFunc
	SetCallDept(dept int) LoggerOptionFunc
}

/**
 * @description: LogrusLogOption
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:25:27
 * @return {*}
 */
type LoggerOption struct {
	Options map[string]LoggerOptionVal
}

/**
 * @description: LoggerOptionVal option相对应的数据
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:25:34
 * @return {*}
 */
type LoggerOptionVal struct{ Value interface{} }

/**
 * @description: LoggerOptionFunc 设置option的方法
 * @param {*} map
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:25:56
 * @return {*}
 */
type LoggerOptionFunc func(map[string]LoggerOptionVal) error

/**
 * @description: 定义option的key
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:11:16
 * @return {*}
 */
var LOGRUS_ERROR_KEY = "errorKey"
var OPTION_LOGLEVEL = "log_level"
var OPTION_IS_SET_LEVEL = "is_set_log_level"
var OPTION_IS_SET_LEVEL_STATUS = false
var OPTION_WITH_FIELDS = "log_with_fields"
var OPTION_ISREPORTCALLER = "is_reportcaller"
var OPTION_OUTPUT = "logrus_output"
var OPTION_FORMATTER_TYPE = "formatter_type"
var OPTION_FORMATTER_DISABLETIME = "formatter_disabletime"
var OPTION_FORMATTER_DISABLEHTMLESCAP = "formatter_disablehtmlescap"
var OPTION_CALLER_DEPT = "log_caller_dept"

/**
 * @description: logrus 所有的日志级别
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:14:04
 * @return {*}
 */
var LOGRUS_ALL_LEVEL = map[Level]logrus.Level{
	Level(PanicLevel): logrus.PanicLevel,
	Level(FatalLevel): logrus.FatalLevel,
	Level(ErrorLevel): logrus.ErrorLevel,
	Level(WarnLevel):  logrus.WarnLevel,
	Level(InfoLevel):  logrus.InfoLevel,
	Level(DebugLevel): logrus.DebugLevel,
	Level(TraceLevel): logrus.TraceLevel,
}

/**
 * @description: logrus 日志格式定义
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:12:04
 * @return {*}
 */
var LOGRUS_FORMATTER_JSON = "json"
var LOGRUS_FORMATTER_TEXT = "text"
var LOGRUS_FORMATTER_ALL = map[string]string{
	LOGRUS_FORMATTER_JSON: LOGRUS_FORMATTER_JSON,
	LOGRUS_FORMATTER_TEXT: LOGRUS_FORMATTER_TEXT,
}

/**
 * @description: SetLogrusLogOptions
 * @param {[]LoggerOptionFunc} optionFuncs
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:36:38
 * @return {*}
 */
func (l *LoggerOption) SetOptions(options map[string]LoggerOptionVal) *LoggerOption {
	loggerOption := &LoggerOption{
		Options: options,
	}
	return loggerOption
}
