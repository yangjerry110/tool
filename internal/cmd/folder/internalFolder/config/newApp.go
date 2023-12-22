/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:46:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:46:29
 * @Description:
 */
package config

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/internal/cmd/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:46:37
 * @return {*}
 */
func (n *NewApp) New() error {

	// protoPath
	path := fmt.Sprintf("%s/%s/%s", config.ProjectPathConf.Path, "internal", "config")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}