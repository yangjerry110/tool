/*
* @Author: Jerry.Yang
* @Date: 2023-12-05 15:51:33
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-15 14:43:04
* @Description: Account service
*/
package test

import (
	"context"

	"github.com/yangjerry110/tool/test/testvo"
)

type AccountService interface {
	GetAccount(ctx context.Context, inputVo *testvo.GetAccountReq) (*testvo.GetAccountResp, error)
}

type Account struct{}

/**
* @description: GetAccount
* @param {context.Context} ctx
* @param {*GetAccountReq} inputVo
* @author: Jerry.Yang
* @date: 2023-12-05 15:51:33
* @return {*}
 */
func (a *Account) GetAccount(ctx context.Context, inputVo *testvo.GetAccountReq) (*testvo.GetAccountResp, error) {

	/**
	* @step
	* @result
	**/
	result := &testvo.GetAccountResp{}

	return result, nil
}
