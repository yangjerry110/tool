/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-31 14:21:59
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-31 14:25:58
 * @FilePath: /toolerrors/engine.go
 * @Description: 错误处理引擎模块，提供统一的错误处理接口和默认错误处理引擎的配置。
 * 通过 `errorInterface` 接口实现错误处理的扩展性和灵活性，支持全局默认错误处理引擎的设置和获取。
 */
package toolerrors

// errorInterface 定义错误处理接口，用于统一错误处理的行为
type errorInterface interface {
	New(message string) error                                 // 创建一个新的错误
	NewError(err error) error                                 // 包装一个现有的错误
	WithStack() errorInterface                                // 添加堆栈信息
	WithFields(name string, value interface{}) errorInterface // 添加附加字段
	GetError() error                                          // 获取错误对象
	Error() string                                            // 实现 error 接口，返回错误信息
	String() string                                           // 返回错误的字符串表示
}

/**
 * @description: defaultErrorsEngine 全局变量，存储默认的错误处理引擎
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:21:59
 */
var defaultErrorsEngine errorInterface

/**
 * @description: SetToolErrorsEnginee 设置全局默认的错误处理引擎
 * @param {errorInterface} ErrorInterface 错误处理引擎实例
 * @return {errorInterface} 返回设置后的错误处理引擎
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:24:05
 */
func SetToolErrorsEnginee(ErrorInterface errorInterface) errorInterface {
	defaultErrorsEngine = ErrorInterface
	return defaultErrorsEngine
}

/**
 * @description: toolErrorsEnginee 返回默认的错误处理引擎实例
 * @return {errorInterface} 返回默认的错误处理引擎
 * @author: Jerry.Yang
 * @date: 2024-05-31 14:25:58
 */
func toolErrorsEnginee() errorInterface {
	return &toolError{} // 返回一个默认的 toolError 实例
}
