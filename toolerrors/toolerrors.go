/*
 * @Author: Jerry.Yang
 * @Date: 2024-06-06 15:49:57
 * @LastEditors: yangjie04 yangjie04@qutoutiao.net
 * @LastEditTime: 2025-03-10 22:34:31
 * @FilePath: /toolerrors/toolerror.go
 * @Description: 错误处理实现模块，提供具体的错误创建、包装、堆栈信息添加、附加字段等功能。
 * 通过 `toolError` 结构体实现 `errorInterface` 接口，支持错误堆栈信息和附加字段的添加。
 */
package toolerrors

import (
	"github.com/pkg/errors"
)

// toolError 结构体，实现 errorInterface 接口
type toolError struct {
	isStack bool                   // 是否添加堆栈信息
	err     error                  // 错误对象
	fields  map[string]interface{} // 附加字段
}

/**
 * @description: New 创建一个新的错误
 * @param {string} message 错误信息
 * @return {error} 返回创建的错误对象
 * @author: Jerry.Yang
 * @date: 2024-06-06 15:49:57
 */
func (e *toolError) New(message string) error {
	e.err = errors.New(message)
	return e.WithStack().GetError()
}

/**
 * @description: NewError 包装一个现有的错误
 * @param {error} err 需要包装的错误对象
 * @return {error} 返回包装后的错误对象
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:46:42
 */
func (e *toolError) NewError(err error) error {
	e.err = err
	return e.WithStack().GetError()
}

/**
 * @description: WithStack 为错误添加堆栈信息
 * @return {errorInterface} 返回支持堆栈信息的错误处理接口实例
 * @author: Jerry.Yang
 * @date: 2024-06-06 15:50:14
 */
func (e *toolError) WithStack() errorInterface {
	e.isStack = true
	return e
}

/**
 * @description: WithFields 为错误添加附加字段
 * @param {string} name 字段名称
 * @param {interface{}} value 字段值
 * @return {errorInterface} 返回支持附加字段的错误处理接口实例
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:10:59
 */
func (e *toolError) WithFields(name string, value interface{}) errorInterface {
	// 如果附加字段为空，初始化 map
	if e.fields == nil {
		e.fields = make(map[string]interface{})
	}

	// 添加附加字段
	e.fields[name] = value
	return e
}

/**
 * @description: Error 获取错误的字符串表示
 * @return {string} 返回错误信息字符串
 * @author: Jerry.Yang
 * @date: 2024-06-06 15:56:45
 */
func (e *toolError) Error() string {
	return e.GetError().Error()
}

/**
 * @description: String 获取错误的字符串表示（与 Error 方法功能相同）
 * @return {string} 返回错误信息字符串
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:15:13
 */
func (e *toolError) String() string {
	return e.GetError().Error()
}

/**
 * @description: GetError 获取最终的错误对象
 * @return {error} 返回处理后的错误对象
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:08:47
 */
func (e *toolError) GetError() error {
	var err error

	// 如果附加字段不为空，为错误添加附加字段信息
	if len(e.fields) != 0 {
		for fieldName, fieldVal := range e.fields {
			e.err = errors.WithMessagef(e.err, "%s=%v", fieldName, fieldVal)
		}
	}

	// 如果需要添加堆栈信息，为错误添加堆栈信息
	if e.isStack {
		err = errors.WithStack(e.err)
	}
	return err
}
