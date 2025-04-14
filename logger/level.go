/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-12 14:42:28
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 10:24:11
 * @Description: level
 */
package logger

import "github.com/yangjerry110/tool/logger/internal/logger"

/**
 * @description: level conf
 * @author: Jerry.Yang
 * @date: 2024-04-10 16:41:23
 * @return {*}
 */
var (
	// PanicLevel
	PanicLevel = logger.PanicLevel
	// FatalLevel
	FatalLevel = logger.FatalLevel
	// ErrorLevel
	ErrorLevel = logger.ErrorLevel
	// WarnLevel
	WarnLevel = logger.WarnLevel
	// InfoLevel
	InfoLevel = logger.InfoLevel
	// DebugLevel
	DebugLevel = logger.DebugLevel
	// TraceLevel
	TraceLevel = logger.TraceLevel
	// UnkownLevel
	UnknowLevel = logger.UnknowLevel
)
