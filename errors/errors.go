/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 15:22:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-07 11:13:25
 * @Description: errors
 */
package errors

import "github.com/yangjerry110/tool/internal/toolErrors"

/**
 * @description: New
 * @param {string} message
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:11:31
 * @return {*}
 */
func New(message string) error {
	return toolErrors.New(message)
}

/**
 * @description: NewError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:11:42
 * @return {*}
 */
func NewError(err error) error {
	return toolErrors.NewError(err)
}

/**
 * @description: WithStack
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:05
 * @return {*}
 */
func WithStack() toolErrors.ErrorInterface {
	return toolErrors.WithStack()
}

/**
 * @description: WithFields
 * @param {string} name
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:25
 * @return {*}
 */
func WithFields(name string, value interface{}) toolErrors.ErrorInterface {
	return toolErrors.WithFields(name, value)
}

/**
 * @description: Error
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:43
 * @return {*}
 */
func Error() string {
	return toolErrors.Error()
}

/**
 * @description: String
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:13:00
 * @return {*}
 */
func String() string {
	return toolErrors.String()
}
