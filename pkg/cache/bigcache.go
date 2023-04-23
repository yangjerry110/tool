/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-25 17:41:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:26:14
 * @Description: bigcache
 */
package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache"
	"github.com/go-redis/redis/v8"
	mytoolBigCache "github.com/yangjerry110/tool/cache/bigcache"
)

type CachePkgBigCache struct {
	BigCache  *bigcache.BigCache
	CacheVal  interface{}
	CacheErr  error
	Eviction  time.Duration
	CleanTime time.Duration
}

/**
 * @description: Client
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:26:13
 * @return {*}
 */
func (b *CachePkgBigCache) Client(filePath string, fileName string) CachePkgInterface {
	cacheBig := &mytoolBigCache.BigCache{}
	CreateCache(cacheBig).CacheInterface.Client(filePath, fileName)
	b.BigCache = cacheBig.BigCache
	return b
}

/**
 * @description: CreateDefaultClient
 * @author: Jerry.Yang
 * @date: 2022-10-25 17:47:28
 * @return {*}
 */
func (b *CachePkgBigCache) CreateDefaultClient() CachePkgInterface {
	cacheBig := &mytoolBigCache.BigCache{
		Eviction:  b.Eviction,
		CleanTime: b.CleanTime,
	}
	cacheBig.CreateDefaultClient()
	b.BigCache = cacheBig.BigCache
	return b
}

// 检查配置
func (b *CachePkgBigCache) CheckConfig() error { return nil }

// 获取结果和，错误
func (b *CachePkgBigCache) Result() (interface{}, error) {
	return b.CacheVal, b.CacheErr
}

// 获取结果
func (b *CachePkgBigCache) GetResult() interface{} {
	return b.CacheVal
}

// 获取错误
func (b *CachePkgBigCache) GetErr() error {
	return b.CacheErr
}

// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
func (b *CachePkgBigCache) Set(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheBig := &mytoolBigCache.BigCache{BigCache: b.BigCache}
	CreateCache(cacheBig).CacheInterface.Set(key, value, expiration)
	b.CacheVal = cacheBig.BigCacheVal
	b.CacheErr = cacheBig.BigCacheErr
	return b
}
func (b *CachePkgBigCache) SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return b
}

// 查询数据库中名称为key的value值
func (b *CachePkgBigCache) Get(key string) CachePkgInterface {
	cacheBig := &mytoolBigCache.BigCache{BigCache: b.BigCache}
	CreateCache(cacheBig).CacheInterface.Get(key)
	b.CacheVal = cacheBig.BigCacheVal
	b.CacheErr = cacheBig.BigCacheErr
	return b
}
func (b *CachePkgBigCache) GetContext(ctx context.Context, key string) CachePkgInterface { return b }

// 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
func (b *CachePkgBigCache) SetNX(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return b
}

// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
func (b *CachePkgBigCache) MGet(keys ...string) CachePkgInterface { return b }
func (b *CachePkgBigCache) MGetContext(ctx context.Context, keys ...string) CachePkgInterface {
	return b
}

// 批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
func (b *CachePkgBigCache) MSet(pairs ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) MsetContext(ctx context.Context, pairs ...interface{}) CachePkgInterface {
	return b
}

// Incr函数每次加一,key对应的值必须是整数或nil
// 否则会报错incr key1: ERR value is not an integer or out of range
func (b *CachePkgBigCache) Incr(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) IncrContext(ctx context.Context, key string) CachePkgInterface { return b }

// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
func (b *CachePkgBigCache) IncrBy(key string, value int64) CachePkgInterface { return b }
func (b *CachePkgBigCache) IncrByContext(ctx context.Context, key string, value int64) CachePkgInterface {
	return b
}

// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
func (b *CachePkgBigCache) IncrByFloat(key string, value float64) CachePkgInterface { return b }
func (b *CachePkgBigCache) IncrByFloatContext(ctx context.Context, key string, value float64) CachePkgInterface {
	return b
}

// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
func (b *CachePkgBigCache) Decr(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) DecrContext(ctx context.Context, key string) CachePkgInterface { return b }

// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
func (b *CachePkgBigCache) DecrBy(key string, decrement int64) CachePkgInterface { return b }
func (b *CachePkgBigCache) DecrByContext(ctx context.Context, key string, decrement int64) CachePkgInterface {
	return b
}

