/*
 * @Author: Jerry.Yang
 * @Date: 2022-10-10 15:36:04
 * @LastEditors: Jerry.Yang
 * @LastEditTime: 2022-11-10 18:28:51
 * @Description: ali dingding notice
 */
package ali

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	mytoolCommon "github.com/yangjerry110/tool/common"
	mytoolHttp "github.com/yangjerry110/tool/http"
)

/**
 * @description: NotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:21:31
 * @return {*}
 */
func (a *AliDingdingNotice) NotifyMessage() (bool, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	a.CheckParams()

	/**
	 * @step
	 * @初始化一个通道,获取钉钉请求参数
	 **/
	notifyAliDingdingReqChan := make(chan []byte)
	notifyAliDingdingReqErrChan := make(chan error)

	/**
	 * @step
	 * @获取accessToken
	 **/
	aliCommon := mytoolCommon.AliCommon{AppId: a.AppId, AppKey: a.AppKey, AppSecret: a.AppSecret, RedisConfPath: a.RedisConfPath, RedisConfName: a.RedisConfName}
	accessToken, err := aliCommon.GetAccessToken()
	if err != nil {
		return false, err
	}
	a.AccessToken = accessToken

	/**
	 * @step
	 * @formatMsg
	 **/
	go func() {
		switch a.MsgType {
		case "text":
			notifyMessage := NotifyTextMessage{AliDingdingNotice: a}
			notifyReq, notifyReqErr := notifyMessage.FormatNotifyMessage()
			notifyAliDingdingReqChan <- notifyReq
			notifyAliDingdingReqErrChan <- notifyReqErr
		case "markdown":
			notifyMessage := NotifyMarkdownMessage{AliDingdingNotice: a}
			notifyReq, notifyReqErr := notifyMessage.FormatNotifyMessage()
			notifyAliDingdingReqChan <- notifyReq
			notifyAliDingdingReqErrChan <- notifyReqErr
		case "image":
			notifyMessage := NotifyMediaMessage{AliDingdingNotice: a}
			notifyReq, notifyReqErr := notifyMessage.FormatNotifyMessage()
			notifyAliDingdingReqChan <- notifyReq
			notifyAliDingdingReqErrChan <- notifyReqErr
		case "link":
			notifyMessage := NotifyLinkMessage{AliDingdingNotice: a}
			notifyReq, notifyReqErr := notifyMessage.FormatNotifyMessage()
			notifyAliDingdingReqChan <- notifyReq
			notifyAliDingdingReqErrChan <- notifyReqErr
		case "file":
			notifyMessage := NotifyFileMessage{AliDingdingNotice: a}
			notifyReq, notifyReqErr := notifyMessage.FormatNotifyMessage()
			notifyAliDingdingReqChan <- notifyReq
			notifyAliDingdingReqErrChan <- notifyReqErr
		case "voice":
			notifyMessage := NotifyVoiceMessage{AliDingdingNotice: a}
			notifyReq, notifyReqErr := notifyMessage.FormatNotifyMessage()
			notifyAliDingdingReqChan <- notifyReq
			notifyAliDingdingReqErrChan <- notifyReqErr
		case "news":
			notifyMessage := NotifyNewsMessage{AliDingdingNotice: a}
			notifyReq, notifyReqErr := notifyMessage.FormatNotifyMessage()
			notifyAliDingdingReqChan <- notifyReq
			notifyAliDingdingReqErrChan <- notifyReqErr
		default:
			notifyAliDingdingReqChan <- nil
			notifyAliDingdingReqErrChan <- errors.New("AliDingdingNotice Err : no match msgType!")
		}
	}()

	/**
	 * @step
	 * @获取请求钉钉通知的参数
	 **/
	notifyAliDingdingReq := <-notifyAliDingdingReqChan

	/**
	 * @step
	 * @获取组装参数的错误
	 **/
	notifyAliDingdingReqErr := <-notifyAliDingdingReqErrChan
	if notifyAliDingdingReqErr != nil {
		return false, notifyAliDingdingReqErr
	}

	/**
	 * @step
	 * @判断参数
	 **/
	if notifyAliDingdingReq == nil {
		return false, errors.New("AliDingdingNotice Err : AliDingdingNoticeReq is nil")
	}

	/**
	 * @step
	 * @发送通知
	 **/
	notifyMsgResult, err := a.DoNotify(accessToken, notifyAliDingdingReq)
	return notifyMsgResult, err
}

