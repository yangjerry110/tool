/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-08 11:17:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-18 17:07:33
 * @Description: cache errors
 */
package errors

import (
	"errors"
	"fmt"
)

// all errors
var (

	// redis cache client
	ErrRedisClientConfIsNotExistByClientName = errors.New("err : redis_cache client is not exist conf by clientName")
	ErrRedisClientIsNotExistByClientName     = errors.New("err : redis_cache client is not exist by clientName")

	// config
	// config path
	ErrPathConfigIsNotConfigPath = errors.New("err : path_config is not configPath")
	// config yaml
	ErrYamlConfIsNoFilePath = errors.New("err : yaml_conf is no filePath")
	ErrYamlConfIsNoFileName = errors.New("err : yaml_conf is no fileName")
	ErrYamlConfIsNoConfData = errors.New("err : yaml_conf is no confData")

	// gormDb conf
	ErrGormDbConfIsNotExist = errors.New("err : gormDb_conf is not exist")
	// gormDb Client
	ErrGormDbClientIsNotExist = errors.New("err : gormDb_client is not exist")

	// perm
	// decrty perm
	ErrRsaPermNoPermPath  = errors.New("rsaPerm err : permPath must be not empty")
	ErrRsaPermNoInputFile = errors.New("rsaPerm err : inputStr must be not empty")
	ErrRsaPermDecrtyFail  = errors.New("rsaPerm err : decrty fail")
	// encrty perm
	ErrRsaPermEncrtyFail = errors.New("rsaPerm err : encrty fail")

	// cmd
	// cmd command
	ErrCmdCommandNoCliApp     = errors.New("cmd err : command no clieApp")
	ErrCmdCommandNoCliContext = errors.New("cmd err : command no CliContext")
	// cmd service
	ErrCmdServiceNoImportProjectPath = errors.New("service err : no projectImportPath")
	// cmd confing
	ErrCmdConfNoProtobufName = errors.New("conf err : no protobufName")
	ErrCmdConfNoAppName      = errors.New("conf err :  no appName")
	// service
	// protocGenToolService
	ErrProtocGenToolServiceNoPlugin       = errors.New("protocGenTool service Err : no plugin")
	ErrProtocGenToolServiceNoFiles        = errors.New("protocGenTool service Err : no files")
	ErrProtocGenToolServiceNoFile         = errors.New("protocGenTool service Err : no file")
	ErrProtocGenToolServiceNoGenerate     = errors.New("protocGenTool service Err : no generate")
	ErrProtocGenToolServiceNoMethods      = errors.New("protocGenTool service Err : no methods")
	ErrProtocGenToolServiceNoMethod       = errors.New("protocGenTool service Err : no method")
	ErrProtocGenToolServiceNoOptions      = errors.New("protocGenTool service Err : no options")
	ErrProtocGenToolServiceNoExtensions   = errors.New("protocGenTool service Err : no extensions")
	ErrProtocGenToolServiceNoHttpRule     = errors.New("protocGenTool service Err : no httpRule")
	ErrProtocGenToolServiceNoHttpRuleType = errors.New("protocGenTool service Err : no httpRule type")
	// cmd template
	ErrCmdTemplateNoPath = errors.New("cmd err : template no path")
	ErrCmdTemplateNoName = errors.New("cmd err :  template no name")

	// router
	// gin router
	ErrGinRouterIsNoDefault = errors.New("gin router Err : is no default")
)

// ErrRsaPermNoPrivatePath
func ErrRsaPermNoPrivatePath(privatePath string) error {
	errMsg := fmt.Sprintf("permPath is err; %s is not exist!", privatePath)
	return errors.New(errMsg)
}

// ErrRsaPermNoPublicPath
func ErrRsaPermNoPublicPath(publicPath string) error {
	errMsg := fmt.Sprintf("permPath is err; %s is not exist!", publicPath)
	return errors.New(errMsg)
}
