/*
* @Author: Jerry.Yang
* @Date: 2024-05-30 15:17:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-30 19:22:40
* @Description: errorsEnginee
*/
package errors

import (
	"github.com/yangjerry110/tool/internal/toolErrors"
)

/**
 * @description: defaultErrorsEnginee
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:21:27
 * @return {*}
 */
var defaultErrorsEnginee toolErrors.ErrorInterface

/**
 * @description: SetErrorsEnginee
 * @param {toolErrors.ErrorInterface} errorsEnginee
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:21:20
 * @return {*}
 */
func SetErrorsEnginee(errorsEnginee toolErrors.ErrorInterface) toolErrors.ErrorInterface {
	defaultErrorsEnginee = errorsEnginee
	return defaultErrorsEnginee
}

/**
 * @description: Error
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:21:12
 * @return {*}
 */
func toolError() toolErrors.ErrorInterface {

	/**
	 * @step
	 * @judge defaultErrorsEngine
	 * @if == nil
	 **/
	if defaultErrorsEnginee == nil {
		return &toolErrors.ToolError{}
	}
	return defaultErrorsEnginee
}
