/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-10 14:28:43
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 10:28:26
 * @Description: logrus
 */
package logrus

import (
	"fmt"
	"io"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/yangjerry110/tool/logger/internal/logger"
)

type Logrus struct {
	*logrus.Entry
	logger           *logrus.Logger
	reportCaller     bool
	enableHTMLEscape bool
	err              error
}

// logrus new
//
// New
// Date 2024-04-11 10:25:57
// Author Jerry.Yang
func (l *Logrus) Init() logger.LoggerInterface {

	// new logrus
	l.logger = logrus.New()

	//  set no lock
	l.logger.SetNoLock()

	// set entry
	l.Entry = logrus.NewEntry(l.logger)
	return l
}

// logrus trace
//
// Trace
// Date 2024-04-12 14:26:41
// Author Jerry.Yang
func (l *Logrus) Trace(args ...interface{}) {
	l.logWithStack(3, logger.TraceLevel, "", args...)
}

// logrus tracef
//
// Tracef
// Date 2024-04-12 14:26:41
// Author Jerry.Yang
func (l *Logrus) Tracef(format string, args ...interface{}) {
	l.logWithStack(3, logger.TraceLevel, format, args...)
}

// logrus Debug
//
// Debug
// Date 2024-04-11 10:33:44
// Author Jerry.Yang
func (l *Logrus) Debug(args ...interface{}) {
	l.logWithStack(3, logger.DebugLevel, "", args...)
}

// logrus Debugf
//
// Debugf
// Date 2024-04-11 10:35:02
// Author Jerry.Yang
func (l *Logrus) Debugf(format string, args ...interface{}) {
	l.logWithStack(3, logger.DebugLevel, format, args...)
}

// logrus info
//
// Info
// Date 2024-04-11 10:33:44
// Author Jerry.Yang
func (l *Logrus) Info(args ...interface{}) {
	l.logWithStack(3, logger.InfoLevel, "", args...)
}

// logrus Infof
//
// Infof
// Date 2024-04-11 10:35:02
// Author Jerry.Yang
func (l *Logrus) Infof(format string, args ...interface{}) {
	l.logWithStack(3, logger.InfoLevel, format, args...)
}

// logrus Print
//
// Print
// Date 2024-04-11 10:33:44
// Author Jerry.Yang
func (l *Logrus) Print(args ...interface{}) {
	l.logWithStack(3, logger.InfoLevel, "", args...)
}

// logrus Printf
//
// Printf
// Date 2024-04-11 10:35:02
// Author Jerry.Yang
func (l *Logrus) Printf(format string, args ...interface{}) {
	l.logWithStack(3, logger.InfoLevel, format, args...)
}

// logrus Error
//
// Error
// Date 2024-04-11 10:33:44
// Author Jerry.Yang
func (l *Logrus) Error(args ...interface{}) {
	l.logWithStack(3, logger.ErrorLevel, "", args...)
}

// logrus Errorf
//
// Errorf
// Date 2024-04-11 10:35:02
// Author Jerry.Yang
func (l *Logrus) Errorf(format string, args ...interface{}) {
	l.logWithStack(3, logger.ErrorLevel, format, args...)
}

// logrus Warn
//
// Warn
// Date 2024-04-11 10:33:44
// Author Jerry.Yang
func (l *Logrus) Warn(args ...interface{}) {
	l.logWithStack(3, logger.WarnLevel, "", args...)
}

// logrus Warnf
//
// Warnf
// Date 2024-04-11 10:35:02
// Author Jerry.Yang
func (l *Logrus) Warnf(format string, args ...interface{}) {
	l.logWithStack(3, logger.WarnLevel, format, args...)
}

// logrus Fatal
//
// Fatal
// Date 2024-04-11 10:33:44
// Author Jerry.Yang
func (l *Logrus) Fatal(args ...interface{}) {
	l.logWithStack(3, logger.FatalLevel, "", args...)
}

// logrus Fatalf
//
// Fatalf
// Date 2024-04-11 10:35:02
// Author Jerry.Yang
func (l *Logrus) Fatalf(format string, args ...interface{}) {
	l.logWithStack(3, logger.FatalLevel, format, args...)
}

// logrus Panic
//
// Panic
// Date 2024-04-11 10:33:44
// Author Jerry.Yang
func (l *Logrus) Panic(args ...interface{}) {
	l.logWithStack(3, logger.PanicLevel, "", args...)
}

// logrus Panicf
//
// Panicf
// Date 2024-04-11 10:35:02
// Author Jerry.Yang
func (l *Logrus) Panicf(format string, args ...interface{}) {
	l.logWithStack(3, logger.PanicLevel, format, args...)
}

// logrus level
//
// SetLevel
// Date 2024-04-10 16:33:19
// Author Jerry.Yang
func (l *Logrus) SetLevel(level logger.Level) error {

	// get logrus level
	logrusLevel, err := l.getLogrusLevel(level)
	if err != nil {
		return err
	}

	// set logrus level
	l.Entry.Logger.SetLevel(logrusLevel)
	return nil
}

