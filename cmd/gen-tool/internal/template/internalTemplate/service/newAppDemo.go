/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:57:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-28 16:09:51
 * @Description:
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
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
		Time              string
		ImportProjectPath string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()
	data.ImportProjectPath = config.ProjectImportPathConf.ImportPath

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
	"{{.ImportProjectPath}}/vo/protobuf"
)

type Demo struct{}

/**
 * @description: RouterName
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
func (d *Demo) RouterName() string {
	return "demo"
}

/**
 * @description: MustRouterHTTPService
 * @author: Jerry.Yang
 * @date: {{.Time}}
 * @return {*}
 */
func (d *Demo) MustRouterHTTPService() {}

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
