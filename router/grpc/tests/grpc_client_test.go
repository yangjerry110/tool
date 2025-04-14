package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/yangjerry110/tool/router/grpc"
)

func TestGrpcClient(t *testing.T) {
	grpcClientConfig := &grpc.Config{}
	grpcClientConfig.Endpoints = []string{"172.25.128.247:50051"}
	ctx := context.Background()
	grpcClientPool, err := grpc.GetGrpcClientPool(ctx, grpcClientConfig)
	if err != nil {
		fmt.Printf("GetGrpcClientPool Err : %+v \r\n", err)
		return
	}

	defer func() {
		if err := grpcClientPool.Close(); err != nil {
			panic(fmt.Sprintf("关闭连接池时出错: %v", err))
		}
	}()

	_, err = grpcClientPool.GetClient()
	if err != nil {
		fmt.Printf("GetGrpcClientPool GetClient Err : %+v \r\n", err)
	}
}
