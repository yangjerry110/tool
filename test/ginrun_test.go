/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-20 11:07:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-30 14:12:51
 * @Description: ginrun_test
 */
package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router"
	"github.com/yangjerry110/tool/internal/router/gin"
)

func TestGinRun(t *testing.T) {

	// SetConfigPath
	if err := conf.CreateConf(&conf.Path{ConfigPath: "/data/app/gopath/src/tool-api/internal/config/yamlConfig"}).SetConfig(); err != nil {
		fmt.Printf("err : %+v", err)
		fmt.Print("\r\n")
		return
	}

	// set gin Conf
	if err := conf.CreateConf(&gin.Gin{}).SetConfig(); err != nil {
		fmt.Printf("err : %+v", err)
		fmt.Print("\r\n")
		return
	}

	for {

		fmt.Printf("routerConf : %+v", router.RouteConf)
		fmt.Print("\r\n")

		time.Sleep(10 * time.Second)
	}

}
