// +build windows,386,cgo

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

//AddLog 增加运行日志
func AddLog(priority int32, logtype, reason string) int32 {
	return int32(C.CQ_addLog(
		C.int32_t(priority),
		cString(logtype),
		cString(reason),
	))
}

//SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 {
	return int32(C.CQ_sendPrivateMsg( 
		C.int64_t(qq), cString(msg),
	))
}

//SendGroupMsg 发送群消息
func SendGroupMsg(group int64, msg string) int32 {
	return int32(C.CQ_sendGroupMsg(
		C.int64_t(group), cString(msg),
	))
}

//SendDiscussMsg 发送讨论组消息
func SendDiscussMsg(discuss int64, msg string) int32 {
	return int32(C.CQ_sendDiscussMsg(
		C.int64_t(discuss), cString(msg),
	))
}

//SendLike 发送赞
func SendLike(qq int64) int32 {
	return int32(C.CQ_sendLike(C.int64_t(qq)))
}

//SendLike2 发送赞2
func SendLike2(qq int64, times int32) int32 {
	return int32(C.CQ_sendLikeV2(
		C.int64_t(qq), C.int32_t(times),
	))
}

//GetCookies 获取cookies
func GetCookies() string {
	return goString(C.CQ_getCookies())
}
