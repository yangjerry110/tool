/*
* @Author: Jerry.Yang
* @Date: 2024-05-30 15:17:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-28 11:25:00
* @Description: errorsEnginee
*/
package toolErrors

import (
	"github.com/yangjerry110/tool/internal/toolErrors"
)

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
