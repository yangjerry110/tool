/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 14:32:55
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-31 11:46:35
 * @Description: errors
 */
package toolErrors

type ErrorInterface interface {
	New(err string) ErrorInterface
	NewError(err error) ErrorInterface
	WithPackage() ErrorInterface
	WithFile() ErrorInterface
	WithFunc() ErrorInterface
	WithLineNo() ErrorInterface
	WithFields(fieldName string, fieldVal interface{}) ErrorInterface
	WithCallFuncName(funcName string) ErrorInterface
	WithError(err error) ErrorInterface
	Error() string
}
