/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-22 16:13:48
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 16:16:00
 * @Description: Test get config
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/tool/internal/conf"
	"github.com/yangjerry110/tool/internal/router/gin"
)

func TestGetConfig(t *testing.T) {

	if err := conf.CreateConf(&gin.Gin{}).SetConfig(); err != nil {
		fmt.Printf("err : %+v", err)
		fmt.Print("\r\n")
		return
	}

	fmt.Printf("pathConfig : %+v", conf.PathConfig)
	fmt.Print("\r\n")

}
