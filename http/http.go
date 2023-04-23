/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 15:20:15
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 16:32:45
 * @Description: http
 */
package http

import "io"

type HttpInterface interface {
	Request() error
}

type HttpClient struct {
	Method  string           // 请求方式
	Url     string           // 请求url
	Body    io.Reader        // 请求体
	Options []HttpOptionFunc // 参数(超时时间等等)
	Output  interface{}      // 返回数据
}

type HttpOptions struct{}

/**
 * @step
 * @定义optionVal
 **/
type HttpOption struct {
	Value interface{}
}

/**
 * @step
 * @定义options
 **/
type HttpOptionFunc func(map[string]HttpOption) error
