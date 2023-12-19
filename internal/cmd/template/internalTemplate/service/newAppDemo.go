/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:57:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 17:08:50
 * @Description:
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppDemoService struct{}

/**
 * @description: NewDemo
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:28:41
 * @return {*}
 */
func (n *NewAppDemoService) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	filePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "demoService.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:28:33
 * @return {*}
 */
func (n *NewAppDemoService) getTemplate() string {
	return `
	/*
* @Author: Jerry.Yang
* @Date: {{.Time}}
* @LastEditors: Jerry.Yang
* @LastEditTime: {{.Time}}
* @Description: Demo service
 */
package service

import (
	"context"
	"git.qutoutiao.net/ee/tool-api/vo/protobuf"
)

type DemoService interface {
	AddDemo(ctx context.Context, inputVo *protobuf.AddDemoReq) (*protobuf.Empty, error)
	DeleteDemo(ctx context.Context, inputVo *protobuf.DeleteDemoReq) (*protobuf.Empty, error)
	UpdateDemo(ctx context.Context, inputVo *protobuf.UpdateDemoReq) (*protobuf.Empty, error)
	GetDemo(ctx context.Context, inputVo *protobuf.GetDemoReq) (*protobuf.GetDemoResp, error)
}

type Demo struct{}

/**
* @description: AddDemo
* @param {context.Context} ctx
* @param {*protobuf.AddDemoReq} inputVo
* @author: Jerry.Yang
* @date: {{.Time}}
* @return {*}
 */
func (d *Demo) AddDemo(ctx context.Context, inputVo *protobuf.AddDemoReq) (*protobuf.Empty, error) {

	/**
	* @step
	* @result
	**/
	result := &protobuf.Empty{}
	return result, nil
}

/**
* @description: DeleteDemo
* @param {context.Context} ctx
* @param {*protobuf.DeleteDemoReq} inputVo
* @author: Jerry.Yang
* @date: {{.Time}}
* @return {*}
 */
func (d *Demo) DeleteDemo(ctx context.Context, inputVo *protobuf.DeleteDemoReq) (*protobuf.Empty, error) {

	/**
	* @step
	* @result
	**/
	result := &protobuf.Empty{}
	return result, nil
}

/**
* @description: UpdateDemo
* @param {context.Context} ctx
* @param {*protobuf.UpdateDemoReq} inputVo
* @author: Jerry.Yang
* @date: {{.Time}}
* @return {*}
 */
func (d *Demo) UpdateDemo(ctx context.Context, inputVo *protobuf.UpdateDemoReq) (*protobuf.Empty, error) {

	/**
	* @step
	* @result
	**/
	result := &protobuf.Empty{}
	return result, nil
}

/**
* @description: GetDemo
* @param {context.Context} ctx
* @param {*protobuf.GetDemoReq} inputVo
* @author: Jerry.Yang
* @date: {{.Time}}
* @return {*}
 */
func (d *Demo) GetDemo(ctx context.Context, inputVo *protobuf.GetDemoReq) (*protobuf.GetDemoResp, error) {

	/**
	* @step
	* @result
	**/
	result := &protobuf.GetDemoResp{}
	return result, nil
}
`
}
