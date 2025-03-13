/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 11:12:09
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 16:39:13
 * @Description:
 */
package cache

import (
	"fmt"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
	"github.com/yangjerry110/tool/cmd/gen-tool/internal/template"
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
	return `package cache

	import (
		"github.com/go-redis/redis/v8"
		"github.com/yangjerry110/tool/cache"
	)
	
	/**
	 * @description: CreateRedisClient
	 * @param {string} redisName
	 * @author: Jerry.Yang
	 * @date: {{.Time}}
	 * @return {*}
	 */
	func CreateRedisClient(redisName string) *redis.Client {
	
		// Get Client
		client, err := cache.CreateRedisCache().GetClient(redisName)
		if err != nil {
			panic(err)
		}
		return client
	}
	`
}
