/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 15:51:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 15:51:21
 * @Description:
 */
package folder

import (
	"os"

	"github.com/yangjerry110/tool/internal/cmd/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 15:51:45
 * @return {*}
 */
func (n *NewApp) New() error {

	// Action MkDir
	err := os.MkdirAll(config.ProjectPathConf.Path, 0777)
	if err != nil {
		return err
	}
	return nil
}
