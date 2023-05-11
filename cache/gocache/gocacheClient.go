/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-09 14:23:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-09 15:48:12
 * @Description: gocache client
 */
package gocache

import (
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	toolCache "github.com/yangjerry110/tool/cache"
	"github.com/yangjerry110/tool/conf"
)

type GoCache struct {
	GoCacheVal        interface{}
	GoCacheErr        error
	GoCache           *cache.Cache
	DefaultExpireTime time.Duration
	DefaultDeleteTime time.Duration
}

/**
 * @description: defaultGoCache
 * @author: Jerry.Yang
 * @date: 2023-05-09 15:48:27
 * @return {*}
 */
var defaultGoCache = &GoCache{}

/**
 * @description: CreateGoCache
 * @author: Jerry.Yang
 * @date: 2023-05-09 15:35:01
 * @return {*}
 */
func CreateGoCache() toolCache.CacheInterface {
	return defaultGoCache
}

/**
 * @description: Client
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2023-05-09 14:38:59
 * @return {*}
 */
func (g *GoCache) Client(filePath string, fileName string) toolCache.CacheInterface {

	/**
	 * @step
	 * @获取配置
	 **/
	g.GetGoCaheClient(filePath, fileName)
	return g
}

/**
 * @description: CreateDefaultClient
 * @author: Jerry.Yang
 * @date: 2023-05-09 14:49:39
 * @return {*}
 */
func (g *GoCache) CreateDefaultClient() toolCache.CacheInterface {

	/**
	 * @step
	 * @判断是否已经初始化过了
	 **/
	if g.GoCache != nil {
		return g
	}

	/**
	 * @step
	 * @获取默认的配置
	 **/
	err := g.GetDefaultConfig()
	if err != nil {
		g.GoCacheErr = err
		return g
	}

	/**
	 * @step
	 * @cache
	 * @进行赋值
	 **/
	g.GoCache = cache.New(g.DefaultExpireTime, g.DefaultDeleteTime)
	return g
}

/**
 * @description: GetDefaultConfig
 * @author: Jerry.Yang
 * @date: 2023-05-09 14:47:34
 * @return {*}
 */
func (g *GoCache) GetDefaultConfig() error {
	g.DefaultExpireTime = 2 * time.Minute
	g.DefaultDeleteTime = 10 * time.Minute
	return nil
}

/**
 * @description: GetGoCaheClient
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2023-05-09 14:36:55
 * @return {*}
 */
func (g *GoCache) GetGoCaheClient(filePath string, fileName string) *cache.Cache {

	/**
	 * @step
	 * @判断是否已经初始化过了
	 **/
	if g.GoCache != nil {
		return g.GoCache
	}

	/**
	 * @step
	 * @渲染配置
	 **/
	yamlConf := conf.YamlConf{FilePath: filePath, FileName: fileName, Conf: &g}
	err := yamlConf.GetConf(g)
	if err != nil {
		g.GoCacheErr = err
		return nil
	}

	/**
	 * @step
	 * @checkConf
	 **/
	err = g.CheckConfig()
	if err != nil {
		g.GoCacheErr = err
		return nil
	}

	/**
	 * @step
	 * @进行赋值
	 **/
	g.GoCache = cache.New(g.DefaultExpireTime, g.DefaultDeleteTime)
	return g.GoCache
}

/**
 * @description: CheckConfig
 * @author: Jerry.Yang
 * @date: 2023-05-09 14:45:00
 * @return {*}
 */
func (g *GoCache) CheckConfig() error {

	/**
	 * @step
	 * @进行判断
	 **/
	if g.DefaultDeleteTime == 0 {
		return errors.New("gocache Err : defaultDeleteTime is empty")

	}

	/**
	 * @step
	 * @判断expireTime
	 **/
	if g.DefaultExpireTime == 0 {
		return errors.New("gocache Err : defaultExpireTime is empty")
	}
	return nil
}
