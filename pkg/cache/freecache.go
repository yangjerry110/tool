/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-25 18:01:24
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:26:29
 * @Description: freecache
 */
package cache

import (
	"context"
	"time"

	"github.com/coocood/freecache"
	"github.com/go-redis/redis/v8"

	mytoolFreecache "github.com/yangjerry110/tool/cache/freecache"
)

type CachePkgFreeCache struct {
	FreeCache *freecache.Cache
	CacheVal  interface{}
	CacheErr  error
	Size      int
	Timer     freecache.Timer
}

/**
 * @description: Client
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:26:28
 * @return {*}
 */
func (c *CachePkgFreeCache) Client(filePath string, fileName string) CachePkgInterface {
	cacheFreeCache := &mytoolFreecache.FreeCache{}
	CreateCache(cacheFreeCache).CacheInterface.Client(filePath, fileName)
	c.FreeCache = cacheFreeCache.FreeCache
	return c
}

/**
 * @description: CreateDefaultClient
 * @author: Jerry.Yang
 * @date: 2022-10-25 17:47:28
 * @return {*}
 */
func (c *CachePkgFreeCache) CreateDefaultClient() CachePkgInterface {
	cacheFreeCache := &mytoolFreecache.FreeCache{
		Size:  c.Size,
		Timer: c.Timer,
	}
	cacheFreeCache.CreateDefaultClient()
	c.FreeCache = cacheFreeCache.FreeCache
	return c
}

// 检查配置
func (c *CachePkgFreeCache) CheckConfig() error { return nil }

// 获取结果和，错误
func (c *CachePkgFreeCache) Result() (interface{}, error) {
	return c.CacheVal, c.CacheErr
}

// 获取结果
func (c *CachePkgFreeCache) GetResult() interface{} {
	return c.CacheVal
}

// 获取错误
func (c *CachePkgFreeCache) GetErr() error {
	return c.CacheErr
}

// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
func (c *CachePkgFreeCache) Set(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheFreeCache := &mytoolFreecache.FreeCache{FreeCache: c.FreeCache}
	CreateCache(cacheFreeCache).CacheInterface.Set(key, value, expiration)
	c.CacheVal = cacheFreeCache.FreeCacheVal
	c.CacheErr = cacheFreeCache.FreeCacheErr
	return c
}
func (c *CachePkgFreeCache) SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return c
}

// 查询数据库中名称为key的value值
func (c *CachePkgFreeCache) Get(key string) CachePkgInterface {
	cacheFreeCache := &mytoolFreecache.FreeCache{FreeCache: c.FreeCache}
	CreateCache(cacheFreeCache).CacheInterface.Get(key)
	c.CacheVal = cacheFreeCache.FreeCacheVal
	c.CacheErr = cacheFreeCache.FreeCacheErr
	return c
}
func (c *CachePkgFreeCache) GetContext(ctx context.Context, key string) CachePkgInterface { return c }

// 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
func (c *CachePkgFreeCache) SetNX(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return c
}

// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
func (c *CachePkgFreeCache) MGet(keys ...string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) MGetContext(ctx context.Context, keys ...string) CachePkgInterface {
	return c
}

// 批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
func (c *CachePkgFreeCache) MSet(pairs ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) MsetContext(ctx context.Context, pairs ...interface{}) CachePkgInterface {
	return c
}

// Incr函数每次加一,key对应的值必须是整数或nil
// 否则会报错incr key1: ERR value is not an integer or out of range
func (c *CachePkgFreeCache) Incr(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) IncrContext(ctx context.Context, key string) CachePkgInterface { return c }

// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
func (c *CachePkgFreeCache) IncrBy(key string, value int64) CachePkgInterface { return c }
func (c *CachePkgFreeCache) IncrByContext(ctx context.Context, key string, value int64) CachePkgInterface {
	return c
}

// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
func (c *CachePkgFreeCache) IncrByFloat(key string, value float64) CachePkgInterface { return c }
func (c *CachePkgFreeCache) IncrByFloatContext(ctx context.Context, key string, value float64) CachePkgInterface {
	return c
}

// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
func (c *CachePkgFreeCache) Decr(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) DecrContext(ctx context.Context, key string) CachePkgInterface { return c }

// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
func (c *CachePkgFreeCache) DecrBy(key string, decrement int64) CachePkgInterface { return c }
func (c *CachePkgFreeCache) DecrByContext(ctx context.Context, key string, decrement int64) CachePkgInterface {
	return c
}

