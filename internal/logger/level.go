/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-10 14:36:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-10 17:10:29
 * @Description: level
 */
package logger

import (
	"github.com/yangjerry110/tool/internal/errors"
)

/**
 * @description: Level
 * @author: Jerry.Yang
 * @date: 2024-04-10 16:41:35
 * @return {*}
 */
type Level string

/**
 * @description: level conf
 * @author: Jerry.Yang
 * @date: 2024-04-10 16:41:23
 * @return {*}
 */
var (
	// PanicLevel
	PanicLevel Level = "panic"
	// FatalLevel
	FatalLevel Level = "fatal"
	// ErrorLevel
	ErrorLevel Level = "error"
	// WarnLevel
	WarnLevel Level = "warn"
	// InfoLevel
	InfoLevel Level = "info"
	// DebugLevel
	DebugLevel Level = "debug"
	// TraceLevel
	TraceLevel Level = "trace"
	// UnkownLevel
	UnknowLevel Level = "unknow"
)

// LevelString isValid
//
// IsValid
// Date 2024-04-10 15:10:41
// Author Jerry.Yang
func (l Level) IsValid() error {

	// get Level
	level := l.Level()

	// if == unkownLevel
	// return err
	if level == UnknowLevel {
		return errors.ErrLoggerLevelUnknowLevel
	}
	return nil
}

// LevelString Level
//
// Level
// Date 2024-04-10 14:45:36
// Author Jerry.Yang
func (l Level) Level() Level {
	switch l {
	case PanicLevel:
		return PanicLevel
	case FatalLevel:
		return FatalLevel
	case ErrorLevel:
		return ErrorLevel
	case WarnLevel:
		return WarnLevel
	case InfoLevel:
		return InfoLevel
	case DebugLevel:
		return DebugLevel
	case TraceLevel:
		return TraceLevel
	// if no case
	// return debug Level
	default:
		return UnknowLevel
	}
}

// level string
//
// String
// Date 2024-04-10 17:10:27
// Author Jerry.Yang
func (l Level) String() string {
	switch l {
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	case ErrorLevel:
		return "error"
	case WarnLevel:
		return "warn"
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	case TraceLevel:
		return "trace"
	// if no case
	// return debug Level
	default:
		return "unknow"
	}
}
