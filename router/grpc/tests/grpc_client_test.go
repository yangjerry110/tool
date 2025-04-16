/*
 * @Author: Jerry.Yang
 * @Date: 2025-04-14 21:46:02
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:47:31
 * @Description:
 */
package tests

import (
	"log"
	"testing"
	"time"

	toolRouterGrpc "github.com/yangjerry110/tool/router/grpc"
)

func TestGrpcClient(t *testing.T) {
	// grpcClientConfig := &grpc.Config{}
	// grpcClientConfig.ServiceName = "necmdb.grpc.server"
	// grpcClientConfig.Endpoints = []string{"127.0.0.1:12001"}
	// ctx := context.Background()
	// grpcClientPool, err := grpc.GetGrpcClientPool(ctx, grpcClientConfig)
	// if err != nil {
	// 	fmt.Printf("GetGrpcClientPool Err : %+v \r\n", err)
	// 	return
	// }
	// defer grpcClientPool.Close()

	// // defer func() {
	// // 	if err := grpcClientPool.Close(); err != nil {
	// // 		panic(fmt.Sprintf("关闭连接池时出错: %v", err))
	// // 	}
	// // }()

	// grpcClient, err := grpcClientPool.GetClient()
	// if err != nil {
	// 	fmt.Printf("GetGrpcClientPool GetClient Err : %+v \r\n", err)
	// 	return
	// }

	// 创建客户端管理器
	// serviceName := "necmdb.grpc.server"
	// etcdAddrs := []string{"127.0.0.1:12001"}
	// grpcClient, err := toolRouterGrpc.InitGRPCClient(etcdAddrs)
	// if err != nil {
	// 	fmt.Printf("init grpc client err : %+v \r\n", err)
	// 	return
	// }
	// defer grpcClient.Close()

	// resolver.Register(&toolRouterGrpc.StaticResolver{})
	// 1. 初始化连接池
	pool, err := toolRouterGrpc.New(
		"newcmdb.grpc.server",
		[]string{"127.0.0.1:12001"},
		toolRouterGrpc.WithPoolSize(6),
		toolRouterGrpc.WithHealthCheckInterval(10*time.Second),
		toolRouterGrpc.WithBalancerPolicy("round_robin"),
		toolRouterGrpc.WithDialTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to create client pool: %v", err)
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
