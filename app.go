// 空插件
package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/v3/cqp"
)

func main() { cqp.Main() }

func init() {
	// AppID 需要修改为你的插件的AppID
	cqp.AppID = "your.app.id"
	cqp.PrivateMsg = onPrivateMsg
	cqp.Enable = func() int32 {
		return 0
	}
}

// TODO: 恢复空插件
func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return 0
}
