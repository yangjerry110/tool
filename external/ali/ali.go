/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 17:05:18
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:28:43
 * @Description: ali
 */
package ali

type AliInterface interface {
	Upload() (string, error)
}

type AliDingdingNoticeInterface interface {
	NotifyMessage() (bool, error)
	FormatNotifyMessage() ([]byte, error)
	CheckParams() error
	DoNotify(accessToken string, noticeReq []byte) (bool, error)
}

/**
 * @description: Ali
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:29
 * @return {*}
 */
type Ali struct{}

/**
 * @description: AliOssUpload
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:09
 * @return {*}
 */
type AliOssUpload struct {
	AccessKeyId     string
	AccessKeySecret string
	EndPoint        string
	Bucket          string
	FileName        string
	FileType        string
	FileData        string
	DownloadDoamin  string
}

/**
 * @description: AliOssUploadFromLocalFile
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:16
 * @return {*}
 */
type AliOssUploadFromLocalFile struct {
	AliOssUpload
	LocalFilePath string
}

/**
 * @description: AliOssUpLoadFromFileUrl
 * @author: Jerry.Yang
 * @date: 2022-09-26 18:44:20
 * @return {*}
 */
type AliOssUpLoadFromFileUrl struct {
	AliOssUpload
	FileUrl string
}

type AliUploadMedia struct {
	FileType         string
	MediaType        string
	MediaData        string
	AccessToken      string
	DingdingFilePath string
}

type AliUploadLink struct {
	Link          string
	PcSlide       bool
	AgentId       string
	CropId        string
	ContainerType string
	RedirectType  string
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
	BtnJsonList      []AliDingdingNoticeBtnJson
}

/**
 * @description: AliDingdingNoticeBtnJson
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:19:49
 * @return {*}
 */
type AliDingdingNoticeBtnJson struct {
	Title     string
	ActionUrl string
}

/**
 * @description: NotifyTextMessage
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:24:33
 * @return {*}
 */
type NotifyTextMessage struct{ *AliDingdingNotice }

/**
 * @description:NotifyMarkdownMessage
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:57:18
 * @return {*}
 */
type NotifyMarkdownMessage struct{ *AliDingdingNotice }

/**
 * @description: NotifyMediaMessage
 * @author: Jerry.Yang
 * @date: 2022-10-10 17:06:28
 * @return {*}
 */
type NotifyMediaMessage struct{ *AliDingdingNotice }

/**
 * @description: NotifyVoiceMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:05:47
 * @return {*}
 */
type NotifyVoiceMessage struct{ *AliDingdingNotice }

/**
 * @description: NotifyFileMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:09:17
 * @return {*}
 */
type NotifyFileMessage struct{ *AliDingdingNotice }

/**
 * @description: NotifyLinkMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:11:29
 * @return {*}
 */
type NotifyLinkMessage struct{ *AliDingdingNotice }

/**
 * @description: NotifyNewsMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:18:01
 * @return {*}
 */
type NotifyNewsMessage struct{ *AliDingdingNotice }
