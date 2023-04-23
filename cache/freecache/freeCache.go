/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-26 15:14:22
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-11 17:22:25
 * @Description: freecache
 */
package freecache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/cache"
)

// 获取结果和，错误
func (f *FreeCache) Result() (interface{}, error) {
	return f.FreeCacheVal, f.FreeCacheErr
}

// 获取结果
func (f *FreeCache) GetResult() interface{} {
	return f.FreeCacheVal
}

// 获取错误
func (f *FreeCache) GetErr() error {
	return f.FreeCacheErr
}

// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
func (f *FreeCache) Set(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	f.FreeCacheErr = f.FreeCache.Set([]byte(key), []byte(fmt.Sprintf("%v", value)), int(expiration))
	return f
}

func (f *FreeCache) SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return f
}

// 查询数据库中名称为key的value值
func (f *FreeCache) Get(key string) cache.CacheInterface {
	f.FreeCacheVal, f.FreeCacheErr = f.FreeCache.Get([]byte(key))
	return f
}

func (f *FreeCache) GetContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
func (f *FreeCache) SetNX(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return f
}
func (f *FreeCache) SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return f
}

// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
func (f *FreeCache) MGet(keys ...string) cache.CacheInterface                             { return f }
func (f *FreeCache) MGetContext(ctx context.Context, keys ...string) cache.CacheInterface { return f }

// 批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
func (f *FreeCache) MSet(pairs ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) MsetContext(ctx context.Context, pairs ...interface{}) cache.CacheInterface {
	return f
}

// Incr函数每次加一,key对应的值必须是整数或nil
// 否则会报错incr key1: ERR value is not an integer or out of range
func (f *FreeCache) Incr(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) IncrContext(ctx context.Context, key string) cache.CacheInterface { return f }

// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
func (f *FreeCache) IncrBy(key string, value int64) cache.CacheInterface { return f }
func (f *FreeCache) IncrByContext(ctx context.Context, key string, value int64) cache.CacheInterface {
	return f
}

// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
func (f *FreeCache) IncrByFloat(key string, value float64) cache.CacheInterface { return f }
func (f *FreeCache) IncrByFloatContext(ctx context.Context, key string, value float64) cache.CacheInterface {
	return f
}

// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
func (f *FreeCache) Decr(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) DecrContext(ctx context.Context, key string) cache.CacheInterface { return f }

// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
func (f *FreeCache) DecrBy(key string, decrement int64) cache.CacheInterface { return f }
func (f *FreeCache) DecrByContext(ctx context.Context, key string, decrement int64) cache.CacheInterface {
	return f
}

// 删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
func (f *FreeCache) Del(key string) cache.CacheInterface {
	result := f.FreeCache.Del([]byte(key))
	if !result {
		f.FreeCacheErr = errors.New(fmt.Sprintf("freecache Err : del %s fail", key))
	}
	return f
}

func (f *FreeCache) DelContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 设置key的过期时间,单位秒
func (f *FreeCache) Expire(key string, expiration time.Duration) cache.CacheInterface { return f }
func (f *FreeCache) ExpireContext(ctx context.Context, key string, expiration time.Duration) cache.CacheInterface {
	return f
}

// 给数据库中名称为key的string值追加value
func (f *FreeCache) Append(key, value string) cache.CacheInterface { return f }
func (f *FreeCache) AppendContext(ctx context.Context, key string, value string) cache.CacheInterface {
	return f
}

// 根据key和field字段设置，field字段的值
func (f *FreeCache) HSet(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return f
}
func (f *FreeCache) HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return f
}

// 根据key和field字段，查询field字段的值
func (f *FreeCache) HGet(key string, field string) cache.CacheInterface { return f }
func (f *FreeCache) HGetContext(ctx context.Context, key string, field string) cache.CacheInterface {
	return f
}

// 如果field字段不存在，则设置hash字段值
func (f *FreeCache) HSetNX(key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return f
}
func (f *FreeCache) HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return f
}

// 根据key和多个字段名和字段值，批量设置hash字段值
func (f *FreeCache) HMSet(key string, pairs ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) HMsetContext(ctx context.Context, key string, pairs ...interface{}) cache.CacheInterface {
	return f
}

// 根据key和多个字段名，批量查询多个hash字段值
func (f *FreeCache) HMGet(key string, fields ...string) cache.CacheInterface { return f }
func (f *FreeCache) HMGetContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	return f
}

// 根据key和field字段，累加字段的数值
func (f *FreeCache) HIncrBy(key string, field string, value int64) cache.CacheInterface { return f }
func (f *FreeCache) HIncrByContext(ctx context.Context, key string, field string, value int64) cache.CacheInterface {
	return f
}

// 根据key和field字段，累加字段的数值
func (f *FreeCache) HIncrByFloat(key string, field string, value float64) cache.CacheInterface {
	return f
}
func (f *FreeCache) HIncrByFloatContext(ctx context.Context, key string, field string, value float64) cache.CacheInterface {
	return f
}

// 根据key和字段名，删除hash字段，支持批量删除hash字段
func (f *FreeCache) HDel(key string, fields ...string) cache.CacheInterface { return f }
func (f *FreeCache) HDelContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	return f
}

