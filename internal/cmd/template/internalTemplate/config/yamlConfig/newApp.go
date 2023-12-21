/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-18 17:33:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 11:11:25
 * @Description: yamlConfig newApp
 */
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
 * @description: NewAppRedis
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:01:39
 * @return {*}
 */
type NewAppRedis struct{}

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
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:02:09
 * @return {*}
 */
func (n *NewAppRedis) New() error {
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/config/yamlConfig")
	return template.SaveTemplate(filePath, "redis.yaml", n.getTemplate(), nil, "yaml")
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

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:11:10
 * @return {*}
 */
func (n *NewAppRedis) getTemplate() string {
	return `redis:
	addr: 127.0.0.1:6379  #[MUST]redis地址
	password: "password"          #redis密码
	database: 0           #redis db index, {default: 0}
   #    dial_timeout: 5000    #连接超时时间，单位: millisecond, {default: 5000}
   #    read_timeout: 1000    #读超时时间，单位: millisecond, {default: 1000}
   #    write_timeout: 1000   #写超时时间，单位: millisecond, {default: 1000}
   #    max_retries: 1        #最大重试次数, {default: 1}
   #    pool_size: 0          #最大连接数大小, {default: runtime.NumCPU*10}
   #    min_idle_conns: 0     #一直保持的空闲连接数(无论是否有请求),一般为0即可 {default: 0}`
}
