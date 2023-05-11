/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-09 14:22:30
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-09 15:45:25
 * @Description: gocache
 */
package gocache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/cache"
)

// 获取结果和，错误
func (g *GoCache) Result() (interface{}, error) {
	return g.GoCacheVal, g.GoCacheErr
}

// 获取结果
func (g *GoCache) GetResult() interface{} {
	return g.GoCacheVal
}

// 获取错误
func (g *GoCache) GetErr() error {
	return g.GoCacheErr
}

// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
func (g *GoCache) Set(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	g.GoCache.Set(key, value, g.DefaultExpireTime)
	return g
}

func (g *GoCache) SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return g
}

// 查询数据库中名称为key的value值
func (g *GoCache) Get(key string) cache.CacheInterface {
	val, ok := g.GoCache.Get(key)

	fmt.Printf("\r\n key = %s; ok  = %+v \r\n", key, ok)
	fmt.Printf("\r\n key = %s; val = %+v \r\n", key, val)

	if !ok {
		g.GoCacheErr = fmt.Errorf("gocache err : key %s is not found", key)
	}
	g.GoCacheVal = val
	return g
}

func (g *GoCache) GetContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
func (g *GoCache) SetNX(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return g
}
func (g *GoCache) SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return g
}

// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
func (g *GoCache) MGet(keys ...string) cache.CacheInterface                             { return g }
func (g *GoCache) MGetContext(ctx context.Context, keys ...string) cache.CacheInterface { return g }

// 批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
func (g *GoCache) MSet(pairs ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) MsetContext(ctx context.Context, pairs ...interface{}) cache.CacheInterface {
	return g
}

// Incr函数每次加一,key对应的值必须是整数或nil
// 否则会报错incr key1: ERR value is not an integer or out of range
func (g *GoCache) Incr(key string) cache.CacheInterface                             { return g }
func (g *GoCache) IncrContext(ctx context.Context, key string) cache.CacheInterface { return g }

// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
func (g *GoCache) IncrBy(key string, value int64) cache.CacheInterface { return g }
func (g *GoCache) IncrByContext(ctx context.Context, key string, value int64) cache.CacheInterface {
	return g
}

// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
func (g *GoCache) IncrByFloat(key string, value float64) cache.CacheInterface { return g }
func (g *GoCache) IncrByFloatContext(ctx context.Context, key string, value float64) cache.CacheInterface {
	return g
}

// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
func (g *GoCache) Decr(key string) cache.CacheInterface                             { return g }
func (g *GoCache) DecrContext(ctx context.Context, key string) cache.CacheInterface { return g }

// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
func (g *GoCache) DecrBy(key string, decrement int64) cache.CacheInterface { return g }
func (g *GoCache) DecrByContext(ctx context.Context, key string, decrement int64) cache.CacheInterface {
	return g
}

// 删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
func (g *GoCache) Del(key string) cache.CacheInterface {
	g.GoCache.Delete(key)
	return g
}

func (g *GoCache) DelContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 设置key的过期时间,单位秒
func (g *GoCache) Expire(key string, expiration time.Duration) cache.CacheInterface { return g }
func (g *GoCache) ExpireContext(ctx context.Context, key string, expiration time.Duration) cache.CacheInterface {
	return g
}

// 给数据库中名称为key的string值追加value
func (g *GoCache) Append(key, value string) cache.CacheInterface { return g }
func (g *GoCache) AppendContext(ctx context.Context, key string, value string) cache.CacheInterface {
	return g
}

// 根据key和field字段设置，field字段的值
func (g *GoCache) HSet(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return g
}
func (g *GoCache) HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return g
}

// 根据key和field字段，查询field字段的值
func (g *GoCache) HGet(key string, field string) cache.CacheInterface { return g }
func (g *GoCache) HGetContext(ctx context.Context, key string, field string) cache.CacheInterface {
	return g
}

// 如果field字段不存在，则设置hash字段值
func (g *GoCache) HSetNX(key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return g
}
func (g *GoCache) HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return g
}

// 根据key和多个字段名和字段值，批量设置hash字段值
func (g *GoCache) HMSet(key string, pairs ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) HMsetContext(ctx context.Context, key string, pairs ...interface{}) cache.CacheInterface {
	return g
}

// 根据key和多个字段名，批量查询多个hash字段值
func (g *GoCache) HMGet(key string, fields ...string) cache.CacheInterface { return g }
func (g *GoCache) HMGetContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	return g
}

// 根据key和field字段，累加字段的数值
func (g *GoCache) HIncrBy(key string, field string, value int64) cache.CacheInterface { return g }
func (g *GoCache) HIncrByContext(ctx context.Context, key string, field string, value int64) cache.CacheInterface {
	return g
}

// 根据key和field字段，累加字段的数值
func (g *GoCache) HIncrByFloat(key string, field string, value float64) cache.CacheInterface {
	return g
}
func (g *GoCache) HIncrByFloatContext(ctx context.Context, key string, field string, value float64) cache.CacheInterface {
	return g
}

// 根据key和字段名，删除hash字段，支持批量删除hash字段
func (g *GoCache) HDel(key string, fields ...string) cache.CacheInterface { return g }
func (g *GoCache) HDelContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	return g
}

// 根据key返回所有字段名
func (g *GoCache) HKeys(key string) cache.CacheInterface                             { return g }
func (g *GoCache) HKeysContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 根据key，查询hash的字段数量
func (g *GoCache) HLen(key string) cache.CacheInterface                             { return g }
func (g *GoCache) HlenContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 检测hash字段名是否存在。
func (g *GoCache) HExists(key string, field string) cache.CacheInterface { return g }
func (g *GoCache) HExistsContext(ctx context.Context, key string, field string) cache.CacheInterface {
	return g
}

