/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-10 15:41:46
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:27:52
 * @Description: ali
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
 * @description: GetAccessToken
 * @author: Jerry.Yang
 * @date: 2022-10-10 15:46:33
 * @return {*}
 */
func (a *AliCommon) GetAccessToken() (string, error) {

	/**
	 * @step
	 * @定义key
	 **/
	key := fmt.Sprintf("%s_dingding_access_token", a.AppId)

	/**
	 * @step
	 * @获取缓存里面的accessToken
	 **/
	cacheAccessTokenInterface, err := cache.Client(a.RedisConfPath, a.RedisConfName).Get(key).Result()
	cacheAccessTokenStr := cacheAccessTokenInterface.(string)
	if err == nil && cacheAccessTokenStr != "" {
		return cacheAccessTokenStr, nil
	}

	/**
	 * @step
	 * @获取accessToken的url
	 **/
	url := fmt.Sprintf("https://oapi.dingtalk.com/gettoken?appkey=%s&appsecret=%s", a.AppKey, a.AppSecret)

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
	err = cache.Client(a.RedisConfPath, a.RedisConfName).Set(key, resp.AccessToken, 7100*time.Second).GetErr()
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}
