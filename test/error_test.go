/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 15:40:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-07 11:09:00
 * @Description: TestErr
 */
package test

import (
	"fmt"
	"testing"

	toolErrors "github.com/yangjerry110/tool/errors"
	// pkgErr "github.com/pkg/errors"
)

func TestError(t *testing.T) {

	err := testError()
	fmt.Printf("TestError err : %+v", err)
	fmt.Print("\r\n")

}

func testError() error {
	err := testTestError()
	return err
}

func testTestError() error {
	err := Err
	return toolErrors.NewError(err)
}

var Err = toolErrors.New("Err : Test Err")