// 从列表左边插入数据
func (g *GoCache) LPush(key string, value ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) LPushContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return g
}

// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
func (g *GoCache) LPushX(key string, value ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) LPushXContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return g
}

// 从列表的右边删除第一个数据，并返回删除的数据
func (g *GoCache) RPop(key string) cache.CacheInterface                             { return g }
func (g *GoCache) RPopContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 从列表右边插入数据
func (g *GoCache) RPush(key string, values ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) RPushContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	return g
}

// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
func (g *GoCache) RPushX(key string, values ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) RPushXContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	return g
}

// 从列表左边删除第一个数据，并返回删除的数据
func (g *GoCache) LPop(key string) cache.CacheInterface                             { return g }
func (g *GoCache) LPopContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 返回列表的大小
func (g *GoCache) LLen(key string) cache.CacheInterface                             { return g }
func (g *GoCache) LLenContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 返回列表的一个范围内的数据，也可以返回全部数据
func (g *GoCache) LRange(key string, start int64, stop int64) cache.CacheInterface { return g }
func (g *GoCache) LRangeContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	return g
}

// 删除列表中的数据
func (g *GoCache) LRem(key string, count int64, value interface{}) cache.CacheInterface { return g }
func (g *GoCache) LRemContext(ctx context.Context, key string, count int64, value interface{}) cache.CacheInterface {
	return g
}

// 根据索引坐标，查询列表中的数据
func (g *GoCache) LIndex(key string, index int64) cache.CacheInterface { return g }
func (g *GoCache) LIndexContext(ctx context.Context, key string, index int64) cache.CacheInterface {
	return g
}

// 在指定位置插入数据
func (g *GoCache) LInsert(key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	return g
}
func (g *GoCache) LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	return g
}

// 添加集合元素
func (g *GoCache) SAdd(key string, value ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) SAddContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return g
}

// 获取集合元素个数
func (g *GoCache) SCard(key string) cache.CacheInterface                             { return g }
func (g *GoCache) SCardContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 判断元素是否在集合中
func (g *GoCache) SIsMember(key string, value interface{}) cache.CacheInterface { return g }
func (g *GoCache) SIsMemberContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	return g
}

// 获取集合中所有的元素
func (g *GoCache) SMembers(key string) cache.CacheInterface                             { return g }
func (g *GoCache) SMembersContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 删除集合元素
func (g *GoCache) SRem(key string, value interface{}) cache.CacheInterface { return g }
func (g *GoCache) SRemContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	return g
}

// 随机返回集合中的元素，并且删除返回的元素
func (g *GoCache) SPop(key string) cache.CacheInterface                             { return g }
func (g *GoCache) SPopContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 随机返回集合中的元素，并且删除返回的元素
func (g *GoCache) SPopN(key string, num int64) cache.CacheInterface { return g }
func (g *GoCache) SPopNContext(ctx context.Context, key string, num int64) cache.CacheInterface {
	return g
}

// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (g *GoCache) ZAdd(key string, redisZ *redis.Z) cache.CacheInterface { return g }
func (g *GoCache) ZAddContext(ctx context.Context, key string, redisZ *redis.Z) cache.CacheInterface {
	return g
}

// 返回集合元素个数
func (g *GoCache) ZCard(key string) cache.CacheInterface                             { return g }
func (g *GoCache) ZCardContext(ctx context.Context, key string) cache.CacheInterface { return g }

// 统计某个分数范围内的元素个数
func (g *GoCache) ZCount(key string, min string, max string) cache.CacheInterface { return g }
func (g *GoCache) ZCountContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	return g
}

// 增加元素的分数
func (g *GoCache) ZIncrBy(key string, incr float64, member string) cache.CacheInterface { return g }
func (g *GoCache) ZIncrByContext(ctx context.Context, key string, incr float64, member string) cache.CacheInterface {
	return g
}

// 返回集合中某个索引范围的元素，根据分数从小到大排序
func (g *GoCache) ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface { return g }
func (g *GoCache) ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return g
}

// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
func (g *GoCache) ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return g
}
func (g *GoCache) ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return g
}

// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
func (g *GoCache) ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return g
}
func (g *GoCache) ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return g
}

// 删除集合元素
func (g *GoCache) ZRem(key string, members ...interface{}) cache.CacheInterface { return g }
func (g *GoCache) ZRemContext(ctx context.Context, key string, members ...interface{}) cache.CacheInterface {
	return g
}

// 根据索引范围删除元素
func (g *GoCache) ZRemRangeByRank(key string, start int64, stop int64) cache.CacheInterface {
	return g
}
func (g *GoCache) ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	return g
}

// 根据分数范围删除元素
func (g *GoCache) ZRemRangeByScore(key string, min string, max string) cache.CacheInterface {
	return g
}
func (g *GoCache) ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	return g
}

// 查询元素对应的分数
func (g *GoCache) ZScore(key string, member string) cache.CacheInterface { return g }
func (g *GoCache) ZScoreContext(ctx context.Context, key string, member string) cache.CacheInterface {
	return g
}

// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (g *GoCache) ZRank(key string, member string) cache.CacheInterface { return g }
func (g *GoCache) ZRankContext(ctx context.Context, key string, member string) cache.CacheInterface {
	return g
}
