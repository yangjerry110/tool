/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:41:54
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 16:54:07
 * @Description: newApp
 */
package proto

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:42:14
 * @return {*}
 */
func (n *NewApp) New() error {

	// protoPath
	path := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "protobuf")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
