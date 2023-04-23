/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-26 11:41:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-26 16:05:12
 * @Description: bigcache
 */
package bigcache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/cache"
)

// 获取结果和，错误
func (b *BigCache) Result() (interface{}, error) {
	return b.BigCacheVal, b.BigCacheErr
}

// 获取结果
func (b *BigCache) GetResult() interface{} {
	return b.BigCacheVal
}

// 获取错误
func (b *BigCache) GetErr() error {
	return b.BigCacheErr
}

// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
func (b *BigCache) Set(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	b.BigCacheErr = b.BigCache.Set(key, []byte(fmt.Sprintf("%v", value)))
	return b
}

func (b *BigCache) SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return b
}

// 查询数据库中名称为key的value值
func (b *BigCache) Get(key string) cache.CacheInterface {
	b.BigCacheVal, b.BigCacheErr = b.BigCache.Get(key)
	return b
}

func (b *BigCache) GetContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
func (b *BigCache) SetNX(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return b
}
func (b *BigCache) SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return b
}

// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
func (b *BigCache) MGet(keys ...string) cache.CacheInterface                             { return b }
func (b *BigCache) MGetContext(ctx context.Context, keys ...string) cache.CacheInterface { return b }

// 批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
func (b *BigCache) MSet(pairs ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) MsetContext(ctx context.Context, pairs ...interface{}) cache.CacheInterface {
	return b
}

// Incr函数每次加一,key对应的值必须是整数或nil
// 否则会报错incr key1: ERR value is not an integer or out of range
func (b *BigCache) Incr(key string) cache.CacheInterface                             { return b }
func (b *BigCache) IncrContext(ctx context.Context, key string) cache.CacheInterface { return b }

// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
func (b *BigCache) IncrBy(key string, value int64) cache.CacheInterface { return b }
func (b *BigCache) IncrByContext(ctx context.Context, key string, value int64) cache.CacheInterface {
	return b
}

// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
func (b *BigCache) IncrByFloat(key string, value float64) cache.CacheInterface { return b }
func (b *BigCache) IncrByFloatContext(ctx context.Context, key string, value float64) cache.CacheInterface {
	return b
}

// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
func (b *BigCache) Decr(key string) cache.CacheInterface                             { return b }
func (b *BigCache) DecrContext(ctx context.Context, key string) cache.CacheInterface { return b }

// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
func (b *BigCache) DecrBy(key string, decrement int64) cache.CacheInterface { return b }
func (b *BigCache) DecrByContext(ctx context.Context, key string, decrement int64) cache.CacheInterface {
	return b
}

// 删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
func (b *BigCache) Del(key string) cache.CacheInterface {
	b.BigCacheErr = b.BigCache.Delete(key)
	return b
}

func (b *BigCache) DelContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 设置key的过期时间,单位秒
func (b *BigCache) Expire(key string, expiration time.Duration) cache.CacheInterface { return b }
func (b *BigCache) ExpireContext(ctx context.Context, key string, expiration time.Duration) cache.CacheInterface {
	return b
}

// 给数据库中名称为key的string值追加value
func (b *BigCache) Append(key, value string) cache.CacheInterface { return b }
func (b *BigCache) AppendContext(ctx context.Context, key string, value string) cache.CacheInterface {
	return b
}

// 根据key和field字段设置，field字段的值
func (b *BigCache) HSet(key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return b
}
func (b *BigCache) HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return b
}

// 根据key和field字段，查询field字段的值
func (b *BigCache) HGet(key string, field string) cache.CacheInterface { return b }
func (b *BigCache) HGetContext(ctx context.Context, key string, field string) cache.CacheInterface {
	return b
}

// 如果field字段不存在，则设置hash字段值
func (b *BigCache) HSetNX(key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return b
}
func (b *BigCache) HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) cache.CacheInterface {
	return b
}

// 根据key和多个字段名和字段值，批量设置hash字段值
func (b *BigCache) HMSet(key string, pairs ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) HMsetContext(ctx context.Context, key string, pairs ...interface{}) cache.CacheInterface {
	return b
}

// 根据key和多个字段名，批量查询多个hash字段值
func (b *BigCache) HMGet(key string, fields ...string) cache.CacheInterface { return b }
func (b *BigCache) HMGetContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	return b
}

// 根据key和field字段，累加字段的数值
func (b *BigCache) HIncrBy(key string, field string, value int64) cache.CacheInterface { return b }
func (b *BigCache) HIncrByContext(ctx context.Context, key string, field string, value int64) cache.CacheInterface {
	return b
}

// 根据key和field字段，累加字段的数值
func (b *BigCache) HIncrByFloat(key string, field string, value float64) cache.CacheInterface {
	return b
}
func (b *BigCache) HIncrByFloatContext(ctx context.Context, key string, field string, value float64) cache.CacheInterface {
	return b
}

// 根据key和字段名，删除hash字段，支持批量删除hash字段
func (b *BigCache) HDel(key string, fields ...string) cache.CacheInterface { return b }
func (b *BigCache) HDelContext(ctx context.Context, key string, fields ...string) cache.CacheInterface {
	return b
}

// 根据key返回所有字段名
func (b *BigCache) HKeys(key string) cache.CacheInterface                             { return b }
func (b *BigCache) HKeysContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 根据key，查询hash的字段数量
func (b *BigCache) HLen(key string) cache.CacheInterface                             { return b }
func (b *BigCache) HlenContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 检测hash字段名是否存在。
func (b *BigCache) HExists(key string, field string) cache.CacheInterface { return b }
func (b *BigCache) HExistsContext(ctx context.Context, key string, field string) cache.CacheInterface {
	return b
}

