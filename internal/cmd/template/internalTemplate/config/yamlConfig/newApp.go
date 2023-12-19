package yamlconfig

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

/**
 * @description: NewAppDatabase
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:31:13
 * @return {*}
 */
type NewAppDatabase struct{}

/**
 * @description: NewAppLogger
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:31:13
 * @return {*}
 */
type NewAppLogger struct{}

/**
 * @description: NewAppRouter
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:31:13
 * @return {*}
 */
type NewAppRouter struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:30:44
 * @return {*}
 */
func (n *NewAppDatabase) New() error {
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config/yamlConfig")
	return template.SaveTemplate(filePath, "database.yaml", n.getTemplate(), nil, "yaml")
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:30:44
 * @return {*}
 */
func (n *NewAppLogger) New() error {
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config/yamlConfig")
	return template.SaveTemplate(filePath, "logger.yaml", n.getTemplate(), nil, "yaml")
}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:30:44
 * @return {*}
 */
func (n *NewAppRouter) New() error {
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config/yamlConfig")
	return template.SaveTemplate(filePath, "router.yaml", n.getTemplate(), nil, "yaml")
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:30:57
 * @return {*}
 */
func (n *NewAppDatabase) getTemplate() string {
	return `master: 
	dsn: "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"`
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:30:57
 * @return {*}
 */
func (n *NewAppLogger) getTemplate() string {
	return `level: debug
callerDept: 4`
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:30:57
 * @return {*}
 */
func (n *NewAppRouter) getTemplate() string {
	return `addr: ":12000"`
}
