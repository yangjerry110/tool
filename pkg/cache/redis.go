/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-13 16:46:14
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:23:20
 * @Description: redis cache
 */
package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	mytoolRedis "github.com/yangjerry110/tool/cache/redis"
)

type CachePkgRedis struct {
	RedisClient            *redis.Client
	CacheErr               error
	CacheVal               interface{}
	Network                string `yaml:"network"`                 //网络类型，支持：tcp，unix；默认tcp
	Addr                   string `yaml:"addr"`                    //网络地址，ip:port，如：172.0.0.1:6379
	Passwd                 string `yaml:"password"`                //密码
	DB                     int    `yaml:"database"`                //redis database，默认0；当前已不推荐使用多DB，该配置只为兼容一些存量系统多DB的使用
	DialTimeout            int    `yaml:"dial_timeout"`            //连接超时时间，默认1000ms
	ReadTimeout            int    `yaml:"read_timeout"`            //socket 读超时时间，默认100ms
	WriteTimeout           int    `yaml:"write_timeout"`           //socket 写超时时间，默认100ms
	PoolSize               int    `yaml:"pool_size"`               //连接池最大数量，默认200
	PoolTimeout            int    `yaml:"pool_timeout"`            //从连接池获取连接超时时间，默认ReadTimeout + 1000ms
	MinIdleConns           int    `yaml:"min_idle_conns"`          //连接池最小空闲连接数，默认30
	MaxRetries             int    `yaml:"max_retries"`             //重试次数，默认0
	TraceIncludeNotFound   bool   `yaml:"trace_include_not_found"` //是否将key NotFound error作为错误记录在trace中，默认为否
	MetricsIncludeNotFound bool   `yaml:"metrics_include_not_found"`
}

// 创建链接
func (c *CachePkgRedis) Client(filePath string, fileName string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{}
	CreateCache(cacheRedis).CacheInterface.Client(filePath, fileName)
	c.RedisClient = cacheRedis.RedisClient
	return c
}

/**
 * @description: CreateDefaultClient
 * @author: Jerry.Yang
 * @date: 2022-10-26 15:08:27
 * @return {*}
 */
func (c *CachePkgRedis) CreateDefaultClient() CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{
		Network:                c.Network,
		Addr:                   c.Addr,
		Passwd:                 c.Passwd,
		DB:                     c.DB,
		DialTimeout:            c.DialTimeout,
		ReadTimeout:            c.ReadTimeout,
		WriteTimeout:           c.WriteTimeout,
		PoolSize:               c.PoolSize,
		PoolTimeout:            c.PoolTimeout,
		MinIdleConns:           c.MinIdleConns,
		MaxRetries:             c.MaxRetries,
		TraceIncludeNotFound:   c.TraceIncludeNotFound,
		MetricsIncludeNotFound: c.MetricsIncludeNotFound,
	}
	CreateCache(cacheRedis).CacheInterface.CreateDefaultClient()
	c.RedisClient = cacheRedis.RedisClient
	return c
}

// 检查配置
func (c *CachePkgRedis) CheckConfig() error {
	return nil
}

// 获取结果和，错误
func (c *CachePkgRedis) Result() (interface{}, error) {
	return c.CacheVal, c.CacheErr
}

// 获取结果
func (c *CachePkgRedis) GetResult() interface{} {
	return c.CacheVal
}

// 获取错误
func (c *CachePkgRedis) GetErr() error {
	return c.CacheErr
}

