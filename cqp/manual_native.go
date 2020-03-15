// +build !websocket

package cqp

import (
	"fmt"
	"runtime/debug"
)

func Main() {}

// GetFriendList 获取好友列表
func GetFriendList() []FriendInfo {
	raw := getRawFriendList(false)
	list, err := UnpackFriendList(raw)
	if err != nil {
		panic(unpackError{API: "好友列表", Raw: raw, Err: err})
	}
	return list
}

func GetGroupInfo(group int64, noCache bool) GroupDetail {
	raw := getRawGroupInfo(group, noCache)
	info, err := UnpackGroupInfo(raw)
	if err != nil {
		panic(unpackError{API: "群信息", Raw: raw, Err: err})
	}
	return info
}

// GetGroupList 获取群列表
func GetGroupList() []GroupInfo {
	raw := getRawGroupList()
	list, err := UnpackGroupList(raw)
	if err != nil {
		panic(unpackError{API: "群列表", Raw: raw, Err: err})
	}
	return list
}

// GetGroupMemberInfo 获取群成员信息
func GetGroupMemberInfo(group, qq int64, noCache bool) GroupMember {
	raw := getRawGroupMemberInfoV2(group, qq, noCache)
	member, err := UnpackGroupMemberInfo(raw)
	if err != nil {
		panic(unpackError{API: "群成员信息", Raw: raw, Err: err})
	}
	return member
}

// GetGroupMemberList 获取群成员列表
func GetGroupMemberList(group int64) []GroupMember {
	raw := getRawGroupMemberList(group)
	list, err := UnpackGroupMemberList(raw)
	if err != nil {
		panic(unpackError{API: "群成员列表", Raw: raw, Err: err})
	}
	return list
}

// GetStrangerInfo 获取陌生人信息
// noCache指定是否使用缓存
func GetStrangerInfo(qq int64, noCache bool) StrangerInfo {
	raw := getRawStrangerInfo(qq, noCache)
	info, err := UnpackStrangerInfo(raw)
	if err != nil {
		panic(unpackError{API: "陌生人信息", Raw: raw, Err: err})
	}
	return info
}

// unpackError 当解码酷Q返回的数据出错时可能会被某些API返回
type unpackError struct {
	Err error
	API string
	Raw string
}

func (u *unpackError) Error() string {
	return "cqp: 内部错误，酷Q返回的" + u.API + "格式不正确: " + u.Err.Error()
}

func (u *unpackError) Unwrap() error {
	return u.Err
}

// 捕获panic并调用AddLog(Fatal)
func panicToFatal() {
	if v := recover(); v != nil {
		// 在这里调用debug.Stack()获取调用栈
		AddLog(Fatal, "panic", fmt.Sprintf("%v\n%s", v, debug.Stack()))
	}
}
