/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-26 15:45:00
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-26 15:48:51
 * @Description: new app demo interface service
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppDemoInterfaceService struct{}

// New
//
// 新建
// author Jerry.Yang
// date 2024-04-26 15:48:35
func (n *NewAppDemoInterfaceService) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time              string
		ImportProjectPath string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()
	data.ImportProjectPath = config.ProjectImportPathConf.ImportPath

	filePath := fmt.Sprintf("%s/internal/service/interfaceService", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "demoInterfaceService.go", n.getTemplate(), data)
}

// getTemplate
//
// 获取template
// author Jerry.Yang
// date 2024-04-26 15:48:18
func (n *NewAppDemoInterfaceService) getTemplate() string {
	return `
	/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: Demo service
	*/
	package interfaceService

	import (
		"context"

		"{{.ImportProjectPath}}/vo/protobuf"
	)

	type DemoService interface {
		AddDemo(ctx context.Context, inputVo *protobuf.AddDemoReq) (*protobuf.Empty, error)
		DeleteDemo(ctx context.Context, inputVo *protobuf.DeleteDemoReq) (*protobuf.Empty, error)
		UpdateDemo(ctx context.Context, inputVo *protobuf.UpdateDemoReq) (*protobuf.Empty, error)
		GetDemo(ctx context.Context, inputVo *protobuf.GetDemoReq) (*protobuf.GetDemoResp, error)
	}
	`
}
