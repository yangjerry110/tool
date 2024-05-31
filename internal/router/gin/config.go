/*
* @Author: Jerry.Yang
* @Date: 2023-12-13 18:39:32
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 17:34:57
* @Description: gin conf
*/
package gin

import (
	toolredis "github.com/yangjerry110/tool/internal/cache/redis"
	"github.com/yangjerry110/tool/internal/conf"
	gormdb "github.com/yangjerry110/tool/internal/db/gormDb"
	"github.com/yangjerry110/tool/internal/router"
)

/**
 * @description: SetConfig
 * @author: Jerry.Yang
 * @date: 2023-12-14 10:47:37
 * @return {*}
 */
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
	if err := conf.CreateConf(&router.Config{}).SetConfig(); err != nil {
		return err
	}
	return nil
}
