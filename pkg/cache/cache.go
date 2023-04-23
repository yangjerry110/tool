/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-13 16:48:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:24:10
 * @Description: cache base
 */
package cache

import (
	"context"
	"time"

	"github.com/coocood/freecache"
	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/cache"
)

type CachePkgInterface interface {
	// 创建链接
	Client(filePath string, fileName string) CachePkgInterface
	CreateDefaultClient() CachePkgInterface
	// 检查配置
	CheckConfig() error
	// 获取结果和，错误
	Result() (interface{}, error)
	// 获取结果
	GetResult() interface{}
	// 获取错误
	GetErr() error
	// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
	Set(key string, value interface{}, expiration time.Duration) CachePkgInterface
	SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface
	// 查询数据库中名称为key的value值
	Get(key string) CachePkgInterface
	GetContext(ctx context.Context, key string) CachePkgInterface
	//如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
	SetNX(key string, value interface{}, expiration time.Duration) CachePkgInterface
	SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface
	// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
	MGet(keys ...string) CachePkgInterface
	MGetContext(ctx context.Context, keys ...string) CachePkgInterface
	//批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
	MSet(pairs ...interface{}) CachePkgInterface
	MsetContext(ctx context.Context, pairs ...interface{}) CachePkgInterface
	// Incr函数每次加一,key对应的值必须是整数或nil
	// 否则会报错incr key1: ERR value is not an integer or out of range
	Incr(key string) CachePkgInterface
	IncrContext(ctx context.Context, key string) CachePkgInterface
	// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
	IncrBy(key string, value int64) CachePkgInterface
	IncrByContext(ctx context.Context, key string, value int64) CachePkgInterface
	// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
	IncrByFloat(key string, value float64) CachePkgInterface
	IncrByFloatContext(ctx context.Context, key string, value float64) CachePkgInterface
	// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
	Decr(key string) CachePkgInterface
	DecrContext(ctx context.Context, key string) CachePkgInterface
	// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
	DecrBy(key string, decrement int64) CachePkgInterface
	DecrByContext(ctx context.Context, key string, decrement int64) CachePkgInterface
	//删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
	Del(key string) CachePkgInterface
	DelContext(ctx context.Context, key string) CachePkgInterface
	// 设置key的过期时间,单位秒
	Expire(key string, expiration time.Duration) CachePkgInterface
	ExpireContext(ctx context.Context, key string, expiration time.Duration) CachePkgInterface
	// 给数据库中名称为key的string值追加value
	Append(key, value string) CachePkgInterface
	AppendContext(ctx context.Context, key string, value string) CachePkgInterface
	// 根据key和field字段设置，field字段的值
	HSet(key string, value interface{}, expiration time.Duration) CachePkgInterface
	HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface
	// 根据key和field字段，查询field字段的值
	HGet(key string, field string) CachePkgInterface
	HGetContext(ctx context.Context, key string, field string) CachePkgInterface
	// 如果field字段不存在，则设置hash字段值
	HSetNX(key string, field string, value interface{}, expiration time.Duration) CachePkgInterface
	HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) CachePkgInterface
	// 根据key和多个字段名和字段值，批量设置hash字段值
	HMSet(key string, pairs ...interface{}) CachePkgInterface
	HMsetContext(ctx context.Context, key string, pairs ...interface{}) CachePkgInterface
	// 根据key和多个字段名，批量查询多个hash字段值
	HMGet(key string, fields ...string) CachePkgInterface
	HMGetContext(ctx context.Context, key string, fields ...string) CachePkgInterface
	// 根据key和field字段，累加字段的数值
	HIncrBy(key string, field string, value int64) CachePkgInterface
	HIncrByContext(ctx context.Context, key string, field string, value int64) CachePkgInterface
	// 根据key和field字段，累加字段的数值
	HIncrByFloat(key string, field string, value float64) CachePkgInterface
	HIncrByFloatContext(ctx context.Context, key string, field string, value float64) CachePkgInterface
	// 根据key和字段名，删除hash字段，支持批量删除hash字段
	HDel(key string, fields ...string) CachePkgInterface
	HDelContext(ctx context.Context, key string, fields ...string) CachePkgInterface
	// 根据key返回所有字段名
	HKeys(key string) CachePkgInterface
	HKeysContext(ctx context.Context, key string) CachePkgInterface
	// 根据key，查询hash的字段数量
	HLen(key string) CachePkgInterface
	HlenContext(ctx context.Context, key string) CachePkgInterface
	// 检测hash字段名是否存在。
	HExists(key string, field string) CachePkgInterface
	HExistsContext(ctx context.Context, key string, field string) CachePkgInterface
	// 从列表左边插入数据
	LPush(key string, value ...interface{}) CachePkgInterface
	LPushContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface
	// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
	LPushX(key string, value ...interface{}) CachePkgInterface
	LPushXContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface
	// 从列表的右边删除第一个数据，并返回删除的数据
	RPop(key string) CachePkgInterface
	RPopContext(ctx context.Context, key string) CachePkgInterface
	// 从列表右边插入数据
	RPush(key string, values ...interface{}) CachePkgInterface
	RPushContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface
	// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
	RPushX(key string, values ...interface{}) CachePkgInterface
	RPushXContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface
	// 从列表左边删除第一个数据，并返回删除的数据
	LPop(key string) CachePkgInterface
	LPopContext(ctx context.Context, key string) CachePkgInterface
	// 返回列表的大小
	LLen(key string) CachePkgInterface
	LLenContext(ctx context.Context, key string) CachePkgInterface
	// 返回列表的一个范围内的数据，也可以返回全部数据
	LRange(key string, start int64, stop int64) CachePkgInterface
	LRangeContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface
	// 删除列表中的数据
	LRem(key string, count int64, value interface{}) CachePkgInterface
	LRemContext(ctx context.Context, key string, count int64, value interface{}) CachePkgInterface
	// 根据索引坐标，查询列表中的数据
	LIndex(key string, index int64) CachePkgInterface
	LIndexContext(ctx context.Context, key string, index int64) CachePkgInterface
	// 在指定位置插入数据
	LInsert(key string, op string, pivot interface{}, value interface{}) CachePkgInterface
	LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) CachePkgInterface
	// 添加集合元素
	SAdd(key string, value ...interface{}) CachePkgInterface
	SAddContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface
	// 获取集合元素个数
	SCard(key string) CachePkgInterface
	SCardContext(ctx context.Context, key string) CachePkgInterface
	// 判断元素是否在集合中
	SIsMember(key string, value interface{}) CachePkgInterface
	SIsMemberContext(ctx context.Context, key string, value interface{}) CachePkgInterface
	// 获取集合中所有的元素
	SMembers(key string) CachePkgInterface
	SMembersContext(ctx context.Context, key string) CachePkgInterface
	// 删除集合元素
	SRem(key string, value interface{}) CachePkgInterface
	SRemContext(ctx context.Context, key string, value interface{}) CachePkgInterface
	// 随机返回集合中的元素，并且删除返回的元素
	SPop(key string) CachePkgInterface
	SPopContext(ctx context.Context, key string) CachePkgInterface
	// 随机返回集合中的元素，并且删除返回的元素
	SPopN(key string, num int64) CachePkgInterface
	SPopNContext(ctx context.Context, key string, num int64) CachePkgInterface
	// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
	ZAdd(key string, redisZ *redis.Z) CachePkgInterface
	ZAddContext(ctx context.Context, key string, redisZ *redis.Z) CachePkgInterface
	// 返回集合元素个数
	ZCard(key string) CachePkgInterface
	ZCardContext(ctx context.Context, key string) CachePkgInterface
	// 统计某个分数范围内的元素个数
	ZCount(key string, min string, max string) CachePkgInterface
	ZCountContext(ctx context.Context, key string, min string, max string) CachePkgInterface
	// 增加元素的分数
	ZIncrBy(key string, incr float64, member string) CachePkgInterface
	ZIncrByContext(ctx context.Context, key string, incr float64, member string) CachePkgInterface
	// 返回集合中某个索引范围的元素，根据分数从小到大排序
	ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface
	ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface
	// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
	ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface
	ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface
	// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
	ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface
	ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface
	// 删除集合元素
	ZRem(key string, members ...interface{}) CachePkgInterface
	ZRemContext(ctx context.Context, key string, members ...interface{}) CachePkgInterface
	// 根据索引范围删除元素
	ZRemRangeByRank(key string, start int64, stop int64) CachePkgInterface
	ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface
	// 根据分数范围删除元素
	ZRemRangeByScore(key string, min string, max string) CachePkgInterface
	ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) CachePkgInterface
	// 查询元素对应的分数
	ZScore(key string, member string) CachePkgInterface
	ZScoreContext(ctx context.Context, key string, member string) CachePkgInterface
	// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
	ZRank(key string, member string) CachePkgInterface
	ZRankContext(ctx context.Context, key string, member string) CachePkgInterface
}