// 删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
func (b *CachePkgBigCache) Del(key string) CachePkgInterface {
	cacheBig := &mytoolBigCache.BigCache{BigCache: b.BigCache}
	CreateCache(cacheBig).CacheInterface.Del(key)
	b.CacheVal = cacheBig.BigCacheVal
	b.CacheErr = cacheBig.BigCacheErr
	return b
}
func (b *CachePkgBigCache) DelContext(ctx context.Context, key string) CachePkgInterface { return b }

// 设置key的过期时间,单位秒
func (b *CachePkgBigCache) Expire(key string, expiration time.Duration) CachePkgInterface { return b }
func (b *CachePkgBigCache) ExpireContext(ctx context.Context, key string, expiration time.Duration) CachePkgInterface {
	return b
}

// 给数据库中名称为key的string值追加value
func (b *CachePkgBigCache) Append(key, value string) CachePkgInterface { return b }
func (b *CachePkgBigCache) AppendContext(ctx context.Context, key string, value string) CachePkgInterface {
	return b
}

// 根据key和field字段设置，field字段的值
func (b *CachePkgBigCache) HSet(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	return b
}

// 根据key和field字段，查询field字段的值
func (b *CachePkgBigCache) HGet(key string, field string) CachePkgInterface { return b }
func (b *CachePkgBigCache) HGetContext(ctx context.Context, key string, field string) CachePkgInterface {
	return b
}

// 如果field字段不存在，则设置hash字段值
func (b *CachePkgBigCache) HSetNX(key string, field string, value interface{}, expiration time.Duration) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) CachePkgInterface {
	return b
}

// 根据key和多个字段名和字段值，批量设置hash字段值
func (b *CachePkgBigCache) HMSet(key string, pairs ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) HMsetContext(ctx context.Context, key string, pairs ...interface{}) CachePkgInterface {
	return b
}

// 根据key和多个字段名，批量查询多个hash字段值
func (b *CachePkgBigCache) HMGet(key string, fields ...string) CachePkgInterface { return b }
func (b *CachePkgBigCache) HMGetContext(ctx context.Context, key string, fields ...string) CachePkgInterface {
	return b
}

// 根据key和field字段，累加字段的数值
func (b *CachePkgBigCache) HIncrBy(key string, field string, value int64) CachePkgInterface { return b }
func (b *CachePkgBigCache) HIncrByContext(ctx context.Context, key string, field string, value int64) CachePkgInterface {
	return b
}

// 根据key和field字段，累加字段的数值
func (b *CachePkgBigCache) HIncrByFloat(key string, field string, value float64) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) HIncrByFloatContext(ctx context.Context, key string, field string, value float64) CachePkgInterface {
	return b
}

// 根据key和字段名，删除hash字段，支持批量删除hash字段
func (b *CachePkgBigCache) HDel(key string, fields ...string) CachePkgInterface { return b }
func (b *CachePkgBigCache) HDelContext(ctx context.Context, key string, fields ...string) CachePkgInterface {
	return b
}

// 根据key返回所有字段名
func (b *CachePkgBigCache) HKeys(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) HKeysContext(ctx context.Context, key string) CachePkgInterface { return b }

// 根据key，查询hash的字段数量
func (b *CachePkgBigCache) HLen(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) HlenContext(ctx context.Context, key string) CachePkgInterface { return b }

// 检测hash字段名是否存在。
func (b *CachePkgBigCache) HExists(key string, field string) CachePkgInterface { return b }
func (b *CachePkgBigCache) HExistsContext(ctx context.Context, key string, field string) CachePkgInterface {
	return b
}

// 从列表左边插入数据
func (b *CachePkgBigCache) LPush(key string, value ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) LPushContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	return b
}

// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
func (b *CachePkgBigCache) LPushX(key string, value ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) LPushXContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	return b
}

// 从列表的右边删除第一个数据，并返回删除的数据
func (b *CachePkgBigCache) RPop(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) RPopContext(ctx context.Context, key string) CachePkgInterface { return b }

// 从列表右边插入数据
func (b *CachePkgBigCache) RPush(key string, values ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) RPushContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface {
	return b
}

// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
func (b *CachePkgBigCache) RPushX(key string, values ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) RPushXContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface {
	return b
}

// 从列表左边删除第一个数据，并返回删除的数据
func (b *CachePkgBigCache) LPop(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) LPopContext(ctx context.Context, key string) CachePkgInterface { return b }

// 返回列表的大小
func (b *CachePkgBigCache) LLen(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) LLenContext(ctx context.Context, key string) CachePkgInterface { return b }

// 返回列表的一个范围内的数据，也可以返回全部数据
func (b *CachePkgBigCache) LRange(key string, start int64, stop int64) CachePkgInterface { return b }
func (b *CachePkgBigCache) LRangeContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface {
	return b
}

