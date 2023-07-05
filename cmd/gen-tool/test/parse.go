/*
* @Author: Jerry.Yang
* @Date: 2023-05-26 14:11:00
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-26 14:38:56
* @Description: Demo service
*/
package test

import (
	"context"

	"github.com/yangjerry110/tool/cmd/gen-tool/test/protobuf"
)

type DemoService interface {
	AddDemo(ctx context.Context, inputVo *protobuf.AddDemoReq) (*protobuf.Empty, error)
	DeleteDemo(ctx context.Context, inputVo *protobuf.DeleteDemoReq) (*protobuf.Empty, error)
	UpdateDemo(ctx context.Context, inputVo *protobuf.UpdateDemoReq) (*protobuf.Empty, error)
	GetDemo(ctx context.Context, inputVo *protobuf.GetDemoReq) (*protobuf.GetDemoResp, error)
	TestFour(ctx context.Context, inputVo *protobuf.GetTestFourReq) (*protobuf.GetTestFourResp, error)
}

type Demo struct{}

/**
 * @description: AddDemo
 * @param {context.Context} ctx
 * @param {*protobuf.AddDemoReq} inputVo
 * @author: Jerry.Yang
 * @date: 2023-05-26 14:11:00
 * @return {*}
 */
func (d *Demo) AddDemo(ctx context.Context, inputVo *protobuf.AddDemoReq) (*protobuf.Empty, error) {

	/**
	 * @step
	 * @result
	 **/
	result := &protobuf.Empty{}
	return result, nil
}

/**
 * @description: DeleteDemo
 * @param {context.Context} ctx
 * @param {*protobuf.DeleteDemoReq} inputVo
 * @author: Jerry.Yang
 * @date: 2023-05-26 14:11:00
 * @return {*}
 */
func (d *Demo) DeleteDemo(ctx context.Context, inputVo *protobuf.DeleteDemoReq) (*protobuf.Empty, error) {

	/**
	 * @step
	 * @result
	 **/
	result := &protobuf.Empty{}
	return result, nil
}

/**
 * @description: UpdateDemo
 * @param {context.Context} ctx
 * @param {*protobuf.UpdateDemoReq} inputVo
 * @author: Jerry.Yang
 * @date: 2023-05-26 14:11:00
 * @return {*}
 */
func (d *Demo) UpdateDemo(ctx context.Context, inputVo *protobuf.UpdateDemoReq) (*protobuf.Empty, error) {

	/**
	 * @step
	 * @result
	 **/
	result := &protobuf.Empty{}
	return result, nil
}

/**
 * @description: GetDemo
 * @param {context.Context} ctx
 * @param {*protobuf.GetDemoReq} inputVo
 * @author: Jerry.Yang
 * @date: 2023-05-26 14:11:00
 * @return {*}
 */
func (d *Demo) GetDemo(ctx context.Context, inputVo *protobuf.GetDemoReq) (*protobuf.GetDemoResp, error) {

	/**
	 * @step
	 * @result
	 **/
	result := &protobuf.GetDemoResp{}
	return result, nil
}

/**
 * @description: TestFour
 * @param {context.Context} ctx
 * @param {*protobuf.GetTestFourReq} inputVo
 * @author: Jerry.Yang
 * @date: 2023-05-26 14:11:00
 * @return {*}
 */
func (d *Demo) TestFour(ctx context.Context, inputVo *protobuf.GetTestFourReq) (*protobuf.GetTestFourResp, error) {

	/**
	 * @step
	 * @result
	 **/
	result := &protobuf.GetTestFourResp{}
	return result, nil
}