// 根据key返回所有字段名
func (f *FreeCache) HKeys(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) HKeysContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 根据key，查询hash的字段数量
func (f *FreeCache) HLen(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) HlenContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 检测hash字段名是否存在。
func (f *FreeCache) HExists(key string, field string) cache.CacheInterface { return f }
func (f *FreeCache) HExistsContext(ctx context.Context, key string, field string) cache.CacheInterface {
	return f
}

// 从列表左边插入数据
func (f *FreeCache) LPush(key string, value ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) LPushContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return f
}

// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
func (f *FreeCache) LPushX(key string, value ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) LPushXContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return f
}

// 从列表的右边删除第一个数据，并返回删除的数据
func (f *FreeCache) RPop(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) RPopContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 从列表右边插入数据
func (f *FreeCache) RPush(key string, values ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) RPushContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	return f
}

// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
func (f *FreeCache) RPushX(key string, values ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) RPushXContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	return f
}

// 从列表左边删除第一个数据，并返回删除的数据
func (f *FreeCache) LPop(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) LPopContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 返回列表的大小
func (f *FreeCache) LLen(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) LLenContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 返回列表的一个范围内的数据，也可以返回全部数据
func (f *FreeCache) LRange(key string, start int64, stop int64) cache.CacheInterface { return f }
func (f *FreeCache) LRangeContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	return f
}

// 删除列表中的数据
func (f *FreeCache) LRem(key string, count int64, value interface{}) cache.CacheInterface { return f }
func (f *FreeCache) LRemContext(ctx context.Context, key string, count int64, value interface{}) cache.CacheInterface {
	return f
}

// 根据索引坐标，查询列表中的数据
func (f *FreeCache) LIndex(key string, index int64) cache.CacheInterface { return f }
func (f *FreeCache) LIndexContext(ctx context.Context, key string, index int64) cache.CacheInterface {
	return f
}

// 在指定位置插入数据
func (f *FreeCache) LInsert(key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	return f
}
func (f *FreeCache) LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	return f
}

// 添加集合元素
func (f *FreeCache) SAdd(key string, value ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) SAddContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return f
}

// 获取集合元素个数
func (f *FreeCache) SCard(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) SCardContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 判断元素是否在集合中
func (f *FreeCache) SIsMember(key string, value interface{}) cache.CacheInterface { return f }
func (f *FreeCache) SIsMemberContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	return f
}

// 获取集合中所有的元素
func (f *FreeCache) SMembers(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) SMembersContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 删除集合元素
func (f *FreeCache) SRem(key string, value interface{}) cache.CacheInterface { return f }
func (f *FreeCache) SRemContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	return f
}

// 随机返回集合中的元素，并且删除返回的元素
func (f *FreeCache) SPop(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) SPopContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 随机返回集合中的元素，并且删除返回的元素
func (f *FreeCache) SPopN(key string, num int64) cache.CacheInterface { return f }
func (f *FreeCache) SPopNContext(ctx context.Context, key string, num int64) cache.CacheInterface {
	return f
}

// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (f *FreeCache) ZAdd(key string, redisZ *redis.Z) cache.CacheInterface { return f }
func (f *FreeCache) ZAddContext(ctx context.Context, key string, redisZ *redis.Z) cache.CacheInterface {
	return f
}

// 返回集合元素个数
func (f *FreeCache) ZCard(key string) cache.CacheInterface                             { return f }
func (f *FreeCache) ZCardContext(ctx context.Context, key string) cache.CacheInterface { return f }

// 统计某个分数范围内的元素个数
func (f *FreeCache) ZCount(key string, min string, max string) cache.CacheInterface { return f }
func (f *FreeCache) ZCountContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	return f
}

// 增加元素的分数
func (f *FreeCache) ZIncrBy(key string, incr float64, member string) cache.CacheInterface { return f }
func (f *FreeCache) ZIncrByContext(ctx context.Context, key string, incr float64, member string) cache.CacheInterface {
	return f
}

// 返回集合中某个索引范围的元素，根据分数从小到大排序
func (f *FreeCache) ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return f
}
func (f *FreeCache) ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return f
}

// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
func (f *FreeCache) ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return f
}
func (f *FreeCache) ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return f
}

// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
func (f *FreeCache) ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return f
}
func (f *FreeCache) ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return f
}

// 删除集合元素
func (f *FreeCache) ZRem(key string, members ...interface{}) cache.CacheInterface { return f }
func (f *FreeCache) ZRemContext(ctx context.Context, key string, members ...interface{}) cache.CacheInterface {
	return f
}

// 根据索引范围删除元素
func (f *FreeCache) ZRemRangeByRank(key string, start int64, stop int64) cache.CacheInterface {
	return f
}
func (f *FreeCache) ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	return f
}

// 根据分数范围删除元素
func (f *FreeCache) ZRemRangeByScore(key string, min string, max string) cache.CacheInterface {
	return f
}
func (f *FreeCache) ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	return f
}

// 查询元素对应的分数
func (f *FreeCache) ZScore(key string, member string) cache.CacheInterface { return f }
func (f *FreeCache) ZScoreContext(ctx context.Context, key string, member string) cache.CacheInterface {
	return f
}

// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (f *FreeCache) ZRank(key string, member string) cache.CacheInterface { return f }
func (f *FreeCache) ZRankContext(ctx context.Context, key string, member string) cache.CacheInterface {
	return f
}
