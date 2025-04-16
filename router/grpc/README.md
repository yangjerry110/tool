<!--
 * @Author: Jerry.Yang
 * @Date: 2025-04-16 14:41:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-04-16 15:43:14
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
grpc.Init(
	"grpc.server",
	[]string{"127.0.0.1:12001"},
	grpc.WithPoolSize(6),
	grpc.WithHealthCheckInterval(10*time.Second),
	grpc.WithBalancerPolicy("round_robin"),
	grpc.WithDialTimeout(10*time.Second),
)

clientPool, err := grpc.GetClientPool()
if err != nil {
	log.Fatalf("Failed to get client pool: %v", err)
	return
}
defer clientPool.Close()

// 2. 获取连接并创建客户端
conn, err := clientPool.GetConn("grpc.server")
if err != nil {
	log.Fatalf("No healthy connection: %v", err)
	return
}
defer conn.Close()
```

## 测试

```bash
go test ./tests
```