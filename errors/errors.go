/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 15:22:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-31 10:33:42
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
	return toolError().New(err)
}

/**
 * @description: WithPackage
 * @param {string} packageName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:31
 * @return {*}
 */
func WithPackage() toolErrors.ErrorInterface {
	return toolError().WithPackage()
}

/**
 * @description: WithFile
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:19:31
 * @return {*}
 */
func WithFile() toolErrors.ErrorInterface {
	return toolError().WithFile()
}

/**
 * @description: WithFunc
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:41
 * @return {*}
 */
func WithFunc() toolErrors.ErrorInterface {
	return toolError().WithFunc()
}

/**
 * @description: WithLineNo
 * @param {int} lineNo
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:20:23
 * @return {*}
 */
func WithLineNo() toolErrors.ErrorInterface {
	return toolError().WithLineNo()
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
	return toolError().WithFields(fieldName, fieldVal)
}

/**
 * @description: WithCallFuncName
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:17:29
 * @return {*}
 */
func WithCallFuncName(funcName string) toolErrors.ErrorInterface {
	return toolError().WithCallFuncName(funcName)
}

/**
 * @description: WithError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:18:39
 * @return {*}
 */
func WithError(err string) toolErrors.ErrorInterface {
	return toolError().WithError(err)
}

/**
 * @description: Error
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:01:02
 * @return {*}
 */
func Error() string {
	return toolError().Error()
}