/**
 * @description: FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:21:40
 * @return {*}
 */
func (n *NotifyTextMessage) FormatNotifyMessage() ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义请求参数结构
	 **/
	type NotifyContentMsg struct {
		Content string `json:"content"`
	}

	type NotifyMsg struct {
		MsgType string           `json:"msgtype"`
		Text    NotifyContentMsg `json:"text"`
	}

	/**
	 * @step
	 * @渲染msg
	 **/
	reqMsg := NotifyMsg{
		MsgType: "text",
		Text:    NotifyContentMsg{Content: n.AliDingdingNotice.Msg},
	}

	/**
	 * @step
	 * @定义notifyMessage
	 **/
	notifyMessage := map[string]interface{}{
		"agent_id":    n.AliDingdingNotice.AgentId,
		"to_all_user": n.AliDingdingNotice.ToAllUser,
		"msg":         reqMsg,
	}

	/**
	 * @step
	 * @userIds
	 **/
	if n.AliDingdingNotice.UserIds != "" {
		notifyMessage["userid_list"] = n.AliDingdingNotice.UserIds
	}

	/**
	 * @step
	 * @deptIds
	 **/
	if n.AliDingdingNotice.DeptIds != "" {
		notifyMessage["dept_id_list"] = n.AliDingdingNotice.DeptIds
	}

	/**
	 * @step
	 * @渲染notfiyMessage
	 **/
	jsonReq, err := json.Marshal(notifyMessage)

	/**
	 * @step
	 * @err
	 **/
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

/**
 * @description: FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:57:50
 * @return {*}
 */
func (n *NotifyMarkdownMessage) FormatNotifyMessage() ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义请求参数结构
	 **/
	type NotifyContentMsg struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}

	type NotifyMsg struct {
		MsgType  string           `json:"msgtype"`
		Markdown NotifyContentMsg `json:"markdown"`
	}

	/**
	 * @step
	 * @渲染msg
	 **/
	reqMsg := NotifyMsg{
		MsgType:  "markdown",
		Markdown: NotifyContentMsg{Title: n.AliDingdingNotice.Title, Text: n.AliDingdingNotice.Msg},
	}

	/**
	 * @step
	 * @定义notifyMessage
	 **/
	notifyMessage := map[string]interface{}{
		"agent_id":    n.AliDingdingNotice.AgentId,
		"to_all_user": n.AliDingdingNotice.ToAllUser,
		"msg":         reqMsg,
	}

	/**
	 * @step
	 * @userIds
	 **/
	if n.AliDingdingNotice.UserIds != "" {
		notifyMessage["userid_list"] = n.AliDingdingNotice.UserIds
	}

	/**
	 * @step
	 * @deptIds
	 **/
	if n.AliDingdingNotice.DeptIds != "" {
		notifyMessage["dept_id_list"] = n.AliDingdingNotice.DeptIds
	}

	/**
	 * @step
	 * @渲染notfiyMessage
	 **/
	jsonReq, err := json.Marshal(notifyMessage)

	/**
	 * @step
	 * @err
	 **/
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

/**
 * @description: FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-10 17:06:59
 * @return {*}
 */
