package cqp

// #include "events.h"
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

//export _on_start
func _on_start() int32 {
	return 0
}

//export _on_exit
func _on_exit() int32 {
	return 0
}

//export _on_private_msg
func _on_private_msg(subType, msgID int32, fromQQ int64, msg *C.char, font int32) int32 {
	if PrivateMsg == nil {
		return 0
	}
	return PrivateMsg(subType, msgID, fromQQ, goString(msg), font)
}

//export _on_group_msg
func _on_group_msg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg *C.char, font int32) int32 {
	if GroupMsg == nil {
		return 0
	}
	return GroupMsg(subType, msgID, fromGroup, fromQQ, goString(fromAnonymous), goString(msg), font)
}

//export _on_discuss_msg
func _on_discuss_msg(sub_type, msg_id int32, from_discuss, from_qq int64, msg *C.char, font int32) int32 {
	if DiscussMsg == nil {
		return 0
	}
	return DiscussMsg(sub_type, msg_id, from_discuss, from_qq, goString(msg), font)
}

//export _on_group_upload
func _on_group_upload(sub_type, send_time int32, from_group, from_qq int64, file *C.char) int32 {
	if GroupUpload == nil {
		return 0
	}
	return GroupUpload(sub_type, send_time, from_group, from_qq, goString(file))
}

//export _on_group_admin
func _on_group_admin(sub_type, send_time int32, from_group, being_operate_qq int64) int32 {
	if GroupAdmin == nil {
		return 0
	}
	return GroupAdmin(sub_type, send_time, from_group, being_operate_qq)
}

//export _on_group_member_decrease
func _on_group_member_decrease(sub_type, send_time int32, from_group, from_qq, being_operate_qq int64) int32 {
	if GroupMemberDecrease == nil {
		return 0
	}
	return GroupMemberDecrease(sub_type, send_time, from_group, from_qq, being_operate_qq)
}

//export _on_group_member_increase
func _on_group_member_increase(sub_type, send_time int32, from_group, from_qq, being_operate_qq int64) int32 {
	if GroupMemberIncrease == nil {
		return 0
	}
	return GroupMemberIncrease(sub_type, send_time, from_group, from_qq, being_operate_qq)
}

//export _on_friend_add
func _on_friend_add(sub_type, send_time int32, from_qq int64) int32 {
	if FriendAdd == nil {
		return 0
	}
	return FriendAdd(sub_type, send_time, from_qq)
}

//export _on_friend_request
func _on_friend_request(sub_type, send_time int32, from_qq int64, msg, response_flag *C.char) int32 {
	if FriendRequest == nil {
		return 0
	}
	return FriendRequest(sub_type, send_time, from_qq, goString(msg), goString(response_flag))
}

//export _on_group_request
func _on_group_request(sub_type, send_time int32, from_group, from_qq int64, msg, response_flag *C.char) int32 {
	if GroupRequest == nil {
		return 0
	}
	return GroupRequest(sub_type, send_time, from_group, from_qq, goString(msg), goString(response_flag))
}

// AppID 插件的AppID。
// 务必在init函数内为这个变量赋值，酷Q加载插件时会读取这个值。
// 为了保证其唯一性，酷Q定义了AppID的规则，即开发者域名反写.应用英文名。
// AppID中仅允许数字、字母、短横线（-）、下划线（_）、点(.)，
// 不允许出现其他字符（如空格等），同时其中域名反写部分的字母全部统一使用小写字母。
var AppID string

// Enable 在插件启动时被调用
var Enable func() int32

// Disable 在插件禁用时被调用
var Disable func() int32

var Start func() int32
var Exit func() int32

// PrivateMsg 在收到私聊消息时被调用。
// subType为子类型，可选的值有，11:来自好友 1:来自在线状态 2:来自群 3:来自讨论组。
// 若返回非0值，消息将被拦截，最高优先不可拦截。
var PrivateMsg func(subType, msgID int32, fromQQ int64, msg string, font int32) int32

// GroupMsg 在收到群聊消息时被调用
// subType目前固定为1
var GroupMsg func(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32

var DiscussMsg func(subType, msgID int32, fromDiscuss, fromQQ int64, msg string, font int32) int32
var GroupUpload func(subType, sendTime int32, fromGroup, fromQQ int64, file string) int32
var GroupAdmin func(subType, sendTime int32, fromGroup, QQ int64) int32
var GroupMemberDecrease func(subType, sendTime int32, fromGroup, fromQQ, beingOperateQQ int64) int32
var GroupMemberIncrease func(subType, sendTime int32, fromGroup, fromQQ, beingOperateQQ int64) int32
var FriendAdd func(subType, sendTime int32, fromQQ int64) int32
var FriendRequest func(subType, sendTime int32, fromQQ int64, msg, responseFlag string) int32
var GroupRequest func(subType, sendTime int32, fromGroup, fromQQ int64, msg, responseFlag string) int32
