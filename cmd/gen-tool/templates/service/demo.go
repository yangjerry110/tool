/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 16:09:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:57:34
 * @Description: Demo
 */
package service

import "github.com/yangjerry110/tool/cmd/gen-tool/templates"

type DemoService interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Demo struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:10:02
 * @return {*}
 */
func (d *Demo) SaveTemplate(path string, projectPath string) error {
	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
		Time        string
	}

	data := &Data{ProjectPath: projectPath, Time: templates.CreateCommonTemplate().GetFormatNowTime()}
	return templates.CreateCommonTemplate().SaveTemplate(path, "demoService.go", d.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 16:10:47
 * @return {*}
 */
func (d *Demo) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: Demo service
	*/
   package service
   
   import (
	   "context"
	   "{{.ProjectPath}}/vo/input"
	   "{{.ProjectPath}}/vo/output"
   )
   
   type DemoService interface {
	   Demo(ctx context.Context, inputVo *input.Demo) (*output.Demo, error)
   }
   
   type Demo struct{}
   
   /**
	* @description: Demo
	* @param {context.Context} ctx
	* @param {*input.Demo} inputVo
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func (t *Demo) Demo(ctx context.Context, inputVo *input.Demo) (*output.Demo, error) {
   
	   /**
		* @step
		* @result
		**/
	   output := &output.Demo{}
	   return output, nil
   }
   `
}
