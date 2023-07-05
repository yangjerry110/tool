/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-26 10:34:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-26 10:52:48
 * @Description: mod
 */
package templates

type ModTemplate interface {
	SaveTemplate(path string, projectPath string) error
	GetTemplate() string
}

type Mod struct{}

/**
 * @description: SaveTemplate
 * @param {string} path
 * @param {string} projectPath
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:40:06
 * @return {*}
 */
func (m *Mod) SaveTemplate(path string, projectPath string) error {

	/**
	 * @step
	 * @定义渲染的数据
	 **/
	type Data struct {
		ProjectPath string
	}

	data := &Data{ProjectPath: projectPath}
	return CreateCommonTemplate().SaveTemplate(path, "go.mod", m.GetTemplate(), data, "gomod")
}

/**
 * @description: GetTemplate
 * @author: Jerry.Yang
 * @date: 2023-04-26 10:36:18
 * @return {*}
 */
func (m *Mod) GetTemplate() string {
	return `module {{.ProjectPath}}

go 1.20`
}