// 给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
func (c *CachePkgRedis) Set(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.Set(key, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

func (c *CachePkgRedis) SetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SetContext(ctx, key, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 查询数据库中名称为key的value值
func (c *CachePkgRedis) Get(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.Get(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

func (c *CachePkgRedis) GetContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.GetContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
func (c *CachePkgRedis) SetNX(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SetNX(key, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

func (c *CachePkgRedis) SetNXContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SetNXContext(ctx, key, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 批量查询key的值。比如redisDb.MGet("name1","name2","name3")
func (c *CachePkgRedis) MGet(keys ...string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.MGet(keys...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

func (c *CachePkgRedis) MGetContext(ctx context.Context, keys ...string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.MGetContext(ctx, keys...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
func (c *CachePkgRedis) MSet(pairs ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.MSet(pairs...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) MsetContext(ctx context.Context, pairs ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.MsetContext(ctx, pairs...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// Incr函数每次加一,key对应的值必须是整数或nil
// 否则会报错incr key1: ERR value is not an integer or out of range
func (c *CachePkgRedis) Incr(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.Incr(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) IncrContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.IncrContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
func (c *CachePkgRedis) IncrBy(key string, value int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.IncrBy(key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) IncrByContext(ctx context.Context, key string, value int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.IncrByContext(ctx, key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
func (c *CachePkgRedis) IncrByFloat(key string, value float64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.IncrByFloat(key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) IncrByFloatContext(ctx context.Context, key string, value float64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.IncrByFloatContext(ctx, key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
func (c *CachePkgRedis) Decr(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.Decr(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) DecrContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.DecrContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
func (c *CachePkgRedis) DecrBy(key string, decrement int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.DecrBy(key, decrement)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) DecrByContext(ctx context.Context, key string, decrement int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.DecrByContext(ctx, key, decrement)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
func (c *CachePkgRedis) Del(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.Del(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) DelContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.DelContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 设置key的过期时间,单位秒
func (c *CachePkgRedis) Expire(key string, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.Expire(key, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ExpireContext(ctx context.Context, key string, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ExpireContext(ctx, key, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 给数据库中名称为key的string值追加value
func (c *CachePkgRedis) Append(key string, value string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.Append(key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) AppendContext(ctx context.Context, key string, value string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.AppendContext(ctx, key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key和field字段设置，field字段的值
func (c *CachePkgRedis) HSet(key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HSet(key, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HSetContext(ctx context.Context, key string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HSetContext(ctx, key, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key和field字段，查询field字段的值
func (c *CachePkgRedis) HGet(key string, field string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HGet(key, field)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HGetContext(ctx context.Context, key string, field string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HGetContext(ctx, key, field)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 如果field字段不存在，则设置hash字段值
func (c *CachePkgRedis) HSetNX(key string, field string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HSetNX(key, field, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HSetNXContext(ctx context.Context, key string, field string, value interface{}, expiration time.Duration) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HSetNXContext(ctx, key, field, value, expiration)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key和多个字段名和字段值，批量设置hash字段值
func (c *CachePkgRedis) HMSet(key string, pairs ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HMSet(key, pairs...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HMsetContext(ctx context.Context, key string, pairs ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HMsetContext(ctx, key, pairs...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key和多个字段名，批量查询多个hash字段值
func (c *CachePkgRedis) HMGet(key string, fields ...string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HMGet(key, fields...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HMGetContext(ctx context.Context, key string, fields ...string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HMGetContext(ctx, key, fields...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key和field字段，累加字段的数值
func (c *CachePkgRedis) HIncrBy(key string, field string, value int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HIncrBy(key, field, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HIncrByContext(ctx context.Context, key string, field string, value int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HIncrByContext(ctx, key, field, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key和field字段，累加字段的数值
func (c *CachePkgRedis) HIncrByFloat(key string, field string, value float64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HIncrByFloat(key, field, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HIncrByFloatContext(ctx context.Context, key string, field string, value float64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HIncrByFloatContext(ctx, key, field, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key和字段名，删除hash字段，支持批量删除hash字段
func (c *CachePkgRedis) HDel(key string, fields ...string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HDel(key, fields...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HDelContext(ctx context.Context, key string, fields ...string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HDelContext(ctx, key, fields...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key返回所有字段名
func (c *CachePkgRedis) HKeys(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HKeys(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HKeysContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HKeysContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据key，查询hash的字段数量
func (c *CachePkgRedis) HLen(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HLen(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HlenContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HlenContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 检测hash字段名是否存在。
func (c *CachePkgRedis) HExists(key string, field string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HExists(key, field)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) HExistsContext(ctx context.Context, key string, field string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.HExistsContext(ctx, key, field)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 从列表左边插入数据
func (c *CachePkgRedis) LPush(key string, value ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LPush(key, value...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LPushContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LPushContext(ctx, key, value...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 跟LPush的区别是，仅当列表存在的时候才插入数据,用法完全一样。
func (c *CachePkgRedis) LPushX(key string, value ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LPushX(key, value...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LPushXContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LPushXContext(ctx, key, value...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 从列表的右边删除第一个数据，并返回删除的数据
func (c *CachePkgRedis) RPop(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.RPop(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) RPopContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.RPopContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 从列表右边插入数据
func (c *CachePkgRedis) RPush(key string, values ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.RPush(key, values...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) RPushContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.RPushContext(ctx, key, values...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 跟RPush的区别是，仅当列表存在的时候才插入数据, 他们用法一样
func (c *CachePkgRedis) RPushX(key string, values ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.RPushX(key, values...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) RPushXContext(ctx context.Context, key string, values ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.RPushXContext(ctx, key, values...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 从列表左边删除第一个数据，并返回删除的数据
func (c *CachePkgRedis) LPop(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LPop(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LPopContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LPopContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 返回列表的大小
func (c *CachePkgRedis) LLen(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LLen(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LLenContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LLenContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 返回列表的一个范围内的数据，也可以返回全部数据
func (c *CachePkgRedis) LRange(key string, start int64, stop int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LRange(key, start, stop)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LRangeContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LRangeContext(ctx, key, start, stop)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 删除列表中的数据
func (c *CachePkgRedis) LRem(key string, count int64, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LRem(key, count, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LRemContext(ctx context.Context, key string, count int64, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LRemContext(ctx, key, count, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据索引坐标，查询列表中的数据
func (c *CachePkgRedis) LIndex(key string, index int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LIndex(key, index)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LIndexContext(ctx context.Context, key string, index int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LIndexContext(ctx, key, index)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 在指定位置插入数据
func (c *CachePkgRedis) LInsert(key string, op string, pivot interface{}, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LInsert(key, op, pivot, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) LInsertContext(ctx context.Context, key string, op string, pivot interface{}, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.LInsertContext(ctx, key, op, pivot, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 添加集合元素
func (c *CachePkgRedis) SAdd(key string, value ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SAdd(key, value...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) SAddContext(ctx context.Context, key string, value ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SAddContext(ctx, key, value...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 获取集合元素个数
func (c *CachePkgRedis) SCard(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SCard(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) SCardContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SCardContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 判断元素是否在集合中
func (c *CachePkgRedis) SIsMember(key string, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SIsMember(key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) SIsMemberContext(ctx context.Context, key string, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SIsMemberContext(ctx, key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 获取集合中所有的元素
func (c *CachePkgRedis) SMembers(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SMembers(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) SMembersContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SMembersContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 删除集合元素
func (c *CachePkgRedis) SRem(key string, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SRem(key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) SRemContext(ctx context.Context, key string, value interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SRemContext(ctx, key, value)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 随机返回集合中的元素，并且删除返回的元素
func (c *CachePkgRedis) SPop(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SPop(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) SPopContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SPopContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 随机返回集合中的元素，并且删除返回的元素
func (c *CachePkgRedis) SPopN(key string, num int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SPopN(key, num)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) SPopNContext(ctx context.Context, key string, num int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.SPopNContext(ctx, key, num)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
func (c *CachePkgRedis) ZAdd(key string, redisZ *redis.Z) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZAdd(key, redisZ)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZAddContext(ctx context.Context, key string, redisZ *redis.Z) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZAddContext(ctx, key, redisZ)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 返回集合元素个数
func (c *CachePkgRedis) ZCard(key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZCard(key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZCardContext(ctx context.Context, key string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZCardContext(ctx, key)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 统计某个分数范围内的元素个数
func (c *CachePkgRedis) ZCount(key string, min string, max string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZCount(key, min, max)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZCountContext(ctx context.Context, key string, min string, max string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZCountContext(ctx, key, min, max)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 增加元素的分数
func (c *CachePkgRedis) ZIncrBy(key string, incr float64, member string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZIncrBy(key, incr, member)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZIncrByContext(ctx context.Context, key string, incr float64, member string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZIncrByContext(ctx, key, incr, member)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 返回集合中某个索引范围的元素，根据分数从小到大排序
func (c *CachePkgRedis) ZRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRangeByScore(key, zRangeBy)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRangeByScoreContext(ctx, key, zRangeBy)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// ZRevRange用法跟ZRange一样，区别是ZRevRange的结果是按分数从大到小排序。
func (c *CachePkgRedis) ZRevRangeByScore(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRevRangeByScore(key, zRangeBy)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZRevRangeByScoreContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRevRangeByScoreContext(ctx, key, zRangeBy)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
func (c *CachePkgRedis) ZRangeByScoreWithScores(key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRangeByScoreWithScores(key, zRangeBy)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZRangeByScoreWithScoresContext(ctx context.Context, key string, zRangeBy *redis.ZRangeBy) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRangeByScoreWithScoresContext(ctx, key, zRangeBy)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 删除集合元素
func (c *CachePkgRedis) ZRem(key string, members ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRem(key, members...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZRemContext(ctx context.Context, key string, members ...interface{}) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRemContext(ctx, key, members...)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据索引范围删除元素
func (c *CachePkgRedis) ZRemRangeByRank(key string, start int64, stop int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRemRangeByRank(key, start, stop)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZRemRangeByRankContext(ctx context.Context, key string, start int64, stop int64) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRemRangeByRankContext(ctx, key, start, stop)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据分数范围删除元素
func (c *CachePkgRedis) ZRemRangeByScore(key string, min string, max string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRemRangeByScore(key, min, max)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZRemRangeByScoreContext(ctx context.Context, key string, min string, max string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRemRangeByScoreContext(ctx, key, min, max)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 查询元素对应的分数
func (c *CachePkgRedis) ZScore(key string, member string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZScore(key, member)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZScoreContext(ctx context.Context, key string, member string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZScoreContext(ctx, key, member)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}

// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (c *CachePkgRedis) ZRank(key string, member string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRank(key, member)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
func (c *CachePkgRedis) ZRankContext(ctx context.Context, key string, member string) CachePkgInterface {
	cacheRedis := &mytoolRedis.RedisCache{RedisClient: c.RedisClient}
	CreateCache(cacheRedis).CacheInterface.ZRankContext(ctx, key, member)
	c.CacheVal = cacheRedis.RedisVal
	c.CacheErr = cacheRedis.RedisErr
	return c
}
