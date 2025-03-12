/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-12 10:24:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 10:25:54
 * @Description: errors
 */
package errors

import "github.com/yangjerry110/tool/toolerrors"

var (
	ErrLoggerFormatterUnkownFormatter = toolerrors.New("logger Err : logger formatter unkown formatter")
	ErrLoggerLevelUnknowLevel         = toolerrors.New("logger Err : logger level unkown level")
)
