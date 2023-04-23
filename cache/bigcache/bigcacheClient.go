/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-25 17:28:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-11 16:07:27
 * @Description: bigcache
 */
package bigcache

import (
	"errors"
	"time"

	"github.com/allegro/bigcache"
	"github.com/yangjerry110/tool/cache"
	"github.com/yangjerry110/tool/conf"
)

type BigCache struct {
	BigCache    *bigcache.BigCache
	BigCacheVal interface{}
	BigCacheErr error
	Eviction    time.Duration
	CleanTime   time.Duration
}

const (
	// offset64 FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset64 = 14695981039346656037
	// prime64 FNVa prime value. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	prime64 = 1099511628211
)

/**
 * @description: Client
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:27:00
 * @return {*}
 */
func (b *BigCache) Client(filePath string, fileName string) cache.CacheInterface {
	yamlConf := conf.YamlConf{FilePath: filePath, FileName: fileName, Conf: &b}
	err := yamlConf.GetConf(b)
	if err != nil {
		b.BigCacheErr = err
		return b
	}

	/**
	 * @step
	 * @check
	 **/
	err = b.CheckConfig()
	if err != nil {
		b.BigCacheErr = err
		return b
	}

	b.BigCache, b.BigCacheErr = bigcache.NewBigCache(b.DefaultConfig(b.Eviction, b.CleanTime))
	return b
}

/**
 * @description: CreateDefaultClient
 * @param {time.Duration} eviction
 * @param {time.Duration} cleanTime
 * @author: Jerry.Yang
 * @date: 2022-10-25 17:40:38
 * @return {*}
 */
func (b *BigCache) CreateDefaultClient() cache.CacheInterface {
	b.BigCache, b.BigCacheErr = bigcache.NewBigCache(b.DefaultConfig(b.Eviction, b.CleanTime))
	return b
}

/**
 * @description: DefaultConfig
 * @param {time.Duration} eviction
 * @param {time.Duration} cleanTime
 * @author: Jerry.Yang
 * @date: 2022-10-25 17:40:46
 * @return {*}
 */
func (b *BigCache) DefaultConfig(eviction time.Duration, cleanTime time.Duration) bigcache.Config {
	return bigcache.Config{
		Shards:             1024,
		LifeWindow:         eviction,
		CleanWindow:        cleanTime,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		Verbose:            true,
		Hasher:             b.NewDefaultHasher(),
		HardMaxCacheSize:   0,
		Logger:             bigcache.DefaultLogger(),
	}
}

/**
 * @description: CheckConfig
 * @author: Jerry.Yang
 * @date: 2022-10-26 14:44:22
 * @return {*}
 */
func (b *BigCache) CheckConfig() error {
	if b.Eviction == 0 {
		return errors.New("bigCache check eviction is not set")
	}
	return nil
}

/**
 * @description: NewDefaultHasher
 * @author: Jerry.Yang
 * @date: 2022-10-25 17:40:53
 * @return {*}
 */
func (b *BigCache) NewDefaultHasher() bigcache.Hasher {
	return b
}

/**
 * @description: Sum64
 * @param {string} key
 * @author: Jerry.Yang
 * @date: 2022-10-25 17:41:03
 * @return {*}
 */
func (b *BigCache) Sum64(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}
	return hash
}
