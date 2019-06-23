package cqp

// #include "cq.h"
import "C"

//export _appinfo
func _appinfo() *C.char { return C.CString("9," + AppID) }

//export _on_enable
func _on_enable() int32 {
	if Enable == nil {
		return 0
	}
	return Enable()
}

//export _on_disable
func _on_disable() int32 {
	if Disable == nil {
		return 0
	}
	return Disable()
}

//export _on_private_msg
func _on_private_msg(subType, msgID int32, fromQQ int64, msg *C.char, font int32) int32 {
	if PrivateMsg == nil {
		return 0
	}
	return PrivateMsg(subType, msgID, fromQQ, goString(msg), font)
}

var (
	// AppID 插件的AppID。
	// 务必在init函数内为这个变量赋值，酷Q加载插件时会读取这个值。
	// 为了保证其唯一性，酷Q定义了AppID的规则，即开发者域名反写.应用英文名。
	// AppID中仅允许数字、字母、短横线（-）、下划线（_）、点(.)，
	// 不允许出现其他字符（如空格等），同时其中域名反写部分的字母全部统一使用小写字母。
	AppID string

	// Enable 在插件启动时被调用
	Enable func() int32

	// Disable 在插件禁用时被调用
	Disable func() int32

	// PrivateMsg 在收到私聊消息时被调用。
	// subType为子类型，可选的值有，11:来自好友 1:来自在线状态 2:来自群 3:来自讨论组。
	// 若返回非0值，消息将被拦截，最高优先不可拦截。
	PrivateMsg func(subType, msgID int32, fromQQ int64, msg string, font int32) int32
)