// 删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
func (c *CachePkgFreeCache) Del(key string) CachePkgInterface {
	cacheFreeCache := &mytoolFreecache.FreeCache{FreeCache: c.FreeCache}
	CreateCache(cacheFreeCache).CacheInterface.Del(key)
	c.CacheVal = cacheFreeCache.FreeCacheVal
	c.CacheErr = cacheFreeCache.FreeCacheErr
	return c
}
func (c *CachePkgFreeCache) DelContext(ctx context.Context, key string) CachePkgInterface { return c }

// 设置key的过期时间,单位秒
func (c *CachePkgFreeCache) Expire(key string, expiration time.Duration) CachePkgInterface { return c }
func (c *CachePkgFreeCache) ExpireContext(ctx context.Context, key string, expiration time.Duration) CachePkgInterface {
	return c
}

// 给数据库中名称为key的string值追加value
func (c *CachePkgFreeCache) Append(key, value string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) AppendContext(ctx context.Context, key string, value string) CachePkgInterface {
	return c
}

// 根据key和field字段设置，field字段的值
func (c *CachePkgFreeCache) HSet(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return c
}

// 根据key和field字段，查询field字段的值
func (c *CachePkgFreeCache) HGet(key string, field string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) HGetContext(ctx context.Context, key string, field string) CachePkgInterface {
	return c
}

// 如果field字段不存在，则设置hash字段值
func (c *CachePkgFreeCache) HSetNX(key string, field string, value interface{}, expiration time.Duration) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) CachePkgInterface {
	return c
}

// 根据key和多个字段名和字段值，批量设置hash字段值
func (c *CachePkgFreeCache) HMSet(key string, pairs ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) HMsetContext(ctx context.Context, key string, pairs ...interface{}) CachePkgInterface {
	return c
}

// 根据key和多个字段名，批量查询多个hash字段值
func (c *CachePkgFreeCache) HMGet(key string, fields ...string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) HMGetContext(ctx context.Context, key string, fields ...string) CachePkgInterface {
	return c
}

// 根据key和field字段，累加字段的数值
func (c *CachePkgFreeCache) HIncrBy(key string, field string, value int64) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) HIncrByContext(ctx context.Context, key string, field string, value int64) CachePkgInterface {
	return c
}

// 根据key和field字段，累加字段的数值
func (c *CachePkgFreeCache) HIncrByFloat(key string, field string, value float64) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) HIncrByFloatContext(ctx context.Context, key string, field string, value float64) CachePkgInterface {
	return c
}

// 根据key和字段名，删除hash字段，支持批量删除hash字段
func (c *CachePkgFreeCache) HDel(key string, fields ...string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) HDelContext(ctx context.Context, key string, fields ...string) CachePkgInterface {
	return c
}

// 根据key返回所有字段名
func (c *CachePkgFreeCache) HKeys(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) HKeysContext(ctx context.Context, key string) CachePkgInterface { return c }

// 根据key，查询hash的字段数量
func (c *CachePkgFreeCache) HLen(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) HlenContext(ctx context.Context, key string) CachePkgInterface { return c }

// 检测hash字段名是否存在。
func (c *CachePkgFreeCache) HExists(key string, field string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) HExistsContext(ctx context.Context, key string, field string) CachePkgInterface {
	return c
}

// 从列表左边插入数据
func (c *CachePkgFreeCache) LPush(key string, value ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) LPushContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	return c
}

// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
func (c *CachePkgFreeCache) LPushX(key string, value ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) LPushXContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	return c
}

// 从列表的右边删除第一个数据，并返回删除的数据
func (c *CachePkgFreeCache) RPop(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) RPopContext(ctx context.Context, key string) CachePkgInterface { return c }

// 从列表右边插入数据
func (c *CachePkgFreeCache) RPush(key string, values ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) RPushContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface {
	return c
}

// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
func (c *CachePkgFreeCache) RPushX(key string, values ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) RPushXContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface {
	return c
}

// 从列表左边删除第一个数据，并返回删除的数据
func (c *CachePkgFreeCache) LPop(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) LPopContext(ctx context.Context, key string) CachePkgInterface { return c }

// 返回列表的大小
func (c *CachePkgFreeCache) LLen(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) LLenContext(ctx context.Context, key string) CachePkgInterface { return c }

// 返回列表的一个范围内的数据，也可以返回全部数据
func (c *CachePkgFreeCache) LRange(key string, start int64, stop int64) CachePkgInterface { return c }
func (c *CachePkgFreeCache) LRangeContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface {
	return c
}

