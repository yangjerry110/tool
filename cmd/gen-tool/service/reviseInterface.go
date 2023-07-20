/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-25 15:48:22
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-07-17 20:00:10
 * @Description: reviseInterface
 */
package service

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"github.com/yangjerry110/tool/cmd/gen-tool/errors"
)

type ReviseInterfaceService interface {
	GetReviseInterface(filePath string, httpRuleMap map[string]*HttpRule) (string, error)
}

type ReviseInterface struct {
	HttpRuleMap map[string]*HttpRule
}

/**
 * @description: ReviseInterface
 * @param {string} filePath
 * @author: Jerry.Yang
 * @date: 2023-05-25 15:47:30
 * @return {*}
 */
func (r *ReviseInterface) GetReviseInterface(filePath string, httpRuleMap map[string]*HttpRule) (string, error) {

	/**
	 * @step
	 * @定义参数，返回等等
	 **/
	fileContent := ""
	fileBytes := []byte{}

	/**
	 * @step
	 * @判断httpRuleMap
	 **/
	if len(httpRuleMap) == 0 {
		return "", errors.ErrAstNodeInterfaceHttpRuleMapIsEmpty
	}

	/**
	 * @step
	 * @赋值
	 **/
	r.HttpRuleMap = httpRuleMap

	/**
	 * @step
	 * @newFileSet
	 **/
	newFileSet := token.NewFileSet()

	/**
	 * @step
	 * @解析文件内容
	 **/
	parseFileObj, err := parser.ParseFile(newFileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// ast.Print(newFileSet, parseFileObj)

	/**
	 * @step
	 * @执行ast walk方法，循环执行每个自己定义的vistor
	 * @其实也不需要执行自己的定义的vistor，Inspect函数，后一个函数执行，也是重载了visit方法的
	 **/
	ast.Walk(r, parseFileObj)
	//ast.Inspect(parseFileObj, r.AstInspectAction)

	/**
	 * @step
	 * @获取渲染之后的文件内容
	 **/
	buffer := bytes.NewBuffer(fileBytes)
	err = format.Node(buffer, newFileSet, parseFileObj)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @赋值fileContent
	 **/
	fileContent = buffer.String()

	// /**
	//  * @step
	//  * @替换内容
	//  **/
	// fileContent = strings.ReplaceAll(fileContent, "protobuf.error", "error")
	return fileContent, nil
}

/**
 * @description: Visit
 * @param {ast.Node} node
 * @author: Jerry.Yang
 * @date: 2023-05-25 15:34:23
 * @return {*}
 */
func (r *ReviseInterface) Visit(astNode ast.Node) (w ast.Visitor) {

	/**
	 * @step
	 * @判断是否是nil
	 **/
	if astNode == nil {
		//CreateCommonService().PrintfErr(errors.ErrAstNodeIsEmpty, "ReviseInterface Visit")
		return r
	}

	/**
	 * @step
	 * @根据node的不同的类型，做不同的判断
	 * @判断是否是interface类型，然后根据上面解析到的funcNames，进行判断
	 * @把不在funcNames里面的funcName加入到interface到
	 * @根据funcName获取到对应的inputReqName，outputRespName
	 **/
	switch nodeType := astNode.(type) {
	// 接口类型
	case *ast.InterfaceType:

		/**
		 * @step
		 * @获取interfaceType
		 **/
		interfaceType := astNode.(*ast.InterfaceType)
		if err := r.Action(nodeType, interfaceType); err != nil {
			CreateCommonService().PrintfErr(err, "ReviseInterface Visit Action")
			return r
		}
	}
	return r
}

/**
 * @description: AstInspectAction
 * @param {ast.Node} node
 * @author: Jerry.Yang
 * @date: 2023-05-25 16:13:15
 * @return {*}
 */
func (r *ReviseInterface) AstInspectAction(astNode ast.Node) bool {

	/**
	 * @step
	 * @判断是否是nil
	 **/
	if astNode == nil {
		//CreateCommonService().PrintfErr(errors.ErrAstNodeIsEmpty, "AstInspectAction")
		return false
	}

	/**
	 * @step
	 * @根据node的不同的类型，做不同的判断
	 * @判断是否是interface类型，然后根据上面解析到的funcNames，进行判断
	 * @把不在funcNames里面的funcName加入到interface到
	 * @根据funcName获取到对应的inputReqName，outputRespName
	 **/
	switch nodeType := astNode.(type) {
	// 接口类型
	case *ast.InterfaceType:

		/**
		 * @step
		 * @获取interfaceType
		 **/
		interfaceType := astNode.(*ast.InterfaceType)
		if err := r.Action(nodeType, interfaceType); err != nil {
			CreateCommonService().PrintfErr(err, "AstInspectAction")
			return false
		}
	}
	return true
}

/**
 * @description: Action
 * @param {ast.Node} nodeType
 * @param {*ast.InterfaceType} interfaceType
 * @author: Jerry.Yang
 * @date: 2023-05-25 15:59:10
 * @return {*}
 */
func (r *ReviseInterface) Action(nodeType ast.Node, interfaceType *ast.InterfaceType) error {

	/**
	 * @step
	 * @判断接口的method
	 **/
	if interfaceType.Methods == nil || interfaceType.Methods.List == nil || len(interfaceType.Methods.List) == 0 {
		return errors.ErrAstNodeInterfaceMethodIsEmpty
	}

	/**
	 * @step
	 * @GetExistInterfaceMethodsNames
	 **/
	existInterfaceMethodNames, err := r.GetExistInterfaceMethodNames(interfaceType)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @获取需要增加的fields
	 **/
	appendAstFields, err := r.GetAppendAstFields(interfaceType, existInterfaceMethodNames)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @执行添加
	 **/
	interfaceType.Methods.List = append(interfaceType.Methods.List, appendAstFields...)
	return nil
}

/**
 * @description: GetAppendAstFields
 * @param {*ast.InterfaceType} interfaceType
 * @param {map[string]string} existInterfaceMethodNames
 * @author: Jerry.Yang
 * @date: 2023-05-25 17:22:29
 * @return {*}
 */
func (r *ReviseInterface) GetAppendAstFields(interfaceType *ast.InterfaceType, existInterfaceMethodNames map[string]string) ([]*ast.Field, error) {

	/**
	 * @step
	 * @定义需要增加astFields
	 **/
	appendAstFields := []*ast.Field{}

	/**
	 * @step
	 * @判断httpRuleMap
	 **/
	if len(r.HttpRuleMap) == 0 {
		return nil, errors.ErrAstNodeInterfaceHttpRuleMapIsEmpty
	}

	/**
	 * @step
	 * @循环httpRuleMap
	 **/
	for existMethodName, httpRule := range r.HttpRuleMap {

		/**
		 * @step
		 * @判断existInterfaceMethodNames是否存在
		 **/
		_, isExist := existInterfaceMethodNames[existMethodName]
		if isExist {
			continue
		}

		/**
		 * @step
		 * @设置当前method的参数,paramFields
		 **/
		paramsFields := []*ast.Field{}
		paramsFields = append(paramsFields, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent("ctx"),
			},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("context"),
				Sel: ast.NewIdent("Context"),
			},
		})
		paramsFields = append(paramsFields, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent("inputVo"),
			},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("*protobuf"),
				Sel: ast.NewIdent(httpRule.InputName),
			},
		})

		/**
		 * @step
		 * @设置method的参数，resultFields
		 **/
		resultFields := []*ast.Field{}
		resultFields = append(resultFields, &ast.Field{
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("*protobuf"),
				Sel: ast.NewIdent(httpRule.OutputName),
			},
		})
		resultFields = append(resultFields, &ast.Field{
			Type: ast.NewIdent("error"),
		})

		/**
		 * @step
		 * @设置method
		 **/
		methodFields := []*ast.Field{}
		methodFieldNames := []*ast.Ident{}
		methodFieldNames = append(methodFieldNames, &ast.Ident{
			Name: httpRule.FuncName,
			Obj:  ast.NewObj(ast.Var, httpRule.FuncName),
		})
		methodFields = append(methodFields, &ast.Field{
			Names: methodFieldNames,
			Type: &ast.FuncType{
				Func: methodFieldNames[0].End() + 1,
				Params: &ast.FieldList{
					List: paramsFields,
				},
				Results: &ast.FieldList{
					List: resultFields,
				},
			},
		})

		/**
		 * @step
		 * @执行添加
		 **/
		appendAstFields = append(appendAstFields, methodFields...)
	}
	return appendAstFields, nil
}

/**
 * @description: GetExistInterfaceMethodNames
 * @param {*ast.InterfaceType} interfaceType
 * @author: Jerry.Yang
 * @date: 2023-05-25 17:20:40
 * @return {*}
 */
func (r *ReviseInterface) GetExistInterfaceMethodNames(interfaceType *ast.InterfaceType) (map[string]string, error) {

	/**
	 * @step
	 * @返回
	 **/
	existInterfaceMethodNames := map[string]string{}

	/**
	 * @step
	 * @判断哪里是需要增加的
	 **/
	for _, methodInfo := range interfaceType.Methods.List {

		/**
		 * @step
		 * @定义methodName
		 **/
		methodName := ""

		/**
		 * @step
		 * @获取names的第一个
		 **/
		if len(methodInfo.Names) != 0 {
			for _, methodInfoName := range methodInfo.Names {
				methodName = methodInfoName.Name
			}
		}

		/**
		 * @step
		 * @进行赋值
		 **/
		existInterfaceMethodNames[methodName] = methodName
	}
	return existInterfaceMethodNames, nil
}
