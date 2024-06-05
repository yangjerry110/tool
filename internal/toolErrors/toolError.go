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
	stackTrace   string
	callFuncName string
	lineNo       int
	fields       map[string]interface{}
	errmsg       string
	error        error
}

/**
 * @description: init
 * @author: Jerry.Yang
 * @date: 2024-06-05 16:12:35
 * @return {*}
 */
func (e *ToolError) init() {
	e = &ToolError{}
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
	 * @init
	 **/
	e.init()

	/**
	 * @step
	 * @default withFunc
	 **/
	e.WithStackTrace().WithErrMsg(err)
	return e
}

/**
 * @description: NewError
 * @param {error} err
 * @author: Jerry.Yang
 * @date: NewError
 * @return {*}
 */
func (e *ToolError) NewError(err error) error {

	/**
	 * @step
	 * @init
	 **/
	e.init()

	/**
	 * @step
	 * @default withFunc
	 **/
	e.WithStackTrace().WithError(err)
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

	/**
	 * @step
	 * @get funcName
	 **/
	pc, _, _, ok := runtime.Caller(runtimeDept)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		e.funcName = funcName
	}
	return e
}

/**
 * @description: WithStackTrace
 * @author: Jerry.Yang
 * @date: 2024-06-03 11:14:16
 * @return {*}
 */
func (e *ToolError) WithStackTrace() ErrorInterface {

	/**
	 * @step
	 * @Only the first 10 layers of the survey stack are retrieved
	 **/
	pc := make([]uintptr, 10)

	/**
	 * @step
	 * @get runtimeDept
	 **/
	runtimeDept := e.getRuntimeDept()
	n := runtime.Callers(runtimeDept+2, pc)
	frames := runtime.CallersFrames(pc[:n])

	/**
	 * @step
	 * @for
	 * @get stackTrace
	 **/
	var stackTrace strings.Builder
	for {
		frame, more := frames.Next()
		stackTrace.WriteString(fmt.Sprintf("%s  \r\n  %s:%d   ", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}

	/**
	 * @step
	 * @set stackTrace
	 **/
	e.stackTrace = stackTrace.String()
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
	 * @getRuntimeDept
	 **/
	e.getRuntimeDept()

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
func (e *ToolError) WithError(err error) ErrorInterface {
	e.error = err
	return e
}

/**
 * @description: WithErrMsg
 * @param {string} err
 * @author: Jerry.Yang
 * @date: 2024-06-05 15:56:07
 * @return {*}
 */
func (e *ToolError) WithErrMsg(err string) ErrorInterface {
	e.errmsg = err
	return e
}

/**
 * @description: SetRuntimeDept
 * @param {int} runtimeDept
 * @author: Jerry.Yang
 * @date: 2024-05-31 15:03:22
 * @return {*}
 */
func (e *ToolError) SetRuntimeDept(runtimeDept int) ErrorInterface {
	e.runtimeDept = runtimeDept
	return e
}

/**
 * @description: GetError
 * @author: Jerry.Yang
 * @date: 2024-05-31 15:57:29
 * @return {*}
 */
func (e *ToolError) GetError() ErrorInterface {

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
		errMsg = fmt.Sprintf("packageName : %s;  \r\n  ", e.packageName)
	}

	/**
	 * @step
	 * @fileName
	 **/
	if e.fileName != "" {
		errMsg = fmt.Sprintf("%sflieName : %s;  \r\n  ", errMsg, e.fileName)
	}

	/**
	 * @step
	 * @funcName
	 **/
	if e.funcName != "" {
		errMsg = fmt.Sprintf("%sfuncName : %s;  \r\n  ", errMsg, e.funcName)
	}

	/**
	 * @step
	 * @lineNo
	 **/
	if e.lineNo != 0 {
		errMsg = fmt.Sprintf("%slineNo:%d;  \r\n  ", errMsg, e.lineNo)
	}

	/**
	 * @step
	 * @callfuncName
	 **/
	if e.callFuncName != "" {
		errMsg = fmt.Sprintf("%scallFuncName : %s;  \r\n  ", errMsg, e.callFuncName)
	}

	/**
	 * @step
	 * @fields
	 **/
	if len(e.fields) != 0 {
		for fieldName, fieldVal := range e.fields {
			errMsg = fmt.Sprintf("%s%s = %s;  \r\n  ", errMsg, fieldName, fieldVal)
		}
	}

	/**
	 * @step
	 * @error
	 **/
	if e.error != nil {
		errMsg = fmt.Sprintf("%s%s  \r\n  ", errMsg, e.error)
	}

	/**
	 * @step
	 * @errMsg
	 **/
	if e.errmsg != "" {
		errMsg = fmt.Sprintf("%s%s \r\n", errMsg, e.errmsg)
	}

	/**
	 * @step
	 * @stackTrace
	 **/
	if e.stackTrace != "" {
		errMsg = fmt.Sprintf("%sstackTrace:  \r\n  %s", errMsg, e.stackTrace)
	}

	/**
	 * @step
	 * @set error
	 **/
	e.errmsg = errMsg
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
	return e.GetError().String()
}

/**
 * @description: String
 * @author: Jerry.Yang
 * @date: 2024-05-31 15:43:35
 * @return {*}
 */
func (e *ToolError) String() string {
	return e.errmsg
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
		return e.runtimeDept + 1
	}

	/**
	 * @step
	 * @init isGetRuntime
	 * @if == false
	 * @return 2
	 **/
	if !e.isGetRuntime {
		e.isGetRuntime = true
		e.runtimeDept = 2
		return e.runtimeDept
	}
	return e.runtimeDept
}
