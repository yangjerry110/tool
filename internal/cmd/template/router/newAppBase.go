/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 14:51:23
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 16:36:17
 * @Description: newApp base router
 */
package router

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppBaseRouter struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:53:24
 * @return {*}
 */
func (n *NewAppBaseRouter) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time       string
		ImportPath string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()
	data.ImportPath = config.ProjectImportPathConf.ImportPath

	filePath := fmt.Sprintf("%s/router", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "base.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 14:52:38
 * @return {*}
 */
func (n *NewAppBaseRouter) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: base
	*/
   package router
   
   import (
	   "{{.ImportPath}}/internal/config"
	   "github.com/yangjerry110/tool/router"
   )
   
   /**
	* @description: RunRouter
	* @author: Jerry.Yang
	* @date: {{.Time}}
	* @return {*}
	*/
   func RunRouter() {
   
		// Register Qiwei
		if err := router.CreateGinRouter().Register("qiwei", &Qiwei{}); err != nil {
			panic(err)
		}

		// Run router
		if err := router.CreateGinRouter().Run(&config.Config{}); err != nil {
			panic(err)
		}
   }
   `
}
