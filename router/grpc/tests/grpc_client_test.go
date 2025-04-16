/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-14 21:46:02
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:55:44
 * @Description:
 */
package tests

import (
	"log"
	"testing"
	"time"

	"github.com/yangjerry110/tool/router/grpc"
)

func TestGrpcClient(t *testing.T) {

	// 1. 初始化连接池
	pool, err := grpc.New(
		"grpc.server",
		[]string{"127.0.0.1:12001"},
		grpc.WithPoolSize(6),
		grpc.WithHealthCheckInterval(10*time.Second),
		grpc.WithBalancerPolicy("round_robin"),
		grpc.WithDialTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// 2. 获取连接并创建客户端
	conn, err := pool.GetConn()
	if err != nil {
		log.Fatalf("No healthy connection: %v", err)
		return
	}
	defer conn.Close()
}
