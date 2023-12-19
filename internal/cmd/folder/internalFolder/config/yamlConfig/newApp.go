/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:47:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:47:26
 * @Description: newAPp
 */
package yamlconfig

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/internal/cmd/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:47:34
 * @return {*}
 */
func (n *NewApp) New() error {

	// protoPath
	path := fmt.Sprintf("%s/%s/%s/%s", config.ProjectPathConf.Path, "internal", "config", "yamlConfig")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
