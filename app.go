// 空插件
package main

import (
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

//go:generate cqcfg -c .
// cqp: 名称: 插件名称
// cqp: 版本: 1.0.0:1
// cqp: 作者: 插件作者姓名
// cqp: 简介: 您插件的简介
func main() { cqp.Main() }

func init() {
	// AppID 需要修改为你的插件的AppID
	cqp.AppID = "your.app.id"
	cqp.Enable = func() int32 {
		return 0
	}
}
