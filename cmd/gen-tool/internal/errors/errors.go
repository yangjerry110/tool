/*
 * @Author: Jerry.Yang
 * @Date: 2025-03-12 17:03:06
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-12 17:26:53
 * @Description: errors
 */
package errors

import "github.com/yangjerry110/tool/toolerrors"

var (
	ErrCommandNoCliContext       = toolerrors.New("gen-tool Err : command no cli context")
	ErrCommandNoCliApp           = toolerrors.New("gen-tool Err : command no cli app")
	ErrConfigNoImportProjectPath = toolerrors.New("gen-tool Err : config no import path")
	ErrConfigNoProtobufName      = toolerrors.New("gen-tool Err : config no protobuf name")
	ErrConfigTemplateNoName      = toolerrors.New("gen-tool Err : config template no name")
	ErrConfigNoAppName           = toolerrors.New("gen-tool Err : config no app name")
	ErrConfigNoCliContext        = toolerrors.New("gen-tool Err : config no cli context")
	ErrConfigNoModelName         = toolerrors.New("gen-tool Err : config no model name")
	ErrTemplateNoPath            = toolerrors.New("gen-tool Err : template no path")
)
