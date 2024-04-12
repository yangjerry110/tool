/*
 * @Author: Jerry.Yang
 * @Date: 2024-04-11 15:28:12
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2024-04-12 15:13:09
 * @Description: config
 */
package logger

import (
	"github.com/yangjerry110/tool/conf"
	"github.com/yangjerry110/tool/internal/logger/logrus"
)

// Create logrus option conf
//
// CreateLogrusOptionConf
// Date 2024-04-11 15:28:53
// Author Jerry.Yang
func CreateLogrusOptionConf() error {

	// set config
	if err := conf.CreateConf(&logrus.LogrusOption{}).SetConfig(); err != nil {
		return err
	}

	// set logrus conf
	// set level
	if err := SetLevel(logrus.LogrusOptionConf.Level); err != nil {
		return err
	}

	// set formatter
	if err := SetFormatter(logrus.LogrusOptionConf.Formatter); err != nil {
		return err
	}

	// set ReportCaller
	if err := SetReportCaller(logrus.LogrusOptionConf.ReportCaller); err != nil {
		return err
	}

	// SetEnableHTMLEscape
	if err := SetEnableHTMLEscape(logrus.LogrusOptionConf.EnableHTMLEscape); err != nil {
		return err
	}
	return nil
}
