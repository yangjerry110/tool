/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-27 18:26:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-16 17:16:37
 * @Description: logger
 */
package logger

import (
	"github.com/yangjerry110/tool/logger"
)

type LoggerPkgInterface interface {
	CreatePkgInterface(loggerPkgInterface LoggerPkgInterface) *LoggerPkg
	CreateInterface(loggerInterface logger.LoggerInterface) *LoggerPkg
	CreateOptionInterface(loggerInterface logger.LoggerOptionInterface) *LoggerPkg
	CreateOptionPkgInterface(loggerOptionsPkgInterface LoggerOptionPkgInterface) *LoggerPkg
	SetOptions(options []logger.LoggerOptionFunc) LoggerPkgInterface
	SetDefaultOptions() LoggerPkgInterface
	SetLoggerOptions() LoggerPkgInterface
	Log(args ...interface{}) error                               // 记录对应级别的日志
	Logf(format string, args ...interface{}) error               // 记录对应级别的日志
	Info(args ...interface{}) error                              // 记录 InfoLevel 级别的日志
	Infof(format string, args ...interface{}) error              // 格式化并记录infof级别的日志
	Trace(args ...interface{}) error                             // 记录 TraceLevel 级别的日志
	Tracef(format string, args ...interface{}) error             // 格式化并记录 TraceLevel 级别的日志
	Debug(args ...interface{}) error                             // 记录 DebugLevel 级别的日志
	Debugf(format string, args ...interface{}) error             // 格式化并记录 DebugLevel 级别的日志
	Warn(args ...interface{}) error                              // 记录 WarnLevel 级别的日志
	Warnf(format string, args ...interface{}) error              // 格式化并记录 WarnLevel 级别的日志
	Error(args ...interface{}) error                             // 记录 ErrorLevel 级别的日志
	Errorf(format string, args ...interface{}) error             // 格式化并记录 ErrorLevel 级别的日志
	Fatal(args ...interface{}) error                             // 记录 FatalLevel 级别的日志
	Fatalf(format string, args ...interface{}) error             // 格式化并记录 FatalLevel 级别的日志
	Panic(args ...interface{}) error                             // 记录 PanicLevel 级别的日志
	Panicf(format string, args ...interface{}) error             // 格式化并记录 PanicLevel 级别的日志
	WithField(key string, value interface{}) LoggerPkgInterface  // 为日志添加一个上下文数据
	WithFields(fields map[string]interface{}) LoggerPkgInterface // 为日志添加多个上下文数据
	WithError(err error) LoggerPkgInterface                      // 为日志添加标准错误上下文数据
}

type LoggerPkg struct {
	DefaultOptionFuns        []logger.LoggerOptionFunc
	OptionFuns               []logger.LoggerOptionFunc
	Options                  map[string]logger.LoggerOptionVal
	LoggerPkgInterface       LoggerPkgInterface
	LoggerInterface          logger.LoggerInterface
	LoggerOptionInterface    logger.LoggerOptionInterface
	LoggerOptionPkgInterface LoggerOptionPkgInterface
}

// Level type
type Level uint32

/**
 * @description: CreateLogger
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:05
 * @return {*}
 */
var defaultLogger = CreateLogger(logrusLogPkg)

/**
 * @description: CreatePkgInterface
 * @param {LoggerPkgInterface} loggerPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:12
 * @return {*}
 */
func CreatePkgInterface(loggerPkgInterface LoggerPkgInterface) *LoggerPkg {
	return &LoggerPkg{LoggerPkgInterface: loggerPkgInterface}
}

/**
 * @description: CreateInterface
 * @param {logger.LoggerInterface} loggerInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:18
 * @return {*}
 */
func CreateInterface(loggerInterface logger.LoggerInterface) *LoggerPkg {
	return &LoggerPkg{LoggerInterface: loggerInterface}
}

/**
 * @description: CreateOptionInterface
 * @param {logger.LoggerOptionInterface} loggerOptionInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:25
 * @return {*}
 */
func CreateOptionInterface(loggerOptionInterface logger.LoggerOptionInterface) *LoggerPkg {
	return &LoggerPkg{LoggerOptionInterface: loggerOptionInterface}
}

/**
 * @description: CreateOptionPkgInterface
 * @param {LoggerOptionPkgInterface} loggerOptionPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-30 14:43:31
 * @return {*}
 */
func CreateOptionPkgInterface(loggerOptionPkgInterface LoggerOptionPkgInterface) *LoggerPkg {
	return &LoggerPkg{LoggerOptionPkgInterface: loggerOptionPkgInterface}
}

/**
 * @description: CreateLogger
 * @param {LoggerPkgInterface} loggerPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:31
 * @return {*}
 */
func CreateLogger(loggerPkgInterface LoggerPkgInterface) LoggerPkgInterface {
	return loggerPkgInterface
}

/**
 * @description: SetLogger
 * @param {LoggerPkgInterface} LoggerPkgInterface
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:12:38
 * @return {*}
 */
func SetLogger(LoggerPkgInterface LoggerPkgInterface) LoggerPkgInterface {
	defaultLogger = CreateLogger(LoggerPkgInterface)
	return LoggerPkgInterface
}

/**
 * @description: SetOptions
 * @param {[]logger.LoggerOptionFunc} options
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:49:02
 * @return {*}
 */
func SetOptions(options []logger.LoggerOptionFunc) LoggerPkgInterface {
	return defaultLogger.SetOptions(options)
}

/**
 * @description: Log
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:06
 * @return {*}
 */
func Log(args ...interface{}) {
	defaultLogger.Log(args...)
}

/**
 * @description: Logf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:13
 * @return {*}
 */
func Logf(format string, args ...interface{}) {
	defaultLogger.Logf(format, args...)
}

/**
 * @description: Trace
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:19
 * @return {*}
 */
func Trace(args ...interface{}) {
	defaultLogger.Trace(args...)
}

/**
 * @description: Tracef
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:26
 * @return {*}
 */
func Tracef(format string, args ...interface{}) {
	defaultLogger.Tracef(format, args...)
}

/**
 * @description: Debug
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:34
 * @return {*}
 */
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

/**
 * @description: Debugf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:41
 * @return {*}
 */
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

/**
 * @description: Info
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:47
 * @return {*}
 */
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

/**
 * @description: Infof
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:13:54
 * @return {*}
 */
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

/**
 * @description: Warn
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:01
 * @return {*}
 */
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

/**
 * @description: Warnf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:07
 * @return {*}
 */
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

/**
 * @description: Error
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:14
 * @return {*}
 */
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

/**
 * @description: Errorf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:21
 * @return {*}
 */
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

/**
 * @description: Fatal
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:28
 * @return {*}
 */
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

/**
 * @description: Fatalf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:35
 * @return {*}
 */
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

/**
 * @description: Panic
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:42
 * @return {*}
 */
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

/**
 * @description: Panicf
 * @param {string} format
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:49
 * @return {*}
 */
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

/**
 * @description: WithField
 * @param {string} key
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:14:56
 * @return {*}
 */
func WithField(key string, value interface{}) LoggerPkgInterface {
	return defaultLogger.WithField(key, value)
}

/**
 * @description: WithFields
 * @param {map[string]interface{}} fields
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:15:03
 * @return {*}
 */
func WithFields(fields map[string]interface{}) LoggerPkgInterface {
	return defaultLogger.WithFields(fields)
}

/**
 * @description: WithError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2022-09-29 19:15:11
 * @return {*}
 */
func WithError(err error) LoggerPkgInterface {
	return defaultLogger.WithError(err)
}
