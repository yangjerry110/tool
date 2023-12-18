/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-15 11:30:05
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-18 16:50:06
 * @Description:
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/tool/internal/cmd/config"
	newfilesetservice "github.com/yangjerry110/tool/internal/cmd/service/newFileSetService"
)

func TestNewFileSet(t *testing.T) {

	filePath := "/data/app/gopath/src/tool/test/"
	fileName := "testNewFile.go"

	protocHttpRules := []*config.ProtocHttpRule{}
	protocHttpRule := &config.ProtocHttpRule{}
	protocHttpRule.Description = "testtest"
	protocHttpRule.FuncName = "getAccount"
	protocHttpRule.Method = "POST"
	protocHttpRule.InputName = "getAccountReq"
	protocHttpRule.OutputName = "getAccountResp"
	protocHttpRules = append(protocHttpRules, protocHttpRule)
	config.ProtocHttpRules = protocHttpRules

	if err := newfilesetservice.CreateNewFileSetService(&newfilesetservice.Service{}).NewFileSet(filePath, fileName); err != nil {
		fmt.Printf("err : %+v", err)
		fmt.Print("\r\n")
		return
	}

}
