/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-22 14:05:52
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-10-26 17:36:20
 * @Description: ali oss
 */
package ali

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	aliyunOss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	mytoolHttp "github.com/yangjerry110/tool/http"
)

/**
 * @description: AliOssUpload Upload
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:08:04
 * @return {*}
 */
func (a *AliOssUpload) Upload() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := a.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建client
	 **/
	client, err := a.CreateClient()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建bucket
	 **/
	bucket, err := a.CreateBucket(client)

	/**
	 * @step
	 * @objectName
	 **/
	objectName := fmt.Sprintf("%d_%s.%s", time.Now().Unix(), uuid.New().String(), a.FileType)

	/**
	 * @step
	 * @option
	 **/
	options := []aliyunOss.Option{
		aliyunOss.ObjectACL(aliyunOss.ACLPublicRead),
	}

	/**
	 * @step
	 * @解密base64数据
	 **/
	decodFileData, err := base64.StdEncoding.DecodeString(a.FileData)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @上传bytes数据到oss
	 **/
	err = bucket.PutObject(objectName, bytes.NewReader(decodFileData), options...)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取实际访问的oss地址
	 **/
	ossUrl := fmt.Sprintf("%s/%s", a.DownloadDoamin, objectName)
	return ossUrl, nil
}

/**
 * @description: AliOssUploadFormLocalFile Upload
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:34:56
 * @return {*}
 */
func (a AliOssUploadFromLocalFile) Upload() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := a.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建client
	 **/
	client, err := a.CreateClient()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建bucket
	 **/
	bucket, err := a.CreateBucket(client)

	/**
	 * @step
	 * @objectName
	 **/
	objectName := fmt.Sprintf("%d_%s.%s", time.Now().Unix(), uuid.New().String(), a.FileType)

	/**
	 * @step
	 * @option
	 **/
	options := []aliyunOss.Option{
		aliyunOss.ObjectACL(aliyunOss.ACLPublicRead),
	}

	/**
	 * @step
	 * @获取文件数据
	 **/
	decodFileData, err := a.GetFileData()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @上传bytes数据到oss
	 **/
	err = bucket.PutObject(objectName, bytes.NewReader(decodFileData), options...)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取实际访问的oss地址
	 **/
	ossUrl := fmt.Sprintf("%s/%s", a.DownloadDoamin, objectName)
	return ossUrl, nil
}

/**
 * @description: AliOssUpLoadFormFileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:38:03
 * @return {*}
 */
func (a *AliOssUpLoadFromFileUrl) Upload() (string, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := a.CheckParams()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建client
	 **/
	client, err := a.CreateClient()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @创建bucket
	 **/
	bucket, err := a.CreateBucket(client)

	/**
	 * @step
	 * @objectName
	 **/
	objectName := fmt.Sprintf("%d_%s.%s", time.Now().Unix(), uuid.New().String(), a.FileType)

	/**
	 * @step
	 * @option
	 **/
	options := []aliyunOss.Option{
		aliyunOss.ObjectACL(aliyunOss.ACLPublicRead),
	}

	/**
	 * @step
	 * @获取文件数据
	 **/
	decodFileData, err := a.GetFileData()
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @上传bytes数据到oss
	 **/
	err = bucket.PutObject(objectName, bytes.NewReader(decodFileData), options...)
	if err != nil {
		return "", err
	}

	/**
	 * @step
	 * @获取实际访问的oss地址
	 **/
	ossUrl := fmt.Sprintf("%s/%s", a.DownloadDoamin, objectName)
	return ossUrl, nil
}

/**
 * @description: AliUploadMedia upload
 * @author: Jerry.Yang
 * @date: 2022-10-10 17:10:40
 * @return {*}
 */
