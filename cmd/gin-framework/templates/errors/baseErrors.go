/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-24 17:14:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 17:17:03
 * @Description: base
 */
package errors

/**
 * @description: CreateError
 * @param {...ErrorsInterface} ErrorsInterfaces
 * @author: Jerry.Yang
 * @date: 2023-04-24 17:17:08
 * @return {*}
 */
func CreateError(ErrorsInterfaces ...ErrorsInterface) ErrorsInterface {
	if len(ErrorsInterfaces) == 0 {
		return &Errors{}
	}
	return ErrorsInterfaces[0]
}
