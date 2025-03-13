/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-12 15:57:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-13 14:39:19
 * @Description: test
 */
package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/router"
)

func TestHttp(T *testing.T) {
	confPath := "/data/app/go/src/tool/router/test/config/yamlConfig"
	err := conf.CreatePathConf(confPath).SetConfig()
	if err != nil {
		fmt.Printf("CreateConfigPathConf err : %+v\r\n", err)
		return
	}

	err = router.SetHTTPRouterConfig().SetConfig()
	if err != nil {
		fmt.Printf("SetHTTPRouterConfig err : %+v\r\n", err)
		return
	}

	router.RegisterHTTP(&testRouter{}).
		RegisterHTTPService(&testService{}).
		RunHTTP(router.SetHTTPRouterConfig())

	time.Sleep(10 * time.Minute)

}

type testRouterHttpServer interface {
	router.RouterHTTPService
	TestRouterFunc(ctx context.Context) string
}

type testRouter struct {
	HttpServer testRouterHttpServer
}

func (t *testRouter) RouterName() string {
	return "testRouter"
}

func (t *testRouter) RegisterHTTP(ginEngine gin.IRouter) {
	ginEngine.GET("/testPing", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "success")
	})
	ginEngine.GET("testFunc", t.testRouterFunc)
}

func (t *testRouter) testRouterFunc(ctx *gin.Context) {
	msg := t.HttpServer.TestRouterFunc(ctx)
	ctx.String(http.StatusOK, msg)
}

func (t *testRouter) RegisterHTTPService(RouterHTTPService router.RouterHTTPService) {
	t.HttpServer = RouterHTTPService.(testRouterHttpServer)
}

type testService struct{}

func (*testService) TestRouterFunc(ctx context.Context) string {
	return "testService"
}

func (*testService) RouterName() string {
	return "testRouter"
}

func (*testService) MustRouterHTTPService() {}
