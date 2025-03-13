/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-21 14:14:56
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-21 14:15:47
 * @Description:
 */
package query

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
)

type NewAppQuery struct{}

/**
 * @description: New
 * @author: Jerry.Yang
 * @date: 2023-12-21 14:15:49
 * @return {*}
 */
func (n *NewAppQuery) New() error {
	// protoPath
	path := fmt.Sprintf("%s/%s/%s", config.ProjectPathConf.Path, "internal", "query")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
