/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-21 18:50:50
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:29:49
 * @Description: qiwei upload
 */
package qiwei

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	mytoolCommon "github.com/yangjerry110/tool/common"
	mytoolHttp "github.com/yangjerry110/tool/http"
)

/**
 * @description: UploadMedia
 * @author: Jerry.Yang
 * @date: 2022-09-21 18:56:23
 * @return {*}
 */
func (q *QiweiNotice) Upload() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := q.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取accessToken
	 **/
	qiweiCommon := mytoolCommon.QiweiCommon{AppId: q.AppId, CropId: q.CropId, CropSecret: q.CropSecret, RedisConfPath: q.RedisConfPath, RedisConfName: q.RedisConfName}
	accessToken, err := qiweiCommon.GetAccessToken()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @解密数据
	 **/
	decodMediaData, err := base64.StdEncoding.DecodeString(q.MediaData)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @构建body参数
	 **/
	body := &bytes.Buffer{}

	/**
	 * @step
	 * @实例化multipart
	 **/
	writer := multipart.NewWriter(body)

	/**
	 * @step
	 * @创建name
	 **/
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), uuid.New().String())

	/**
	 * @step
	 * @创建path
	 **/
	filePath := fmt.Sprintf("%s.%s", q.QiweiFilePath+fileName, q.MediaType)

	/**
	 * @step
	 * @创建multipart 文件字段
	 * @创建一个文件的write对象
	 **/
	formFileWriter, err := writer.CreateFormFile(fileName, filepath.Base(filePath))
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @写入传入的文件数据
	 **/
	formFileWriter.Write(decodMediaData)

	/**
	 * @step
	 * @close
	 **/
	err = writer.Close()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @构建请求url
	 **/
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", accessToken, q.MediaType)

	/**
	 * @step
	 * @构建返回体结构体
	 **/
	type UploadMediaOutput struct {
		ErrCode   int32  `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
		Type      string `json:"type"`
		MediaId   string `json:"media_id"`
		CreatedAt string `json:"created_at"`
	}

	/**
	 * @step
	 * @发送请求
	 **/
	resp := &UploadMediaOutput{}
	/**
	 * @step
	 * @设置httpOption
	 **/
	mytoolHttpOptions := mytoolHttp.HttpOptions{}
	httpOptions := []mytoolHttp.HttpOptionFunc{
		mytoolHttpOptions.SetHeaders(map[string]string{
			"Content-Type": "multipart/form-data",
		}),
	}

	/**
	 * @step
	 * @组装请求参数
	 **/
	httpClient := mytoolHttp.HttpClient{
		Method:  "POST",
		Url:     url,
		Body:    body,
		Options: httpOptions,
		Output:  resp,
	}
	httpClient.Request()

	/**
	 * @step
	 * @判断结果
	 **/
	if resp.ErrCode != 0 {
		return "", errors.New(resp.ErrMsg)
	}

	return resp.MediaId, nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-26 17:42:45
 * @return {*}
 */
func (q *QiweiNotice) CheckParamsUpload() error {
	/**
	 * @step
	 * @判断appId
	**/
	if q.AppId == "" {
		return errors.New("QiweiUploadMedia Err : appId is empty!")
	}

	/**
	 * @step
	 * @判断qiweiFilePath
	 **/
	if q.QiweiFilePath == "" {
		return errors.New("QiweiUploadMedia Err : qiweiFilePath is empty!")
	}
	return nil
}
