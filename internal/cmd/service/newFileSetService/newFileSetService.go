/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-14 19:11:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-18 16:37:12
 * @Description:
 */
package newfilesetservice

import "go/ast"

type NewFileSetService interface {
	NewFileSet(filePath string, fileName string) error
	Visit(node ast.Node) ast.Visitor
	Inspect(node ast.Node) bool
}

/**
 * @description: CreateNewFileSetService
 * @param {NewFileSetService} NewFileSetService
 * @author: Jerry.Yang
 * @date: 2023-12-14 19:13:06
 * @return {*}
 */
func CreateNewFileSetService(NewFileSetService NewFileSetService) NewFileSetService {
	return NewFileSetService
}
