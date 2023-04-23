/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 18:28:08
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-09 19:39:02
 * @Description: logrus
 */
package logger

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

/**
 * @description: LogrusLog
 * @author: Jerry.Yang
 * @date: 2022-09-27 17:45:17
 * @return {*}
 */
type LogrusLog struct {
	Logger         *logrus.Entry    // entry
	CallDept       int              // 记录调用者的层级
	WithFields     logrus.Fields    // fields
	LogLevel       logrus.Level     // 日志级别
	IsReportCaller bool             // 是否在日志中添加方法名和文件
	Output         io.Writer        // 输出对象
	Formatter      logrus.Formatter // 日志输入方式
	Options        *LogrusOption    // option
	LoggerErr      error
}

/**
 * @description: WriteLog
 * @param {...interface{}} args
 * @author: Jerry.Yang
 * @date: 2022-09-27 18:51:31
 * @return {*}
 */
func (l *LogrusLog) WriteLog(args ...interface{}) error {
	/**
	 * @step
	 * @检查参数
	 **/
	l.CheckParams()
	if l.LoggerErr != nil {
		l.CreateDefaultLog()
		l.Logger.Log(l.LogLevel, "WriteLog Err : %v", l.LoggerErr)
		return l.LoggerErr
	}

	levelIsTrue := l.Logger.Logger.IsLevelEnabled(l.Logger.Level)
	if !levelIsTrue {
		l.Logger.Log(l.LogLevel, "WriteLog Err : WriteLog Err : logLevel is invalid")
		return errors.New("WriteLog Err : logLevel is invalid")
	}

	/**
	 * @step
	 * @添加wirthFields
	 **/
	if l.WithFields != nil {
		if l.IsReportCaller {
			if err, ok := l.WithFields[LOGRUS_ERROR_KEY].(interface {
				Stack() []string
			}); ok {
				l.WithFields["err.stack"] = strings.Join(err.Stack(), ";")
			}
		}
		l.Logger.WithFields(l.WithFields).Log(l.LogLevel, args...)
		return nil
	}
	l.Logger.Log(l.LogLevel, args...)
	return nil
}

/**
 * @description: CreateDefaultLog
 * @author: Jerry.Yang
 * @date: 2022-09-29 16:17:35
 * @return {*}
 */
func (l *LogrusLog) CreateDefaultLog() LoggerInterface {
	logrusOptionFuncs := []LoggerOptionFunc{}

	/**
	 * @step
	 * @判断哪些需要赋值默认配置的
	 * @level
	 **/
	if l.Options.Options[OPTION_LOGLEVEL].Value == nil {
		logrusOptionFuncs = append(logrusOptionFuncs, l.Options.SetLevel(WarnLevel))
	}

	/**
	 * @step
	 * @formatter
	 **/
	if l.Options.Options[OPTION_FORMATTER_TYPE].Value == nil {
		logrusOptionFuncs = append(logrusOptionFuncs, l.Options.SetFormatter(LOGRUS_FORMATTER_JSON))
	}

	/**
	 * @step
	 * @check isReportcaller
	 **/
	if l.Options.Options[OPTION_ISREPORTCALLER].Value == nil {
		logrusOptionFuncs = append(logrusOptionFuncs, l.Options.SetIsReportcaller(true))
	}

	/**
	 * @step
	 * @check output
	 **/
	if l.Options.Options[OPTION_OUTPUT].Value == nil {
		logrusOptionFuncs = append(logrusOptionFuncs, l.Options.SetOutput(os.Stdout))
	}

	/**
	 * @step
	 * @check FormatterDisableHtmlEscap
	 **/
	if l.Options.Options[OPTION_FORMATTER_DISABLEHTMLESCAP].Value == nil {
		logrusOptionFuncs = append(logrusOptionFuncs, l.Options.SetFormatterDisableHtmlEscap(true))
	}

	/**
	 * @step
	 * @假如没有需要加入的配置
	 * @则说明已经加过了，则直接跳过
	 **/
	if len(logrusOptionFuncs) == 0 {
		return l
	}

	l.Options.SetOptions(logrusOptionFuncs)
	l.SetLevel().SetFormatter().SetOutput().SetLogger()
	return l
}

/**
 * @description: SetLogger
 * @author: Jerry.Yang
 * @date: 2022-09-27 17:45:02
 * @return {*}
 */
func (l *LogrusLog) SetLogger() LoggerInterface {

	/**
	 * @step
	 * @创建logrus
	 **/
	logrusNew := logrus.New()

	/**
	 * @step
	 * @设置日志级别
	 **/
	logrusNew.SetLevel(l.LogLevel)

	/**
	 * @step
	 * @获取设置format
	 **/
	logrusNew.SetFormatter(l.Formatter)

	/**
	 * @step
	 * @设置output
	 **/
	logrusNew.SetOutput(l.Output)

	/**
	 * @step
	 * @设置no lock
	 **/
	logrusNew.SetNoLock()

	/**
	 * @step
	 * @设置SetReportCaller
	 **/
	if l.IsReportCaller {
		//logrusNew.SetReportCaller(l.IsReportCaller)

		/**
		 * @step
		 * @记录调用者的信息
		 **/
		l.SetCaller()
	}

	/**
	 * @step
	 * @设置logger
	 **/
	l.Logger = logrus.NewEntry(logrusNew)
	return l
}

