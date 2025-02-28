/*
 * @Author: Jerry.Yang
 * @Date: 2024-05-30 15:40:37
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-28 11:24:37
 * @Description: TestErr
 */
package test

import (
	"fmt"
	"testing"

	"github.com/yangjerry110/tool/toolErrors"
	// pkgErr "github.com/pkg/errors"
)

func TestError(t *testing.T) {

	err := toolErrors.New("test-111111111")
	// fmt.Printf("TestError err : %+v", err)
	// fmt.Print("\r\n")

	// fmt.Printf("TestErr err.Error : %+v", err.Error())
	// fmt.Print("\r\n")

	newErr := toolErrors.NewError(err)
	fmt.Printf("TestError newErr : %+v", newErr)
	fmt.Print("\r\n")

	fmt.Printf("TestErr newErr.Error : %+v", newErr.Error())
	fmt.Print("\r\n")

}