// 删除列表中的数据
func (c *CachePkgFreeCache) LRem(key string, count int64, value interface{}) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) LRemContext(ctx context.Context, key string, count int64, value interface{}) CachePkgInterface {
	return c
}

// 根据索引坐标，查询列表中的数据
func (c *CachePkgFreeCache) LIndex(key string, index int64) CachePkgInterface { return c }
func (c *CachePkgFreeCache) LIndexContext(ctx context.Context, key string, index int64) CachePkgInterface {
	return c
}

// 在指定位置插入数据
func (c *CachePkgFreeCache) LInsert(key string, op string, pivot interface{}, value interface{}) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) CachePkgInterface {
	return c
}

// 添加集合元素
func (c *CachePkgFreeCache) SAdd(key string, value ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) SAddContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	return c
}

// 获取集合元素个数
func (c *CachePkgFreeCache) SCard(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) SCardContext(ctx context.Context, key string) CachePkgInterface { return c }

// 判断元素是否在集合中
func (c *CachePkgFreeCache) SIsMember(key string, value interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) SIsMemberContext(ctx context.Context, key string, value interface{}) CachePkgInterface {
	return c
}

// 获取集合中所有的元素
func (c *CachePkgFreeCache) SMembers(key string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) SMembersContext(ctx context.Context, key string) CachePkgInterface {
	return c
}

// 删除集合元素
func (c *CachePkgFreeCache) SRem(key string, value interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) SRemContext(ctx context.Context, key string, value interface{}) CachePkgInterface {
	return c
}

// 随机返回集合中的元素，并且删除返回的元素
func (c *CachePkgFreeCache) SPop(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) SPopContext(ctx context.Context, key string) CachePkgInterface { return c }

// 随机返回集合中的元素，并且删除返回的元素
func (c *CachePkgFreeCache) SPopN(key string, num int64) CachePkgInterface { return c }
func (c *CachePkgFreeCache) SPopNContext(ctx context.Context, key string, num int64) CachePkgInterface {
	return c
}

// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (c *CachePkgFreeCache) ZAdd(key string, redisZ *redis.Z) CachePkgInterface { return c }
func (c *CachePkgFreeCache) ZAddContext(ctx context.Context, key string, redisZ *redis.Z) CachePkgInterface {
	return c
}

// 返回集合元素个数
func (c *CachePkgFreeCache) ZCard(key string) CachePkgInterface                             { return c }
func (c *CachePkgFreeCache) ZCardContext(ctx context.Context, key string) CachePkgInterface { return c }

// 统计某个分数范围内的元素个数
func (c *CachePkgFreeCache) ZCount(key string, min string, max string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) ZCountContext(ctx context.Context, key string, min string, max string) CachePkgInterface {
	return c
}

// 增加元素的分数
func (c *CachePkgFreeCache) ZIncrBy(key string, incr float64, member string) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) ZIncrByContext(ctx context.Context, key string, incr float64, member string) CachePkgInterface {
	return c
}

// 返回集合中某个索引范围的元素，根据分数从小到大排序
func (c *CachePkgFreeCache) ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return c
}

// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
func (c *CachePkgFreeCache) ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return c
}

// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
func (c *CachePkgFreeCache) ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return c
}

// 删除集合元素
func (c *CachePkgFreeCache) ZRem(key string, members ...interface{}) CachePkgInterface { return c }
func (c *CachePkgFreeCache) ZRemContext(ctx context.Context, key string, members ...interface{}) CachePkgInterface {
	return c
}

// 根据索引范围删除元素
func (c *CachePkgFreeCache) ZRemRangeByRank(key string, start int64, stop int64) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface {
	return c
}

// 根据分数范围删除元素
func (c *CachePkgFreeCache) ZRemRangeByScore(key string, min string, max string) CachePkgInterface {
	return c
}
func (c *CachePkgFreeCache) ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) CachePkgInterface {
	return c
}

// 查询元素对应的分数
func (c *CachePkgFreeCache) ZScore(key string, member string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) ZScoreContext(ctx context.Context, key string, member string) CachePkgInterface {
	return c
}

// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (c *CachePkgFreeCache) ZRank(key string, member string) CachePkgInterface { return c }
func (c *CachePkgFreeCache) ZRankContext(ctx context.Context, key string, member string) CachePkgInterface {
	return c
}