func (n *NotifyMediaMessage) FormatNotifyMessage() ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @获取mediaId
	 **/
	aliUploadMedia := AliUploadMedia{
		FileType:    "image",
		MediaType:   n.AliDingdingNotice.MediaType,
		MediaData:   n.AliDingdingNotice.MediaData,
		AccessToken: n.AliDingdingNotice.AccessToken,
	}
	mediaId, err := aliUploadMedia.Upload()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义请求参数结构
	 **/
	type NotifyContentMsg struct {
		MediaId string `json:"media_id"`
	}

	type NotifyMsg struct {
		MsgType string           `json:"msgtype"`
		Image   NotifyContentMsg `json:"image"`
	}

	/**
	 * @step
	 * @渲染msg
	 **/
	reqMsg := NotifyMsg{
		MsgType: "image",
		Image:   NotifyContentMsg{MediaId: mediaId},
	}

	/**
	 * @step
	 * @定义notifyMessage
	 **/
	notifyMessage := map[string]interface{}{
		"agent_id":    n.AliDingdingNotice.AgentId,
		"to_all_user": n.AliDingdingNotice.ToAllUser,
		"msg":         reqMsg,
	}

	/**
	 * @step
	 * @userIds
	 **/
	if n.AliDingdingNotice.UserIds != "" {
		notifyMessage["userid_list"] = n.AliDingdingNotice.UserIds
	}

	/**
	 * @step
	 * @deptIds
	 **/
	if n.AliDingdingNotice.DeptIds != "" {
		notifyMessage["dept_id_list"] = n.AliDingdingNotice.DeptIds
	}

	/**
	 * @step
	 * @渲染notfiyMessage
	 **/
	jsonReq, err := json.Marshal(notifyMessage)

	/**
	 * @step
	 * @err
	 **/
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

/**
 * @description: FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:06:15
 * @return {*}
 */
func (n *NotifyVoiceMessage) FormatNotifyMessage() ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @获取mediaId
	 **/
	aliUploadMedia := AliUploadMedia{
		MediaType:   n.AliDingdingNotice.MediaType,
		MediaData:   n.AliDingdingNotice.MediaData,
		AccessToken: n.AliDingdingNotice.AccessToken,
	}
	mediaId, err := aliUploadMedia.Upload()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义请求参数结构
	 **/
	type NotifyContentMsg struct {
		MediaId  string `json:"media_id"`
		Duration string `json:"duration"`
	}

	type NotifyMsg struct {
		MsgType string           `json:"msgtype"`
		Voice   NotifyContentMsg `json:"voice"`
	}

	/**
	 * @step
	 * @渲染msg
	 **/
	reqMsg := NotifyMsg{
		MsgType: "image",
		Voice:   NotifyContentMsg{MediaId: mediaId, Duration: n.Duration},
	}

	/**
	 * @step
	 * @定义notifyMessage
	 **/
	notifyMessage := map[string]interface{}{
		"agent_id":    n.AliDingdingNotice.AgentId,
		"to_all_user": n.AliDingdingNotice.ToAllUser,
		"msg":         reqMsg,
	}

	/**
	 * @step
	 * @userIds
	 **/
	if n.AliDingdingNotice.UserIds != "" {
		notifyMessage["userid_list"] = n.AliDingdingNotice.UserIds
	}

	/**
	 * @step
	 * @deptIds
	 **/
	if n.AliDingdingNotice.DeptIds != "" {
		notifyMessage["dept_id_list"] = n.AliDingdingNotice.DeptIds
	}

	/**
	 * @step
	 * @渲染notfiyMessage
	 **/
	jsonReq, err := json.Marshal(notifyMessage)

	/**
	 * @step
	 * @err
	 **/
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

/**
 * @description: NotifyFileMessage FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:10:00
 * @return {*}
 */
func (n *NotifyFileMessage) FormatNotifyMessage() ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @获取mediaId
	 **/
	aliUploadMedia := AliUploadMedia{
		MediaType:   n.AliDingdingNotice.MediaType,
		MediaData:   n.AliDingdingNotice.MediaData,
		AccessToken: n.AliDingdingNotice.AccessToken,
	}
	mediaId, err := aliUploadMedia.Upload()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义请求参数结构
	 **/
	type NotifyContentMsg struct {
		MediaId string `json:"media_id"`
	}

	type NotifyMsg struct {
		MsgType string           `json:"msgtype"`
		File    NotifyContentMsg `json:"file"`
	}

	/**
	 * @step
	 * @渲染msg
	 **/
	reqMsg := NotifyMsg{
		MsgType: "image",
		File:    NotifyContentMsg{MediaId: mediaId},
	}

	/**
	 * @step
	 * @定义notifyMessage
	 **/
	notifyMessage := map[string]interface{}{
		"agent_id":    n.AliDingdingNotice.AgentId,
		"to_all_user": n.AliDingdingNotice.ToAllUser,
		"msg":         reqMsg,
	}

	/**
	 * @step
	 * @userIds
	 **/
	if n.AliDingdingNotice.UserIds != "" {
		notifyMessage["userid_list"] = n.AliDingdingNotice.UserIds
	}

	/**
	 * @step
	 * @deptIds
	 **/
	if n.AliDingdingNotice.DeptIds != "" {
		notifyMessage["dept_id_list"] = n.AliDingdingNotice.DeptIds
	}

	/**
	 * @step
	 * @渲染notfiyMessage
	 **/
	jsonReq, err := json.Marshal(notifyMessage)

	/**
	 * @step
	 * @err
	 **/
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

