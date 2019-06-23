// 您的应用的代码写在这里
package main

import "github.com/Tnze/CoolQ-Golang-SDK/cqp"

// AppID 需要修改为你的插件的appid
const AppID = "your.app.id"

// Enable 在插件启动时被调用一次
func Enable() int32 {
	return 0
}

// PrivateMsg 在收到私聊消息时被调用
// - subType	子类型，11/来自好友 1/来自在线状态 2/来自群 3/来自讨论组
// - msgId		消息ID
// - fromQQ		来源QQ
// - msg		消息内容
// - font		字体
// 返回非零值,消息将被拦截,最高优先不可拦截
func PrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {

}

func main(){
	
}