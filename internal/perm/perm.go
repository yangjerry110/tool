/*
 * @Author: Jerry.Yang
 * @Date: 2023-12-11 11:35:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-12-11 11:47:05
 * @Description: perm
 */
package perm

type PermInterface interface {
	CreatePerm(byteSize int32, permPath string) (bool, error)
	Encrty(permPath string, inputStr string) (string, error)
	DoRsaEncrty(permPath string, inputStr string) (string, error)
	Decrty(permPath string, inputStr string) (string, error)
	DoRsaDecrty(permPath string, inputStr string) (string, error)
}
