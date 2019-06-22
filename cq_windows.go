package main

// #include <stdlib.h>
// #include "cq.h"
import "C"
import ()

func main() {}

var _AppInfo = C.CString("9," + AppID)

//export _appinfo
func _appinfo() *C.char { return _AppInfo }

//export _on_enable
func _on_enable() int32 { return Enable() }

//AddLog 增加运行日志
func AddLog(priority int32, logtype, reason string) int32 {
	return int32(C.CQ_addLog(
		C.int32_t(priority),
		C.CString(logtype),
		C.CString(reason),
	))
}

//SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 {
	return int32(C.CQ_sendPrivateMsg(
		C.int64_t(qq), C.CString(msg),
	))
}
