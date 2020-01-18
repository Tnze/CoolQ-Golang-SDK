// 空插件
package main

import "github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"

func main() { cqp.Main() }

func init() {
	// AppID 需要修改为你的插件的AppID
	cqp.AppID = "your.app.id"
}
