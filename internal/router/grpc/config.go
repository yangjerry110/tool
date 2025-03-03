/*
 * @Author: Jerry.Yang
 * @Date: 2024-08-19 10:43:53
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-03 16:44:12
 * @Description:
 */
package grpc

import (
	toolredis "github.com/yangjerry110/tool/internal/cache/redis"
	"github.com/yangjerry110/tool/internal/conf"
	gormdb "github.com/yangjerry110/tool/internal/db/gormDb"
	"github.com/yangjerry110/tool/internal/router"
)

type GrpcConfig struct{}

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2024-08-19 10:48:57
 * @return {*}
 */
func (g *GrpcConfig) SetConfig() error {

	// Set config path
	if err := conf.CreateConf(&conf.ConfigPath{}).SetConfig(); err != nil {
		return err
	}

	// Set db config
	if err := conf.CreateConf(&gormdb.GormDbConfig{}).SetConfig(); err != nil {
		return err
	}

	// Set redis config
	if err := conf.CreateConf(&toolredis.RedisConf{}).SetConfig(); err != nil {
		return err
	}

	// Set gin config
	if err := conf.CreateConf(&router.GrpcConfig{}).SetConfig(); err != nil {
		return err
	}
	return nil
}
