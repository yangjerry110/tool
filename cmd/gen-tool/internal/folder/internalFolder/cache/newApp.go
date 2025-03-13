/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 11:14:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 11:15:08
 * @Description: internal cache
 */
package cache

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
)

type NewAppCache struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 11:15:07
 * @return {*}
 */
func (n *NewAppCache) New() error {

	// protoPath
	path := fmt.Sprintf("%s/%s/%s", config.ProjectPathConf.Path, "internal", "cache")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
