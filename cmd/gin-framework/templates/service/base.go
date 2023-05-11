/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:02:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 15:25:50
 * @Description: base
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gin-framework/templates"
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
	return templates.CreateCommonTemplate().SaveTemplate(path, "base.go", b.GetTemplate(), nil)
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
	}

	/**
	 * @step
	 * @渲染参数
	 **/
	data := &Data{ServiceName: serviceName, ServiceNameUp: templates.CreateCommonTemplate().FirstUpper(serviceName)}
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
	* @Date: 2023-04-21 16:56:17
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-23 14:27:27
	* @Description: base
	*/
   package service
   
   /**
	* @description: CreateBeforeStartService
	* @param {...BeforeStartService} BeforeStartServices
	* @author: Jerry.Yang
	* @date: 2023-04-21 17:30:56
	* @return {*}
	*/
   func CreateBeforeStartService(BeforeStartServices ...BeforeStartService) BeforeStartService {
	   if len(BeforeStartServices) == 0 {
		   return &BeforeStart{}
	   }
	   return BeforeStartServices[0]
   }
   
   /**
	* @description: CreateTestService
	* @param {...TestService} TestServices
	* @author: Jerry.Yang
	* @date: 2023-04-23 14:27:34
	* @return {*}
	*/
   func CreateTestService(TestServices ...TestService) TestService {
	   if len(TestServices) == 0 {
		   return &Test{}
	   }
	   return TestServices[0]
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
	* @description: Create{{.ServiceNameUp}}Dao
	* @param {...{{.ServiceNameUp}}Dao} {{.ServiceName}}Daos
	* @author: Jerry.Yang
	* @date: 2023-04-24 17:02:59
	* @return {*}
	*/
   func Create{{.ServiceNameUp}}Dao({{.ServiceName}}Daos ...{{.ServiceNameUp}}Dao) {{.ServiceNameUp}}Dao {
	   if len({{.ServiceName}}Daos) == 0 {
		   return &{{.ServiceNameUp}}{}
	   }
	   return {{.ServiceName}}Daos[0]
   }`
}
