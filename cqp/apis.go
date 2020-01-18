// +build windows,386,cgo,!websocket

package cqp

// #include "apis.h"
import "C"
import (
	"fmt"
	sc "golang.org/x/text/encoding/simplifiedchinese"
)

// Main 原生插件不使用
func Main() {}

func cString(str string) *C.char {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return C.CString(gbstr)
}

func goString(str *C.char) string {
	utf8str, _ := sc.GB18030.NewDecoder().String(C.GoString(str))
	return utf8str
}

func cBool(b bool) C.int32_t {
	if b {
		return 1
	}
	return 0
}

// AddLog 增加运行日志
// 	priority 是Log的优先级，请使用cqp预定义好的值。
// 	logType 是日志类型，酷Q日志窗口将将其显示在日志本身的前一列。
// 	reason 是日志内容
func AddLog(p Priority, logType, reason string) int32 {
	return int32(C.CQ_addLog(
		C.int32_t(p),
		cString(logType),
		cString(reason),
	))
}

// CanSendImage 能否发送图片
func CanSendImage() bool {
	return C.CQ_canSendImage() != 0
}

// CanSendRecord 能否发送语音
func CanSendRecord() bool {
	return C.CQ_canSendRecord() != 0
}

// DeleteMsg 撤回消息
func DeleteMsg(msgID int64) int32 {
	return int32(C.CQ_deleteMsg(C.int64_t(msgID)))
}

// GetAppDir 取应用目录
// 返回的路径末尾带"\"，一般用法如下，所以不用关心是否有斜线
// 	os.Open(filepath.Join(cqp.GetAppDir(), "data.db"))
func GetAppDir() string {
	return goString(C.CQ_getAppDirectory())
}

// GetCookies 获取cookies
//
// 需要严格授权
func GetCookies() string {
	return goString(C.CQ_getCookies())
}

// GetCookiesV2 获取cookies V2
//
// 需要严格授权
func GetCookiesV2(domain string) string {
	return goString(C.CQ_getCookiesV2(cString(domain)))
}

// GetCSRFToken 获取CSRF Token
//
// 需要严格授权
func GetCSRFToken() int32 {
	return int32(C.CQ_getCsrfToken())
}

// GetFriendList 获取好友列表
func GetFriendList() []FriendInfo {
	raw := goString(C.CQ_getFriendList(cBool(false)))
	list, err := UnpackFriendList(raw)
	if err != nil {
		panic(fmt.Errorf("cqp: 内部错误，酷Q返回的好友列表格式不正确: %v", err))
	}
	return list
}

func GetGroupInfo(group int64, noCatch bool) GroupDetail {
	raw := goString(C.CQ_getGroupInfo(C.int64_t(group), cBool(noCatch)))
	info, err := UnpackGroupInfo(raw)
	if err != nil {
		panic(fmt.Errorf("cqp: 内部错误，酷Q返回的群信息格式不正确: %v", err))
	}
	return info
}

// GetGroupList 获取群列表
func GetGroupList() []GroupInfo {
	raw := goString(C.CQ_getGroupList())
	list, err := UnpackGroupList(raw)
	if err != nil {
		panic(fmt.Errorf("cqp: 内部错误，酷Q返回的群列表格式不正确: %v", err))
	}
	return list
}

// GetGroupMemberInfo 获取群成员信息
func GetGroupMemberInfo(group, qq int64, noCatch bool) GroupMember {
	raw := goString(C.CQ_getGroupMemberInfoV2(
		C.int64_t(group), C.int64_t(qq), cBool(noCatch),
	))
	member, err := UnpackGroupMemberInfo(raw)
	if err != nil {
		panic(fmt.Errorf("cqp: 内部错误，酷Q返回的群成员信息格式不正确: %v", err))
	}
	return member
}

// GetGroupMemberList 获取群成员列表
func GetGroupMemberList(group int64) []GroupMember {
	raw := goString(C.CQ_getGroupMemberList(C.int64_t(group)))
	list, err := UnpackGroupMemberList(raw)
	if err != nil {
		panic(fmt.Errorf("cqp: 内部错误，酷Q返回的群成员列表格式不正确: %v", err))
	}
	return list
}

// GetImage 获取图片
// 参数为CQ码内容，返回值为图片的文件路径
func GetImage(file string) string {
	return goString(C.CQ_getImage(cString(file)))
}

//GetLoginNick 获取登录号昵称
func GetLoginNick() string {
	return goString(C.CQ_getLoginNick())
}

//GetLoginQQ 获取登陆号QQ
func GetLoginQQ() int64 {
	return int64(C.CQ_getLoginQQ())
}

// GetRecord 获取语音
// file参数为CQ码内容，format为插件所需格式，返回值应该是文件路径
//
// 需要严格授权
func GetRecord(file, format string) string {
	return goString(C.CQ_getRecord(
		cString(file), cString(format),
	))
}

// GetRecordV2 获取语音
// 应该同GetRecord
//
// 需要严格授权
func GetRecordV2(file, format string) string {
	return goString(C.CQ_getRecordV2(
		cString(file), cString(format),
	))
}

