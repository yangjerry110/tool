/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-25 16:21:29
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 16:48:13
 * @Description: common
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
)

type CommonService interface {
	PrintfErr(err error, errMsgs ...string)
}

type Common struct{}

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
 * @description: PrintErr
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2023-05-25 16:23:05
 * @return {*}
 */
func (c *Common) PrintfErr(err error, errMsgs ...string) {
	if len(errMsgs) == 0 {
		fmt.Printf("\033[1;31;40m ERROR | %s | %+v\033[0m\n", templates.CreateCommonTemplate().GetFormatNowTime(), err)
		return
	}
	errMsgFirst := errMsgs[0]
	printErrMsg := fmt.Sprintf("%s err : %+v", errMsgFirst, err)
	fmt.Printf("\033[1;31;40m ERROR | %s | %s\033[0m\n", templates.CreateCommonTemplate().GetFormatNowTime(), printErrMsg)
}
