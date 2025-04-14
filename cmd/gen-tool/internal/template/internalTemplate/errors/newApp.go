/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:07:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 16:46:39
 * @Description: newApp errors
 */
package errors

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type NewAppError struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:09:49
 * @return {*}
 */
func (n *NewAppError) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	filePath := fmt.Sprintf("%s/internal/errors", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "errors.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:09:41
 * @return {*}
 */
func (n *NewAppError) getTemplate() string {
	return `
/*
 * @Author: Jerry.Yang
 * @Date: {{.Time}}
 * @LastEditors: Jerry.Yang
 * @LastEditTime: {{.Time}}
 * @Description: error
 */
package errors

import (
	"errors"
)

/**
 * @description: DefaultErrCode
 * @author: Jerry.Yang 
 * @date: {{.Time}}
 * @return {*}
 */
var DefaultErrCode = -1

/**
 * @description: Err_Default_Msg
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
var Err_Default_Msg = errors.New("err : err")

/**
 * @description: ErrCodes
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
var ErrCodes = map[error]int32{
	Err_Default_Msg: int32(DefaultErrCode),
}`
}
