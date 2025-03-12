package interfaceService

import (
	"fmt"
	"os"

	"github.com/yangjerry110/tool/cmd/gen-tool/internal/config"
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
	path := fmt.Sprintf("%s/%s/%s/interfaceService", config.ProjectPathConf.Path, "internal", "service")

	// Action MkDir
	err := os.MkdirAll(path, 0777)
	if err != nil {
		return err
	}
	return nil
}
