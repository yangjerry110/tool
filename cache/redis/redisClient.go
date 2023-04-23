/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-12 14:24:55
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-11 16:07:09
 * @Description: client
 */
package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/cache"
	"github.com/yangjerry110/tool/conf"
)

type RedisCache struct {
	RedisClient            *redis.Client
	RedisErr               error
	RedisVal               interface{}
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

/**
 * @description: Client
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:27:09
 * @return {*}
 */
func (r *RedisCache) Client(filePath string, fileName string) cache.CacheInterface {
	yamlConf := conf.YamlConf{FilePath: filePath, FileName: fileName, Conf: &r}
	err := yamlConf.GetConf(r)
	if err != nil {
		r.RedisErr = err
		return r
	}

	/**
	 * @step
	 * @check
	 **/
	err = r.CheckConfig()
	if err != nil {
		r.RedisErr = err
		return r
	}

	/**
	 * @step
	 * @创建链接
	 **/
	redisClient := redis.NewClient(&redis.Options{
		Network:      r.Network,
		Addr:         r.Addr,
		Password:     r.Passwd,
		DB:           r.DB,
		DialTimeout:  time.Duration(r.DialTimeout),
		ReadTimeout:  time.Duration(r.ReadTimeout),
		WriteTimeout: time.Duration(r.WriteTimeout),
		PoolSize:     r.PoolSize,
		PoolTimeout:  time.Duration(r.PoolTimeout),
		MinIdleConns: r.MinIdleConns,
		MaxRetries:   r.MaxRetries,
	})
	r.RedisClient = redisClient
	return r
}

/**
 * @description: CreateDefaultClient
 * @author: Jerry.Yang
 * @date: 2022-10-12 11:21:56
 * @return {*}
 */
func (r *RedisCache) CreateDefaultClient() cache.CacheInterface {

	/**
	 * @step
	 * @创建链接
	 **/
	redisClient := redis.NewClient(&redis.Options{
		Network:      r.Network,
		Addr:         r.Addr,
		Password:     r.Passwd,
		DB:           r.DB,
		DialTimeout:  time.Duration(r.DialTimeout),
		ReadTimeout:  time.Duration(r.ReadTimeout),
		WriteTimeout: time.Duration(r.WriteTimeout),
		PoolSize:     r.PoolSize,
		PoolTimeout:  time.Duration(r.PoolTimeout),
		MinIdleConns: r.MinIdleConns,
		MaxRetries:   r.MaxRetries,
	})
	r.RedisClient = redisClient
	return r
}

/**
 * @description: CheckConfig
 * @author: Jerry.Yang
 * @date: 2022-10-25 18:59:35
 * @return {*}
 */
func (r *RedisCache) CheckConfig() error {
	if r.Addr == "" {
		return errors.New("redisClient Err : addr is not set!")
	}
	return nil
}
