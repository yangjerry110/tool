/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 14:29:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 14:31:00
 * @Description: perm
 */
package perm

import (
	"github.com/yangjerry110/tool/internal/perm"
	rsaperm "github.com/yangjerry110/tool/internal/perm/rsaPerm"
)

/**
 * @description: default perm
 * @author: Jerry.Yang
 * @date: 2023-12-11 14:29:49
 * @return {*}
 */
var DefaultPerm = &rsaperm.RsaPerm{}

/**
 * @description: CreatePerm
 * @param {...perm.PermInterface} PermInterfaces
 * @author: Jerry.Yang
 * @date: 2023-12-11 14:30:34
 * @return {*}
 */
func CreatePerm(PermInterfaces ...perm.PermInterface) perm.PermInterface {
	if len(PermInterfaces) == 0 {
		return &rsaperm.RsaPerm{}
	}
	return PermInterfaces[0]
}
