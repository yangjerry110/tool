/*
* @Author: Jerry.Yang
* @Date: 2024-05-30 15:17:53
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 20:52:02
* @Description: errorsEnginee
*/
package toolerrors

/**
 * @description: default Error enginee
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:21:59
 * @return {*}
 */
var defaultErrorsEngine errorInterface

/**
 * @description: SetErrorsEnginee
 * @param {ErrorInterface} ErrorInterface
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:24:05
 * @return {*}
 */
func SetToolErrorsEnginee(ErrorInterface errorInterface) errorInterface {
	defaultErrorsEngine = ErrorInterface
	return defaultErrorsEngine
}

/**
 * @description: toolErrorsEnginee
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:25:58
 * @return {*}
 */
func toolErrorsEnginee() errorInterface {
	return &toolError{}
}
