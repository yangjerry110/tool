/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 17:04:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-04-24 15:18:17
 * @Description: errors
 */
package errors

import "errors"

var ErrProjectNameIsEmpty = errors.New("err : projectname is empty; please check")
var ErrProjectPathIsEmpty = errors.New("err : projectpath is empty; please check")

var ErrTemplateSavePathIsEmpty = errors.New("err : templates Err => path is empty")
var ErrTemplateSavePathFmtErr = errors.New("err : template Err => format is err")