// 从列表左边插入数据
func (b *BigCache) LPush(key string, value ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) LPushContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return b
}

// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
func (b *BigCache) LPushX(key string, value ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) LPushXContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return b
}

// 从列表的右边删除第一个数据，并返回删除的数据
func (b *BigCache) RPop(key string) cache.CacheInterface                             { return b }
func (b *BigCache) RPopContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 从列表右边插入数据
func (b *BigCache) RPush(key string, values ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) RPushContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	return b
}

// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
func (b *BigCache) RPushX(key string, values ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) RPushXContext(ctx context.Context, key string, values ...interface{}) cache.CacheInterface {
	return b
}

// 从列表左边删除第一个数据，并返回删除的数据
func (b *BigCache) LPop(key string) cache.CacheInterface                             { return b }
func (b *BigCache) LPopContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 返回列表的大小
func (b *BigCache) LLen(key string) cache.CacheInterface                             { return b }
func (b *BigCache) LLenContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 返回列表的一个范围内的数据，也可以返回全部数据
func (b *BigCache) LRange(key string, start int64, stop int64) cache.CacheInterface { return b }
func (b *BigCache) LRangeContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	return b
}

// 删除列表中的数据
func (b *BigCache) LRem(key string, count int64, value interface{}) cache.CacheInterface { return b }
func (b *BigCache) LRemContext(ctx context.Context, key string, count int64, value interface{}) cache.CacheInterface {
	return b
}

// 根据索引坐标，查询列表中的数据
func (b *BigCache) LIndex(key string, index int64) cache.CacheInterface { return b }
func (b *BigCache) LIndexContext(ctx context.Context, key string, index int64) cache.CacheInterface {
	return b
}

// 在指定位置插入数据
func (b *BigCache) LInsert(key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	return b
}
func (b *BigCache) LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) cache.CacheInterface {
	return b
}

// 添加集合元素
func (b *BigCache) SAdd(key string, value ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) SAddContext(ctx context.Context, key string, value ...interface{}) cache.CacheInterface {
	return b
}

// 获取集合元素个数
func (b *BigCache) SCard(key string) cache.CacheInterface                             { return b }
func (b *BigCache) SCardContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 判断元素是否在集合中
func (b *BigCache) SIsMember(key string, value interface{}) cache.CacheInterface { return b }
func (b *BigCache) SIsMemberContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	return b
}

// 获取集合中所有的元素
func (b *BigCache) SMembers(key string) cache.CacheInterface                             { return b }
func (b *BigCache) SMembersContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 删除集合元素
func (b *BigCache) SRem(key string, value interface{}) cache.CacheInterface { return b }
func (b *BigCache) SRemContext(ctx context.Context, key string, value interface{}) cache.CacheInterface {
	return b
}

// 随机返回集合中的元素，并且删除返回的元素
func (b *BigCache) SPop(key string) cache.CacheInterface                             { return b }
func (b *BigCache) SPopContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 随机返回集合中的元素，并且删除返回的元素
func (b *BigCache) SPopN(key string, num int64) cache.CacheInterface { return b }
func (b *BigCache) SPopNContext(ctx context.Context, key string, num int64) cache.CacheInterface {
	return b
}

// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (b *BigCache) ZAdd(key string, redisZ *redis.Z) cache.CacheInterface { return b }
func (b *BigCache) ZAddContext(ctx context.Context, key string, redisZ *redis.Z) cache.CacheInterface {
	return b
}

// 返回集合元素个数
func (b *BigCache) ZCard(key string) cache.CacheInterface                             { return b }
func (b *BigCache) ZCardContext(ctx context.Context, key string) cache.CacheInterface { return b }

// 统计某个分数范围内的元素个数
func (b *BigCache) ZCount(key string, min string, max string) cache.CacheInterface { return b }
func (b *BigCache) ZCountContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	return b
}

// 增加元素的分数
func (b *BigCache) ZIncrBy(key string, incr float64, member string) cache.CacheInterface { return b }
func (b *BigCache) ZIncrByContext(ctx context.Context, key string, incr float64, member string) cache.CacheInterface {
	return b
}

// 返回集合中某个索引范围的元素，根据分数从小到大排序
func (b *BigCache) ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface { return b }
func (b *BigCache) ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return b
}

// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
func (b *BigCache) ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return b
}
func (b *BigCache) ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return b
}

// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
func (b *BigCache) ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return b
}
func (b *BigCache) ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) cache.CacheInterface {
	return b
}

// 删除集合元素
func (b *BigCache) ZRem(key string, members ...interface{}) cache.CacheInterface { return b }
func (b *BigCache) ZRemContext(ctx context.Context, key string, members ...interface{}) cache.CacheInterface {
	return b
}

// 根据索引范围删除元素
func (b *BigCache) ZRemRangeByRank(key string, start int64, stop int64) cache.CacheInterface {
	return b
}
func (b *BigCache) ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) cache.CacheInterface {
	return b
}

// 根据分数范围删除元素
func (b *BigCache) ZRemRangeByScore(key string, min string, max string) cache.CacheInterface {
	return b
}
func (b *BigCache) ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) cache.CacheInterface {
	return b
}

// 查询元素对应的分数
func (b *BigCache) ZScore(key string, member string) cache.CacheInterface { return b }
func (b *BigCache) ZScoreContext(ctx context.Context, key string, member string) cache.CacheInterface {
	return b
}

// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (b *BigCache) ZRank(key string, member string) cache.CacheInterface { return b }
func (b *BigCache) ZRankContext(ctx context.Context, key string, member string) cache.CacheInterface {
	return b
}
