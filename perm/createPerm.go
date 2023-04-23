/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-20 11:17:29
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-09-26 16:28:03
 * @Description: createPerm
 */
package perm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/**
 * @description: CreatePermFunc
 * @param {int32} byteSize
 * @param {string} permPath
 * @author: Jerry.Yang
 * @date: 2022-09-20 11:24:42
 * @return {*}
 */
func (c *PermRsa) CreatePerm(byteSize int32, permPath string) (bool, error) {

	/**
	 * @step
	 * @判断bytesize
	 **/
	if byteSize == 0 {
		byteSize = 1024
	}

	/**
	 * @step
	 * @判断permPath
	 **/
	if permPath == "" {
		dir, err := os.Getwd()
		if err != nil {
			return false, err
		}
		permPath = dir
	}

	/**
	 * @step
	 * @生成私钥
	 **/
	privateKey, err := rsa.GenerateKey(rand.Reader, int(byteSize))
	if err != nil {
		return false, err
	}

	/**
	 * @step
	 * @通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	 **/
	derText := x509.MarshalPKCS1PrivateKey(privateKey)

	/**
	 * @step
	 * @要组织一个pem.Block(base64编码)
	 **/
	block := pem.Block{
		Type:  "rsa private key", // 这个地方写个字符串就行
		Bytes: derText,
	}

	/**
	 * @step
	 * @pem编码
	 **/
	file, err := os.Create(permPath + "/private.pem")
	if err != nil {
		return false, err
	}
	pem.Encode(file, &block)
	file.Close()

	/**
	 * @step
	 * @从公钥中提取私钥
	 * @使用x509标准序列化
	 **/
	publicKey := privateKey.PublicKey
	derstream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return false, err
	}

	/**
	 * @step
	 * @定义block
	 **/
	block = pem.Block{
		Type:  "rsa public key",
		Bytes: derstream,
	}

	/**
	 * @step
	 * @创建公钥文件
	 **/
	file, err = os.Create(permPath + "/public.pem")
	if err != nil {
		return false, err
	}
	pem.Encode(file, &block)
	file.Close()
	return true, nil
}
