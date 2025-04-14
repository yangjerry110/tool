/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-10 15:02:11
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 10:26:08
 * @Description: formatter
 */
package logger

import "github.com/yangjerry110/tool/logger/internal/errors"

/**
 * @description: formatter
 * @author: Jerry.Yang
 * @date: 2024-04-10 15:00:40
 * @return {*}
 */
var JsonFormatter Formatter = "json"
var TextFormatter Formatter = "text"
var UnknowFormatter Formatter = "unkown"

/**
 * @description: Formatter
 * @author: Jerry.Yang
 * @date: 2024-04-10 15:08:19
 * @return {*}
 */
type Formatter string

// Formatter IsValid
//
// isValid
// Date 2024-04-10 15:08:03
// Author Jerry.Yang
func (f Formatter) IsValid() error {

	// get formatter
	formatter := f.Formatter()

	// if formatter == unkown
	if formatter == UnknowFormatter {
		return errors.ErrLoggerFormatterUnkownFormatter
	}
	return nil
}

// Formatter
//
// Formatter
// Date 2024-04-10 16:45:19
// Author Jerry.Yang
func (f Formatter) Formatter() Formatter {
	switch f {
	case JsonFormatter:
		return JsonFormatter
	case TextFormatter:
		return TextFormatter
	default:
		return UnknowFormatter
	}
}

// Formatter string
//
// String
// Date 2024-04-10 15:07:39
// Author Jerry.Yang
func (f Formatter) String() string {
	switch f {
	case JsonFormatter:
		return "json"
	case TextFormatter:
		return "text"
	default:
		return "unkown"
	}
}
