/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:09:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 15:19:56
 * @Description: test
 */
package service

import "github.com/yangjerry110/tool/cmd/gin-framework/templates"

type TestService interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Test struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:10:02
 * @return {*}
 */
func (t *Test) SaveTemplate(path string, projectPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
		Time        string
	}

	data := &Data{ProjectPath: projectPath, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "testService.go", t.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:10:47
 * @return {*}
 */
func (t *Test) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: test service
	*/
   package service
   
   import (
	   "context"
	   "{{.ProjectPath}}/vo/input"
	   "{{.ProjectPath}}/vo/output"
   )
   
   type TestService interface {
	   Test(ctx context.Context, inputVo *input.Test) (*output.Test, error)
   }
   
   type Test struct{}
   
   /**
	* @description: Test
	* @param {context.Context} ctx
	* @param {*input.Test} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (t *Test) Test(ctx context.Context, inputVo *input.Test) (*output.Test, error) {
   
	   /**
		* @step
		* @result
		**/
	   output := &output.Test{}
	   return output, nil
   }
   `
}