/**
 * @description:FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:17:22
 * @return {*}
 */
func (n *NotifyLinkMessage) FormatNotifyMessage() ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @获取mediaId
	 **/
	aliUploadMedia := AliUploadMedia{
		FileType:    n.FileType,
		MediaType:   n.AliDingdingNotice.MediaType,
		MediaData:   n.AliDingdingNotice.MediaData,
		AccessToken: n.AliDingdingNotice.AccessToken,
	}
	mediaId, err := aliUploadMedia.Upload()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @获取跳转链接
	 **/
	aliUploadLink := AliUploadLink{
		Link:          n.MessageUrl,
		CropId:        n.CropId,
		PcSlide:       n.PcSlide,
		AgentId:       n.AgentId,
		ContainerType: n.ContainerType,
		RedirectType:  n.RedirectType,
	}
	messageUrl, err := aliUploadLink.Upload()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义请求参数结构
	 **/
	type NotifyContentMsg struct {
		MessageUrl string `json:"messageUrl"`
		PicUrl     string `json:"picUrl"`
		Title      string `json:"title"`
		Text       string `json:"text"`
	}

	type NotifyMsg struct {
		MsgType string           `json:"msgtype"`
		Link    NotifyContentMsg `json:"link"`
	}

	/**
	 * @step
	 * @渲染msg
	 **/
	reqMsg := NotifyMsg{
		MsgType: "link",
		Link:    NotifyContentMsg{MessageUrl: messageUrl, PicUrl: mediaId, Title: n.Title, Text: n.Msg},
	}

	/**
	 * @step
	 * @定义notifyMessage
	 **/
	notifyMessage := map[string]interface{}{
		"agent_id":    n.AliDingdingNotice.AgentId,
		"to_all_user": n.AliDingdingNotice.ToAllUser,
		"msg":         reqMsg,
	}

	/**
	 * @step
	 * @userIds
	 **/
	if n.AliDingdingNotice.UserIds != "" {
		notifyMessage["userid_list"] = n.AliDingdingNotice.UserIds
	}

	/**
	 * @step
	 * @deptIds
	 **/
	if n.AliDingdingNotice.DeptIds != "" {
		notifyMessage["dept_id_list"] = n.AliDingdingNotice.DeptIds
	}

	/**
	 * @step
	 * @渲染notfiyMessage
	 **/
	jsonReq, err := json.Marshal(notifyMessage)

	/**
	 * @step
	 * @err
	 **/
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

/**
 * @description: FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:18:14
 * @return {*}
 */