type CachePkg struct {
	CacheInterface    cache.CacheInterface
	CachePkgInterface CachePkgInterface
}

/**
 * @description: CreateCache
 * @param {cache.CacheInterface} cacheInterface
 * @author: Jerry.Yang
 * @date: 2022-10-26 11:15:34
 * @return {*}
 */
func CreateCache(cacheInterface cache.CacheInterface) *CachePkg {
	return &CachePkg{CacheInterface: cacheInterface}
}

/**
 * @description: CreatePkgCache
 * @param {CachePkgInterface} CachePkgInterface
 * @author: Jerry.Yang
 * @date: 2022-10-26 11:15:42
 * @return {*}
 */
func CreatePkgCache(CachePkgInterface CachePkgInterface) *CachePkg {
	return &CachePkg{CachePkgInterface: CachePkgInterface}
}

/**
 * @description: SetCache
 * @param {...string} cache
 * @author: Jerry.Yang
 * @date: 2022-10-26 11:16:35
 * @return {*}
 */
func SetCache(cache ...string) CachePkgInterface {
	if len(cache) == 0 {
		cachePkg := &CachePkgRedis{}
		return CreatePkgCache(cachePkg).CachePkgInterface
	}

	// 判断cache
	switch cache[0] {
	case "redis":
		cachePkg := &CachePkgRedis{}
		return CreatePkgCache(cachePkg).CachePkgInterface
	case "bigCache":
		bigCachePkg := &CachePkgBigCache{}
		return CreatePkgCache(bigCachePkg).CachePkgInterface
	}

	cachePkg := &CachePkgRedis{}
	return CreatePkgCache(cachePkg).CachePkgInterface
}

