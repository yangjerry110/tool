/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-11 18:02:06
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:20:26
 * @Description: cache
 */
package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheInterface interface {
	// 创建链接
	Client(filePath string, fileName string) CacheInterface
	CreateDefaultClient() CacheInterface
	// 检查配置
	CheckConfig() error
	// 获取结果和，错误
	Result() (interface{}, error)
	// 获取结果
	GetResult() interface{}
	// 获取错误
	GetErr() error
	// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
	Set(key string, value interface{}, expiration time.Duration) CacheInterface
	SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CacheInterface
	// 查询数据库中名称为key的value值
	Get(key string) CacheInterface
	GetContext(ctx context.Context, key string) CacheInterface
	//如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
	SetNX(key string, value interface{}, expiration time.Duration) CacheInterface
	SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CacheInterface
	// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
	MGet(keys ...string) CacheInterface
	MGetContext(ctx context.Context, keys ...string) CacheInterface
	//批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
	MSet(pairs ...interface{}) CacheInterface
	MsetContext(ctx context.Context, pairs ...interface{}) CacheInterface
	// Incr函数每次加一,key对应的值必须是整数或nil
	// 否则会报错incr key1: ERR value is not an integer or out of range
	Incr(key string) CacheInterface
	IncrContext(ctx context.Context, key string) CacheInterface
	// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
	IncrBy(key string, value int64) CacheInterface
	IncrByContext(ctx context.Context, key string, value int64) CacheInterface
	// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
	IncrByFloat(key string, value float64) CacheInterface
	IncrByFloatContext(ctx context.Context, key string, value float64) CacheInterface
	// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
	Decr(key string) CacheInterface
	DecrContext(ctx context.Context, key string) CacheInterface
	// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
	DecrBy(key string, decrement int64) CacheInterface
	DecrByContext(ctx context.Context, key string, decrement int64) CacheInterface
	//删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
	Del(key string) CacheInterface
	DelContext(ctx context.Context, key string) CacheInterface
	// 设置key的过期时间,单位秒
	Expire(key string, expiration time.Duration) CacheInterface
	ExpireContext(ctx context.Context, key string, expiration time.Duration) CacheInterface
	// 给数据库中名称为key的string值追加value
	Append(key, value string) CacheInterface
	AppendContext(ctx context.Context, key string, value string) CacheInterface
	// 根据key和field字段设置，field字段的值
	HSet(key string, value interface{}, expiration time.Duration) CacheInterface
	HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CacheInterface
	// 根据key和field字段，查询field字段的值
	HGet(key string, field string) CacheInterface
	HGetContext(ctx context.Context, key string, field string) CacheInterface
	// 如果field字段不存在，则设置hash字段值
	HSetNX(key string, field string, value interface{}, expiration time.Duration) CacheInterface
	HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) CacheInterface
	// 根据key和多个字段名和字段值，批量设置hash字段值
	HMSet(key string, pairs ...interface{}) CacheInterface
	HMsetContext(ctx context.Context, key string, pairs ...interface{}) CacheInterface
	// 根据key和多个字段名，批量查询多个hash字段值
	HMGet(key string, fields ...string) CacheInterface
	HMGetContext(ctx context.Context, key string, fields ...string) CacheInterface
	// 根据key和field字段，累加字段的数值
	HIncrBy(key string, field string, value int64) CacheInterface
	HIncrByContext(ctx context.Context, key string, field string, value int64) CacheInterface
	// 根据key和field字段，累加字段的数值
	HIncrByFloat(key string, field string, value float64) CacheInterface
	HIncrByFloatContext(ctx context.Context, key string, field string, value float64) CacheInterface
	// 根据key和字段名，删除hash字段，支持批量删除hash字段
	HDel(key string, fields ...string) CacheInterface
	HDelContext(ctx context.Context, key string, fields ...string) CacheInterface
	// 根据key返回所有字段名
	HKeys(key string) CacheInterface
	HKeysContext(ctx context.Context, key string) CacheInterface
	// 根据key，查询hash的字段数量
	HLen(key string) CacheInterface
	HlenContext(ctx context.Context, key string) CacheInterface
	// 检测hash字段名是否存在。
	HExists(key string, field string) CacheInterface
	HExistsContext(ctx context.Context, key string, field string) CacheInterface
	// 从列表左边插入数据
	LPush(key string, value ...interface{}) CacheInterface
	LPushContext(ctx context.Context, key string, value ...interface{}) CacheInterface
	// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
	LPushX(key string, value ...interface{}) CacheInterface
	LPushXContext(ctx context.Context, key string, value ...interface{}) CacheInterface
	// 从列表的右边删除第一个数据，并返回删除的数据
	RPop(key string) CacheInterface
	RPopContext(ctx context.Context, key string) CacheInterface
	// 从列表右边插入数据
	RPush(key string, values ...interface{}) CacheInterface
	RPushContext(ctx context.Context, key string, values ...interface{}) CacheInterface
	// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
	RPushX(key string, values ...interface{}) CacheInterface
	RPushXContext(ctx context.Context, key string, values ...interface{}) CacheInterface
	// 从列表左边删除第一个数据，并返回删除的数据
	LPop(key string) CacheInterface
	LPopContext(ctx context.Context, key string) CacheInterface
	// 返回列表的大小
	LLen(key string) CacheInterface
	LLenContext(ctx context.Context, key string) CacheInterface
	// 返回列表的一个范围内的数据，也可以返回全部数据
	LRange(key string, start int64, stop int64) CacheInterface
	LRangeContext(ctx context.Context, key string, start int64, stop int64) CacheInterface
	// 删除列表中的数据
	LRem(key string, count int64, value interface{}) CacheInterface
	LRemContext(ctx context.Context, key string, count int64, value interface{}) CacheInterface
	// 根据索引坐标，查询列表中的数据
	LIndex(key string, index int64) CacheInterface
	LIndexContext(ctx context.Context, key string, index int64) CacheInterface
	// 在指定位置插入数据
	LInsert(key string, op string, pivot interface{}, value interface{}) CacheInterface
	LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) CacheInterface
	// 添加集合元素
	SAdd(key string, value ...interface{}) CacheInterface
	SAddContext(ctx context.Context, key string, value ...interface{}) CacheInterface
	// 获取集合元素个数
	SCard(key string) CacheInterface
	SCardContext(ctx context.Context, key string) CacheInterface
	// 判断元素是否在集合中
	SIsMember(key string, value interface{}) CacheInterface
	SIsMemberContext(ctx context.Context, key string, value interface{}) CacheInterface
	// 获取集合中所有的元素
	SMembers(key string) CacheInterface
	SMembersContext(ctx context.Context, key string) CacheInterface
	// 删除集合元素
	SRem(key string, value interface{}) CacheInterface
	SRemContext(ctx context.Context, key string, value interface{}) CacheInterface
	// 随机返回集合中的元素，并且删除返回的元素
	SPop(key string) CacheInterface
	SPopContext(ctx context.Context, key string) CacheInterface
	// 随机返回集合中的元素，并且删除返回的元素
	SPopN(key string, num int64) CacheInterface
	SPopNContext(ctx context.Context, key string, num int64) CacheInterface
	// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
	ZAdd(key string, redisZ *redis.Z) CacheInterface
	ZAddContext(ctx context.Context, key string, redisZ *redis.Z) CacheInterface
	// 返回集合元素个数
	ZCard(key string) CacheInterface
	ZCardContext(ctx context.Context, key string) CacheInterface
	// 统计某个分数范围内的元素个数
	ZCount(key string, min string, max string) CacheInterface
	ZCountContext(ctx context.Context, key string, min string, max string) CacheInterface
	// 增加元素的分数
	ZIncrBy(key string, incr float64, member string) CacheInterface
	ZIncrByContext(ctx context.Context, key string, incr float64, member string) CacheInterface
	// 返回集合中某个索引范围的元素，根据分数从小到大排序
	ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) CacheInterface
	ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CacheInterface
	// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
	// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
	ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) CacheInterface
	ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CacheInterface
	// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
	ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) CacheInterface
	ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CacheInterface
	// 删除集合元素
	ZRem(key string, members ...interface{}) CacheInterface
	ZRemContext(ctx context.Context, key string, members ...interface{}) CacheInterface
	// 根据索引范围删除元素
	ZRemRangeByRank(key string, start int64, stop int64) CacheInterface
	ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) CacheInterface
	// 根据分数范围删除元素
	ZRemRangeByScore(key string, min string, max string) CacheInterface
	ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) CacheInterface
	// 查询元素对应的分数
	ZScore(key string, member string) CacheInterface
	ZScoreContext(ctx context.Context, key string, member string) CacheInterface
	// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
	ZRank(key string, member string) CacheInterface
	ZRankContext(ctx context.Context, key string, member string) CacheInterface
}
