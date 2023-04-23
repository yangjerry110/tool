/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-20 11:30:16
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2023-02-16 17:07:17
 * @Description: encrty 加密
 */
package perm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

/**
 * @description: Encrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-20 11:40:30
 * @return {*}
 */
func (r *PermRsa) Encrty(permPath string, inputStr string) (string, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	if permPath == "" {
		return "", errors.New("permPath must be not empty!")
	}

	if inputStr == "" {
		return "", errors.New("inputStr must be not empty!")
	}

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	publicPermPath := fmt.Sprintf("%s/public.pem", permPath)
	_, err := os.Stat(publicPermPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New(fmt.Sprintf("permPath is err; %s is not exist!", publicPermPath))
		}
		return "", err
	}

	/**
	 * @step
	 * @开始通道
	 **/
	doRascrtyChan := make(chan string)

	/**
	 * @step
	 * @开启协程，获取加密之后的数据
	 **/
	go func() {
		outputStr, err := r.DoRsaEncrty(permPath, inputStr)
		if err != nil {
			doRascrtyChan <- ""
		} else {
			doRascrtyChan <- outputStr
		}
		close(doRascrtyChan)
	}()

	/**
	 * @step
	 * @判断结果
	 **/
	outputStr := <-doRascrtyChan
	if outputStr == "" {
		return "", errors.New("encrty fail!")
	}
	return outputStr, nil
}

/**
 * @description: doRsaEncrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-20 11:40:21
 * @return {*}
 */
func (r *PermRsa) DoRsaEncrty(permPath string, inputStr string) (string, error) {
	/**
	 * @step
	 * @打开公钥文件
	 **/
	file, err := os.Open(fmt.Sprintf("%s/public.pem", permPath))
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @读取文件内容
	 **/
	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()

	/**
	 * @step
	 * @定义参数
	 **/
	publicByte := []byte(buf)
	inputStrByte := []byte(inputStr)

	/**
	 * @step
	 * @pem编码
	 **/
	block, _ := pem.Decode(publicByte)

	/**
	 * @step
	 * @pem编码
	 **/
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @判断是否正常的publicKey
	 **/
	pubKey := pubInterface.(*rsa.PublicKey)

	/**
	 * @step
	 * @使用公钥加密
	 **/
	outputByte, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, inputStrByte)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @base64编码
	 **/
	outputStr := base64.StdEncoding.EncodeToString(outputByte)
	return outputStr, nil
}