// GetStrangerInfo 获取陌生人信息
// noCatch指定是否使用缓存
func GetStrangerInfo(qq int64, noCatch bool) string {
	return goString(C.CQ_getStrangerInfo(
		C.int64_t(qq), cBool(noCatch),
	))
}

// SendDiscussMsg 发送讨论组消息
func SendDiscussMsg(discuss int64, msg string) int32 {
	return int32(C.CQ_sendDiscussMsg(
		C.int64_t(discuss), cString(msg),
	))
}

// SendGroupMsg 发送群消息
func SendGroupMsg(group int64, msg string) int32 {
	return int32(C.CQ_sendGroupMsg(
		C.int64_t(group), cString(msg),
	))
}

// SendLike 发送赞
func SendLike(qq int64) int32 {
	return int32(C.CQ_sendLike(C.int64_t(qq)))
}

// SendLike2 发送赞2
// times指定赞的次数
func SendLike2(qq int64, times int32) int32 {
	return int32(C.CQ_sendLikeV2(
		C.int64_t(qq), C.int32_t(times),
	))
}

// SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 {
	return int32(C.CQ_sendPrivateMsg(
		C.int64_t(qq), cString(msg),
	))
}

// SetDiscussLeave 退出讨论组
func SetDiscussLeave(discuss int64) int32 {
	return int32(C.CQ_setDiscussLeave(C.int64_t(discuss)))
}

//// SetFatal 未知函数，请不要调用
//func SetFatal(errmsg string) int32 {
//	return int32(C.CQ_setFatal(cString(errmsg)))
//}

// SetFriendAddRequest 响应好友申请
// ReqFeedback 请传入好友请求事件(FriendRequest)中收到的responseFlag。
// FeedbackType 是否同意请求，同意:1，拒绝:2。
// remark 好友备注
func SetFriendAddRequest(ReqFeedback string, FeedbackType int32, remark string) int32 {
	return int32(C.CQ_setFriendAddRequest(
		cString(ReqFeedback),
		C.int32_t(FeedbackType),
		cString(remark),
	))
}

func SetGroupAddRequest(ReqFeedback string, ReqType, FeedbackType int32) int32 {
	return int32(C.CQ_setGroupAddRequest(
		cString(ReqFeedback),
		C.int32_t(ReqType), C.int32_t(FeedbackType),
	))
}

func SetGroupAddRequest2(ReqFeedback string, ReqType, FeedbackType int32, reason string) int32 {
	return int32(C.CQ_setGroupAddRequestV2(
		cString(ReqFeedback),
		C.int32_t(ReqType), C.int32_t(FeedbackType),
		cString(reason),
	))
}

//SetGroupAdmin 设置群管理员
func SetGroupAdmin(group, qq int64, admin bool) int32 {
	return int32(C.CQ_setGroupAdmin(
		C.int64_t(group),
		C.int64_t(qq),
		cBool(admin),
	))
}

//SetGroupAnonymous 设置群匿名是否开启
func SetGroupAnonymous(group int64, anonymous bool) int32 {
	return int32(C.CQ_setGroupAnonymous(
		C.int64_t(group),
		cBool(anonymous),
	))
}

//SetGroupAnonymousBan 设置群匿名成员禁言
func SetGroupAnonymousBan(group int64, anonymous string, time int64) int32 {
	return int32(C.CQ_setGroupAnonymousBan(
		C.int64_t(group),
		cString(anonymous),
		C.int64_t(time),
	))
}

//SetGroupBan 设置群成员禁言
func SetGroupBan(group, qq, bantime int64) int32 {
	return int32(C.CQ_setGroupBan(
		C.int64_t(group),
		C.int64_t(qq),
		C.int64_t(bantime),
	))
}

//SetGroupCard 设置群成员名片
func SetGroupCard(group, qq int64, card string) int32 {
	return int32(C.CQ_setGroupCard(
		C.int64_t(group),
		C.int64_t(qq),
		cString(card),
	))
}

//SetGroupKick 将群成员踢出群聊
func SetGroupKick(group, qq int64, rej bool) int32 {
	return int32(C.CQ_setGroupKick(
		C.int64_t(group),
		C.int64_t(qq),
		cBool(rej),
	))
}

// SetGroupLeave 退出群聊
func SetGroupLeave(group int64, dissolve bool) int32 {
	return int32(C.CQ_setGroupLeave(
		C.int64_t(group), cBool(dissolve),
	))
}

//SetGroupSpecialTitle 设置群成员头衔
func SetGroupSpecialTitle(group, qq int64, title string, timeout int64) int32 {
	return int32(C.CQ_setGroupSpecialTitle(
		C.int64_t(group),
		C.int64_t(qq),
		cString(title),
		C.int64_t(timeout),
	))
}

// SetGroupWholeBan 设置全员禁言
func SetGroupWholeBan(group int64, ban bool) int32 {
	return int32(C.CQ_setGroupWholeBan(
		C.int64_t(group), cBool(ban),
	))
}

// SetRestart 未知函数，请不要调用
func SetRestart() int32 {
	return int32(C.CQ_setRestart())
}
