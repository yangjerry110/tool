/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 17:42:36
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:28:19
 * @Description: qiwei common
 */
package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/yangjerry110/tool/http"
	"github.com/yangjerry110/tool/pkg/cache"
)

/**
 * @description: GetQiweiAccessToken
 * @author: Jerry.Yang
 * @date: 2022-09-22 16:02:25
 * @return {*}
 */
func (q *QiweiCommon) GetAccessToken() (string, error) {

	/**
	 * @step
	 * @定义key
	 **/
	key := fmt.Sprintf("%s_qiwei_access_token", q.AppId)

	/**
	 * @step
	 * @获取缓存里面的accessToken
	 **/
	cacheAccessTokenInterface, err := cache.Client(q.RedisConfPath, q.RedisConfPath).Get(key).Result()
	cacheAccessTokenStr := cacheAccessTokenInterface.(string)
	if err == nil && cacheAccessTokenStr != "" {
		return cacheAccessTokenStr, nil
	}

	/**
	 * @step
	 * @获取accessToken的url
	 **/
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", q.CropId, q.CropSecret)

	/**
	 * @step
	 * @定义出参
	 **/
	type GetAccessTokenOutput struct {
		Errcode     int32  `json:"errcode"`
		Errmsg      string `json:"errmsg"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	/**
	 * @step
	 * @获取accessToken
	 **/
	resp := &GetAccessTokenOutput{}
	httpClient := http.HttpClient{
		Method: "GET",
		Url:    url,
		Output: resp,
	}
	httpClient.Request()

	/**
	 * @step
	 * @判断accessToken和错误
	 **/
	if resp.Errcode != 0 {
		return "", errors.New(resp.Errmsg)
	}

	/**
	 * @step
	 * @设置accessToken缓存
	 **/
	err = cache.Client(q.RedisConfPath, q.RedisConfName).Set(key, resp.AccessToken, 7100*time.Second).GetErr()
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}