/**
 * @description: Client
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:24:09
 * @return {*}
 */
func Client(filePath string, fileName string) CachePkgInterface {
	cachePkg := &CachePkgRedis{}
	return CreatePkgCache(cachePkg).CachePkgInterface.Client(filePath, fileName)
}

/**
 * @description: CreateDefaultRedisClient
 * @param {string} addr
 * @author: Jerry.Yang
 * @date: 2022-10-26 15:11:28
 * @return {*}
 */
func CreateDefaultRedisClient(addr string) CachePkgInterface {
	cachePkg := &CachePkgRedis{Addr: addr}
	return CreatePkgCache(cachePkg).CachePkgInterface.CreateDefaultClient()
}

/**
 * @description: ClientBigCache
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:24:35
 * @return {*}
 */
func ClientBigCache(filePath string, fileName string) CachePkgInterface {
	cachePkg := &CachePkgBigCache{}
	return CreatePkgCache(cachePkg).CachePkgInterface.Client(filePath, fileName)
}

/**
 * @description: ClientDefaultBigCache
 * @param {time.Duration} eviction
 * @param {time.Duration} cleanTime
 * @author: Jerry.Yang
 * @date: 2022-10-26 15:11:47
 * @return {*}
 */
func ClientDefaultBigCache(eviction time.Duration, cleanTime time.Duration) CachePkgInterface {
	cachePkg := &CachePkgBigCache{Eviction: eviction, CleanTime: cleanTime}
	return CreatePkgCache(cachePkg).CachePkgInterface.CreateDefaultClient()
}

/**
 * @description: ClientFreeCache
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:24:59
 * @return {*}
 */
func ClientFreeCache(filePath string, fileName string) CachePkgInterface {
	cachePkg := &CachePkgFreeCache{}
	return CreatePkgCache(cachePkg).CachePkgInterface.Client(filePath, fileName)
}

/**
 * @description: ClientDefaultFreeCache
 * @param {int} size
 * @param {freecache.Timer} timer
 * @author: Jerry.Yang
 * @date: 2022-10-26 15:50:45
 * @return {*}
 */
func ClientDefaultFreeCache(size int, timer freecache.Timer) CachePkgInterface {
	cachePkg := &CachePkgFreeCache{Size: size, Timer: timer}
	return CreatePkgCache(cachePkg).CachePkgInterface.CreateDefaultClient()
}
