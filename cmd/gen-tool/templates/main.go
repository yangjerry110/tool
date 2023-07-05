/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-25 17:37:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-26 10:38:27
 * @Description: main
 */
package templates

type MainTemplate interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Main struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:38:49
 * @return {*}
 */
func (m *Main) SaveTemplate(path string, projectPath string) error {

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
	}

	data := &Data{ProjectPath: projectPath}
	return CreateCommonTemplate().SaveTemplate(path, "main.go", m.GetTemplate(), data)
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-25 17:38:08
 * @return {*}
 */
func (m *Main) GetTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: 2023-04-21 14:44:29
	* @LastEditors: Jerry.Yang
	* @LastEditTime: 2023-04-21 17:31:12
	* @Description: main
	*/
   package main
   
   import (
	   "{{.ProjectPath}}/router"
	   "{{.ProjectPath}}/service"
   )
   
   func main() {
   
	   /**
		* @step
		* @启动之前的准备动作
		**/
	   service.CreateBeforeStartService().Preparing()
   
	   /**
		* @step
		* @run router
		**/
	   router.CreateCommonRouter().CreateRouter()
   }
   `
}
