/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 14:29:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-22 16:24:53
 * @Description: perm
 */
package perm

import (
	"github.com/yangjerry110/tool/internal/perm"
	rsaperm "github.com/yangjerry110/tool/internal/perm/rsaPerm"
)

/**
 * @description: CreateRsaPerm
 * @author: Jerry.Yang
 * @date: 2023-12-22 16:24:45
 * @return {*}
 */
func CreateRsaPerm() perm.PermInterface {
	return perm.CreatePerm(&rsaperm.RsaPerm{})
}
