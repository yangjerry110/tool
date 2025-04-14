/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-20 19:28:19
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2025-03-11 10:37:10
 * @Description: decrty 解密
 */
package rsaperm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"fmt"
	"os"

	"github.com/yangjerry110/tool/toolerrors"
)

/**
 * @description: Decrty
 * @param {string} permPath
 * @param {string} inputStr
 * @author: Jerry.Yang
 * @date: 2022-09-20 19:32:21
 * @return {*}
 */
func (r *RsaPerm) Decrty(permPath string, inputStr string) (string, error) {

	/**
	 * @step
	 * @判断参数
	 **/
	if permPath == "" {
		return "", toolerrors.New("perm Err : rsa perm no permPath")
	}

	if inputStr == "" {
		return "", toolerrors.New("perm Err : rsa perm no inputFile")
	}

	/**
	 * @step
	 * @判断文件是否存在
	 **/
	privatePath := fmt.Sprintf("%s/private.pem", permPath)
	_, err := os.Stat(privatePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", toolerrors.WithFields("privatePath", privatePath).New("perm Err : rsa perm no privatePath")
		}
		return "", err
	}

	/**
	 * @step
	 * @创建通道
	 **/
	doRasDecrtyChan := make(chan string)

	/**
	 * @step
	 * @协程处理
	 **/
	go func() {
		outputStr, err := r.DoRsaDecrty(permPath, inputStr)
		if err != nil {
			doRasDecrtyChan <- ""
		} else {
			doRasDecrtyChan <- outputStr
		}
		close(doRasDecrtyChan)
	}()

	/**
	 * @step
	 * @获取通道数据，并且判断
	 **/
	outputStr := <-doRasDecrtyChan
	if outputStr == "" {
		return "", toolerrors.New("perm Err : rsa perm decrtyFail")
	}
	return outputStr, nil
}

/**
 * @description: DoRsaDecrty
 * @param {string} inputStr
 * @param {string} permPath
 * @author: Jerry.Yang
 * @date: 2022-09-20 19:32:12
 * @return {*}
 */
func (r *RsaPerm) DoRsaDecrty(permPath string, inputStr string) (string, error) {
	/**
	 * @step
	 * @进行base64 decode
	 **/
	inputStrByte, err := base64.StdEncoding.DecodeString(inputStr)
	if err != nil {
		return inputStr, err
	}

	/**
	 * @step
	 * @打开公钥文件
	 **/
	file, err := os.Open(fmt.Sprintf("%s/private.perm", permPath))
	if err != nil {
		return inputStr, err
	}

	/**
	 * @step
	 * @读取文件内容
	 **/
	fileInfo, err := file.Stat()
	if err != nil {
		return inputStr, err
	}

	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()

	/**
	 * @step
	 * @pem编码
	 **/
	block, _ := pem.Decode(buf)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @使用私钥解密
	 **/
	outputByte, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, inputStrByte)
	if err != nil {
		return inputStr, err
	}
	return string(outputByte), nil
}
