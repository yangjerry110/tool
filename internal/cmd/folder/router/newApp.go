/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:39:38
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:40:08
 * @Description:
 */
package router

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/internal/cmd/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:40:34
 * @return {*}
 */
func (n *NewApp) New() error {

	// routerPath
	path := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "router")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
