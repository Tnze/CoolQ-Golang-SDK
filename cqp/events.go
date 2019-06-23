// +build windows,386,cgo

package cqp

// #include "cq.h"
import "C"

var _AppInfo = C.CString("9," + AppID)

//export _appinfo
func _appinfo() *C.char { return _AppInfo }

//export _on_enable
func _on_enable() int32 { return Enable() }

//export _on_disable
func _on_disable() int32 { return 0 }

//export _on_private_msg
func _on_private_msg(subType, msgID int32, fromQQ int64, msg *C.char, font int32) int32 {
	return PrivateMsg(subType, msgID, fromQQ, C.GoString(msg), font)
}
