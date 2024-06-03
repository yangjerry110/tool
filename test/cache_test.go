/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-31 14:31:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-31 15:09:33
 * @Description: TestCache
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/tool/cache"
	"github.com/yangjerry110/tool/conf"
)

func TestCache(t *testing.T) {

	if err := conf.CreatePathConf("/data/app/gopath/src/tool/test/yamlConfig").SetConfig(); err != nil {
		fmt.Printf("conf.CreatePathConf err : %+v", err)
		fmt.Print("\r\n")
		return
	}

	if err := cache.CreateRedisConf().SetConfig(); err != nil {
		fmt.Printf("cache.CreateRedisConf().SetConfig err : %+v", err)
		fmt.Print("\r\n")
		return
	}

	if err := cache.CreateRedisCache().CreateAllClient(); err != nil {
		fmt.Printf("cache.CreateRedisCache().CreateAllClient() err : %+v", err)
		fmt.Print("\r\n")
		return
	}

	if err := cache.CreateRedisCache().CreateClient("testRedis"); err != nil {
		fmt.Printf("cache.CreateRedisCache().CreateClient(\"testRedis\") err : %+v", err)
		fmt.Print("\r\n")
		return
	}

}
