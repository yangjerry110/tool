/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:02:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 23:03:44
 * @Description: base
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/templates"
)

type BaseService interface {
	SaveTemplate(path string) error
	GetTemplate() string
	AppendFuncTemplate(path string, serviceName string) error
	GetAppendFuncTemplate() string
}

type Base struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:04:49
 * @return {*}
 */
func (b *Base) SaveTemplate(path string) error {

	/**
	 * @step
	 * @定义数据结构
	 **/
	type Data struct {
		Time string
	}

	data := &Data{Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), data)
}

/**
 * @description: AppendFuncTemplate
 * @param {string} path
 * @param {string} serviceName
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:13:22
 * @return {*}
 */
func (b *Base) AppendFuncTemplate(path string, serviceName string) error {

	/**
	 * @step
	 * @获取base的文件地址
	 **/
	basePath := fmt.Sprintf("%s/%s", path, "base.go")

	/**
	 * @step
	 * @定义需要渲染的数据结构
	 **/
	type Data struct {
		ServiceName   string
		ServiceNameUp string
		Time          string
	}

	/**
	 * @step
	 * @渲染参数
	 **/
	data := &Data{ServiceName: serviceName, ServiceNameUp: templates.CreateCommonTemplate().FirstUpper(serviceName), Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().AppendTemplate(basePath, b.GetAppendFuncTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:04:12
 * @return {*}
 */
func (b *Base) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: base
	*/
   package service
   
   /**
	* @description: CreateBeforeStartService
	* @param {...BeforeStartService} BeforeStartServices
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateBeforeStartService(BeforeStartServices ...BeforeStartService) BeforeStartService {
	   if len(BeforeStartServices) == 0 {
		   return &BeforeStart{}
	   }
	   return BeforeStartServices[0]
   }
   
   /**
	* @description: CreateDemoService
	* @param {...DemoService} DemoServices
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func CreateDemoService(DemoServices ...DemoService) DemoService {
	   if len(DemoServices) == 0 {
		   return &Demo{}
	   }
	   return DemoServices[0]
   }
   `
}

/**
 * @description: GetAppendFuncTemplate
 * @author: Jerry.Yang
 * @date: 2023-05-11 15:24:35
 * @return {*}
 */
func (b *Base) GetAppendFuncTemplate() string {
	return `/**
	* @description: Create{{.ServiceNameUp}}Service
	* @param {...{{.ServiceNameUp}}Service} {{.ServiceName}}Services
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func Create{{.ServiceNameUp}}Service({{.ServiceName}}Services ...{{.ServiceNameUp}}Service) {{.ServiceNameUp}}Service {
	   if len({{.ServiceName}}Services) == 0 {
		   return &{{.ServiceNameUp}}{}
	   }
	   return {{.ServiceName}}Services[0]
   }`
}
