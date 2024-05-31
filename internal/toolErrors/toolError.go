package toolErrors

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type ToolError struct {
	isGetRuntime bool
	runtimeDept  int
	packageName  string
	fileName     string
	funcName     string
	callFuncName string
	lineNo       int
	fields       map[string]interface{}
	error        string
}

/**
 * @description: New
 * @param {string} err
 * @author: Jerry.Yang
 * @date: 2024-05-31 10:32:03
 * @return {*}
 */
func (e *ToolError) New(err string) error {

	/**
	 * @step
	 * @default withFunc
	 **/
	e.runtimeDept = 2
	e.WithFunc().WithError(err)
	return e
}

/**
 * @description: WithPackage
 * @param {string} packageName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:31
 * @return {*}
 */
func (e *ToolError) WithPackage() ErrorInterface {

	/**
	 * @step
	 * @get runtimeDept
	 **/
	runtimeDept := e.getRuntimeDept()

	/**
	 * @step
	 * @Get the data for the upper two levels
	 **/
	_, file, _, ok := runtime.Caller(runtimeDept)

	/**
	 * @step
	 * @get packageName
	 **/
	if ok {
		packageName := filepath.Base(filepath.Dir(file))
		e.packageName = packageName
	}
	return e
}

/**
 * @description: WithFile
 * @param {string} fileName
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:14:40
 * @return {*}
 */
func (e *ToolError) WithFile() ErrorInterface {

	/**
	 * @step
	 * @get runtimeDept
	 **/
	runtimeDept := e.getRuntimeDept()

	/**
	 * @step
	 * @Get the data for the upper two levels
	 **/
	_, fileName, _, ok := runtime.Caller(runtimeDept)
	if ok {
		e.fileName = fileName
	}
	return e
}

/**
 * @description: WithFunc
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:41
 * @return {*}
 */
func (e *ToolError) WithFunc() ErrorInterface {

	/**
	 * @step
	 * @get runtimeDept
	 * @judge runtimeDept
	 * @if exist; get
	 **/
	runtimeDept := e.getRuntimeDept()
	if e.runtimeDept != 0 {
		runtimeDept = e.runtimeDept
	}

	/**
	 * @step
	 * @get funcName
	 **/
	pc, _, _, ok := runtime.Caller(runtimeDept)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		funcNameArr := strings.Split(funcName, ".")
		if len(funcNameArr) > 1 {
			e.funcName = funcNameArr[1]
		}
	}
	return e
}

/**
 * @description: WithLineNo
 * @param {int} lineNo
 * @author: Jerry.Yang
 * @date: 2024-05-30 16:15:02
 * @return {*}
 */
func (e *ToolError) WithLineNo() ErrorInterface {

	/**
	 * @step
	 * @get runtimeDept
	 **/
	runtimeDept := e.getRuntimeDept()

	/**
	 * @step
	 * @set lineNo
	 **/
	_, _, lineNo, ok := runtime.Caller(runtimeDept)
	if ok {
		e.lineNo = lineNo
	}
	return e
}

/**
 * @description: WithCallFuncName
 * @param {string} funcName
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:13:03
 * @return {*}
 */
func (e *ToolError) WithCallFuncName(funcName string) ErrorInterface {

	/**
	 * @step
	 * @set callFuncName
	 **/
	e.callFuncName = funcName
	return e
}

/**
 * @description: WithFields
 * @param {string} fieldName
 * @param {interface{}} fieldVal
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:00:52
 * @return {*}
 */
func (e *ToolError) WithFields(fieldName string, fieldVal interface{}) ErrorInterface {

	/**
	 * @step
	 * @judge e.fields
	 * @if == nil; new
	 **/
	if e.fields == nil {
		e.fields = make(map[string]interface{}, 0)
	}

	/**
	 * @step
	 * @set fields
	 **/
	e.fields[fieldName] = fieldVal
	return e
}

/**
 * @description: WithError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:08:29
 * @return {*}
 */
func (e *ToolError) WithError(err string) ErrorInterface {

	/**
	 * @step
	 * @define
	 **/
	errMsg := ""

	/**
	 * @step
	 * @packageName
	 **/
	if e.packageName != "" {
		errMsg = fmt.Sprintf("%s;", e.packageName)
	}

	/**
	 * @step
	 * @fileName
	 **/
	if e.fileName != "" {
		errMsg = fmt.Sprintf("%s %s;", errMsg, e.fileName)
	}

	/**
	 * @step
	 * @funcName
	 **/
	if e.funcName != "" {
		errMsg = fmt.Sprintf("%s %s;", errMsg, e.funcName)
	}

	/**
	 * @step
	 * @lineNo
	 **/
	if e.lineNo != 0 {
		errMsg = fmt.Sprintf("%s lineNo:%d;", errMsg, e.lineNo)
	}

	/**
	 * @step
	 * @callfuncName
	 **/
	if e.callFuncName != "" {
		errMsg = fmt.Sprintf("%s %s;", errMsg, e.callFuncName)
	}

	/**
	 * @step
	 * @fields
	 **/
	if len(e.fields) != 0 {
		for fieldName, fieldVal := range e.fields {
			errMsg = fmt.Sprintf("%s %s = %s;", errMsg, fieldName, fieldVal)
		}
	}

	/**
	 * @step
	 * @error
	 **/
	if err != "" {
		errMsg = fmt.Sprintf("%s Err : %s", errMsg, err)
	}

	/**
	 * @step
	 * @set error
	 **/
	e.error = errMsg
	return e
}

/**
 * @description: Error
 * @param {error} err
 * @author: Jerry.Yang
 * @date: 2024-05-30 15:01:02
 * @return {*}
 */
func (e *ToolError) Error() string {
	return e.error
}

/**
 * @description: getRuntimeDept
 * @author: Jerry.Yang
 * @date: 2024-05-30 17:49:27
 * @return {*}
 */
func (e *ToolError) getRuntimeDept() int {

	/**
	 * @step
	 * @judge isGetRuntime
	 * @if == true
	 * @return 1
	 **/
	if e.isGetRuntime {
		return 1
	}

	/**
	 * @step
	 * @init isGetRuntime
	 * @if == false
	 * @return 2
	 **/
	if !e.isGetRuntime {
		e.isGetRuntime = true
		return 2
	}
	return 2
}