func (n *NotifyNewsMessage) FormatNotifyMessage() ([]byte, error) {

	/**
	 * @step
	 * @检查参数
	 **/
	err := n.CheckParams()
	if err != nil {
		return nil, err
	}

	/**
	 * @step
	 * @定义请求参数结构
	 **/
	type NotifyOrientationContentBtnJsonListMsg struct {
		Title     string `json:"title"`
		ActionUrl string `json:"action_list"`
	}

	type NotifyOrientationContentMsg struct {
		Title          string `json:"title"`
		Markdown       string `json:"markdown"`
		BtnOrientation string `json:"btn_orientation"`
		BtnJsonList    []NotifyOrientationContentBtnJsonListMsg
	}

	type NotifySingleContentMsg struct {
		Title       string `json:"title"`
		Markdown    string `json:"markdown"`
		SingleTitle string `json:"single_title"`
		SingleUrl   string `json:"single_url"`
	}

	type NotifyOrientationMsg struct {
		MsgType    string                      `json:"msgtype"`
		ActionCard NotifyOrientationContentMsg `json:"action_card"`
	}

	type NotifySingleMsg struct {
		MsgType    string                 `json:"msgtype"`
		ActionCard NotifySingleContentMsg `json:"action_card"`
	}

	/**
	 * @step
	 * @定义notifyMessage
	 **/
	notifyMessage := map[string]interface{}{
		"agent_id":    n.AliDingdingNotice.AgentId,
		"to_all_user": n.AliDingdingNotice.ToAllUser,
	}

	if n.BtnOrientation == "" {

		/**
		 * @step
		 * @渲染msg
		 **/
		reqMsgSingle := NotifySingleMsg{
			MsgType:    "action_card",
			ActionCard: NotifySingleContentMsg{Title: n.Title, Markdown: n.Msg, SingleTitle: n.SingleTitle, SingleUrl: n.SingleUrl},
		}
		notifyMessage["msg"] = reqMsgSingle
	} else {

		/**
		 * @step
		 * @渲染NotifyOrientationContentBtnJsonListMsg
		 **/
		notifyOrientationContentBtnJsonListMsgs := make([]NotifyOrientationContentBtnJsonListMsg, 0)
		for _, item := range n.BtnJsonList {
			notifyOrientationContentBtnJsonListMsgs = append(notifyOrientationContentBtnJsonListMsgs, NotifyOrientationContentBtnJsonListMsg{
				Title:     item.Title,
				ActionUrl: item.ActionUrl,
			})
		}

		reqMsgOrientation := NotifyOrientationMsg{
			MsgType:    "action_card",
			ActionCard: NotifyOrientationContentMsg{Title: n.Title, Markdown: n.Msg, BtnOrientation: n.BtnOrientation, BtnJsonList: notifyOrientationContentBtnJsonListMsgs},
		}
		notifyMessage["msg"] = reqMsgOrientation
	}

	/**
	 * @step
	 * @userIds
	 **/
	if n.AliDingdingNotice.UserIds != "" {
		notifyMessage["userid_list"] = n.AliDingdingNotice.UserIds
	}

	/**
	 * @step
	 * @deptIds
	 **/
	if n.AliDingdingNotice.DeptIds != "" {
		notifyMessage["dept_id_list"] = n.AliDingdingNotice.DeptIds
	}

	/**
	 * @step
	 * @渲染notfiyMessage
	 **/
	jsonReq, err := json.Marshal(notifyMessage)

	/**
	 * @step
	 * @err
	 **/
	if err != nil {
		return nil, err
	}
	return jsonReq, nil
}

/**
 * @description: FormatNotifyMessage
 * @author: Jerry.Yang
 * @date: 2022-10-26 18:16:33
 * @return {*}
 */
func (a AliDingdingNotice) FormatNotifyMessage() ([]byte, error) {
	return nil, nil
}

/**
 * @description: DoNotify
 * @param {string} accessToken
 * @param {[]byte} noticeReq
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:21:49
 * @return {*}
 */
func (a *AliDingdingNotice) DoNotify(accessToken string, noticeReq []byte) (bool, error) {

	/**
	 * @step
	 * @url
	 **/
	url := fmt.Sprintf("https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2?access_token=%s", accessToken)

	/**
	 * @step
	 * @返回数据 struct
	 **/
	type NotifyOutput struct {
		RequestId string `json:"request_id"`
		TaskId    int64  `json:"task_id"`
		ErrCode   int32  `json:"errcode"`
		ErrMsg    string `json:"errmsg"`
	}

	/**
	 * @step
	 * @发送消息
	 **/
	resp := &NotifyOutput{}

	/**
	 * @step
	 * @设置httpOption
	 **/
	mytoolHttpOptions := mytoolHttp.HttpOptions{}
	httpOptions := []mytoolHttp.HttpOptionFunc{
		mytoolHttpOptions.SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}),
	}

	/**
	 * @step
	 * @组装请求参数
	 **/
	httpClient := mytoolHttp.HttpClient{
		Method:  "POST",
		Url:     url,
		Body:    bytes.NewBuffer(noticeReq),
		Options: httpOptions,
		Output:  resp,
	}
	httpClient.Request()

	/**
	 * @step
	 * @判断错误
	 **/
	if resp.ErrCode != 0 {
		return false, errors.New(resp.ErrMsg)
	}
	return true, nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-10-10 16:22:56
 * @return {*}
 */
