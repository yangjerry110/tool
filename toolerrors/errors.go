/*
 * @Author: Jerry.Yang
 * @Date: 2024-06-07 11:11:31
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 22:32:44
 * @FilePath: /toolerrors/errors.go
 * @Description: 错误处理工具模块，提供统一的错误创建、包装、堆栈信息添加、附加字段等功能。
 * 通过 `errorInterface` 接口实现错误处理的扩展性和灵活性，支持全局默认错误处理引擎的配置。
 */
package toolerrors

/**
 * @description: New 创建一个新的错误
 * @param {string} message 错误信息
 * @return {error} 返回创建的错误对象
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:11:31
 */
func New(message string) error {
	return toolErrorsEnginee().New(message)
}

/**
 * @description: NewError 包装一个现有的错误
 * @param {error} err 需要包装的错误对象
 * @return {error} 返回包装后的错误对象
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:11:42
 */
func NewError(err error) error {
	return toolErrorsEnginee().NewError(err)
}

/**
 * @description: WithStack 为错误添加堆栈信息
 * @return {errorInterface} 返回支持堆栈信息的错误处理接口实例
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:05
 */
func WithStack() errorInterface {
	return toolErrorsEnginee().WithStack()
}

/**
 * @description: WithFields 为错误添加附加字段
 * @param {string} name 字段名称
 * @param {interface{}} value 字段值
 * @return {errorInterface} 返回支持附加字段的错误处理接口实例
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:25
 */
func WithFields(name string, value interface{}) errorInterface {
	return toolErrorsEnginee().WithFields(name, value)
}

/**
 * @description: Error 获取错误的字符串表示
 * @return {string} 返回错误信息字符串
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:43
 */
func Error() string {
	return toolErrorsEnginee().Error()
}

/**
 * @description: String 获取错误的字符串表示（与 Error 方法功能相同）
 * @return {string} 返回错误信息字符串
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:13:00
 */
func String() string {
	return toolErrorsEnginee().String()
}
