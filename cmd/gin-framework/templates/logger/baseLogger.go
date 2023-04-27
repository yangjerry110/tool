/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 10:25:34
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 10:36:08
 * @Description: baseLogger
 */
package logger

/**
 * @description: CreateBaseLogger
 * @param {...BaseLogger} BaseLoggers
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:31:09
 * @return {*}
 */
func CreateBaseLogger(BaseLoggers ...BaseLogger) BaseLogger {
	if len(BaseLoggers) == 0 {
		return &Base{}
	}
	return BaseLoggers[0]
}

/**
 * @description: CreateCommonLogger
 * @param {...CommonLogger} CommonLoggers
 * @author: Jerry.Yang
 * @date: 2023-04-25 10:36:16
 * @return {*}
 */
func CreateCommonLogger(CommonLoggers ...CommonLogger) CommonLogger {
	if len(CommonLoggers) == 0 {
		return &Common{}
	}
	return CommonLoggers[0]
}
