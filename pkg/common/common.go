/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 11:00:49
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:38:11
 * @Description: common
 */
package common

import "github.com/yangjerry110/tool/common"

type CommonInterface interface {
	CreateCommonInterface(commonInterface common.CommonInterface) *Common
	GetQiweiAccessToken(appId string, cropId string, cropSecret string) (string, error)
}

type Common struct {
	CommonInterface common.CommonInterface
}

/**
 * @description: CreateCommonInterface
 * @param {common.CommonInterface} commonInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:38:43
 * @return {*}
 */
func CreateCommonInterface(commonInterface common.CommonInterface) *Common {
	return &Common{CommonInterface: commonInterface}
}
