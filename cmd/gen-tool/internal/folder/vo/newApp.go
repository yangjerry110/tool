/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-19 17:01:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-19 17:02:18
 * @Description: vo folder
 */
package vo

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
)

type NewApp struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-19 17:02:58
 * @return {*}
 */
func (n *NewApp) New() error {
	// voPath
	path := fmt.Sprintf("%s/%s", config.ProjectPathConf.Path, "vo")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
