/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 17:01:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 17:02:48
 * @Description: vo protobuf folder
 */
package vo

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/internal/cmd/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 17:02:50
 * @return {*}
 */
func (n *NewApp) New() error {
	// voPath
	path := fmt.Sprintf("%s/%s/%s", config.ProjectPathConf.Path, "vo", "protobuf")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
