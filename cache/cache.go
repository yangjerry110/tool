/*
 * @Author: Jerry.Yang
 * @Date: 2023-11-30 16:16:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-08 11:40:45
 * @Description: cache
 */
package cache

import (
	"github.com/yangjerry110/tool/internal/cache"
	cacheredis "github.com/yangjerry110/tool/internal/cache/redis"
)

/**
 * @description: DefaultCache
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:39:40
 * @return {*}
 */
var DefaultCache = &cacheredis.RedisClient{}

/**
 * @description: CreateCache
 * @param {...cache.Cache} caches
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:39:32
 * @return {*}
 */
func CreateCache(caches ...cache.Cache) cache.Cache {
	if len(caches) == 0 {
		return DefaultCache
	}
	return caches[0]
}
