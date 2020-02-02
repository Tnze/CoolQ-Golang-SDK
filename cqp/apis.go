package cqp

// AddLog 增加运行日志
// 	priority 是Log的优先级，请使用cqp预定义好的值。
// 	logType 是日志类型，酷Q日志窗口将将其显示在日志本身的前一列。
// 	reason 是日志内容
func AddLog(p Priority, logType, reason string) int32 { return addLog(int32(p), logType, reason) }

// CanSendImage 能否发送图片
func CanSendImage() bool { return canSendImage() }

// CanSendRecord 能否发送语音
func CanSendRecord() bool { return canSendRecord() }

// DeleteMsg 撤回消息
func DeleteMsg(msgID int64) int32 { return deleteMsg(msgID) }

// GetAppDir 取应用目录
// 返回的路径末尾带"\"，一般用法如下，所以不用关心是否有斜线
// 	os.Open(filepath.Join(cqp.GetAppDir(), "data.db"))
func GetAppDir() string { return getAppDirectory() }

// GetCookies 获取cookies
//
// 需要严格授权
func GetCookies(domain string) string { return getCookiesV2(domain) }

// GetCSRFToken 获取CSRF Token
//
// 需要严格授权
func GetCSRFToken() int32 { return getCsrfToken() }

// GetImage 获取图片
// 参数为CQ码内容，返回值为图片的文件路径
func GetImage(file string) string { return getImage(file) }

//GetLoginNick 获取登录号昵称
func GetLoginNick() string { return getLoginNick() }

//GetLoginQQ 获取登陆号QQ
func GetLoginQQ() int64 { return getLoginQQ() }

// GetRecord 获取语音
// file参数为CQ码内容，format为插件所需格式，返回语音文件绝对路径
func GetRecord(file, format string) string { return getRecordV2(file, format) }

// SendDiscussMsg 发送讨论组消息
func SendDiscussMsg(discuss int64, msg string) int32 { return sendDiscussMsg(discuss, msg) }

// SendGroupMsg 发送群消息
func SendGroupMsg(group int64, msg string) int32 { return sendGroupMsg(group, msg) }

// SendLike 发送赞
// times指定赞的次数
func SendLike(qq int64, times int32) int32 { return sendLikeV2(qq, times) }

// SendPrivateMsg 发送私聊消息
func SendPrivateMsg(qq int64, msg string) int32 { return sendPrivateMsg(qq, msg) }

// SetDiscussLeave 退出讨论组
func SetDiscussLeave(discuss int64) int32 { return setDiscussLeave(discuss) }

// SetFriendAddRequest 响应好友申请
// code 请传入好友请求事件(FriendRequest)中收到的responseFlag。
// Type 是否同意请求，同意:1，拒绝:2。
// remark 好友备注
func SetFriendAddRequest(code string, Type int32, remark string) int32 {
	return setFriendAddRequest(code, Type, remark)
}
func SetGroupAddRequest(ReqFeedback string, ReqType, FeedbackType int32, reason string) int32 {
	return setGroupAddRequestV2(ReqFeedback, ReqType, FeedbackType, reason)
}

//SetGroupAdmin 设置群管理员
func SetGroupAdmin(group, qq int64, admin bool) int32 { return setGroupAdmin(group, qq, admin) }

//SetGroupAnonymous 设置群匿名是否开启
func SetGroupAnonymous(group int64, anonymous bool) int32 { return setGroupAnonymous(group, anonymous) }

//SetGroupAnonymousBan 设置群匿名成员禁言
func SetGroupAnonymousBan(group int64, anonymous string, time int64) int32 {
	return setGroupAnonymousBan(group, anonymous, time)
}

//SetGroupBan 设置群成员禁言
func SetGroupBan(group, qq, bantime int64) int32 { return setGroupBan(group, qq, bantime) }

//SetGroupCard 设置群成员名片
func SetGroupCard(group, qq int64, card string) int32 { return setGroupCard(group, qq, card) }

//SetGroupKick 将群成员踢出群聊
func SetGroupKick(group, qq int64, rej bool) int32 { return setGroupKick(group, qq, rej) }

// SetGroupLeave 退出群聊
func SetGroupLeave(group int64, dissolve bool) int32 { return setGroupLeave(group, dissolve) }

//SetGroupSpecialTitle 设置群成员头衔
func SetGroupSpecialTitle(group, qq int64, title string, timeout int64) int32 {
	return setGroupSpecialTitle(group, qq, title, timeout)
}

// SetGroupWholeBan 设置全员禁言
func SetGroupWholeBan(group int64, ban bool) int32 { return setGroupWholeBan(group, ban) }
