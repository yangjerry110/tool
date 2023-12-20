/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-14 19:11:06
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-20 17:15:26
 * @Description: service
 */
package newfilesetservice

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"

	"github.com/yangjerry110/tool/internal/cmd/config"
)

type Service struct {
	FileContent string
	AppendFuncs []*config.ProtocHttpRule
}

/**
 * @description: NewFileSet
 * @param {string} filePath
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2023年12月15日11:29:00
 * @return {*}
 */
func (s *Service) NewFileSet(filePath string, fileName string) error {

	// define filePath + fileName
	parseFilePath := fmt.Sprintf("%s/%s", filePath, fileName)

	// Read the source code from the file
	code, err := ioutil.ReadFile(parseFilePath)
	if err != nil {
		return err
	}

	// Parse the source code string
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, parseFilePath, code, parser.ParseComments)
	if err != nil {
		return err
	}

	// Use ast.Walk to modify the AST
	ast.Walk(s, file)
	// Because ast.walk is read-only, inspect needs to be used to modify it
	// Use ast.Inspect to modify this AST
	// ast.Inspect(file, s.Inspect)

	// Gets the contents of the rendered file
	fileBytes := []byte{}
	buffer := bytes.NewBuffer(fileBytes)
	err = format.Node(buffer, fset, file)
	if err != nil {
		return err
	}

	// Set AppendFunc
	if err := s.setAppendFunc(file); err != nil {
		return err
	}

	// Get fileContent
	fileContent := buffer.String()
	s.FileContent = fileContent
	return nil
}

/**
 * @description: Visit
 * @param {ast.Node} node
 * @author: Jerry.Yang
 * @date: 2023-12-15 11:28:48
 * @return {*}
 */
func (s *Service) Visit(node ast.Node) (w ast.Visitor) {

	if interfaceType, ok := node.(*ast.InterfaceType); ok {

		// Define interfaceType methods
		interfaceTypeMethods := []*ast.Field{}

		// Re-render the entire methods according to protocHttpRules
		// Check protocHttpRules
		// If len == 0 ; return false
		if len(config.ProtocHttpRules) == 0 {
			return s
		}

		// For protocHttpRules
		for _, protoHttpRule := range config.ProtocHttpRules {

			// Define astField
			astField := &ast.Field{}

			// Set astFieldName
			astFieldName := &ast.Ident{Name: protoHttpRule.FuncName}

			// Set astFieldType
			astFieldType := &ast.FuncType{}

			// Define astFieldTypeParams
			astFieldTypeParams := []*ast.Field{}

			// ContextParam
			// Set ContextParamNames
			contextParamNames := []*ast.Ident{}
			contextParamNames = append(contextParamNames, &ast.Ident{Name: "ctx"})

			// Set ContextType
			contextParamType := &ast.SelectorExpr{X: ast.NewIdent("context"), Sel: ast.NewIdent("context")}

			// Set ContextParam
			contextParam := &ast.Field{Names: contextParamNames, Type: contextParamType}

			// ReqParam
			// Set ReqParamNames
			reqParamNames := []*ast.Ident{}
			reqParamNames = append(reqParamNames, &ast.Ident{Name: "inputVo"})

			// Set ReqType
			reqParamType := &ast.StarExpr{X: ast.NewIdent(fmt.Sprintf("protobuf.%s", protoHttpRule.InputName))}

			// Set ReqParam
			reqParam := &ast.Field{Names: reqParamNames, Type: reqParamType}

			// Append astFieldTypeParams
			astFieldTypeParams = append(astFieldTypeParams, contextParam)
			astFieldTypeParams = append(astFieldTypeParams, reqParam)

			// Define Results
			astFieldTypeResultParams := []*ast.Field{}

			// // Set outputNames
			// outputNames := []*ast.Ident{}
			// outputNames = append(outputNames, &ast.Ident{Name: ""})

			// Set OutputType
			outputResultParamType := &ast.StarExpr{X: ast.NewIdent(fmt.Sprintf("protobuf.%s", protoHttpRule.OutputName))}

			// Set Output
			// outputResultParam := &ast.Field{Names: outputNames, Type: outputResultParamType}
			outputResultParam := &ast.Field{Type: outputResultParamType}

			// // // Set ErrNames
			// errNames := []*ast.Ident{}
			// errNames = append(errNames, &ast.Ident{Name: "error"})

			// Set ErrType
			errResultType := ast.NewIdent("error")

			// Set Err
			errResutParam := &ast.Field{Type: errResultType}

			// Append astFieldTypeResultParams
			astFieldTypeResultParams = append(astFieldTypeResultParams, outputResultParam)
			astFieldTypeResultParams = append(astFieldTypeResultParams, errResutParam)

			// Set Result
			astFieldType.Results = &ast.FieldList{List: astFieldTypeResultParams}

			// Set Params
			astFieldType.Params = &ast.FieldList{List: astFieldTypeParams}

			// Set Comment
			commentList := []*ast.Comment{}
			commentList = append(commentList, &ast.Comment{Text: protoHttpRule.Description})

			// Set astField
			astField.Names = []*ast.Ident{astFieldName}
			astField.Type = astFieldType
			astField.Comment = &ast.CommentGroup{List: commentList}

			// Append InterfaceMethods
			interfaceTypeMethods = append(interfaceTypeMethods, astField)
		}

		// Set method
		interfaceType.Methods = &ast.FieldList{List: interfaceTypeMethods}
	}

	return s
}

/**
 * @description: setAppendFunc
 * @param {*ast.File} file
 * @author: Jerry.Yang
 * @date: 2023-12-18 16:35:31
 * @return {*}
 */
func (s *Service) setAppendFunc(file *ast.File) error {

	// Define
	// Exist funcs
	existFuncs := map[string]bool{}

	// Iterate through all declarations in the file
	for _, decl := range file.Decls {
		// If it is a function declaration
		if fn, ok := decl.(*ast.FuncDecl); ok {
			// Get the function name
			name := fn.Name.Name

			// Set Exist Funcs
			existFuncs[name] = true
		}
	}

	// If len config.ProtocHttpRules != 0
	// For config.ProtocHttpRules;
	// If not exist; add AppendFunc
	if len(config.ProtocHttpRules) != 0 {
		for _, protoHttpRule := range config.ProtocHttpRules {

			// If isExist
			_, isExist := existFuncs[protoHttpRule.FuncName]
			if !isExist {
				s.AppendFuncs = append(s.AppendFuncs, protoHttpRule)
			}
		}
	}
	return nil
}

/**
 * @description: Inspect
 * @param {ast.Node} node
 * @author: Jerry.Yang
 * @date: 2023-12-15 14:18:42
 * @return {*}
 */
func (s *Service) Inspect(node ast.Node) bool {
	return true
}
