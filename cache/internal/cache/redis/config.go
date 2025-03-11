/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:07:10
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 10:52:39
 * @Description: redis cache
 */
package cacheredis

import (
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/toolerrors"
)

type RedisConf struct{}

/**
 * @description: RedisClientConf
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:10:50
 * @return {*}
 */
type RedisClientConf struct {
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
 * @description: RedisClientConfs
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:01:56
 * @return {*}
 */
var RedisClientConfs = map[string]*RedisClientConf{}

/**
 * @description: SetConfig
 * @param {context.Context} ctx
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2023-12-08 10:59:38
 * @return {*}
 */
func (r *RedisConf) SetConfig() error {

	/**
	 * @step
	 * @返回结果
	 **/
	if err := conf.CreateConf(&conf.Yaml{FilePath: conf.PathConfig.ConfigPath, FileName: "redis.yaml", FileType: "yaml", ConfData: RedisClientConfs}).SetConfig(); err != nil {
		return toolerrors.NewError(err)
	}
	return nil
}
