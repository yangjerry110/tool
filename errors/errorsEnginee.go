/*
* @Author: Jerry.Yang
* @Date: 2024-05-30 15:17:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-03 10:58:37
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
	return toolErrors.SetToolErrorsEnginee(errorsEnginee)
}
