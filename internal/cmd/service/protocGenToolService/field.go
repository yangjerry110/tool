/*
 * @Author: Jerry.Yang
 * @Date: 2025-02-25 10:46:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-02-26 14:36:25
 * @Description: field generate
 */
package protocgentoolservice

import (
	"fmt"
	"os"
	"strings"

	"github.com/yangjerry110/protoc-gen-go/compiler/protogen"
)

type Field struct {
	OsFile  *os.File
	Message *protogen.Message
	Field   *protogen.Field
}

/**
 * @description: field Generate
 * @author: Jerry.Yang
 * @date: 2025-02-26 14:36:23
 * @return {*}
 */
func (f *Field) Generate() error {

	// 为每个 optional 字段生成 Has<Field> 方法
	messageName := f.Message.GoIdent.GoName
	fieldName := f.Field.GoName
	capitalizedField := strings.Title(fieldName) // 转为大写开头的字段名
	content := fmt.Sprintf(`
	func (x *%s) Has%s() bool {
		return x.%s != nil
	}`, messageName, capitalizedField, fieldName)

	// 写入内容到文件
	_, err := f.OsFile.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
