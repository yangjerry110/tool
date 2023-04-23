/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:05:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:41:25
 * @Description: http
 */
package http

import (
	"io"

	"github.com/yangjerry110/tool/http"
)

type HttpClient struct{}

/**
 * @description: HttpRequest
 * @param {string} method
 * @param {string} url
 * @param {io.Reader} body
 * @param {interface{}} output
 * @param {...http.HttpOptionFunc} options
 * @author: Jerry.Yang
 * @date: 2022-09-23 15:01:01
 * @return {*}
 */
func HttpRequest(method string, url string, body io.Reader, output interface{}, options ...http.HttpOptionFunc) error {
	return CreateHttpInterface(&http.HttpClient{Method: method, Url: url, Body: body, Output: &output, Options: options}).HttpInterface.Request()
}
