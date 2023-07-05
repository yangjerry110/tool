/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-25 15:45:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-25 23:33:18
 * @Description: parse service
 */
package service

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/yangjerry110/tool/cmd/gen-tool/errors"
)

type ParseService interface {
	ParseFile(filePath string) (*Parse, error)
	GetNoExistFuncByParseFile(filePath string, httpRuleMap map[string]*HttpRule) ([]string, error)
	GetFuncNamesByAstNode(astNode ast.Node) bool
}

/**
 * @description: Parse
 * @author: Jerry.Yang
 * @date: 2023-05-25 14:27:38
 * @return {*}
 */
type Parse struct {
	FuncNames []string
}

/**
 * @description: GetFuncsByParseFile
 * @param {string} filePath
 * @author: Jerry.Yang
 * @date: 2023-05-25 15:47:55
 * @return {*}
 */
func (p *Parse) ParseFile(filePath string) (*Parse, error) {

	/**
	 * @step
	 * @newFileSet
	 **/
	newFileSet := token.NewFileSet()

	/**
	 * @step
	 * @解析文件内容
	 **/
	parseFileObj, err := parser.ParseFile(newFileSet, filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @遍历语法树
	 **/
	ast.Inspect(parseFileObj, func(astNode ast.Node) bool {
		return p.GetFuncNamesByAstNode(astNode)
	})
	return p, nil
}

/**
 * @description: GetNoExistFuncByParseFile
 * @param {string} filePath
 * @param {map[string]*HttpRule} httpRuleMap
 * @author: Jerry.Yang
 * @date: 2023-05-25 21:34:21
 * @return {*}
 */
func (p *Parse) GetNoExistFuncByParseFile(filePath string, httpRuleMap map[string]*HttpRule) ([]string, error) {

	/**
	 * @step
	 * @返回
	 **/
	noExistFuncNames := []string{}

	/**
	 * @step
	 * @定义参数
	 **/
	fileExistFuncNamesMap := map[string]string{}

	/**
	 * @step
	 * @获取filePath中已经有的funcNames
	 **/
	fileExistFuncNames, err := p.ParseFile(filePath)
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @循环filePath中exist的funcName
	 **/
	if len(fileExistFuncNames.FuncNames) != 0 {
		for _, funcName := range fileExistFuncNames.FuncNames {
			fileExistFuncNamesMap[funcName] = funcName
		}
	}

	/**
	 * @step
	 * @判断下httpRules
	 **/
	if len(httpRuleMap) == 0 {
		return nil, errors.ErrAstNodeInterfaceHttpRuleMapIsEmpty
	}

	/**
	 * @step
	 * @循环httpRules，判断哪里是需要新增的
	 **/
	for _, httpRule := range httpRuleMap {

		/**
		 * @step
		 * @是否存在当前文件中,methodName
		 **/
		_, isExist := fileExistFuncNamesMap[httpRule.FuncName]
		if !isExist {
			noExistFuncNames = append(noExistFuncNames, httpRule.FuncName)
		}
	}
	return noExistFuncNames, nil
}

/**
 * @description: GetFuncNamesByAstNode
 * @param {ast.Node} astNode
 * @author: Jerry.Yang
 * @date: 2023-05-25 16:06:06
 * @return {*}
 */
func (p *Parse) GetFuncNamesByAstNode(astNode ast.Node) bool {

	/**
	 * @step
	 * @判断是否是nil
	 **/
	// if astNode == nil {
	// 	return nil, errors.ErrAstNodeIsEmpty
	// }

	/**
	 * @step
	 * @根据不同node类型，处理
	 **/
	switch astNodeType := astNode.(type) {
	// 函数相关
	case *ast.FuncDecl:

		/**
		 * @step
		 * @获取函数名称
		 **/
		funcName := astNodeType.Name.Name
		p.FuncNames = append(p.FuncNames, funcName)
	}
	return true
}
