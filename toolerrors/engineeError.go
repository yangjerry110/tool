package toolerrors

type errorInterface interface {
	New(message string) error
	NewError(err error) error
	WithStack() errorInterface
	WithFields(name string, value interface{}) errorInterface
	GetError() error
	Error() string
	String() string
}

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
