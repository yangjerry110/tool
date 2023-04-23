/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 17:32:01
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:32:21
 * @Description:
 */
package qiwei

import (
	"github.com/yangjerry110/tool/external/qiwei"
)

type QiweiInterface interface {
	CreateExternalQiweiInterface(externalQiweiInterface qiwei.ExternalQiweiInterface) *Qiwei
	NotifyMessage() (bool, error)
	NotifyMessageBot() (bool, error)
	QiweiUploadMedia(appId string, cropId string, cropSecret string, mediaData string, mediaType string, qiweiFilePath string) (string, error)
}

type Qiwei struct {
	ExternalQiweiInterface qiwei.ExternalQiweiInterface
}

type QiweiNotice struct {
	AppId               string
	MsgType             string
	CropId              string
	CropSecret          string
	AgentId             string
	DepartmentIds       string
	TagIds              string
	UserIds             string
	Safe                int32
	SendMsg             string
	MediaData           string
	MediaType           string
	Title               string
	Description         string
	Url                 string
	PicUrl              string
	EnableIdTrans       int32
	Btntxt              string
	AppletId            string
	AppletPagepath      string
	QiweiFilePath       string
	BotUrl              string
	RedisConfPath       string
	RedisConfName       string
	MentionedList       []string
	MentionedMobileList []string
}

/**
 * @description: CreateInterface
 * @param {qiwei.ExternalQiweiInterface} externalQiweiInterface
 * @author: Jerry.Yang
 * @date: 2022-09-26 17:36:28
 * @return {*}
 */
func CreateExternalQiweiInterface(externalQiweiInterface qiwei.ExternalQiweiInterface) *Qiwei {
	return &Qiwei{ExternalQiweiInterface: externalQiweiInterface}
}

/**
 * @description: NotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:26:10
 * @return {*}
 */
func (q *QiweiNotice) NotifyMessage() (bool, error) {
	return CreateExternalQiweiInterface(&qiwei.QiweiNotice{AppId: q.AppId, MsgType: q.MsgType, CropId: q.CropId, CropSecret: q.CropSecret, AgentId: q.AgentId, DepartmentIds: q.DepartmentIds, TagIds: q.TagIds, UserIds: q.UserIds, Safe: q.Safe, SendMsg: q.SendMsg, MediaData: q.MediaData, MediaType: q.MediaType, Title: q.Title, Description: q.Description, Url: q.Url, PicUrl: q.PicUrl, EnableIdTrans: q.EnableIdTrans, Btntxt: q.Btntxt, AppletId: q.AppletId, AppletPagepath: q.AppletPagepath, QiweiFilePath: q.QiweiFilePath, RedisConfPath: q.RedisConfPath, RedisConfName: q.RedisConfName}).ExternalQiweiInterface.NotifyMessage()
}

/**
 * @description: NotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:26:17
 * @return {*}
 */
func (q *QiweiNotice) NotifyMessageBot() (bool, error) {
	return CreateExternalQiweiInterface(&qiwei.QiweiNotice{MsgType: q.MsgType, BotUrl: q.BotUrl, SendMsg: q.SendMsg, MediaData: q.MediaData, Title: q.Title, Description: q.Description, Url: q.Url, PicUrl: q.PicUrl, MentionedList: q.MentionedList, MentionedMobileList: q.MentionedMobileList}).ExternalQiweiInterface.NotifyMessageBot()
}

/**
 * @description: QiweiUploadMedia
 * @param {string} appId
 * @param {string} cropId
 * @param {string} cropSecret
 * @param {string} mediaData
 * @param {string} mediaType
 * @param {string} qiweiFilePath
 * @author: Jerry.Yang
 * @date: 2022-09-26 16:57:56
 * @return {*}
 */
func QiweiUploadMedia(appId string, cropId string, cropSecret string, mediaData string, mediaType string, qiweiFilePath string, redisConfPath string, redisConfName string) (string, error) {
	return CreateExternalQiweiInterface(&qiwei.QiweiNotice{AppId: appId, CropId: cropId, CropSecret: cropSecret, MediaData: mediaData, MediaType: mediaType, QiweiFilePath: qiweiFilePath, RedisConfPath: redisConfPath, RedisConfName: redisConfName}).ExternalQiweiInterface.Upload()
}
