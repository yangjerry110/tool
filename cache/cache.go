/*
 * @Author: Jerry.Yang
 * @Date: 2023-11-30 16:16:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 16:22:16
 * @Description: cache
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
	return cache.CreateCache(&cacheredis.RedisClient{})
}
