/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:48:03
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:48:30
 * @Description: newApp
 */
package dao

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:48:29
 * @return {*}
 */
func (n *NewApp) New() error {
	// protoPath
	path := fmt.Sprintf("%s/%s/%s", config.ProjectPathConf.Path, "internal", "dao")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
