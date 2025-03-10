package toolerrors

/**
 * @description: New
 * @param {string} message
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:11:31
 * @return {*}
 */
func New(message string) error {
	return toolErrorsEnginee().New(message)
}

/**
 * @description: NewError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:11:42
 * @return {*}
 */
func NewError(err error) error {
	return toolErrorsEnginee().NewError(err)
}

/**
 * @description: WithStack
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:05
 * @return {*}
 */
func WithStack() errorInterface {
	return toolErrorsEnginee().WithStack()
}

/**
 * @description: WithFields
 * @param {string} name
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:25
 * @return {*}
 */
func WithFields(name string, value interface{}) errorInterface {
	return toolErrorsEnginee().WithFields(name, value)
}

/**
 * @description: Error
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:12:43
 * @return {*}
 */
func Error() string {
	return toolErrorsEnginee().Error()
}

/**
 * @description: String
 * @author: Jerry.Yang
 * @date: 2024-06-07 11:13:00
 * @return {*}
 */
func String() string {
	return toolErrorsEnginee().String()
}
