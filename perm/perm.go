/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 14:29:42
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 10:33:49
 * @Description: perm
 */
package perm

import (
	"github.com/yangjerry110/tool/perm/internal/perm"
	rsaperm "github.com/yangjerry110/tool/perm/internal/perm/rsaPerm"
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
