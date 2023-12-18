/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 10:48:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-08 11:29:40
 * @Description: cache redis
 */
package cacheredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yangjerry110/tool/internal/errors"
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
 * @description: CreateClient
 * @param {context.Context} ctx
 * @param {string} clientName
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:25:53
 * @return {*}
 */
func (r *RedisClient) CreateClient(ctx context.Context, clientName string) error {

	// get conf by clientName
	redisConf, isExistredisConf := RedisClientConfs[clientName]

	// redis conf is exist
	// if not exist => err
	if !isExistredisConf {
		return errors.ErrRedisClientConfIsNotExistByClientName
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
 * @param {context.Context} ctx
 * @param {string} clientName
 * @author: Jerry.Yang
 * @date: 2023-12-08 11:22:24
 * @return {*}
 */
func (r *RedisClient) GetClient(ctx context.Context, clientName string) (*redis.Client, error) {

	// judge clientName is exist
	existClient, isExistClient := RedisClients[clientName]
	if !isExistClient {
		return nil, errors.ErrRedisClientIsNotExistByClientName
	}
	return existClient, nil
}
