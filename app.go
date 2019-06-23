// 您的应用的代码写在这里
package main

import (
)

//AppID 需要修改为你的插件的appid
const AppID = "your.app.id"

//Enable 在插件启动时被调用一次
func Enable() int32 {
	AddLog(0, "Debug", "Test logggggger!!!")
	return 0
}
