/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-11 15:18:28
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 14:50:57
 * @Description: logger test
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/internal/logger"
	"github.com/yangjerry110/tool/internal/logger/logrus"
	toolLogger "github.com/yangjerry110/tool/logger"
)

func TestLogger(t *testing.T) {

	conf.CreatePathConf("/data/app/gopath/src/tool/test/yamlConfig").SetConfig()
	err := toolLogger.CreateLogrusOptionConf()

	fmt.Printf("CreateLogrusOptionConf err : %+v", logrus.LogrusOptionConf)
	fmt.Print("\r\n")

	fmt.Printf("CreateLogrusOptionConf err : %+v", err)
	fmt.Print("\r\n")

	err = toolLogger.GetErr()
	fmt.Printf("toolLogger err : %+v", err)
	fmt.Print("\r\n")

	toolLogger.SetLevel(logger.TraceLevel)

	toolLogger.Debug("this is info log")
	toolLogger.Infof("this is info log : %+v", "this is test log msg")
	toolLogger.Trace("this is trace")

	testLogger1()

}

func testLogger1() {

	toolLogger.Debug("TestLogger1 this is info log")
	toolLogger.Infof("TestLogger1 this is info log : %+v", "this is test log msg")
	toolLogger.Trace("TestLogger1 this is trace")
}
