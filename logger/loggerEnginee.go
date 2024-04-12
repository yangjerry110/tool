/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-11 15:09:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 16:13:54
 * @Description: logger enginee
 */
package logger

import (
	internalLogger "github.com/yangjerry110/tool/internal/logger"
	"github.com/yangjerry110/tool/internal/logger/logrus"
)

var loggerEnginee internalLogger.LoggerInterface

// SetLoggerEnginee
//
// SetLoggerEnginee
// Date 2024-04-11 15:08:01
// Author Jerry.Yang
func SetLoggerEnginee(logger internalLogger.LoggerInterface) internalLogger.LoggerInterface {
	loggerEnginee = logger.Init()
	return loggerEnginee
}

// log
//
// log
// Date 2024-04-11 15:08:01
// Author Jerry.Yang
func log() internalLogger.LoggerInterface {

	// if loggerEnginee == nil
	if loggerEnginee == nil {

		// set defaut logger enginee
		loggerEnginee = SetLoggerEnginee(&logrus.Logrus{})
	}
	return loggerEnginee
}
