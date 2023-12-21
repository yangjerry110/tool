/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 11:12:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 15:27:06
 * @Description:
 */
package cache

import (
	"fmt"

	"github.com/yangjerry110/tool/internal/cmd/config"
	"github.com/yangjerry110/tool/internal/cmd/template"
)

type NewAppCacheRedis struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:14:02
 * @return {*}
 */
func (n *NewAppCacheRedis) New() error {

	// The structure that needs to be rendered
	type Data struct {
		Time string
	}

	// Set Data
	data := &Data{}
	data.Time = template.GetFormatNowTime()

	// return
	filePath := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "internal/cache")
	return template.SaveTemplate(filePath, "redis.go", n.getTemplate(), data)
}

/**
 * @description: getTemplate
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:13:49
 * @return {*}
 */
func (n *NewAppCacheRedis) getTemplate() string {
	return `package cache`
}
