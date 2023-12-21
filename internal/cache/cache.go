/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:37:06
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 15:23:06
 * @Description: cache
 */
package cache

import (
	"github.com/go-redis/redis/v8"
)

type Cache interface {
	CreateAllClient() error
	CreateClient(clientName string) error
	GetClient(clientName string) (*redis.Client, error)
}
