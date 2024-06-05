/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 15:22:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-03 14:09:41
 * @Description: errors
 */
package errors

import (
	"github.com/yangjerry110/tool/internal/toolErrors"
)

/**
 * @description: New
 * @param {string} err
 * @author: Jerry.Yang
 * @date: 2024-05-31 10:33:37
 * @return {*}
 */
func New(err string) error {
	return toolErrors.New(err)
}

/**
 * @description: NewError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-31 11:47:07
 * @return {*}
 */
func NewError(err error) error {
	return toolErrors.NewError(err)
}

/**
 * @description: WithPackage
 * @param {string} packageName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:31
 * @return {*}
 */
func WithPackage() toolErrors.ErrorInterface {
	return toolErrors.WithPackage()
}

/**
 * @description: WithFile
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:19:31
 * @return {*}
 */
func WithFile() toolErrors.ErrorInterface {
	return toolErrors.WithFile()
}

/**
 * @description: WithFunc
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:41
 * @return {*}
 */
func WithFunc() toolErrors.ErrorInterface {
	return toolErrors.WithFunc()
}

/**
 * @description: WithStackTrace
 * @author: Jerry.Yang
 * @date: 2024-06-03 14:09:11
 * @return {*}
 */
func WithStackTrace() toolErrors.ErrorInterface {
	return toolErrors.WithStackTrace()
}

/**
 * @description: WithLineNo
 * @param {int} lineNo
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:20:23
 * @return {*}
 */
func WithLineNo() toolErrors.ErrorInterface {
	return toolErrors.WithLineNo()
}

/**
 * @description: WithFields
 * @param {string} paramName
 * @param {interface{}} paramVal
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:52
 * @return {*}
 */
func WithFields(fieldName string, fieldVal interface{}) toolErrors.ErrorInterface {
	return toolErrors.WithFields(fieldName, fieldVal)
}

/**
 * @description: WithCallFuncName
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:17:29
 * @return {*}
 */
func WithCallFuncName(funcName string) toolErrors.ErrorInterface {
	return toolErrors.WithCallFuncName(funcName)
}

/**
 * @description: WithError
 * @param {error} errmsg
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:18:39
 * @return {*}
 */
func WithError(err error) toolErrors.ErrorInterface {
	return toolErrors.WithError(err)
}

/**
 * @description: SetRuntimeDept
 * @param {int} runtimeDept
 * @author: Jerry.Yang
 * @date: 2024-05-31 15:04:15
 * @return {*}
 */
func SetRuntimeDept(runtimeDept int) toolErrors.ErrorInterface {
	return toolErrors.SetRuntimeDept(runtimeDept)
}

/**
 * @description: Error
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:01:02
 * @return {*}
 */
func Error() string {
	return toolErrors.Error()
}
