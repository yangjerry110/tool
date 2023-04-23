package logger

import (
	"errors"
	"io"

	"github.com/sirupsen/logrus"
)

type LogrusOption struct {
	LoggerOption
}

/**
 * @description: SetLevel
 * @param {Level} logLevel
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:09:46
 * @return {*}
 */
func (l *LogrusOption) SetLevel(logLevel Level) LoggerOptionFunc {
	OPTION_IS_SET_LEVEL_STATUS = true
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_LOGLEVEL] = LoggerOptionVal{logLevel}
		return nil
	}
}

/**
 * @description: SetLogrusWithFields
 * @param {logrus.Fields} fields
 * @author: Jerry.Yang
 * @date: 2022-09-29 17:11:26
 * @return {*}
 */
func (l *LogrusOption) SetWithFields(fields map[string]interface{}) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_WITH_FIELDS] = LoggerOptionVal{fields}
		return nil
	}
}

/**
 * @description: SetIsSetLevel
 * @param {bool} status
 * @author: Jerry.Yang
 * @date: 2022-09-29 16:34:28
 * @return {*}
 */
func (l *LogrusLog) SetIsSetLevel(status bool) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_IS_SET_LEVEL] = LoggerOptionVal{status}
		return nil
	}
}

/**
 * @description: SetIsReportcaller
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:13:02
 * @return {*}
 */
func (l *LogrusOption) SetIsReportcaller(isOpen bool) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_ISREPORTCALLER] = LoggerOptionVal{isOpen}
		return nil
	}
}

/**
 * @description: SetOutput
 * @param {io.Writer} output
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:14:15
 * @return {*}
 */
func (l *LogrusOption) SetOutput(output io.Writer) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_OUTPUT] = LoggerOptionVal{output}
		return nil
	}
}

/**
 * @description: SetFormatter
 * @param {string} formatter
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:15:35
 * @return {*}
 */
func (l *LogrusOption) SetFormatter(formatter string) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_FORMATTER_TYPE] = LoggerOptionVal{formatter}
		return nil
	}
}

/**
 * @description: SetFormatterDisableTime
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:35:42
 * @return {*}
 */
func (l *LogrusOption) SetFormatterDisableTime(isOpen bool) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_FORMATTER_DISABLETIME] = LoggerOptionVal{isOpen}
		return nil
	}
}

/**
 * @description: SetFormatterDisableHtmlEscap
 * @param {bool} isOpen
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:49:08
 * @return {*}
 */
func (l *LogrusOption) SetFormatterDisableHtmlEscap(isOpen bool) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_FORMATTER_DISABLEHTMLESCAP] = LoggerOptionVal{isOpen}
		return nil
	}
}

/**
 * @description: SetCallDept
 * @param {int32} dept
 * @author: Jerry.Yang
 * @date: 2022-10-09 18:36:41
 * @return {*}
 */
func (l *LoggerOption) SetCallDept(dept int) LoggerOptionFunc {
	return func(LoggerOptionFunc map[string]LoggerOptionVal) error {
		LoggerOptionFunc[OPTION_CALLER_DEPT] = LoggerOptionVal{dept}
		return nil
	}
}

/**
 * @description: GetLevel
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:43:48
 * @return {*}
 */
func (l *LogrusOption) GetLevel() (Level, error) {

	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的level
	 **/
	setLevelInterface, levelIsExist := options[OPTION_LOGLEVEL]
	if !levelIsExist {
		return 0, errors.New("GetLevel Err : level is not set")
	}

	/**
	 * @step
	 * @获取设置的level
	 **/
	setLevel := setLevelInterface.Value.(Level)

	/**
	 * @step
	 * @判断是否存在
	 **/
	_, setLevelIsExist := LOGRUS_ALL_LEVEL[setLevel]
	if !setLevelIsExist {
		return 0, errors.New("GetLevel Err : level is not exist")
	}
	return setLevel, nil
}

/**
 * @description: GetLogrusWithFields
 * @author: Jerry.Yang
 * @date: 2022-09-29 17:15:21
 * @return {*}
 */
func (l *LogrusOption) GetWithFields() (map[string]interface{}, error) {
	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置
	 **/
	setLogrusWithFields, logrusWithFieldsIsExist := options[OPTION_WITH_FIELDS]
	if !logrusWithFieldsIsExist {
		return nil, errors.New("GetLogrusWithFields Err : with fields is not set")
	}
	return setLogrusWithFields.Value.(map[string]interface{}), nil
}

/**
 * @description: GetLogrusLevel
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:10:47
 * @return {*}
 */
