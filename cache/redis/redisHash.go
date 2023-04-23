/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-12 14:25:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-12 14:41:24
 * @Description: redis hash
 */
package redis

import (
	"context"
	"time"

	"github.com/yangjerry110/tool/cache"
)

/**
 * @description: Set
 * @param {string} key
 * @param {interface{}} value
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:05
 * @return {*}
 */
func (r *RedisCache) HSet(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.HSet(ctx, key, value, expiration).Err()
	return r
}

/**
 * @description: SetContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {interface{}} value
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:13
 * @return {*}
 */
func (r *RedisCache) HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	r.RedisErr = r.RedisClient.HSet(ctx, key, value, expiration).Err()
	return r
}

/**
 * @description: Get
 * @param {string} key
 * @param {string} field
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:25:09
 * @return {*}
 */
func (r *RedisCache) HGet(key string, field string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.HGet(ctx, key, field).Result()
	return r
}

/**
 * @description: GetContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} field
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:30:00
 * @return {*}
 */
func (r *RedisCache) HGetContext(ctx context.Context, key string, field string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.HGet(ctx, key, field).Result()
	return r
}

/**
 * @description: SetNX
 * @param {string} key
 * @param {string} field
 * @param {interface{}} value
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:21
 * @return {*}
 */
func (r *RedisCache) HSetNX(key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.HSetNX(ctx, key, field, value).Err()
	return r
}

/**
 * @description: SetNXContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} field
 * @param {interface{}} value
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:31
 * @return {*}
 */
func (r *RedisCache) HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	r.RedisErr = r.RedisClient.HSetNX(ctx, key, field, value).Err()
	return r
}

/**
 * @description: MSet
 * @param {string} key
 * @param {...interface{}} pairs
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:38
 * @return {*}
 */
func (r *RedisCache) HMSet(key string, pairs ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.HMSet(ctx, key, pairs...).Err()
	return r
}

/**
 * @description: MsetContext
 * @param {string} key
 * @param {context.Context} ctx
 * @param {...interface{}} pairs
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:45
 * @return {*}
 */
func (r *RedisCache) HMsetContext(ctx context.Context, key string, pairs ...interface{}) cache.CacheInterface {
	r.RedisErr = r.RedisClient.HMSet(ctx, key, pairs...).Err()
	return r
}

/**
 * @description: MGet
 * @param {string} key
 * @param {...string} fields
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:52
 * @return {*}
 */
func (r *RedisCache) HMGet(key string, fields ...string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.HMGet(ctx, key, fields...).Result()
	return r
}

/**
 * @description: MGetContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {...string} fields
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:59
 * @return {*}
 */
func (r *RedisCache) HMGetContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.HMGet(ctx, key, fields...).Result()
	return r
}

/**
 * @description: IncrBy
 * @param {string} key
 * @param {string} field
 * @param {int64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:33
 * @return {*}
 */
func (r *RedisCache) HIncrBy(key string, field string, value int64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.HIncrBy(ctx, key, field, value).Err()
	return r
}

/**
 * @description: IncrByContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} field
 * @param {int64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:40
 * @return {*}
 */
func (r *RedisCache) HIncrByContext(ctx context.Context, key string, field string, value int64) cache.CacheInterface {
	r.RedisErr = r.RedisClient.HIncrBy(ctx, key, field, value).Err()
	return r
}

/**
 * @description: IncrByFloat
 * @param {string} key
 * @param {string} field
 * @param {float64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:48
 * @return {*}
 */
func (r *RedisCache) HIncrByFloat(key string, field string, value float64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.HIncrByFloat(ctx, key, field, value).Err()
	return r
}

/**
 * @description: IncrByFloatContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} field
 * @param {float64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:55
 * @return {*}
 */
func (r *RedisCache) HIncrByFloatContext(ctx context.Context, key string, field string, value float64) cache.CacheInterface {
	r.RedisErr = r.RedisClient.HIncrByFloat(ctx, key, field, value).Err()
	return r
}

/**
 * @description: Del
 * @param {string} key
 * @param {...string} fields
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:24
 * @return {*}
 */
func (r *RedisCache) HDel(key string, fields ...string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.HDel(ctx, key, fields...).Err()
	return r
}

/**
 * @description: DelContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} fields
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:31
 * @return {*}
 */
func (r *RedisCache) HDelContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	r.RedisErr = r.RedisClient.HDel(ctx, key, fields...).Err()
	return r
}

/**
 * @description: HKeys
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:38:56
 * @return {*}
 */
func (r *RedisCache) HKeys(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.HKeys(ctx, key).Result()
	return r
}

/**
 * @description: HKeysContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:39:05
 * @return {*}
 */
func (r *RedisCache) HKeysContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.HKeys(ctx, key).Result()
	return r
}

/**
 * @description: HLen
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:40:11
 * @return {*}
 */
func (r *RedisCache) HLen(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.HLen(ctx, key).Result()
	return r
}

/**
 * @description: HlenContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:40:17
 * @return {*}
 */
func (r *RedisCache) HlenContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.HLen(ctx, key).Result()
	return r
}

/**
 * @description: HExists
 * @param {string} key
 * @param {string} field
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:41:32
 * @return {*}
 */
func (r *RedisCache) HExists(key string, field string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.HExists(ctx, key, field).Result()
	return r
}

/**
 * @description: HExistsContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} field
 * @author: Jerry.Yang
 * @date: 2022-10-12 14:41:41
 * @return {*}
 */
func (r *RedisCache) HExistsContext(ctx context.Context, key string, field string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.HExists(ctx, key, field).Result()
	return r
}
