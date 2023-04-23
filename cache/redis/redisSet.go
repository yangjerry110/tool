/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-12 16:24:55
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-12 16:37:27
 * @Description: redis set
 */
package redis

import (
	"context"

	"github.com/yangjerry110/tool/cache"
)

/**
 * @description: SAdd
 * @param {string} key
 * @param {...interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:26:58
 * @return {*}
 */
func (r *RedisCache) SAdd(key string, value ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.SAdd(ctx, key, value...).Result()
	return r
}

/**
 * @description: SAddContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {...interface{}} value
 * @author: Jerry.Yang
 * @date:
 * @return {*}
 */
func (r *RedisCache) SAddContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.SAdd(ctx, key, value...).Result()
	return r
}

/**
 * @description: SCard
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:28:26
 * @return {*}
 */
func (r *RedisCache) SCard(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.SCard(ctx, key).Result()
	return r
}

/**
 * @description: SCardContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:28:32
 * @return {*}
 */
func (r *RedisCache) SCardContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.SCard(ctx, key).Result()
	return r
}

/**
 * @description: SIsMember
 * @param {string} key
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:30:11
 * @return {*}
 */
func (r *RedisCache) SIsMember(key string, value interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.SIsMember(ctx, key, value).Result()
	return r
}

/**
 * @description: SIsMemberContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:30:20
 * @return {*}
 */
func (r *RedisCache) SIsMemberContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.SIsMember(ctx, key, value).Result()
	return r
}

/**
 * @description: SMembers
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:31:43
 * @return {*}
 */
func (r *RedisCache) SMembers(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.SMIsMember(ctx, key).Result()
	return r
}

/**
 * @description: SMembersContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:31:54
 * @return {*}
 */
func (r *RedisCache) SMembersContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.SMIsMember(ctx, key).Result()
	return r
}

/**
 * @description: SRem
 * @param {string} key
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:36:48
 * @return {*}
 */
func (r *RedisCache) SRem(key string, value interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.SRem(ctx, key, value).Result()
	return r
}

/**
 * @description: SRemContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {interface{}} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:36:55
 * @return {*}
 */
func (r *RedisCache) SRemContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.SRem(ctx, key, value).Result()
	return r
}

/**
 * @description: SPop
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:37:04
 * @return {*}
 */
func (r *RedisCache) SPop(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.SPop(ctx, key).Result()
	return r
}

/**
 * @description: SPopContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:37:11
 * @return {*}
 */
func (r *RedisCache) SPopContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.SPop(ctx, key).Result()
	return r
}

/**
 * @description: SPopN
 * @param {string} key
 * @param {int64} num
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:37:17
 * @return {*}
 */
func (r *RedisCache) SPopN(key string, num int64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.SPopN(ctx, key, num).Result()
	return r
}

/**
 * @description: SPopNContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {int64} num
 * @author: Jerry.Yang
 * @date: 2022-10-12 16:37:24
 * @return {*}
 */
func (r *RedisCache) SPopNContext(ctx context.Context, key string, num int64) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.SPopN(ctx, key, num).Result()
	return r
}
