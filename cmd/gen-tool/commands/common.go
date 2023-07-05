/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-08 17:39:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-08 17:42:32
 * @Description: common
 */
package commands

import "sync"

type CommandsCommands interface {
	GetSyncMap() *sync.Map
}

type Common struct{}

/**
 * @description: syncMap
 * @author: Jerry.Yang
 * @date: 2023-05-08 16:43:53
 * @return {*}
 */
var syncMap *sync.Map

/**
 * @description: GetSyncMap
 * @author: Jerry.Yang
 * @date: 2023-05-08 17:41:14
 * @return {*}
 */
func (c *Common) GetSyncMap() *sync.Map {
	if syncMap != nil {
		return syncMap
	}

	syncMap = &sync.Map{}
	return syncMap
}
