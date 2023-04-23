/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-11 18:04:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-13 16:13:48
 * @Description: redis
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
func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.Set(ctx, key, value, expiration).Err()
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
func (r *RedisCache) SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	r.RedisErr = r.RedisClient.Set(ctx, key, value, expiration).Err()
	return r
}

/**
 * @description: Get
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:25:09
 * @return {*}
 */
func (r *RedisCache) Get(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.Get(ctx, key).Result()
	return r
}

/**
 * @description: GetContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:30:00
 * @return {*}
 */
func (r *RedisCache) GetContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.Get(ctx, key).Result()
	return r
}

/**
 * @description: SetNX
 * @param {string} key
 * @param {interface{}} value
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:21
 * @return {*}
 */
func (r *RedisCache) SetNX(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.SetNX(ctx, key, value, expiration).Err()
	return r
}

/**
 * @description: SetNXContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {interface{}} value
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:31
 * @return {*}
 */
func (r *RedisCache) SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	r.RedisErr = r.RedisClient.SetNX(ctx, key, value, expiration).Err()
	return r
}

/**
 * @description: MSet
 * @param {...interface{}} pairs
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:38
 * @return {*}
 */
func (r *RedisCache) MSet(pairs ...interface{}) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.MSet(ctx, pairs...).Err()
	return r
}

/**
 * @description: MsetContext
 * @param {context.Context} ctx
 * @param {...interface{}} pairs
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:45
 * @return {*}
 */
func (r *RedisCache) MsetContext(ctx context.Context, pairs ...interface{}) cache.CacheInterface {
	r.RedisErr = r.RedisClient.MSet(ctx, pairs...).Err()
	return r
}

/**
 * @description: MGet
 * @param {...string} keys
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:52
 * @return {*}
 */
func (r *RedisCache) MGet(keys ...string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisVal, r.RedisErr = r.RedisClient.MGet(ctx, keys...).Result()
	return r
}

/**
 * @description: MGetContext
 * @param {context.Context} ctx
 * @param {...string} keys
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:22:59
 * @return {*}
 */
func (r *RedisCache) MGetContext(ctx context.Context, keys ...string) cache.CacheInterface {
	r.RedisVal, r.RedisErr = r.RedisClient.MGet(ctx, keys...).Result()
	return r
}

/**
 * @description: Incr
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:17
 * @return {*}
 */
func (r *RedisCache) Incr(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.Incr(ctx, key).Err()
	return r
}

/**
 * @description: IncrContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:25
 * @return {*}
 */
func (r *RedisCache) IncrContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisErr = r.RedisClient.Incr(ctx, key).Err()
	return r
}

/**
 * @description: IncrBy
 * @param {string} key
 * @param {int64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:33
 * @return {*}
 */
func (r *RedisCache) IncrBy(key string, value int64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.IncrBy(ctx, key, value).Err()
	return r
}

/**
 * @description: IncrByContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {int64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:40
 * @return {*}
 */
func (r *RedisCache) IncrByContext(ctx context.Context, key string, value int64) cache.CacheInterface {
	r.RedisErr = r.RedisClient.IncrBy(ctx, key, value).Err()
	return r
}

/**
 * @description: IncrByFloat
 * @param {string} key
 * @param {float64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:48
 * @return {*}
 */
func (r *RedisCache) IncrByFloat(key string, value float64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.IncrByFloat(ctx, key, value).Err()
	return r
}

/**
 * @description: IncrByFloatContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {float64} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:23:55
 * @return {*}
 */
func (r *RedisCache) IncrByFloatContext(ctx context.Context, key string, value float64) cache.CacheInterface {
	r.RedisErr = r.RedisClient.IncrByFloat(ctx, key, value).Err()
	return r
}

/**
 * @description: Decr
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:03
 * @return {*}
 */
func (r *RedisCache) Decr(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.Decr(ctx, key).Err()
	return r
}

/**
 * @description: DecrContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-13 16:16:26
 * @return {*}
 */
func (r *RedisCache) DecrContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisErr = r.RedisClient.Decr(ctx, key).Err()
	return r
}

/**
 * @description: DecrBy
 * @param {string} key
 * @param {int64} decrement
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:10
 * @return {*}
 */
func (r *RedisCache) DecrBy(key string, decrement int64) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.DecrBy(ctx, key, decrement).Err()
	return r
}

/**
 * @description: DecrByContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {int64} decrement
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:17
 * @return {*}
 */
func (r *RedisCache) DecrByContext(ctx context.Context, key string, decrement int64) cache.CacheInterface {
	r.RedisErr = r.RedisClient.DecrBy(ctx, key, decrement).Err()
	return r
}

/**
 * @description: Del
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:24
 * @return {*}
 */
func (r *RedisCache) Del(key string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.Del(ctx, key).Err()
	return r
}

/**
 * @description: DelContext
 * @param {context.Context} ctx
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:31
 * @return {*}
 */
func (r *RedisCache) DelContext(ctx context.Context, key string) cache.CacheInterface {
	r.RedisErr = r.RedisClient.Del(ctx, key).Err()
	return r
}

/**
 * @description: Expire
 * @param {string} key
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:38
 * @return {*}
 */
func (r *RedisCache) Expire(key string, expiration time.Duration) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.Expire(ctx, key, expiration).Err()
	return r
}

/**
 * @description: ExpireContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {time.Duration} expiration
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:45
 * @return {*}
 */
func (r *RedisCache) ExpireContext(ctx context.Context, key string, expiration time.Duration) cache.CacheInterface {
	r.RedisErr = r.RedisClient.Expire(ctx, key, expiration).Err()
	return r
}

/**
 * @description: Append
 * @param {string} key
 * @param {string} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:24:52
 * @return {*}
 */
func (r *RedisCache) Append(key string, value string) cache.CacheInterface {
	ctx := context.Background()
	r.RedisErr = r.RedisClient.Append(ctx, key, value).Err()
	return r
}

/**
 * @description: AppendContext
 * @param {context.Context} ctx
 * @param {string} key
 * @param {string} value
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:25:02
 * @return {*}
 */
func (r *RedisCache) AppendContext(ctx context.Context, key string, value string) cache.CacheInterface {
	r.RedisErr = r.RedisClient.Append(ctx, key, value).Err()
	return r
}

/**
 * @description: Result
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:25:16
 * @return {*}
 */
func (r *RedisCache) Result() (interface{}, error) {
	return r.RedisVal, r.RedisErr
}

/**
 * @description: GetResult
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:25:23
 * @return {*}
 */
func (r *RedisCache) GetResult() interface{} {
	return r.RedisVal
}

/**
 * @description: GetErr
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:25:32
 * @return {*}
 */
func (r *RedisCache) GetErr() error {
	return r.RedisErr
}