func (n *NotifyTextMessage) CheckParams() error {
	if n.AliDingdingNotice.Msg == "" {
		return errors.New("NotifyTextMessage CheckParams : msg is not set")
	}
	return nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:04:52
 * @return {*}
 */
func (n *NotifyMediaMessage) CheckParams() error {
	if n.MediaType == "" {
		return errors.New("NotifyMediaMessage CheckParams : mediaType is not set")
	}

	if n.MediaData == "" {
		return errors.New("NotifyMediaMessage CheckParams : mediaData is not set")
	}

	return nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:08:22
 * @return {*}
 */
func (n *NotifyVoiceMessage) CheckParams() error {
	if n.MediaType == "" {
		return errors.New("NotifyVoiceMessage CheckParams : mediaType is not set")
	}

	if n.MediaData == "" {
		return errors.New("NotifyVoiceMessage CheckParams : mediaData is not set")
	}

	if n.Duration == "" {
		return errors.New("NotifyVoiceMessage CheckParams : duration is not set")
	}
	return nil
}

/**
 * @description:
 * @author: Jerry.Yang
 * @date:
 * @return {*}
 */
func (n *NotifyFileMessage) CheckParams() error {
	if n.MediaType == "" {
		return errors.New("NotifyFileMessage CheckParams : mediaType is not set")
	}

	if n.MediaData == "" {
		return errors.New("NotifyFileMessage CheckParams : mediaData is not set")
	}

	return nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:15:46
 * @return {*}
 */
func (n *NotifyLinkMessage) CheckParams() error {
	if n.MessageUrl == "" {
		return errors.New("NotifyLinkMessage CheckParams : messageUrl is not set")
	}

	if n.FileType == "" {
		return errors.New("NotifyLinkMessage CheckParams : fileType is not set")
	}

	if n.MediaType == "" {
		return errors.New("NotifyLinkMessage CheckParams : mediaType is not set")
	}

	if n.MediaData == "" {
		return errors.New("NotifyLinkMessage CheckParams : mediaData is not set")
	}

	if n.ContainerType != "" {
		if n.CropId == "" {
			return errors.New("NotifyLinkMessage CheckParams : CropId is not set")
		}

		if n.RedirectType == "" {
			return errors.New("NotifyLinkMessage CheckParams : RedirectType is not set")
		}
	}

	if n.Title == "" {
		return errors.New("NotifyLinkMessage CheckParams : title is not set")
	}

	if n.Msg == "" {
		return errors.New("NotifyLinkMessage CheckParams : msg is not set")
	}
	return nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-10-11 16:23:23
 * @return {*}
 */
func (n *NotifyNewsMessage) CheckParams() error {
	if n.Title == "" {
		return errors.New("NotifyNewsMessage CheckParams : title is not set")
	}

	if n.Msg == "" {
		return errors.New("NotifyNewsMessage CheckParams : msg is not set")
	}

	if n.BtnOrientation != "" {
		if n.BtnJsonList == nil {
			return errors.New("NotifyNewsMessage CheckParams : btnJsonList is not set")
		}
	} else {
		if n.SingleTitle == "" {
			return errors.New("NotifyNewsMessage CheckParams : singleTitle is not set")
		}

		if n.SingleUrl == "" {
			return errors.New("NotifyNewsMessage CheckParams : singleUrl is not set")
		}

	}

	return nil
}

/**
 * @description: CheckParams
 * @author: Jerry.Yang
 * @date: 2022-10-10 15:41:16
 * @return {*}
 */
func (a *AliDingdingNotice) CheckParams() error {

	if a.AppId == "" {
		return errors.New("AliDingdingNotice CheckParams : appId is not set")
	}

	if a.AppKey == "" {
		return errors.New("AliDingdingNotice CheckParams : appKey is not set")
	}

	if a.AppSecret == "" {
		return errors.New("AliDingdingNotice CheckParams : appSecret is not set")
	}

	if a.RedisConfPath == "" {
		return errors.New("AliDingdingNotice CheckParams : redisConfPath is not set")
	}
	return nil
}
