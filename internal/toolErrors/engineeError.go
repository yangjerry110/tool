/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-31 14:21:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-05 16:20:48
 * @Description: internal error
 */
package toolErrors

/**
 * @description: default Error enginee
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:21:59
 * @return {*}
 */
var defaultErrorsEngine ErrorInterface

/**
 * @description: SetErrorsEnginee
 * @param {ErrorInterface} ErrorInterface
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:24:05
 * @return {*}
 */
func SetToolErrorsEnginee(ErrorInterface ErrorInterface) ErrorInterface {
	defaultErrorsEngine = ErrorInterface
	return defaultErrorsEngine
}

/**
 * @description: toolErrorsEnginee
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:25:58
 * @return {*}
 */
func toolErrorsEnginee() ErrorInterface {
	return &ToolError{}
}
