/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:02:35
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:30:10
 * @Description: common
 */
package common

import "github.com/yangjerry110/tool/common"

type QiweiCommon struct{}

/**
 * @description: GetQiweiAccessToken
 * @param {string} appId
 * @param {string} cropId
 * @param {string} cropSecret
 * @param {string} redisConfPath
 * @param {string} redisConfName
 * @author: Jerry.Yang
 * @date: 2022-11-10 18:30:18
 * @return {*}
 */
func GetQiweiAccessToken(appId string, cropId string, cropSecret string, redisConfPath string, redisConfName string) (string, error) {
	return CreateCommonInterface(&common.QiweiCommon{AppId: appId, CropId: cropId, CropSecret: cropSecret, RedisConfPath: redisConfPath, RedisConfName: redisConfName}).CommonInterface.GetAccessToken()
}
