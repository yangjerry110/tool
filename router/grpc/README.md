<!--
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:41:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 14:43:13
 * @Description: 
-->
# gRPC Connection Pool (Go)

本项目实现了一个支持多服务、线程安全的 gRPC 连接池。

## 功能特点

- 支持按 `serviceName` 注册多个连接池
- 每个连接池只初始化一次（懒加载）
- 使用 Round-Robin 轮询获取连接
- 支持自定义 `Option` 扩展

## 使用示例

```go
import "github.com/yangjerry110/tool/router/grpc"

// 1. 初始化连接池
pool, err := grpc.New(
	"grpc.server",
	[]string{"127.0.0.1:12001"},
	toolRouterGrpc.WithPoolSize(6),
	toolRouterGrpc.WithHealthCheckInterval(10*time.Second),
	toolRouterGrpc.WithBalancerPolicy("round_robin"),
	toolRouterGrpc.WithDialTimeout(10*time.Second),
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
client := protobuf.NewGrpcClient(conn)
defer conn.Close()
```

## 测试

```bash
go test ./tests
```