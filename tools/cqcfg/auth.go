package main

/* //20,  //[敏感]取Cookies	getCookies / getCsrfToken
   //30,  //接收语音			getRecord
   101,  //发送群消息			sendGroupMsg
   103,  //发送讨论组消息		sendDiscussMsg
   106,  //发送私聊消息		sendPrivateMsg
   //110,  //[敏感]发送赞				sendLike
   120,  //置群员移除			setGroupKick
   121,  //置群员禁言			setGroupBan
   122,  //置群管理员			setGroupAdmin
   123,  //置全群禁言			setGroupWholeBan
   124,  //置匿名群员禁言		setGroupAnonymousBan
   125,  //置群匿名设置		setGroupAnonymous
   126,  //置群成员名片		setGroupCard
   //127, //[敏感]置群退出		setGroupLeave
   128,  //置群成员专属头衔	setGroupSpecialTitle
   130,  //取群成员信息		getGroupMemberInfoV2 / getGroupMemberInfo
   131,  //取陌生人信息		getStrangerInfo
   140,  //置讨论组退出		setDiscussLeave
   150,  //置好友添加请求		setFriendAddRequest
   151,  //置群添加请求		setGroupAddRequest
   160,  //取群成员列表		getGroupMemberList
   161,  //取群列表			getGroupList
   180   //撤回消息			deleteMsg*/

var authcode = map[string]int{
	"GetCookies":           20,
	"GetCookiesV2":         20,
	"GetRecord":            30,
	"SendGroupMsg":         101,
	"SendDiscussMsg":       103,
	"SendPrivateMsg":       106,
	"SendLike":             110,
	"SendLikeV2":           110,
	"SetGroupKick":         120,
	"SetGroupBan":          121,
	"SetGroupAdmin":        122,
	"SetGroupWholeBan":     123,
	"SetGroupAnonymousBan": 124,
	"SetGroupAnonymous":    125,
	"SetGroupCard":         126,
	"SetGroupLeave":        127,
	"SetGroupSpecialTitle": 128,
	"GetGroupMemberInfo":   130,
	"GetGroupMemberInfoV2": 130,
	"GetStrangerInfo":      131,
	"SetDiscussLeave":      140,
	"SetFriendAddRequest":  150,
	"SetGroupAddRequest":   151,
	"GetGroupMemberList":   160,
	"GetGroupList":         161,
	"DeleteMsg":            180,
}

func onCallAPI(calls map[string]int) {
	// 这里去除了重复值
	authCodes := make(map[int]bool)
	for i, v := range calls { //添加权限声明
		if code, ok := authcode[i]; ok && v > 0 {
			authCodes[code] = true
		}
	}
	for code, set := range authCodes {
		if set {
			info.Auth = append(info.Auth, code)
		}
	}
}
