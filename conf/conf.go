/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 15:15:25
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-15 14:57:48
 * @Description: conf
 */
package conf

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type ConfInterface interface {
	Init(conf interface{}) ConfInterface
	GetNewConf() error
	GetConf(conf interface{}) error
	GetParseConf() interface{}
	GetErr() error
}

/**
 * @description: Conf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:42
 * @return {*}
 */
type Conf struct {
	FilePath       string
	FileName       string
	FileType       string
	Intervals      time.Duration
	Data           interface{}
	LastModityTime time.Time
	Error          error
}

/**
 * @description: Init
 * @param {interface{}} conf
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:28:28
 * @return {*}
 */
func (c *Conf) Init(conf interface{}) ConfInterface {
	return c
}

/**
 * @description: GetHotUpdateConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:13:19
 * @return {*}
 */
func (c *Conf) GetNewConf() error {

	/**
	 * @step
	 * @获取配置
	 **/
	err := c.GetConf(c.Data)
	if err != nil {
		c.Error = err
		return err
	}

	/**
	 * @step
	 * @添加观察者
	 **/
	//c.AddNotifyer(configStore)

	/**
	 * @step
	 * @添加一个协程，热更新配置文件
	 **/
	go c.RelodConf()
	return nil
}

/**
 * @description: GetConfAtomic
 * @author: Jerry.Yang
 * @date: 2022-11-11 11:26:23
 * @return {*}
 */
func (c *Conf) GetParseConf() interface{} {
	return c.Data
}

/**
 * @description: GetConf
 * @author: Jerry.Yang
 * @date: 2022-11-10 17:15:51
 * @return {*}
 */
func (c *Conf) GetConf(conf interface{}) error {

	/**
	 * @step
	 * @根据type，获取文件内容
	 **/
	switch c.FileType {
	case "yaml":
		yamlConf := YamlConf{FilePath: c.FilePath, FileName: c.FileName, Conf: c.Data}
		yamlConfErr := yamlConf.GetConf(c.Data)
		if yamlConfErr != nil {
			return yamlConfErr
		}
	default:
		return errors.New("GetConf Err : no match fileType")
	}
	return nil
}

/**
 * @description: GetConfTimer
 * @author: Jerry.Yang
 * @date: 2022-11-10 17:15:57
 * @return {*}
 */
func (c *Conf) RelodConf() error {

	/**
	 * @step
	 * @创建一个定时器
	 **/
	timeTickers := time.NewTicker(time.Second * c.Intervals)

	/**
	 * @step
	 * @定时读取配置
	 **/
	for range timeTickers.C {

		/**
		 * @step
		 * @打开文件
		 **/
		fileObj, err := os.Open(fmt.Sprintf("%s/%s", c.FilePath, c.FileName))
		if err != nil {
			//timeTickers.Stop()
			c.Error = err
			return err
		}
		defer fileObj.Close()

		/**
		 * @step
		 * @获取文件详情
		 **/
		fileInfo, err := fileObj.Stat()
		if err != nil {
			//timeTickers.Stop()
			c.Error = err
			return err
		}

		/**
		 * @step
		 * @获取当前文件的修改时间
		 * @当发生更新的时候，执行
		 **/
		if fileInfo.ModTime().Unix() > c.LastModityTime.Unix() {
			/**
			 * @step
			 * @当发生更新的时候
			 * @赋值时间
			 **/
			c.LastModityTime = fileInfo.ModTime()

			/**
			 * @step
			 * @当更新的时候，重新解析
			 **/
			err = c.GetConf(c.Data)
			if err != nil {
				c.Error = err
				return err
			}

			/**
			 * @step
			 * @通知订阅者
			 **/
			// for _, notifyer := range c.NotifyerList {
			// 	notifyer.Callback(c)
			// }
		}

	}
	//timeTickers.Stop()
	return nil
}

/**
 * @description: GetErr
 * @author: Jerry.Yang
 * @date: 2022-11-15 14:58:20
 * @return {*}
 */
func (c *Conf) GetErr() error {
	return c.Error
}
