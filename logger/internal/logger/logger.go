/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-10 11:25:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 15:12:10
 * @Description: logger
 */
package logger

import (
	"io"
)

type LoggerInterface interface {
	Init() LoggerInterface                           // 初始化
	Trace(args ...interface{})                       // 记录 TraceLevel 级别的日志
	Tracef(format string, args ...interface{})       // 格式化并记录 TraceLevel 级别的日志
	Debug(args ...interface{})                       // 记录 DebugLevel 级别的日志
	Debugf(format string, args ...interface{})       // 格式化并记录 DebugLevel 级别的日志
	Info(args ...interface{})                        // 记录 InfoLevel 级别的日志
	Infof(format string, args ...interface{})        // 格式化并记录 InfoLevel 级别的日志
	Print(args ...interface{})                       // 记录 InfoLevel 级别的日志[gorm logger扩展]
	Printf(format string, args ...interface{})       // 格式化并记录 InfoLevel 级别的日志[gorm logger扩展]
	Warn(args ...interface{})                        // 记录 WarnLevel 级别的日志
	Warnf(format string, args ...interface{})        // 格式化并记录 WarnLevel 级别的日志
	Error(args ...interface{})                       // 记录 ErrorLevel 级别的日志
	Errorf(format string, args ...interface{})       // 格式化并记录 ErrorLevel 级别的日志
	Fatal(args ...interface{})                       // 记录 FatalLevel 级别的日志
	Fatalf(format string, args ...interface{})       // 格式化并记录 FatalLevel 级别的日志
	Panic(args ...interface{})                       // 记录 PanicLevel 级别的日志
	Panicf(format string, args ...interface{})       // 格式化并记录 PanicLevel 级别的日志
	SetLevel(level Level) error                      // 设置level
	GetLevel() Level                                 // 获取level
	SetFormatter(formatter Formatter) error          // 设置日志的格式
	GetFormatter() Formatter                         // 获取日志的格式
	SetReportCaller(ReportCaller bool) error         // 设置需要打印函数行数等等
	GetReportCaller() bool                           // 获取设置的打印函数等等
	SetEnableHTMLEscape(enableHTMLEscape bool) error // 设置html是否需要转义
	GetEnableHTMLEscape() bool                       // 获取设置的是否需要html转义
	SetOutput(output io.Writer) error                // 设置输出形式
	GetOutput() io.Writer                            // 获取输出形式
	GetErr() error                                   // 获取错误
}