func (a *AliUploadMedia) Upload() (string, error) {

	/**
	 * @step
	 * @解密数据
	 **/
	decodMediaData, err := base64.StdEncoding.DecodeString(a.MediaData)
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
	filePath := fmt.Sprintf("%s.%s", a.DingdingFilePath+fileName, a.MediaType)

	/**
	 * @step
	 * @创建multipart 文件字段
	 * @创建一个文件的write对象
	 **/
	formFileWriter, err := writer.CreateFormFile("media", filepath.Base(filePath))
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
	url := fmt.Sprintf("https://oapi.dingtalk.com/media/upload?access_token=%s&type=%s", a.AccessToken, a.FileType)

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
			"Content-Type": writer.FormDataContentType(),
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
 * @description: Upload
 * @author: Jerry.Yang
 * @date: 2022-10-26 17:28:45
 * @return {*}
 */
func (a *AliUploadLink) Upload() (string, error) {

	/**
	 * @step
	 * @定义返回
	 **/
	resultUrl := ""

	/**
	 * @step
	 * @对link进行urlencode
	 **/
	encodeUrl := url.QueryEscape(a.Link)

	/**
	 * @step
	 * @判断是否在工作台打开
	 **/
	if a.ContainerType != "" {
		resultUrl = fmt.Sprintf("dingtalk://dingtalkclient/action/openapp?corpid=%s&container_type=%s&app_id=%s&redirect_type=jump&redirect_url=%s", a.CropId, a.ContainerType, a.AgentId, encodeUrl)
		return resultUrl, nil
	}

	resultUrl = fmt.Sprintf("dingtalk://dingtalkclient/page/link?url=%s&pc_slide=%t", encodeUrl, a.PcSlide)
	return resultUrl, nil
}

/**
 * @description: CreateClient
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:12:39
 * @return {*}
 */
func (a *AliOssUpload) CreateClient() (*aliyunOss.Client, error) {
	return aliyunOss.New(a.EndPoint, a.AccessKeyId, a.AccessKeySecret)
}

/**
 * @description: CreateBucket
 * @param {*aliyunOss.Client} aliyunOssClient
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:12:47
 * @return {*}
 */
func (a *AliOssUpload) CreateBucket(aliyunOssClient *aliyunOss.Client) (*aliyunOss.Bucket, error) {
	return aliyunOssClient.Bucket(a.Bucket)
}

/**
 * @description: AliOssUploadFromLocalFile GetFileData
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:33:04
 * @return {*}
 */
func (a *AliOssUploadFromLocalFile) GetFileData() ([]byte, error) {
	/**
	 * @step
	 * @读取本地文件
	 **/
	res, err := os.Open(a.LocalFilePath)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	localFileBuf := bufio.NewReader(res)

	/**
	 * @step
	 * @读取本地文件二进制流
	 **/
	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := localFileBuf.Read(buf)

		/**
		 * @step
		 * @拼接内容
		 **/
		chunks = append(chunks, buf[:n]...)

		/**
		 * @step
		 * @报错
		 **/
		if err != nil && err != io.EOF {
			return nil, err
		}

		/**
		 * @step
		 * @读取完成
		 **/
		if n == 0 {
			break
		}
	}
	return chunks, nil
}

/**
 * @description: AliOssUpLoadFromFileUrl GetFileData
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:36:41
 * @return {*}
 */
func (a *AliOssUpLoadFromFileUrl) GetFileData() ([]byte, error) {

	/**
	 * @step
	 * @获取网络图片相关的资源
	 **/
	res, err := http.Get(a.FileUrl)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	/**
	 * @step
	 * @定义返回内容
	 **/
	var buffer []byte

	/**
	 * @step
	 * @读取网络图片的内容到byte中,每次读取1024
	 **/
	buf := make([]byte, 1024)
	for {
		n, err := res.Body.Read(buf)

		/**
		 * @step
		 * @拼接内容
		 **/
		buffer = append(buffer, buf[:n]...)

		/**
		 * @step
		 * @报错
		 **/
		if err != nil && err != io.EOF {
			return nil, err
		}

		/**
		 * @step
		 * @读取完成
		 **/
		if n == 0 {
			break
		}
	}
	return buffer, nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:43:53
 * @return {*}
 */
func (a *AliOssUpload) CheckParams() error {
	if a.FileName == "" {
		return errors.New("AliOssUpload Err : FileName is empty!")
	}

	if a.FileType == "" {
		return errors.New("AliOssUpload Err : FileType is empty!")
	}

	if a.FileData == "" {
		return errors.New("AliOssUpload Err : FileData is empty!")
	}

	if a.AccessKeyId == "" {
		return errors.New("AliOssUpload Err : AccessKeyId is empty!")
	}

	if a.AccessKeySecret == "" {
		return errors.New("AliOssUpload Err : AccessKeySecret is empty!")
	}

	if a.Bucket == "" {
		return errors.New("AliOssUpload Err : Bucket is empty!")
	}

	if a.EndPoint == "" {
		return errors.New("AliOssUpload Err : EndPoint is empty!")
	}

	if a.DownloadDoamin == "" {
		return errors.New("AliOssUpload Err : DownloadDoamin is empty!")
	}
	return nil
}

/**
 * @description: AliOssUploadFromLocalFile CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:47:00
 * @return {*}
 */
func (a *AliOssUploadFromLocalFile) CheckParams() error {
	err := a.AliOssUpload.CheckParams()
	if err != nil {
		return err
	}

	if a.LocalFilePath == "" {
		return errors.New("AliOssUploadFormLocalFile Err : LocalFilePath is empty!")
	}
	return nil
}

/**
 * @description: AliOssUpLoadFromFileUrl CheckParams
 * @author: Jerry.Yang
 * @date: 2022-09-22 14:47:51
 * @return {*}
 */
func (a *AliOssUpLoadFromFileUrl) CheckParams() error {
	err := a.AliOssUpload.CheckParams()
	if err != nil {
		return err
	}

	if a.FileUrl == "" {
		return errors.New("AliOssUpLoadFormFileUrl Err : FileUrl is empty!")
	}
	return nil
}
