/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-10 14:47:08
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-10 14:48:36
 * @Description: gin conf
 */
package gin

import (
	toolredis "github.com/yangjerry110/tool/internal/cache/redis"
	"github.com/yangjerry110/tool/internal/conf"
	gormdb "github.com/yangjerry110/tool/internal/db/gormDb"
	"github.com/yangjerry110/tool/internal/router"
)

func (g *Gin) SetConfig() error {
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
	if err := conf.CreateConf(&router.GinRouter{}).SetConfig(); err != nil {
		return err
	}
	return nil
}