func (l *LogrusOption) GetLogrusLevel() (logrus.Level, error) {

	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的level
	 **/
	setLevelInterface, levelIsExist := options[OPTION_LOGLEVEL]
	if !levelIsExist {
		return 0, errors.New("GetLevel Err : level is not set")
	}

	/**
	 * @step
	 * @获取设置的level
	 **/
	setLevel := setLevelInterface.Value.(Level)

	/**
	 * @step
	 * @判断是否存在
	 **/
	setLogrusLevel, setLevelIsExist := LOGRUS_ALL_LEVEL[setLevel]
	if !setLevelIsExist {
		return 0, errors.New("GetLevel Err : level is not exist")
	}
	return setLogrusLevel, nil
}

/**
 * @description: GetIsReportcaller
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:55:41
 * @return {*}
 */
func (l *LogrusOption) GetIsReportcaller() (bool, error) {
	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的IsReportcaller
	 **/
	isReportcaller, isReportcallerIsExist := options[OPTION_ISREPORTCALLER]
	if !isReportcallerIsExist {
		return false, errors.New("GetIsReportcaller Err : isReportcaller is not set")
	}
	return isReportcaller.Value.(bool), nil
}

/**
 * @description: GetOutput
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:59:06
 * @return {*}
 */
func (l *LogrusOption) GetOutput() (io.Writer, error) {
	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的output
	 **/
	output, isOutputIsExist := options[OPTION_OUTPUT]
	if !isOutputIsExist {
		return nil, errors.New("GetOutput Err : output is not set")
	}
	return output.Value.(io.Writer), nil
}

/**
 * @description: GetFormatter
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:19:46
 * @return {*}
 */
func (l *LogrusOption) GetFormatterType() (string, error) {
	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的formatter
	 **/
	setFormatterInterface, setFormatterIsExist := options[OPTION_FORMATTER_TYPE]
	if !setFormatterIsExist {
		return "", errors.New("GetFormatter Err : formatter is not set")
	}

	/**
	 * @step
	 * @获取直接的数据
	 **/
	setFormatter := setFormatterInterface.Value.(string)

	/**
	 * @step
	 * @判断是否存在
	 **/
	formatter, formatterIsExist := LOGRUS_FORMATTER_ALL[setFormatter]
	if !formatterIsExist {
		return "", errors.New("GetFormatter Err : formatter is not exist")
	}
	return formatter, nil
}

/**
 * @description: GetDisableTime
 * @author: Jerry.Yang
 * @date: 2022-09-27 16:45:33
 * @return {*}
 */
func (l *LogrusOption) GetDisableTime() (bool, error) {
	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的IsReportcaller
	 **/
	disableTimeInterface, disableTimeIsExist := options[OPTION_FORMATTER_DISABLETIME]
	if !disableTimeIsExist {
		return false, errors.New("GetDisableTime Err : disableTime is not set")
	}
	return disableTimeInterface.Value.(bool), nil
}

/**
 * @description: GetDisableHtmlEscape
 * @author: Jerry.Yang
 * @date: 2022-09-27 17:15:03
 * @return {*}
 */
func (l *LogrusOption) GetDisableHtmlEscape() (bool, error) {
	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的IsReportcaller
	 **/
	disableHTMLEscapeInterface, disableHTMLEscapeIsExist := options[OPTION_FORMATTER_DISABLEHTMLESCAP]
	if !disableHTMLEscapeIsExist {
		return false, errors.New("GetDisableHtmlEscape Err : disableHTMLEscape is not set")
	}
	return disableHTMLEscapeInterface.Value.(bool), nil
}

/**
 * @description: GetCallerDept
 * @author: Jerry.Yang
 * @date: 2022-10-09 18:39:24
 * @return {*}
 */
func (l *LogrusOption) GetCallerDept() (int, error) {
	/**
	 * @step
	 * @获取所有的options
	 **/
	options := l.Options

	/**
	 * @step
	 * @获取设置的IsReportcaller
	 **/
	setVal, setValExist := options[OPTION_CALLER_DEPT]
	if !setValExist {
		return 0, errors.New("GetCallerDept Err : callerDept is not set")
	}
	return setVal.Value.(int), nil
}

/**
 * @description: SetLogrusLogOptions
 * @param {[]LoggerOptionFunc} optionFuncs
 * @author: Jerry.Yang
 * @date: 2022-09-27 15:36:38
 * @return {*}
 */
func (l *LogrusOption) SetOptions(optionFuncs []LoggerOptionFunc) LoggerOptionInterface {
	/**
	 * @step
	 * @定义
	 **/
	setOptions := map[string]LoggerOptionVal{}

	/**
	 * @step
	 * @获取set的functions
	 **/
	for _, option := range optionFuncs {
		if option != nil {
			err := option(setOptions)
			if err != nil {
				return nil
			}
		}
	}
	l.Options = setOptions
	return l
}
