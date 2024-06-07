/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 14:32:55
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-07 11:11:47
 * @Description: errors
 */
package toolErrors

type ErrorInterface interface {
	New(message string) error
	NewError(err error) error
	WithStack() ErrorInterface
	WithFields(name string, value interface{}) ErrorInterface
	GetError() error
	Error() string
	String() string
}

/**
 * @description: New
 * @param {string} message
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:09:52
 * @return {*}
 */
func New(message string) error {
	return toolErrorsEnginee().New(message)
}

/**
 * @description: NewError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:10:06
 * @return {*}
 */
func NewError(err error) error {
	return toolErrorsEnginee().NewError(err)
}

/**
 * @description: WithStack
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:10:20
 * @return {*}
 */
func WithStack() ErrorInterface {
	return toolErrorsEnginee().WithStack()
}

/**
 * @description: WithFields
 * @param {string} name
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date:
 * @return {*}
 */
func WithFields(name string, value interface{}) ErrorInterface {
	return toolErrorsEnginee().WithFields(name, value)
}

/**
 * @description: Error
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:10:47
 * @return {*}
 */
func Error() string {
	return toolErrorsEnginee().Error()
}

/**
 * @description: String
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:11:04
 * @return {*}
 */
func String() string {
	return toolErrorsEnginee().String()
}
