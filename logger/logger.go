/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 18:35:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-09 18:40:48
 * @Description: logger
 */
package logger

type LoggerInterface interface {
	WriteLog(args ...interface{}) error
	CheckParams() LoggerInterface
	SetLogger() LoggerInterface
	SetLevel() LoggerInterface
	SetCaller() LoggerInterface
	SetWithFields() LoggerInterface
	SetFormatter() LoggerInterface
	SetOutput() LoggerInterface
	SetIsReportcaller() LoggerInterface
	SetCallDept() LoggerInterface
}

type Logger struct{}

// Level type
type Level uint32

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)
