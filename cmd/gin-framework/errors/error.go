/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 17:04:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-11 11:26:29
 * @Description: errors
 */
package errors

import "errors"

var ErrProjectNameIsEmpty = errors.New("err : projectname is empty; please check")
var ErrProjectPathIsEmpty = errors.New("err : projectpath is empty; please check")
var ErrImportProjectPathIsEmpty = errors.New("err : importProjectPath is empty")

var ErrTemplateSavePathIsEmpty = errors.New("err : templates Err => path is empty")
var ErrTemplateSavePathFmtErr = errors.New("err : template Err => format is err")

var ErrAppNameIsEmpty = errors.New("err : appname is empty; please check")

var ErrGetInitParams = errors.New("err : get initParams is err; ")

var ErrModelNameIsEmpty = errors.New("err : modelname is empty; please check")
var ErrModelConfigIsEmpty = errors.New("err : modelconfig is empty; please check")

var ErrDaoNameIsEmpty = errors.New("err : daoname is empty; please check")
