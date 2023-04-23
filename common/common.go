/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 15:09:17
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:27:42
 * @Description: common
 */
package common

type CommonInterface interface {
	GetAccessToken() (string, error)
}

type QiweiCommon struct {
	AppId         string
	CropId        string
	CropSecret    string
	RedisConfPath string
	RedisConfName string
}

type AliCommon struct {
	AppId         string
	AppKey        string
	AppSecret     string
	RedisConfPath string
	RedisConfName string
}
