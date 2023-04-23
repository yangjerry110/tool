/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 18:43:31
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:32:42
 * @Description: ali
 */
package ali

import "github.com/yangjerry110/tool/external/ali"

type AliInterface interface {
	CreateAliInterface(aliInterface ali.AliInterface) *Ali
}

/**
 * @description: Ali
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:46:00
 * @return {*}
 */
type Ali struct {
	AliInterface               ali.AliInterface
	AliDingdingNoticeInterface ali.AliDingdingNoticeInterface
}

/**
 * @description: AliDingdingNotice
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:24:19
 * @return {*}
 */
type AliDingdingNotice struct {
	AppId            string
	CropId           string
	AppKey           string
	AppSecret        string
	AgentId          string
	AccessToken      string
	MsgType          string
	UserIds          string
	DeptIds          string
	Title            string
	ToAllUser        bool
	Msg              string
	MediaType        string
	MediaData        string
	Duration         string
	MessageUrl       string
	PicUrl           string
	SingleTitle      string
	SingleUrl        string
	BtnOrientation   string
	RedisConfPath    string
	RedisConfName    string
	DingdingFilePath string
	FileType         string
	PcSlide          bool
	ContainerType    string
	RedirectType     string
	BtnJsonList      []ali.AliDingdingNoticeBtnJson
}

/**
 * @description: CreateAliInterface
 * @param {ali.AliInterface} aliInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:45:53
 * @return {*}
 */
func CreateAliInterface(aliInterface ali.AliInterface) *Ali {
	return &Ali{AliInterface: aliInterface}
}

/**
 * @description: CreateAliDingdingNoticeInterface
 * @param {ali.AliDingdingNoticeInterface} aliDingdingNoticeInterface
 * @author: Jerry.Yang
 * @date: 2022-10-26 18:11:57
 * @return {*}
 */
func CreateAliDingdingNoticeInterface(aliDingdingNoticeInterface ali.AliDingdingNoticeInterface) *Ali {
	return &Ali{AliDingdingNoticeInterface: aliDingdingNoticeInterface}
}

/**
 * @description: AliUploadOss
 * @param {string} accessKeyId
 * @param {string} accessKeySecret
 * @param {string} endPoint
 * @param {string} bucket
 * @param {string} filename
 * @param {string} fileType
 * @param {string} fileData
 * @param {string} downloadDomain
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:51:56
 * @return {*}
 */
func AliUploadOss(accessKeyId string, accessKeySecret string, endPoint string, bucket string, filename string, fileType string, fileData string, downloadDomain string) (string, error) {
	return CreateAliInterface(&ali.AliOssUpload{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret, EndPoint: endPoint, Bucket: bucket, FileName: filename, FileType: fileType, FileData: fileData, DownloadDoamin: downloadDomain}).AliInterface.Upload()
}

/**
 * @description: AliUploadOssFromLocaFile
 * @param {string} accessKeyId
 * @param {string} accessKeySecret
 * @param {string} endPoint
 * @param {string} bucket
 * @param {string} filename
 * @param {string} fileType
 * @param {string} fileData
 * @param {string} downloadDomain
 * @param {string} localFilePath
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:55:44
 * @return {*}
 */
func AliUploadOssFromLocaFile(accessKeyId string, accessKeySecret string, endPoint string, bucket string, filename string, fileType string, fileData string, downloadDomain string, localFilePath string) (string, error) {
	return CreateAliInterface(&ali.AliOssUploadFromLocalFile{AliOssUpload: ali.AliOssUpload{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret, EndPoint: endPoint, Bucket: bucket, FileName: filename, FileType: fileType, FileData: fileData, DownloadDoamin: downloadDomain}, LocalFilePath: localFilePath}).AliInterface.Upload()
}

/**
 * @description: AliUploadOssFromFileUrl
 * @param {string} accessKeyId
 * @param {string} accessKeySecret
 * @param {string} endPoint
 * @param {string} bucket
 * @param {string} filename
 * @param {string} fileType
 * @param {string} fileData
 * @param {string} downloadDomain
 * @param {string} fileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:55:55
 * @return {*}
 */
func AliUploadOssFromFileUrl(accessKeyId string, accessKeySecret string, endPoint string, bucket string, filename string, fileType string, fileData string, downloadDomain string, fileUrl string) (string, error) {
	return CreateAliInterface(&ali.AliOssUpLoadFromFileUrl{AliOssUpload: ali.AliOssUpload{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret, EndPoint: endPoint, Bucket: bucket, FileName: filename, FileType: fileType, FileData: fileData, DownloadDoamin: downloadDomain}, FileUrl: fileUrl}).AliInterface.Upload()
}

/**
 * @description: AliDingdingNotice
 * @author: Jerry.Yang
 * @date: 2022-10-26 18:14:14
 * @return {*}
 */
func (a *AliDingdingNotice) AliDingdingNotice() (bool, error) {
	return CreateAliDingdingNoticeInterface(&ali.AliDingdingNotice{
		AppId: a.AppId, CropId: a.CropId, AppKey: a.AppKey, AppSecret: a.AppSecret, AgentId: a.AgentId, AccessToken: a.AccessToken, MsgType: a.MsgType, UserIds: a.UserIds, DeptIds: a.DeptIds, Title: a.Title, ToAllUser: a.ToAllUser, Msg: a.Msg, MediaType: a.MediaType, MediaData: a.MediaData, Duration: a.Duration, MessageUrl: a.MessageUrl, PicUrl: a.PicUrl, SingleTitle: a.SingleTitle, SingleUrl: a.SingleUrl, BtnOrientation: a.BtnOrientation, RedisConfPath: a.RedisConfPath, DingdingFilePath: a.DingdingFilePath, FileType: a.FileType, PcSlide: a.PcSlide, ContainerType: a.ContainerType, RedirectType: a.RedirectType, BtnJsonList: a.BtnJsonList, RedisConfName: a.RedisConfName}).AliDingdingNoticeInterface.NotifyMessage()
}
