/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 16:08:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-27 15:12:03
 * @Description: options
 */
package http

/**
 * @定义常量
 * @author Jerry.Yang
 * @date 2022-09-21 15:57:57
 **/
var OPTION_OUT_TIME = "outTime"
var OPTION_HEADERS = "headers"

/**
 * @description: SetHeaders
 * @param {map[string]string} value
 * @author: Jerry.Yang
 * @date: 2022-09-21 16:32:35
 * @return {*}
 */
func (h *HttpOptions) SetHeaders(value map[string]string) HttpOptionFunc {
	return func(HttpOptionFunc map[string]HttpOption) error {
		HttpOptionFunc[OPTION_HEADERS] = HttpOption{value}
		return nil
	}
}

/**
 * @description: setTimeOut
 * @param {int} value
 * @author: Jerry.Yang
 * @date: 2022-03-08 11:17:32
 * @return {*}
 */
func (h *HttpOptions) SetOutTime(value int) HttpOptionFunc {
	return func(HttpOptionFunc map[string]HttpOption) error {
		HttpOptionFunc[OPTION_OUT_TIME] = HttpOption{value}
		return nil
	}
}

/**
 * @description: GetHeaders
 * @param {map[string]HttpOption} httpOptions
 * @author: Jerry.Yang
 * @date: 2022-09-21 16:59:40
 * @return {*}
 */
func (h *HttpOptions) GetHeaders(httpOptions map[string]HttpOption) map[string]string {

	/**
	 * @step
	 * @定义
	 **/
	headers := map[string]string{}

	/**
	 * @step
	 * @获取设置的headers
	 **/
	setHeaders, ok := httpOptions[OPTION_HEADERS]
	if ok {
		return setHeaders.Value.(map[string]string)
	}
	return headers
}

/**
 * @description: GetOutTime
 * @param {map[string]HttpOption} httpOptions
 * @author: Jerry.Yang
 * @date: 2022-09-21 16:49:46
 * @return {*}
 */
func (h *HttpOptions) GetOutTime(httpOptions map[string]HttpOption) int {

	/**
	 * @step
	 * @定义初始值
	 **/
	outTime := 60

	/**
	 * @step
	 * @获取设置的outTime
	 **/
	setOutTime, ok := httpOptions[OPTION_OUT_TIME]
	if !ok {
		return outTime
	}

	/**
	 * @step
	 * @获取设置的过期时间的真实的数据
	 **/
	setOutTimeInt := setOutTime.Value.(int)

	/**
	 * @step
	 * @判断设置的过期时间是否有效
	 **/
	if setOutTimeInt <= 0 {
		return outTime
	}
	return setOutTimeInt
}

/**
 * @description: SetOptions
 * @param {[]Option} options
 * @author: Jerry.Yang
 * @date: 2022-03-01 15:44:58
 * @return {*}
 */
func (h *HttpOptions) SetOptions(httpOptionFuncs []HttpOptionFunc) map[string]HttpOption {

	/**
	 * @step
	 * @定义
	 **/
	setHttpOption := map[string]HttpOption{}

	/**
	 * @step
	 * @获取set的functions
	 **/
	for _, option := range httpOptionFuncs {
		if option != nil {
			err := option(setHttpOption)
			if err != nil {
				return nil
			}
		}
	}
	return setHttpOption
}
