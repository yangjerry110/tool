/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-23 15:57:51
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:41:53
 * @Description: perm
 */
package perm

import "github.com/yangjerry110/tool/perm"

type PermInterface interface {
	CreatePermInterface(permInterface perm.PermInterface) *Perm
	RsaCreatePerm(byteSize int32, permPath string) (bool, error)
	RsaDecrty(permPath string, inputStr string) (string, error)
	RsaEncrty(permPath string, inputStr string) (string, error)
}

type Perm struct {
	PermInterface perm.PermInterface
}

/**
 * @description: CreatePermInterface
 * @param {perm.PermInterface} permInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:42:02
 * @return {*}
 */
func CreatePermInterface(permInterface perm.PermInterface) *Perm {
	return &Perm{PermInterface: permInterface}
}
