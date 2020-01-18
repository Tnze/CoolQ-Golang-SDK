// +build windows,386,cgo,!websocket

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
	defer panicToFatal()
	return Enable()
}

//export _on_disable
func _on_disable() int32 {
	if Disable == nil {
		return 0
	}
	defer panicToFatal()
	return Disable()
}

//export _on_start
func _on_start() int32 {
	if Start == nil {
		return 0
	}
	defer panicToFatal()
	return Start()
}

//export _on_exit
func _on_exit() int32 {
	if Exit == nil {
		return 0
	}
	defer panicToFatal()
	return Exit()
}

//export _on_private_msg
func _on_private_msg(subType, msgID int32, fromQQ int64, msg *C.char, font int32) int32 {
	if PrivateMsg == nil {
		return 0
	}
	defer panicToFatal()
	return PrivateMsg(subType, msgID, fromQQ, goString(msg), font)
}

//export _on_group_msg
func _on_group_msg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg *C.char, font int32) int32 {
	if GroupMsg == nil {
		return 0
	}
	defer panicToFatal()
	return GroupMsg(subType, msgID, fromGroup, fromQQ, goString(fromAnonymous), goString(msg), font)
}

//export _on_discuss_msg
func _on_discuss_msg(sub_type, msg_id int32, from_discuss, from_qq int64, msg *C.char, font int32) int32 {
	if DiscussMsg == nil {
		return 0
	}
	defer panicToFatal()
	return DiscussMsg(sub_type, msg_id, from_discuss, from_qq, goString(msg), font)
}

//export _on_group_upload
func _on_group_upload(sub_type, send_time int32, from_group, from_qq int64, file *C.char) int32 {
	if GroupUpload == nil {
		return 0
	}
	defer panicToFatal()
	return GroupUpload(sub_type, send_time, from_group, from_qq, goString(file))
}

//export _on_group_admin
func _on_group_admin(sub_type, send_time int32, from_group, being_operate_qq int64) int32 {
	if GroupAdmin == nil {
		return 0
	}
	defer panicToFatal()
	return GroupAdmin(sub_type, send_time, from_group, being_operate_qq)
}

//export _on_group_member_decrease
func _on_group_member_decrease(sub_type, send_time int32, from_group, from_qq, being_operate_qq int64) int32 {
	if GroupMemberDecrease == nil {
		return 0
	}
	defer panicToFatal()
	return GroupMemberDecrease(sub_type, send_time, from_group, from_qq, being_operate_qq)
}

//export _on_group_member_increase
func _on_group_member_increase(sub_type, send_time int32, from_group, from_qq, being_operate_qq int64) int32 {
	if GroupMemberIncrease == nil {
		return 0
	}
	defer panicToFatal()
	return GroupMemberIncrease(sub_type, send_time, from_group, from_qq, being_operate_qq)
}

//export _on_friend_add
func _on_friend_add(sub_type, send_time int32, from_qq int64) int32 {
	if FriendAdd == nil {
		return 0
	}
	defer panicToFatal()
	return FriendAdd(sub_type, send_time, from_qq)
}

//export _on_friend_request
func _on_friend_request(sub_type, send_time int32, from_qq int64, msg, response_flag *C.char) int32 {
	if FriendRequest == nil {
		return 0
	}
	defer panicToFatal()
	return FriendRequest(sub_type, send_time, from_qq, goString(msg), goString(response_flag))
}

//export _on_group_request
func _on_group_request(sub_type, send_time int32, from_group, from_qq int64, msg, response_flag *C.char) int32 {
	if GroupRequest == nil {
		return 0
	}
	defer panicToFatal()
	return GroupRequest(sub_type, send_time, from_group, from_qq, goString(msg), goString(response_flag))
}

// 捕获panic并调用AddLog(Fatal)
func panicToFatal() {
	if v := recover(); v != nil {
		// 在这里调用debug.Stack()获取调用栈
		AddLog(Fatal, "panic", fmt.Sprintf("%v\n%s", v, debug.Stack()))
	}
}
