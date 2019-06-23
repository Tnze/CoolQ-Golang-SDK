package cqp

// #include "cq.h"
import "C"
import sc "golang.org/x/text/encoding/simplifiedchinese"

//Log优先级（priority）,AddLog的第一个参数
const (
	Debug       = 0
	Info        = 10
	InfoSuccess = 11
	InfoRecv    = 12
	InfoSend    = 13
	Warning     = 20
	Error       = 30
	Fatal       = 40
)

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
func AddLog(priority int32, logtype, reason string) int32 {
	return int32(C.CQ_addLog(
		C.int32_t(priority),
		cString(logtype),
		cString(reason),
	))
}

// SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 {
	return int32(C.CQ_sendPrivateMsg(
		C.int64_t(qq), cString(msg),
	))
}

// SendGroupMsg 发送群消息
func SendGroupMsg(group int64, msg string) int32 {
	return int32(C.CQ_sendGroupMsg(
		C.int64_t(group), cString(msg),
	))
}

// SendDiscussMsg 发送讨论组消息
func SendDiscussMsg(discuss int64, msg string) int32 {
	return int32(C.CQ_sendDiscussMsg(
		C.int64_t(discuss), cString(msg),
	))
}

// SendLike 发送赞
func SendLike(qq int64) int32 {
	return int32(C.CQ_sendLike(C.int64_t(qq)))
}

// SendLike2 发送赞2
func SendLike2(qq int64, times int32) int32 {
	return int32(C.CQ_sendLikeV2(
		C.int64_t(qq), C.int32_t(times),
	))
}

// GetCookies 获取cookies
func GetCookies() string {
	return goString(C.CQ_getCookies())
}

// GetRecord 获取语音
// 需要严格授权
func GetRecord(file, format string) string {
	return goString(C.CQ_getRecord(
		cString(file), cString(format),
	))
}

// GetCSRFToken 获取CSRF Token
// 需要严格授权
func GetCSRFToken() int32 {
	return int32(C.CQ_getCsrfToken())
}

// GetAppDir 取应用目录
// 返回的路径末尾带"\"
func GetAppDir() string {
	return goString(C.CQ_getAppDirectory())
}

//GetLoginQQ 获取登陆号QQ
func GetLoginQQ() int64 {
	return int64(C.CQ_getLoginQQ())
}

//GetLoginNick 获取登录号昵称
func GetLoginNick() string {
	return goString(C.CQ_getLoginNick())
}

//SetGroupKick 将群成员踢出群聊
func SetGroupKick(group, qq int64, rej bool) int32 {
	return int32(C.CQ_setGroupKick(
		C.int64_t(group),
		C.int64_t(qq),
		cBool(rej),
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

//SetGroupAdmin 设置群管理员
func SetGroupAdmin(group, qq int64, admin bool) int32 {
	return int32(C.CQ_setGroupAdmin(
		C.int64_t(group),
		C.int64_t(qq),
		cBool(admin),
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

//SetGroupAnonymousBan 设置群匿名成员禁言
func SetGroupAnonymousBan(group int64, anonymous string, time int64) int32 {
	return int32(C.CQ_setGroupAnonymousBan(
		C.int64_t(group),
		cString(anonymous),
		C.int64_t(time),
	))
}

//SetGroupAnonymous 设置群匿名是否开启
func SetGroupAnonymous(group int64, anonymous bool) int32 {
	return int32(C.CQ_setGroupAnonymous(
		C.int64_t(group),
		cBool(anonymous),
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

// SetGroupLeave 退出群聊
func SetGroupLeave(group int64, dissolve bool) int32 {
	return int32(C.CQ_setGroupLeave(
		C.int64_t(group), cBool(dissolve),
	))
}

func SetDiscussLeave(discuss int64) int32 {
	return int32(C.CQ_setDiscussLeave(C.int64_t(discuss)))
}

func SetFriendAddRequest(ReqFeedback string, FeedbackType int32, remark string) int32 {
	return int32(C.CQ_setFriendAddRequest(
		cString(ReqFeedback),
		C.int32_t(FeedbackType),
		cString(remark),
	))
}

func SetGroupAddequest(ReqFeedback string, ReqType, FeedbackType int32) int32 {
	return int32(C.CQ_setGroupAddRequest(
		cString(ReqFeedback),
		C.int32_t(ReqType), C.int32_t(FeedbackType),
	))
}

func SetGroupAddequest2(ReqFeedback string, ReqType, FeedbackType int32, reason string) int32 {
	return int32(C.CQ_setGroupAddRequestV2(
		cString(ReqFeedback),
		C.int32_t(ReqType), C.int32_t(FeedbackType),
		cString(reason),
	))
}

func SetFatal(errmsg string) int32 {
	return int32(C.CQ_setFatal(cString(errmsg)))
}

func GetGroupMemberInfo(group, qq int64) string {
	return goString(C.CQ_getGroupMemberInfo(
		C.int64_t(group), C.int64_t(qq),
	))
}

func GetGroupMemberInfo2(group, qq int64, noCatch bool) string {
	return goString(C.CQ_getGroupMemberInfoV2(
		C.int64_t(group), C.int64_t(qq), cBool(noCatch),
	))
}

func GetStrangerInfo(qq int64, noCatch bool) string {
	return goString(C.CQ_getStrangerInfo(
		C.int64_t(qq), cBool(noCatch),
	))
}

func GetGroupMemberList(group int64) string {
	return goString(C.CQ_getGroupMemberList(C.int64_t(group)))
}

func GetGroupList() string {
	return goString(C.CQ_getGroupList())
}

func DeleteMsg(msgID int64) int32 {
	return int32(C.CQ_deleteMsg(C.int64_t(msgID)))
}
