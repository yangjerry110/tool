/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-23 16:38:43
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 11:35:59
 * @Description: common
 */
package commands

/**
 * @description: HttpRule
 * @author: Jerry.Yang
 * @date: 2023-05-23 16:39:18
 * @return {*}
 */
type HttpRule struct {
	Description string
	FuncName    string
	Method      string
	Url         string
	InputName   string
	OutputName  string
}

/**
 * @description: Common
 * @author: Jerry.Yang
 * @date: 2023-05-25 11:04:08
 * @return {*}
 */
type Common struct {
	IsFirstCreate bool
	IsAppend      bool
}

/**
 * @description: CommandParams
 * @author: Jerry.Yang
 * @date: 2023-05-25 11:04:16
 * @return {*}
 */
var CommandParams = &Common{}
