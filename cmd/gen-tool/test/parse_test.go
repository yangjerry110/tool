/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-25 18:47:28
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-26 14:41:21
 * @Description: parse Test
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/tool/cmd/gen-tool/service"
)

func TestParse(t *testing.T) {

	// /**
	//  * @step
	//  * @newFileSet
	//  **/
	// newFileSet := token.NewFileSet()

	// /**
	//  * @step
	//  * @解析文件内容
	//  **/
	// parseFileObj, err := parser.ParseFile(newFileSet, "parse.go", nil, 0)

	// fmt.Printf("\r\n Err : %+v \r\n", err)
	// ast.Print(newFileSet, parseFileObj)

	// return

	httpRuleMap := map[string]*service.HttpRule{}
	httpRuleMap["TestTwo"] = &service.HttpRule{
		FuncName:   "TestTwo",
		InputName:  "GetTestTwoReq",
		OutputName: "GetTestTwoResp",
	}

	fileContent, err := service.CreateReviseInterfaceService().GetReviseInterface("parse.go", httpRuleMap)

	fmt.Printf("\r\n Err : %+v \r\n", err)
	fmt.Printf("\r\n FileContent : %+v \r\n", fileContent)
}
