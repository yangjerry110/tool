/*
 * @Author: Jerry.Yang
 * @Date: 2023-05-18 15:36:41
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-05-18 15:37:48
 * @Description: newapi
 */
package commands

import "github.com/golib/cli"

type NewApiCommands interface {
	NewApi(ctx *cli.Context) error
}

type NewApi struct{}

/**
 * @description: NewApi
 * @param {*cli.Context} ctx
 * @author: Jerry.Yang
 * @date: 2023-05-18 15:37:56
 * @return {*}
 */
func (n *NewApi) NewApi(ctx *cli.Context) error {
	return nil
}