/**
 * @description: SetLevel
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:29:10
 * @return {*}
 */
func (l *LogrusLog) SetLevel() LoggerInterface {

	/**
	 * @step
	 * @获取日志logrus日志级别
	 **/
	logrusLevel, err := l.Options.GetLogrusLevel()
	if err != nil {
		l.LoggerErr = err
	}
	l.LogLevel = logrusLevel
	return l
}

/**
 * @description: setWithFields
 * @author: Jerry.Yang
 * @date: 2022-09-29 17:16:41
 * @return {*}
 */
func (l *LogrusLog) SetWithFields() LoggerInterface {

	/**
	 * @step
	 * @获取
	 **/
	withFields, err := l.Options.GetWithFields()
	if err != nil {
		return l
	}
	l.WithFields = withFields

	/**
	 * @step
	 * @获取dept
	 **/
	dept, err := l.Options.GetCallerDept()
	if err != nil {
		return l
	}

	/**
	 * @step
	 * @减dept
	 **/
	l.CallDept = dept - 1
	return l
}

/**
 * @description: SetCallDept
 * @author: Jerry.Yang
 * @date: 2022-10-09 18:40:24
 * @return {*}
 */
func (l *LogrusLog) SetCallDept() LoggerInterface {
	/**
	 * @step
	 * @获取
	 **/
	setval, err := l.Options.GetCallerDept()
	if err != nil {
		return l
	}
	l.CallDept = setval
	return l
}

/**
 * @description: SetCaller
 * @author: Jerry.Yang
 * @date: 2022-10-09 18:31:02
 * @return {*}
 */
func (l *LogrusLog) SetCaller() LoggerInterface {
	_, f, n, ok := runtime.Caller(l.CallDept)
	if !ok {
		return l
	}

	/**
	 * @step
	 * @判断withFields是否为空
	 **/
	if l.WithFields == nil {
		l.WithFields = make(logrus.Fields)
	}

	/**
	 * @step
	 * @拼接调用信息
	 **/
	l.WithFields["file"] = fmt.Sprintf("%s:%d", f, n)
	return l
}

/**
 * @description: SetFormatter
 * @author: Jerry.Yang
 * @date: 2022-09-27 17:16:58
 * @return {*}
 */
func (l *LogrusLog) SetFormatter() LoggerInterface {

	/**
	 * @step
	 * @获取日志的输入格式
	 **/
	logrusFormatterType, err := l.Options.GetFormatterType()
	if err != nil {
		l.LoggerErr = err
	}

	/**
	 * @step
	 * @判断假如是json
	 **/
	if logrusFormatterType == LOGRUS_FORMATTER_JSON {

		/**
		 * @step
		 * @获取是否转义html
		 **/
		disableHtmlScape, err := l.Options.GetDisableHtmlEscape()
		if err != nil {
			l.LoggerErr = err
		}

		/**
		 * @step
		 * @时间
		 **/
		disableTime, _ := l.Options.GetDisableTime()
		l.Formatter = &logrus.JSONFormatter{DisableHTMLEscape: disableHtmlScape, DisableTimestamp: disableTime}
	}

	/**
	 * @step
	 * @假如是text
	 **/
	if logrusFormatterType == LOGRUS_FORMATTER_TEXT {
		l.Formatter = &logrus.TextFormatter{}
	}
	return l
}

/**
 * @description: SetOutput
 * @author: Jerry.Yang
 * @date: 2022-09-27 17:25:15
 * @return {*}
 */
func (l *LogrusLog) SetOutput() LoggerInterface {

	/**
	 * @step
	 * @获取设置的outout
	 **/
	setOutput, err := l.Options.GetOutput()
	if err != nil {
		l.LoggerErr = err
	}
	l.Output = setOutput
	return l
}

/**
 * @description: SetIsReportcaller
 * @author: Jerry.Yang
 * @date: 2022-10-09 16:18:20
 * @return {*}
 */
func (l *LogrusLog) SetIsReportcaller() LoggerInterface {

	/**
	 * @step
	 * @获取设置
	 **/
	setIsReportcaller, err := l.Options.GetIsReportcaller()
	if err != nil {
		return l
	}
	l.IsReportCaller = setIsReportcaller
	return l
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-27 18:24:10
 * @return {*}
 */
func (l *LogrusLog) CheckParams() LoggerInterface {

	/**
	 * @step
	 * @ckeck logLevel
	 **/
	_, logLevelIsExist := LOGRUS_ALL_LEVEL[Level(l.LogLevel)]
	if !logLevelIsExist {
		l.LoggerErr = errors.New("CheckParams Err : logLevel is not exist")
	}

	if !OPTION_IS_SET_LEVEL_STATUS {
		l.LoggerErr = errors.New("CheckParams Err : logLevel is not exist")
	}

	/**
	 * @step
	 * @check formattter
	 **/
	if l.Formatter == nil {
		l.LoggerErr = errors.New("CheckParams Err : formattter is not exist")
	}

	/**
	 * @step
	 * @check output
	 **/
	if l.Output == nil {
		l.LoggerErr = errors.New("CheckParams Err : output is not exist")
	}
	return l
}
