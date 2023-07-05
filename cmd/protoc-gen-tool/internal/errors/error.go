/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-23 15:55:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-24 19:06:49
 * @Description: error
 */
package errors

import "errors"

var Err_Plugins_Files_Is_Empty = errors.New("err : plugins files is empty")

var Err_Http_Rules_Options_Is_Empty = errors.New("err : httpRules options is empty")
var Err_Http_Rules_Extensions_Is_Empty = errors.New("err : httpRules extensions is empty")
var Err_Http_Rules_Is_Empty = errors.New("err : httpRules is empty")
var Err_Http_Rules_Methods_No_Match = errors.New("err : httpRules method is not match")

var Err_Proto_File_Is_Empty = errors.New("err : protoFiles is empty")

var Err_Http_Method_Is_Empty = errors.New("err : httpMethod is empty")
var Err_Http_Url_Is_Empty = errors.New("err : httpUrl is empty")

var Err_Gen_File_isGenerated_False = errors.New("err : genFile isGenerated is false")

var Err_Args_IsFirstCreate_Is_Empty = errors.New("err : args isFirstCreate is empty")
