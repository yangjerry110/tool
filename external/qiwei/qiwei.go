/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-26 17:01:20
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:29:03
 * @Description: qiwei
 */
package qiwei

/**
 * @ExternalQiweiInterface 公用interface
 * @author Jerry.Yang
 * @date 2022-09-26 17:58:46
 **/
type ExternalQiweiInterface interface {
	CheckParams() error
	NotifyMessage() (bool, error)
	FormatNotifyParams() ([]byte, error)
	DoNotify(accessToken string, qiweiMsg []byte) (bool, error)
	NotifyMessageBot() (bool, error)
	DoNotifyBot(accessToken string, qiweiMsg []byte) (bool, error)
	Upload() (string, error)
}

type ExternalQiwei struct{}

/**
 * @step
 * @qiwei notice
 **/
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
 * @step
 * @定义基础的请求结构
 **/
type NotifyMessage struct {
	Touser                 string `json:"touser"`
	Toparty                string `json:"toparty"`
	Totag                  string `json:"totag"`
	Msgtype                string `json:"msgtype"`
	Agentid                string `json:"agentid"`
	EnableIdTrans          int32  `json:"enable_id_trans"`
	EnableDuplicateCheck   int32  `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32  `json:"duplicate_check_interval"`
}

/**
 * @step
 * @定义textMessage的结构
 **/
type NotifyTextMessage struct{ QiweiNotice }

/**
 * @step
 * @定义markdownMessage的结构
 **/
type NotifyMarkdownMessage struct{ QiweiNotice }

/**
 * @step
 * @定义imageMessage的结构
 **/
type NotifyImageMessage struct{ QiweiNotice }

/**
 * @step
 * @定义cardMessage的结构
 **/
type NotifyCardMessage struct{ QiweiNotice }

/**
 * @step
 * @定义newsMessage的结构
 **/
type NotifyNewsMessage struct{ QiweiNotice }

/**
 * @step
 * @定义botTextMessage
 **/
type NotifyBotTextMessage struct{ QiweiNotice }

/**
 * @step
 * @定义botMarkdownMessage
 **/
type NotifyBotMarkdownMessage struct{ QiweiNotice }

/**
 * @step
 * @定义botImageMessage
 **/
type NotifyBotImageMessage struct{ QiweiNotice }

/**
 * @step
 * @定义botNewsMessage
 **/
type NotifyBotNewsMessage struct{ QiweiNotice }
