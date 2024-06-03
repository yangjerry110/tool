/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 15:40:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-06-03 11:25:47
 * @Description: TestErr
 */
package test

import (
	"fmt"
	"testing"

	toolErrors "github.com/yangjerry110/tool/errors"
)

func TestError(t *testing.T) {

	err := testError()
	fmt.Printf("err : %+v", err)
	fmt.Print("\r\n")

}

func testError() error {
	return testTestError()
}

func testTestError() error {

	err := toolErrors.New("Err : test Error")
	return err
}
