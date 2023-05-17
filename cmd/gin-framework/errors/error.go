/*
 * @Author: Jerry.Yang
 * @Date: 2023-04-23 17:04:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-16 17:31:19
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
var ErrControllerNameIsEmpty = errors.New("err : controllername is empty; please check")
var ErrRouterNameIsEmpty = errors.New("err : routername is empty; please check")
var ErrServiceNameIsEmpty = errors.New("err : servicename is empty; please check")
var ErrVoNameIsEmpty = errors.New("err : voname is empty; please check")

var ErrGetInitParams = errors.New("err : get initParams is err; ")

var ErrAppendControllerPathIsEmpyty = errors.New("err : append controllerPath is empty; please check")
var ErrAppendRouterPathIsEmpty = errors.New("err : append routerpath is empty; please check")
var ErrAppendServicePathIsEmpty = errors.New("err : append servicePath is empty; please check")
var ErrAppendVoPathIsEmprty = errors.New("err : append voPath is empty; please check")
var ErrDaoPathIsEmpty = errors.New("err : append daoPath is empty; please check")

var ErrModelNameIsEmpty = errors.New("err : modelname is empty; please check")
var ErrModelConfigIsEmpty = errors.New("err : modelconfig is empty; please check")

var ErrDaoNameIsEmpty = errors.New("err : daoname is empty; please check")
