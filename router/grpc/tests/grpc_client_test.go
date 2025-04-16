/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-14 21:46:02
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 15:41:35
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
	grpc.Init(
		"grpc.server",
		[]string{"127.0.0.1:12001"},
		grpc.WithPoolSize(6),
		grpc.WithHealthCheckInterval(10*time.Second),
		grpc.WithBalancerPolicy("round_robin"),
		grpc.WithDialTimeout(10*time.Second),
	)

	clientPool, err := grpc.GetClientPool("grpc.server")
	if err != nil {
		log.Fatalf("Failed to get client pool: %v", err)
		return
	}
	defer clientPool.Close()

	// 2. 获取连接并创建客户端
	conn, err := clientPool.GetConn()
	if err != nil {
		log.Fatalf("No healthy connection: %v", err)
		return
	}
	defer conn.Close()
}
