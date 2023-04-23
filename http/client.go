/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 15:47:47
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 16:32:49
 * @Description: client 请求连接
 */
package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * @description: HttpRequest
 * @author: Jerry.Yang
 * @date: 2022-09-21 16:02:35
 * @return {*}
 */
func (h *HttpClient) Request() error {

	/**
	 * @step
	 * @创建httpOptions，并且设置
	 **/
	httpOptionsObj := HttpOptions{}
	httpOptions := httpOptionsObj.SetOptions(h.Options)

	/**
	 * @step
	 * @创建client
	 **/
	client := &http.Client{Timeout: time.Duration(httpOptionsObj.GetOutTime(httpOptions)) * time.Second}

	/**
	 * @step
	 * @request
	 **/
	req, err := http.NewRequest(h.Method, h.Url, h.Body)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @获取headers
	 **/
	headers := httpOptionsObj.GetHeaders(httpOptions)

	/**
	 * @step
	 * @添加headers
	 **/
	if len(headers) > 0 {
		for headerKey, headerVal := range headers {
			req.Header.Set(headerKey, headerVal)
		}
	}

	/**
	 * @step
	 * @执行请求
	 **/
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @获取返回的数据
	 **/
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	/**
	 * @step
	 * @渲染返回值
	 **/
	err = json.Unmarshal(data, &h.Output)
	if err != nil {
		return err
	}
	return nil
}
