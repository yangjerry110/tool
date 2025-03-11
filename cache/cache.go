/*
 * @Author: Jerry.Yang
 * @Date: 2023-11-30 16:16:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 17:16:59
 * @Description: cache
 */
package cache

import (
	"github.com/yangjerry110/tool/cache/internal/cache"
	cacheredis "github.com/yangjerry110/tool/cache/internal/cache/redis"
	"github.com/yangjerry110/tool/conf"
)

/**
 * @description: CreateRedisConf
 * @author: Jerry.Yang
 * @date: 2023-12-19 11:41:29
 * @return {*}
 */
func SetRedisConf() conf.Conf {
	return conf.CreateConf(&cacheredis.RedisConf{})
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