// logrus GetLevel
//
// GetLevel
// Date 2024-04-10 16:56:12
// AUthor Jerry.Yang
func (l *Logrus) GetLevel() logger.Level {
	return logger.Level(l.Logger.GetLevel().String())
}

// logrus SetFormatter
//
// SetFormatter
// Date 2024-04-10 17:53:24
// Author Jerry.Yang
func (l *Logrus) SetFormatter(formatter logger.Formatter) error {

	// formatter isValid
	if err := formatter.IsValid(); err != nil {
		return err
	}

	// set formatter
	if formatter == logger.JsonFormatter {
		l.Entry.Logger.SetFormatter(&logrus.JSONFormatter{DisableHTMLEscape: !l.enableHTMLEscape})
	}

	// conf
	LogrusOptionConf.Formatter = formatter
	return nil
}

// logrus GetFormatter
//
// Date 2024-04-10 17:55:54
// Author Jerry.Yang
func (l *Logrus) GetFormatter() logger.Formatter {
	return LogrusOptionConf.Formatter
}

// logrus SetReportCaller
//
// SetReportCall
// Date 2024-04-12 15:04:31
// Author Jerry.Yang
func (l *Logrus) SetReportCaller(ReportCaller bool) error {
	l.reportCaller = ReportCaller
	return nil
}

// logrus GetReportCaller
//
// GetReportCaller
// Date 2024-04-12 15:04:52
// Author Jerry.Yang
func (l *Logrus) GetReportCaller() bool {
	return l.reportCaller
}

// logrus SetEnableHTMLEscape
//
// SetEnableHTMLEscape
// Date 2024-04-12 15:06:05
// Author Jerry.Yang
func (l *Logrus) SetEnableHTMLEscape(enableHTMLEscape bool) error {
	l.enableHTMLEscape = enableHTMLEscape
	return nil
}

// logrus GetEnableHTMLEscape
//
// GetEnableHTMLEscape
// Date 2024-04-12 15:06:22
// Author Jerry.Yang
func (l *Logrus) GetEnableHTMLEscape() bool {
	return l.enableHTMLEscape
}

// logrus setOutput
//
// setOutput
// Date 2024-04-10 15:28:18
// Author Jerry.Yang
func (l *Logrus) SetOutput(output io.Writer) error {
	l.Entry.Logger.Out = output
	return nil
}

// logrus GetOutput
//
// GetOutput
// Date 2024-04-10 15:35:57
// Author Jerry.Yang
func (l *Logrus) GetOutput() io.Writer {
	return l.Logger.Out
}

// logrus GetErr
//
// GetErr
// Date 2024-04-10 17:39:40
// Author Jerry.Yang
func (l Logrus) GetErr() error {
	return l.err
}

// logrus log with stack
//
// logWithStack
// Date 2024-04-10 17:15:34
// Author Jerry.Yang
func (l *Logrus) logWithStack(depth int, level logger.Level, format string, args ...interface{}) error {

	// if err != nil
	if l.err != nil {
		fmt.Printf("logrus Err : %+v", l.err)
		fmt.Print("\r\n")
		return l.err
	}

	// if as Set Level
	setLogrusLevel := l.logger.GetLevel()
	// get logrusLevel by level
	logrusLevel, err := l.getLogrusLevel(level)
	if err != nil {
		return err
	}

	// Check whether the level of the current logrus is smaller than the level to be recorded
	// If the set log level is smaller than the required log level, no log is recorded
	if setLogrusLevel < logrusLevel {
		return nil
	}

	// setRuntimeCaller && ReportCaller
	if l.reportCaller {
		// l.SetReportCaller(LogrusOptionConf.ReportCaller)
		if err := l.setRuntimeCaller(depth); err != nil {
			return err
		}
	}

	// set log
	// if format == ""
	if format == "" {
		l.Entry.Log(logrusLevel, args...)
		return nil
	}

	// set log
	// if formatr != ""
	l.Entry.Logf(logrusLevel, format, args...)
	return nil
}

// logrus setRuntimeCaller
//
// setRuntimeCaller
// Date 2024-04-10 17:32:20
// Author Jerry.Yang
func (l *Logrus) setRuntimeCaller(depth int) error {

	// get runtime call
	_, file, lineNo, ok := runtime.Caller(1 + depth)
	if !ok {
		return nil
	}

	// set withFields
	l.Entry = l.WithField("file", fmt.Sprintf("%s:%d", file, lineNo))
	return nil
}

// get logrus level
//
// getLogrusLevel
// Date 2024-04-10 17:12:39
// Author Jerry.Yang
func (l *Logrus) getLogrusLevel(level logger.Level) (logrus.Level, error) {

	// logrus Level
	var logrusLevel logrus.Level

	// if level isValid
	if err := level.IsValid(); err != nil {
		return logrusLevel, err
	}

	// get level
	return logrus.ParseLevel(level.String())
}
