package main

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
	"GetGroupInfo":		132,
	"SetDiscussLeave":      140,
	"SetFriendAddRequest":  150,
	"SetGroupAddRequest":   151,
	"GetGroupMemberList":   160,
	"GetGroupList":         161,
	"GetFriendList":	162,
	"DeleteMsg":            180,
}

// addAuth 往info里为调用过的cqp.*函数填写权限
func addAuth(calls map[string]int) {
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
