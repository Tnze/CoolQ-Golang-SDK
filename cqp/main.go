package cqp

type Priority int32

// Log优先级（priority, AddLog的第一个参数）
const (
	Debug       Priority = 0
	Info                 = 10
	InfoSuccess          = 11
	InfoRecv             = 12
	InfoSend             = 13
	Warning              = 20
	Error                = 30
	Fatal                = 40
)

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

// GroupRequest 在机器人收到群请求时被调用
// subType为请求子类型, 可选的值有, 1：加群请求 2: 登录号被邀请入群
// msg 加群请求的验证问答, 若subType为2时通常为空字符串
// responseFlag 请求flag, 在调用处理请求api时需传入
// PC端在线时酷Q不会收到群请求
var GroupRequest func(subType, sendTime int32, fromGroup, fromQQ int64, msg, responseFlag string) int32
