/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:37:06
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-08 11:37:19
 * @Description: cache
 */
package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	CreateClient(ctx context.Context, clientName string) error
	GetClient(ctx context.Context, clientName string) (*redis.Client, error)
}
