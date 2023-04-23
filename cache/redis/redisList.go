/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-12 14:42:29
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-12 16:24:43
 * @Description: redis list
 */
package redis

import (
	"context"

	"github.com/yangjerry110/tool/cache"
)

/**
 * @description: LPush
 * @param {string} key
 * @param {...interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:45:10
 * @return {*}
 */
func (r *RedisCache) LPush(key string, value ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LPush(ctx, key, value...).Result()
	return r
}

/**
 * @description: LPushContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {...interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:45:53
 * @return {*}
 */
func (r *RedisCache) LPushContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LPush(ctx, key, value...).Result()
	return r
}

/**
 * @description: LPushX
 * @param {string} key
 * @param {...interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:47:17
 * @return {*}
 */
func (r *RedisCache) LPushX(key string, value ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LPushX(ctx, key, value...).Result()
	return r
}

/**
 * @description: LPushXContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {...interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:47:24
 * @return {*}
 */
func (r *RedisCache) LPushXContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LPushX(ctx, key, value...).Result()
	return r
}

/**
 * @description: RPop
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:48:45
 * @return {*}
 */
func (r *RedisCache) RPop(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.RPop(ctx, key).Result()
	return r
}

/**
 * @description: RPopContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:48:52
 * @return {*}
 */
func (r *RedisCache) RPopContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.RPop(ctx, key).Result()
	return r
}

/**
 * @description: RPush
 * @param {string} key
 * @param {...interface{}} values
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:49:57
 * @return {*}
 */
func (r *RedisCache) RPush(key string, values ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.RPush(ctx, key, values...).Result()
	return r
}

/**
 * @description: RPushContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {...interface{}} values
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:51:17
 * @return {*}
 */
func (r *RedisCache) RPushContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.RPush(ctx, key, values...).Result()
	return r
}

/**
 * @description: RPushX
 * @param {string} key
 * @param {...interface{}} values
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:52:26
 * @return {*}
 */
func (r *RedisCache) RPushX(key string, values ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.RPushX(ctx, key, values...).Result()
	return r
}

/**
 * @description: RPushXContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {...interface{}} values
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:52:34
 * @return {*}
 */
func (r *RedisCache) RPushXContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.RPushX(ctx, key, values...).Result()
	return r
}

/**
 * @description: LPop
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:14:01
 * @return {*}
 */
func (r *RedisCache) LPop(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LPop(ctx, key).Result()
	return r
}

/**
 * @description: LPopContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:14:13
 * @return {*}
 */
func (r *RedisCache) LPopContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LPop(ctx, key).Result()
	return r
}

/**
 * @description: LLen
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:15:19
 * @return {*}
 */
func (r *RedisCache) LLen(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LLen(ctx, key).Result()
	return r
}

/**
 * @description: LLenContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:15:29
 * @return {*}
 */
func (r *RedisCache) LLenContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LLen(ctx, key).Result()
	return r
}

/**
 * @description: LRange
 * @param {string} key
 * @param {int64} start
 * @param {int64} stop
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:17:48
 * @return {*}
 */
func (r *RedisCache) LRange(key string, start int64, stop int64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LRange(ctx, key, start, stop).Result()
	return r
}

/**
 * @description: LRangeContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {int64} start
 * @param {int64} stop
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:17:55
 * @return {*}
 */
func (r *RedisCache) LRangeContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LRange(ctx, key, start, stop).Result()
	return r
}

/**
 * @description: LRem
 * @param {string} key
 * @param {int64} count
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:19:44
 * @return {*}
 */
func (r *RedisCache) LRem(key string, count int64, value interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LRem(ctx, key, count, value).Result()
	return r
}

/**
 * @description: LRemContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {int64} count
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:19:52
 * @return {*}
 */
func (r *RedisCache) LRemContext(ctx context.Context, key string, count int64, value interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LRem(ctx, key, count, value).Result()
	return r
}

/**
 * @description: LIndex
 * @param {string} key
 * @param {int64} index
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:21:07
 * @return {*}
 */
func (r *RedisCache) LIndex(key string, index int64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LIndex(ctx, key, index).Result()
	return r
}

/**
 * @description: LIndexContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {int64} index
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:21:16
 * @return {*}
 */
func (r *RedisCache) LIndexContext(ctx context.Context, key string, index int64) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LIndex(ctx, key, index).Result()
	return r
}

/**
 * @description: LInsert
 * @param {string} key
 * @param {string} op 在什么位置，前还是后before or after
 * @param {interface{}} pivot 在什么数据插入
 * @param {interface{}} value 插入的数据
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:24:00
 * @return {*}
 */
func (r *RedisCache) LInsert(key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.LInsert(ctx, key, op, pivot, value).Result()
	return r
}

/**
 * @description: LInsertContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} op
 * @param {interface{}} pivot
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:24:08
 * @return {*}
 */
func (r *RedisCache) LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.LInsert(ctx, key, op, pivot, value).Result()
	return r
}
