/*
 * @Author: Jerry.Yang
 * @Date: 2024-03-05 14:05:27
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-03-05 14:36:27
 * @Description:
 */
package service

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
)

type NewBaseService struct {
	ProjectImportPath string
	Time              string
}

// new
//
// new
// Author Jerry.Yang
// Date 2024-03-05 14:06:58
func (n *NewBaseService) New() error {
	// saveTemplate
	filePath := fmt.Sprintf("%s/internal/service", config.ProjectPathConf.Path)
	return template.SaveTemplate(filePath, "baseService.go", n.getTemplate(), n, "baseGo")
}

// getTemplate
//
// getTemplate
// Author Jerry.Yang
// Date 2024-03-05 14:06:41
func (n *NewBaseService) getTemplate() string {
	return `/*
	* @Author: Jerry.Yang
	* @Date: {{.Time}}
	* @LastEditors: Jerry.Yang
	* @LastEditTime: {{.Time}}
	* @Description: base
	*/
   package service

   import "{{.ProjectImportPath}}/internal/service/interfaceService"
	`
}
