package main

import (
	"go/ast"
)

type event struct {
	Name     string `json:"name"`
	Function string `json:"function"`
	Type     int    `json:"type"`
	Priority int    `json:"priority"`
	ID       int    `json:"id"`
}

func onSetEvent(name string, rhs ast.Expr) { //记录事件注册
	switch name {
	case "Enable":
		info.Events = append(info.Events, event{
			ID:       1003,
			Type:     1003,
			Name:     "插件启用",
			Priority: 20000,
			Function: "EVENT_ON_ENABLE",
		})
	case "Disable":
		info.Events = append(info.Events, event{
			ID:       1004,
			Type:     1004,
			Name:     "插件停用",
			Priority: 20000,
			Function: "EVENT_ON_DISABLE",
		})
	case "Start":
		info.Events = append(info.Events, event{
			ID:       1001,
			Type:     1001,
			Name:     "酷Q启动",
			Priority: 20000,
			Function: "EVENT_ON_START",
		})
	case "Exit":
		info.Events = append(info.Events, event{
			ID:       1002,
			Type:     1002,
			Name:     "酷Q退出",
			Priority: 20000,
			Function: "EVENT_ON_EXIT",
		})
	case "PrivateMsg":
		info.Events = append(info.Events, event{
			ID:       1,
			Type:     21,
			Name:     "私聊消息",
			Function: "EVENT_ON_PRIVATE_MSG",
			Priority: 20000,
		})

	case "GroupMsg":
		info.Events = append(info.Events, event{
			ID:       2,
			Type:     2,
			Name:     "群消息",
			Function: "EVENT_ON_GROUP_MSG",
			Priority: 20000,
		})

	case "DiscussMsg":
		info.Events = append(info.Events, event{
			ID:       3,
			Type:     4,
			Name:     "讨论组消息",
			Function: "EVENT_ON_DISCUSS_MSG",
			Priority: 20000,
		})

	case "GroupUpload":
		info.Events = append(info.Events, event{
			ID:       4,
			Type:     11,
			Name:     "群文件上传",
			Function: "EVENT_ON_GROUP_UPLOAD",
			Priority: 20000,
		})
	case "GroupAdmin":
		info.Events = append(info.Events, event{
			ID:       5,
			Type:     101,
			Name:     "群管理员变动",
			Function: "EVENT_ON_GROUP_ADMIN",
			Priority: 20000,
		})
	case "GroupMemberDecrease":
		info.Events = append(info.Events, event{
			ID:       6,
			Type:     102,
			Name:     "群成员减少",
			Function: "EVENT_ON_GROUP_MEMBER_DECREASE",
			Priority: 20000,
		})
	case "GroupMemberIncrease":
		info.Events = append(info.Events, event{
			ID:       7,
			Type:     103,
			Name:     "群成员增加",
			Function: "EVENT_ON_GROUP_MEMBER_INCREASE",
			Priority: 20000,
		})
	case "FriendAdd":
		info.Events = append(info.Events, event{
			ID:       10,
			Type:     201,
			Name:     "好友添加",
			Function: "EVENT_ON_FRIEND_ADD",
			Priority: 20000,
		})
	case "FriendRequest":
		info.Events = append(info.Events, event{
			ID:       8,
			Type:     301,
			Name:     "加好友请求",
			Function: "EVENT_ON_FRIEND_REQUEST",
			Priority: 20000,
		})
	case "GroupRequest":
		info.Events = append(info.Events, event{
			ID:       9,
			Type:     302,
			Name:     "加群请求／邀请",
			Function: "EVENT_ON_GROUP_REQUEST",
			Priority: 20000,
		})
	}
}
