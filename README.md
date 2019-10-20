# CoolQ Golang SDK
这是一个Native 酷Q插件 Go语言SDK  
[![GoDoc](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp?status.svg)](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp)

把Go代码编译成dll，而不是基于http接口的网络调用，大大提高程序运行效率。  
如果使用中遇到问题，请大胆提issue**喜欢也请Star支持哦** 

工具自动分析源码，无需手动编写app.json

## 食用方法
1. 使用[模板](https://github.com/Tnze/CoolQ-Golang-Plugin)创建你的插件👉[![UseTemplate](https://img.shields.io/badge/-Use_Template-success)](https://github.com/Tnze/CoolQ-Golang-Plugin/generate)
1. 查看模板的README
1. 参考[官方教程](https://d.cqp.me/Pro/开发/快速入门)学习插件调试、打包等方法

## 代码
```go
package main

import "github.com/Tnze/CoolQ-Golang-SDK/cqp"

//go:generate cqcfg -c .
// cqp: 名称: <插件名>
// cqp: 版本: 1.0.0:0
// cqp: 作者: <作者>
// cqp: 简介: <简介>
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "me.cqp.tnze.demo" // TODO: 修改为这个插件的ID
	cqp.PrivateMsg = onPrivateMsg
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return 0
}
```