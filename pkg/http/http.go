/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 14:52:39
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:41:16
 * @Description: http
 */
package http

import (
	"io"

	"github.com/yangjerry110/tool/http"
)

type HttpInterface interface {
	CreateHttpInterface(httpInterface http.HttpInterface) http.HttpInterface
	HttpRequest(method string, url string, body io.Reader, output interface{}, options ...http.HttpOptionFunc) error
}

type Http struct {
	HttpInterface http.HttpInterface
}

/**
 * @description: CreateHttpInterface
 * @param {http.HttpInterface} httpInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:40:57
 * @return {*}
 */
func CreateHttpInterface(httpInterface http.HttpInterface) *Http {
	return &Http{HttpInterface: httpInterface}
}
