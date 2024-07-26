/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-31 11:17:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-07-26 14:25:22
 * @Description:
 */
package toolErrors

import (
	"github.com/pkg/errors"
)

type ToolError struct {
	isStack bool
	err     error
	fields  map[string]interface{}
}

/**
 * @description: New
 * @param {string} message
 * @author: Jerry.Yang
 * @date: 2024-06-06 15:49:57
 * @return {*}
 */
func (e *ToolError) New(message string) error {
	e.err = errors.New(message)
	return e.WithStack().GetError()
}

/**
 * @description: NewError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:46:42
 * @return {*}
 */
func (e *ToolError) NewError(err error) error {
	e.err = err
	return e.WithStack().GetError()
}

/**
 * @description: WithStack
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-06-06 15:50:14
 * @return {*}
 */
func (e *ToolError) WithStack() ErrorInterface {
	e.isStack = true
	return e
}

/**
 * @description: WithFields
 * @param {string} name
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:10:59
 * @return {*}
 */
func (e *ToolError) WithFields(name string, value interface{}) ErrorInterface {

	/**
	 * @step
	 * @judge fields
	 * @if == nil
	 * @make
	 **/
	if e.fields == nil {
		e.fields = make(map[string]interface{})
	}

	/**
	 * @step
	 * @set fields
	 **/
	e.fields[name] = value
	return e
}

/**
 * @description: Error
 * @author: Jerry.Yang
 * @date: 2024-06-06 15:56:45
 * @return {*}
 */
func (e *ToolError) Error() string {
	return e.GetError().Error()
}

/**
 * @description: String
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:15:13
 * @return {*}
 */
func (e *ToolError) String() string {
	return e.GetError().Error()
}

/**
 * @description: getError
 * @author: Jerry.Yang
 * @date: 2024-06-06 16:08:47
 * @return {*}
 */
func (e *ToolError) GetError() error {

	/**
	 * @step
	 * @定义
	 **/
	var err error

	/**
	 * @step
	 * @judge fields
	 * @if len != 0
	 * @set
	 **/
	if len(e.fields) != 0 {
		for fieldName, fieldVal := range e.fields {
			e.err = errors.WithMessagef(e.err, fieldName, fieldVal)
		}
	}

	/**
	 * @step
	 * @judge isStack
	 **/
	if e.isStack {
		err = errors.WithStack(e.err)
	}
	return err
}
