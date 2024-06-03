/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 10:48:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-05-31 15:16:56
 * @Description: cache redis
 */
package cacheredis

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/internal/errors"
	"github.com/yangjerry110/tool/internal/toolErrors"
)

type RedisClient struct{}

/**
 * @description: RedisClients
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:26:02
 * @return {*}
 */
var RedisClients = map[string]*redis.Client{}

/**
 * @description: CreateAllClient
 * @author: Jerry.Yang
 * @date: 2023-12-21 15:24:53
 * @return {*}
 */
func (r *RedisClient) CreateAllClient() error {

	// If RedisClientConfs == 0
	// Return
	if len(RedisClientConfs) == 0 {
		return nil
	}

	// For RedisClientConfs
	for redisClientName := range RedisClientConfs {
		if err := r.CreateClient(redisClientName); err != nil {
			return toolErrors.NewError(err)
		}
	}
	return nil
}

/**
 * @description: CreateClient
 * @param {string} clientName
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:25:53
 * @return {*}
 */
func (r *RedisClient) CreateClient(clientName string) error {

	// get conf by clientName
	redisConf, isExistredisConf := RedisClientConfs[clientName]

	// redis conf is exist
	// if not exist => err
	if !isExistredisConf {
		return toolErrors.WithFields("clientName", clientName).WithFields("testFieldName", "testFieldVal").NewError(errors.ErrRedisClientConfIsNotExistByClientName)
	}

	// create client by client name
	// by redisConf
	redisClient := redis.NewClient(&redis.Options{
		Network:      redisConf.Network,
		Addr:         redisConf.Addr,
		Password:     redisConf.Passwd,
		DB:           redisConf.DB,
		DialTimeout:  time.Duration(redisConf.DialTimeout),
		ReadTimeout:  time.Duration(redisConf.ReadTimeout),
		WriteTimeout: time.Duration(redisConf.WriteTimeout),
		PoolSize:     redisConf.PoolSize,
		PoolTimeout:  time.Duration(redisConf.PoolTimeout),
		MinIdleConns: redisConf.MinIdleConns,
		MaxRetries:   redisConf.MaxRetries,
	})

	// set redisClient RedisClients
	RedisClients[clientName] = redisClient
	return nil
}

/**
 * @description: GetClient
 * @param {string} clientName
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:22:24
 * @return {*}
 */
func (r *RedisClient) GetClient(clientName string) (*redis.Client, error) {

	// judge clientName is exist
	existClient, isExistClient := RedisClients[clientName]
	if !isExistClient {
		return nil, toolErrors.NewError(errors.ErrRedisClientIsNotExistByClientName)
	}
	return existClient, nil
}
