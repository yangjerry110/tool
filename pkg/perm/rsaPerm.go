/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 16:20:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 18:42:24
 * @Description: ras
 */
package perm

import "github.com/yangjerry110/tool/perm"

type RasPerm struct{}

/**
 * @description: CreateRsaPerm
 * @param {int32} byteSize
 * @param {string} permPath
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:40:11
 * @return {*}
 */
func CreateRsaPerm(byteSize int32, permPath string) (bool, error) {
	return CreatePermInterface(&perm.PermRsa{}).PermInterface.CreatePerm(byteSize, permPath)
}

/**
 * @description: DecrtyRsa
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:41:07
 * @return {*}
 */
func DecrtyRsa(permPath string, inputStr string) (string, error) {
	return CreatePermInterface(&perm.PermRsa{}).PermInterface.Decrty(permPath, inputStr)
}

/**
 * @description: EncrtyRsa
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:41:56
 * @return {*}
 */
func EncrtyRsa(permPath string, inputStr string) (string, error) {
	return CreatePermInterface(&perm.PermRsa{}).PermInterface.Encrty(permPath, inputStr)
}