// 删除列表中的数据
func (b *CachePkgBigCache) LRem(key string, count int64, value interface{}) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) LRemContext(ctx context.Context, key string, count int64, value interface{}) CachePkgInterface {
	return b
}

// 根据索引坐标，查询列表中的数据
func (b *CachePkgBigCache) LIndex(key string, index int64) CachePkgInterface { return b }
func (b *CachePkgBigCache) LIndexContext(ctx context.Context, key string, index int64) CachePkgInterface {
	return b
}

// 在指定位置插入数据
func (b *CachePkgBigCache) LInsert(key string, op string, pivot interface{}, value interface{}) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) CachePkgInterface {
	return b
}

// 添加集合元素
func (b *CachePkgBigCache) SAdd(key string, value ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) SAddContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	return b
}

// 获取集合元素个数
func (b *CachePkgBigCache) SCard(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) SCardContext(ctx context.Context, key string) CachePkgInterface { return b }

// 判断元素是否在集合中
func (b *CachePkgBigCache) SIsMember(key string, value interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) SIsMemberContext(ctx context.Context, key string, value interface{}) CachePkgInterface {
	return b
}

// 获取集合中所有的元素
func (b *CachePkgBigCache) SMembers(key string) CachePkgInterface { return b }
func (b *CachePkgBigCache) SMembersContext(ctx context.Context, key string) CachePkgInterface {
	return b
}

// 删除集合元素
func (b *CachePkgBigCache) SRem(key string, value interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) SRemContext(ctx context.Context, key string, value interface{}) CachePkgInterface {
	return b
}

// 随机返回集合中的元素，并且删除返回的元素
func (b *CachePkgBigCache) SPop(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) SPopContext(ctx context.Context, key string) CachePkgInterface { return b }

// 随机返回集合中的元素，并且删除返回的元素
func (b *CachePkgBigCache) SPopN(key string, num int64) CachePkgInterface { return b }
func (b *CachePkgBigCache) SPopNContext(ctx context.Context, key string, num int64) CachePkgInterface {
	return b
}

// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (b *CachePkgBigCache) ZAdd(key string, redisZ *redis.Z) CachePkgInterface { return b }
func (b *CachePkgBigCache) ZAddContext(ctx context.Context, key string, redisZ *redis.Z) CachePkgInterface {
	return b
}

// 返回集合元素个数
func (b *CachePkgBigCache) ZCard(key string) CachePkgInterface                             { return b }
func (b *CachePkgBigCache) ZCardContext(ctx context.Context, key string) CachePkgInterface { return b }

// 统计某个分数范围内的元素个数
func (b *CachePkgBigCache) ZCount(key string, min string, max string) CachePkgInterface { return b }
func (b *CachePkgBigCache) ZCountContext(ctx context.Context, key string, min string, max string) CachePkgInterface {
	return b
}

// 增加元素的分数
func (b *CachePkgBigCache) ZIncrBy(key string, incr float64, member string) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) ZIncrByContext(ctx context.Context, key string, incr float64, member string) CachePkgInterface {
	return b
}

// 返回集合中某个索引范围的元素，根据分数从小到大排序
func (b *CachePkgBigCache) ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return b
}

// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
func (b *CachePkgBigCache) ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return b
}

// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
func (b *CachePkgBigCache) ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	return b
}

// 删除集合元素
func (b *CachePkgBigCache) ZRem(key string, members ...interface{}) CachePkgInterface { return b }
func (b *CachePkgBigCache) ZRemContext(ctx context.Context, key string, members ...interface{}) CachePkgInterface {
	return b
}

// 根据索引范围删除元素
func (b *CachePkgBigCache) ZRemRangeByRank(key string, start int64, stop int64) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface {
	return b
}

// 根据分数范围删除元素
func (b *CachePkgBigCache) ZRemRangeByScore(key string, min string, max string) CachePkgInterface {
	return b
}
func (b *CachePkgBigCache) ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) CachePkgInterface {
	return b
}

// 查询元素对应的分数
func (b *CachePkgBigCache) ZScore(key string, member string) CachePkgInterface { return b }
func (b *CachePkgBigCache) ZScoreContext(ctx context.Context, key string, member string) CachePkgInterface {
	return b
}

// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (b *CachePkgBigCache) ZRank(key string, member string) CachePkgInterface { return b }
func (b *CachePkgBigCache) ZRankContext(ctx context.Context, key string, member string) CachePkgInterface {
	return b
}
