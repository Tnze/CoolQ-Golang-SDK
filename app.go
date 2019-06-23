// 您的应用的代码写在这里
package main

import "github.com/Tnze/CoolQ-Golang-SDK/cqp"

func main() {} //仅为编译需要，不要在这里写代码，它们不会被执行

func init() {
	// AppID 需要修改为你的插件的appid
	cqp.AppID = "your.app.id"
	cqp.Enable = Enable
}

func Enable() int32 {
	cqp.AddLog(cqp.Info, `\u4fe1\u606f`, "\u4fe1\u606f")//TODO 解决中文乱码的问题
	return 0
}
