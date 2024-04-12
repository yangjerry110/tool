/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-11 10:37:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 15:12:30
 * @Description: logger
 */
package logger

import (
	"io"

	"github.com/yangjerry110/tool/internal/logger"
)

// 记录 TraceLevel 级别的日志
func Trace(args ...interface{}) {
	log().Trace(args...)
}

// 格式化并记录 TraceLevel 级别的日志
func Tracef(format string, args ...interface{}) {
	log().Tracef(format, args...)
}

// 记录 DebugLevel 级别的日志
func Debug(args ...interface{}) {
	log().Debug(args...)
}

// 格式化并记录 DebugLevel 级别的日志
func Debugf(format string, args ...interface{}) {
	log().Debugf(format, args...)
}

// 记录 InfoLevel 级别的日志
func Info(args ...interface{}) {
	log().Info(args...)
}

// 格式化并记录 InfoLevel 级别的日志
func Infof(format string, args ...interface{}) {
	log().Infof(format, args...)
}

// 记录 InfoLevel 级别的日志[gorm logger扩展]
func Print(args ...interface{}) {
	log().Print(args...)
}

// 格式化并记录 InfoLevel 级别的日志[gorm logger扩展]
func Printf(format string, args ...interface{}) {
	log().Printf(format, args...)
}

// 记录 WarnLevel 级别的日志
func Warn(args ...interface{}) {
	log().Warn(args...)
}

// 格式化并记录 WarnLevel 级别的日志
func Warnf(format string, args ...interface{}) {
	log().Warnf(format, args...)
}

// 记录 ErrorLevel 级别的日志
func Error(args ...interface{}) {
	log().Error(args...)
}

// 格式化并记录 ErrorLevel 级别的日志
func Errorf(format string, args ...interface{}) {
	log().Errorf(format, args...)
}

// 记录 FatalLevel 级别的日志
func Fatal(args ...interface{}) {
	log().Fatal(args...)
}

// 格式化并记录 FatalLevel 级别的日志
func Fatalf(format string, args ...interface{}) {
	log().Fatalf(format, args...)
}

// 记录 PanicLevel 级别的日志
func Panic(args ...interface{}) {
	log().Panic(args...)
}

// 格式化并记录 PanicLevel 级别的日志
func Panicf(format string, args ...interface{}) {
	log().Panicf(format, args...)
}

// 设置level
func SetLevel(level logger.Level) error {
	return log().SetLevel(level)
}

// 获取level
func GetLevel() logger.Level {
	return log().GetLevel()
}

// 设置日志的格式
func SetFormatter(formatter logger.Formatter) error {
	return log().SetFormatter(formatter)
}

// 获取日志的格式
func GetFormatter() logger.Formatter {
	return log().GetFormatter()
}

// 设置需要打印函数行数等等
func SetReportCaller(ReportCaller bool) error {
	return log().SetReportCaller(ReportCaller)
}

// 获取设置的打印函数等等
func GetReportCaller() bool {
	return log().GetReportCaller()
}

// 设置html是否需要转义
func SetEnableHTMLEscape(enableHTMLEscape bool) error {
	return log().SetEnableHTMLEscape(enableHTMLEscape)
}

// 获取设置的是否需要html转义
func GetEnableHTMLEscape() bool {
	return log().GetEnableHTMLEscape()
}

// 设置输出形式
func SetOutput(output io.Writer) error {
	return log().SetOutput(output)
}

// 获取输出形式
func GetOutput() io.Writer {
	return log().GetOutput()
}

// 获取错误
func GetErr() error {
	return log().GetErr()
}
