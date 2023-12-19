/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 11:39:44
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 16:38:38
 * @Description:
 */
package cache

import (
	"github.com/yangjerry110/tool/internal/cache"
	cacheredis "github.com/yangjerry110/tool/internal/cache/redis"
	"github.com/yangjerry110/tool/internal/conf"
)

/**
 * @description: CreateRedisConf
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:41:29
 * @return {*}
 */
func CreateRedisConf() error {
	return conf.CreateConf(&cacheredis.RedisConf{}).SetConfig()
}

/**
 * @description: CreateRedisCache
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:42:54
 * @return {*}
 */
func CreateRedisCache() cache.Cache {
	return CreateCache(&cacheredis.RedisClient{})
}
