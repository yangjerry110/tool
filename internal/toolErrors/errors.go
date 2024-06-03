/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 14:32:55
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-03 11:15:43
 * @Description: errors
 */
package toolErrors

type ErrorInterface interface {
	New(err string) ErrorInterface
	NewError(err error) ErrorInterface
	WithPackage() ErrorInterface
	WithFile() ErrorInterface
	WithFunc() ErrorInterface
	WithStackTrace() ErrorInterface
	WithLineNo() ErrorInterface
	WithFields(fieldName string, fieldVal interface{}) ErrorInterface
	WithCallFuncName(funcName string) ErrorInterface
	WithError(err error) ErrorInterface
	GetError() ErrorInterface
	SetRuntimeDept(runtimeDept int) ErrorInterface
	Error() string
	String() string
}

/**
 * @description: New
 * @param {string} err
 * @author: Jerry.Yang
 * @date: 2024-05-31 10:33:37
 * @return {*}
 */
func New(err string) ErrorInterface {
	return toolErrorsEnginee().New(err)
}

/**
 * @description: NewError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-31 11:47:07
 * @return {*}
 */
func NewError(err error) ErrorInterface {
	return toolErrorsEnginee().NewError(err)
}

/**
 * @description: WithPackage
 * @param {string} packageName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:31
 * @return {*}
 */
func WithPackage() ErrorInterface {
	return toolErrorsEnginee().WithPackage()
}

/**
 * @description: WithFile
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:19:31
 * @return {*}
 */
func WithFile() ErrorInterface {
	return toolErrorsEnginee().WithFile()
}

/**
 * @description: WithFunc
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:41
 * @return {*}
 */
func WithFunc() ErrorInterface {
	return toolErrorsEnginee().WithFunc()
}

/**
 * @description: WithStackTrace
 * @author: Jerry.Yang
 * @date: 2024-06-03 11:15:31
 * @return {*}
 */
func WithStackTrace() ErrorInterface {
	return toolErrorsEnginee().WithStackTrace()
}

/**
 * @description: WithLineNo
 * @param {int} lineNo
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:20:23
 * @return {*}
 */
func WithLineNo() ErrorInterface {
	return toolErrorsEnginee().WithLineNo()
}

/**
 * @description: WithFields
 * @param {string} paramName
 * @param {interface{}} paramVal
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:52
 * @return {*}
 */
func WithFields(fieldName string, fieldVal interface{}) ErrorInterface {
	return toolErrorsEnginee().WithFields(fieldName, fieldVal)
}

/**
 * @description: WithCallFuncName
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:17:29
 * @return {*}
 */
func WithCallFuncName(funcName string) ErrorInterface {
	return toolErrorsEnginee().WithCallFuncName(funcName)
}

/**
 * @description: WithError
 * @param {error} errmsg
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:18:39
 * @return {*}
 */
func WithError(err error) ErrorInterface {
	return toolErrorsEnginee().WithError(err)
}

/**
 * @description: SetRuntimeDept
 * @param {int} runtimeDept
 * @author: Jerry.Yang
 * @date: 2024-05-31 15:04:15
 * @return {*}
 */
func SetRuntimeDept(runtimeDept int) ErrorInterface {
	return toolErrorsEnginee().SetRuntimeDept(runtimeDept)
}

/**
 * @description: Error
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:01:02
 * @return {*}
 */
func Error() string {
	return toolErrorsEnginee().Error()
}
