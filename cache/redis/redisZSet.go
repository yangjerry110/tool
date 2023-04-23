/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-12 16:37:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-12 17:50:34
 * @Description: 有序集合
 */
package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/cache"
)

/**
 * @description: ZAdd
 * @param {string} key
 * @param {*redis.Z} redisZ
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:40:19
 * @return {*}
 */
func (r *RedisCache) ZAdd(key string, redisZ *redis.Z) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZAdd(ctx, key, redisZ).Result()
	return r
}

/**
 * @description: ZAddContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {*redis.Z} redisZ
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:40:25
 * @return {*}
 */
func (r *RedisCache) ZAddContext(ctx context.Context, key string, redisZ *redis.Z) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZAdd(ctx, key, redisZ).Result()
	return r
}

/**
 * @description: ZCard
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:40:31
 * @return {*}
 */
func (r *RedisCache) ZCard(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZCard(ctx, key).Result()
	return r
}

/**
 * @description: ZCardContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:40:37
 * @return {*}
 */
func (r *RedisCache) ZCardContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZCard(ctx, key).Result()
	return r
}

/**
 * @description: ZCount
 * @param {string} key
 * @param {string} min
 * @param {string} max
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:40:44
 * @return {*}
 */
func (r *RedisCache) ZCount(key string, min string, max string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZCount(ctx, key, min, max).Result()
	return r
}

/**
 * @description: ZCountContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} min
 * @param {string} max
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:40:50
 * @return {*}
 */
func (r *RedisCache) ZCountContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZCount(ctx, key, min, max).Result()
	return r
}

/**
 * @description: ZIncrBy
 * @param {string} key
 * @param {float64} incr
 * @param {string} member
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:40:56
 * @return {*}
 */
func (r *RedisCache) ZIncrBy(key string, incr float64, member string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZIncrBy(ctx, key, incr, member).Result()
	return r
}

/**
 * @description: ZIncrByContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {float64} incr
 * @param {string} member
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:41:02
 * @return {*}
 */
func (r *RedisCache) ZIncrByContext(ctx context.Context, key string, incr float64, member string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZIncrBy(ctx, key, incr, member).Result()
	return r
}

/**
 * @description: ZRangeByScore
 * @param {string} key
 * @param {*redis.ZRangeBy} zRangeBy
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:42:30
 * @return {*}
 */
func (r *RedisCache) ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZRangeByScore(ctx, key, zRangeBy).Result()
	return r
}

/**
 * @description: ZRangeByScoreContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {*redis.ZRangeBy} zRangeBy
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:42:37
 * @return {*}
 */
func (r *RedisCache) ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZRangeByScore(ctx, key, zRangeBy).Result()
	return r
}

/**
 * @description: ZRevRangeByScore
 * @param {string} key
 * @param {*redis.ZRangeBy} zRangeBy
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:43:37
 * @return {*}
 */
func (r *RedisCache) ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZRevRangeByScore(ctx, key, zRangeBy).Result()
	return r
}

/**
 * @description: ZRevRangeByScoreContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {*redis.ZRangeBy} zRangeBy
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:43:44
 * @return {*}
 */
func (r *RedisCache) ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZRevRangeByScore(ctx, key, zRangeBy).Result()
	return r
}

/**
 * @description: ZRangeByScoreWithScores
 * @param {string} key
 * @param {*redis.ZRangeBy} zRangeBy
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:44:25
 * @return {*}
 */
func (r *RedisCache) ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZRangeByScoreWithScores(ctx, key, zRangeBy).Result()
	return r
}

/**
 * @description: ZRangeByScoreWithScoresContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {*redis.ZRangeBy} zRangeBy
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:44:32
 * @return {*}
 */
func (r *RedisCache) ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZRangeByScoreWithScores(ctx, key, zRangeBy).Result()
	return r
}

/**
 * @description: ZRem
 * @param {string} key
 * @param {...interface{}} members
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:49:16
 * @return {*}
 */
func (r *RedisCache) ZRem(key string, members ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZRem(ctx, key, members...).Result()
	return r
}

/**
 * @description: ZRemContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {...interface{}} members
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:49:23
 * @return {*}
 */
func (r *RedisCache) ZRemContext(ctx context.Context, key string, members ...interface{}) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZRem(ctx, key, members...).Result()
	return r
}

/**
 * @description: ZRemRangeByRank
 * @param {string} key
 * @param {int64} start
 * @param {int64} stop
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:49:40
 * @return {*}
 */
func (r *RedisCache) ZRemRangeByRank(key string, start int64, stop int64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZRemRangeByRank(ctx, key, start, stop).Result()
	return r
}

/**
 * @description: ZRemRangeByRankContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {int64} start
 * @param {int64} stop
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:49:49
 * @return {*}
 */
func (r *RedisCache) ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZRemRangeByRank(ctx, key, start, stop).Result()
	return r
}

/**
 * @description: ZRemRangeByScore
 * @param {string} key
 * @param {string} min
 * @param {string} max
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:49:58
 * @return {*}
 */
func (r *RedisCache) ZRemRangeByScore(key string, min string, max string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZRemRangeByScore(ctx, key, min, max).Result()
	return r
}

/**
 * @description: ZRemRangeByScoreContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} min
 * @param {string} max
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:50:04
 * @return {*}
 */
func (r *RedisCache) ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZRemRangeByScore(ctx, key, min, max).Result()
	return r
}

/**
 * @description: ZScore
 * @param {string} key
 * @param {string} member
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:50:12
 * @return {*}
 */
func (r *RedisCache) ZScore(key string, member string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZScore(ctx, key, member).Result()
	return r
}

/**
 * @description: ZScoreContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} member
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:50:19
 * @return {*}
 */
func (r *RedisCache) ZScoreContext(ctx context.Context, key string, member string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZScore(ctx, key, member).Result()
	return r
}

/**
 * @description: ZRank
 * @param {string} key
 * @param {string} member
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:50:25
 * @return {*}
 */
func (r *RedisCache) ZRank(key string, member string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.ZRank(ctx, key, member).Result()
	return r
}

/**
 * @description: ZRankContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} member
 * @author: Jerry.Yang
 * @date: 2022-10-12 17:50:33
 * @return {*}
 */
func (r *RedisCache) ZRankContext(ctx context.Context, key string, member string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.ZRank(ctx, key, member).Result()
	return r
}
