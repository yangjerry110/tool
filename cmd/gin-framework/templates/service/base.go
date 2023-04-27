/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:02:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-25 16:05:21
 * @Description: base
 */
package service

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type BaseService interface {
	SaveTemplate(path string) error
	GetTemplate() string
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
