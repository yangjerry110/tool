/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-25 17:50:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-11 16:07:21
 * @Description:
 */
package freecache

import (
	"errors"
	"time"

	"github.com/coocood/freecache"
	"github.com/yangjerry110/tool/cache"
	"github.com/yangjerry110/tool/conf"
)

type FreeCache struct {
	FreeCache    *freecache.Cache
	FreeCacheVal interface{}
	FreeCacheErr error
	Size         int
	Timer        freecache.Timer
}

/**
 * @description: Client
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:26:49
 * @return {*}
 */
func (f *FreeCache) Client(filePath string, fileName string) cache.CacheInterface {
	yamlConf := conf.YamlConf{FilePath: filePath, FileName: fileName, Conf: &f}
	err := yamlConf.GetConf(f)
	if err != nil {
		f.FreeCacheErr = err
		return f
	}

	/**
	 * @step
	 * @check
	 **/
	err = f.CheckConfig()
	if err != nil {
		f.FreeCacheErr = err
		return f
	}
	f.FreeCache = freecache.NewCacheCustomTimer(f.Size, f.Timer)
	return f
}

/**
 * @description: CreateClient
 * @param {int} size
 * @param {freecache.Timer} timer
 * @author: Jerry.Yang
 * @date: 2022-10-25 18:01:15
 * @return {*}
 */
func (f *FreeCache) CreateDefaultClient() cache.CacheInterface {
	f.FreeCache = freecache.NewCacheCustomTimer(f.Size, f.Timer)
	return f
}

/**
 * @description: CheckConfig
 * @author: Jerry.Yang
 * @date: 2022-10-26 15:21:29
 * @return {*}
 */
func (f *FreeCache) CheckConfig() error {
	if f.Size == 0 {
		return errors.New("freecache check : size is not set")
	}
	return nil
}

/**
 * @description: FreeCache
 * @author: Jerry.Yang
 * @date: 2022-10-25 18:06:53
 * @return {*}
 */
func (f *FreeCache) Now() uint32 {
	return uint32(time.Now().Add(60 * time.Minute).Unix())
}
