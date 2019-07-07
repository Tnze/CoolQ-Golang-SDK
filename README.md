# CoolQ Golang SDK
Native 酷Q插件 Go语言SDK  

**详细的文档！** [![GoDoc](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp?status.svg)](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp)

通过直接把Go代码编译成dll，省去从前Go语言SDK的网络调用过程，大大提高程序运行效率。  
该项目将会继续完善，完成一个简单易用的酷Q**Go语言SDK** 。

由于制作该SDK工作量较大，部分API和EVENT没有测试，如果使用中遇到问题，请大胆提issue～

> 喜欢要记得Star哦

TODO list:
- [x] 成功编译成dll
- [x] 导出函数供酷Q调用
- [x] 调用酷Q提供的函数
- [x] 编写使用文档
- [x] 导出全部API
- [x] 导出全部Event
- [ ] 编写更详细的注释
- [x] 解析群成员列表和群成员信息

# 食用方法

> 只要`import "github.com/Tnze/CoolQ-Golang-SDK/cqp"`并且使用恰当的编译命令即可成功编译出酷Q插件，`app.go`只是一个最简单的demo

1. 先clone该项目
2. 检查是否安装了[go语言编译器](https://golang.google.cn)：`go version`
3. 检查是否安装了[gcc编译器](http://tdm-gcc.tdragon.net)（cgo需要gcc编译器）：`gcc --version`
4. 运行`.\build.bat`编译，检查是否有生成*app.dll*
5. 将*app.dll*和*app.json*复制到酷Q目录下，检查是否能成功加载
6. 在*app.go* 内编写你的插件，然后重新编译、测试
7. 插件打包等方法与其他SDK相同

ps: 我不知道并发调用酷Q的API会发生什么，不知道酷Q内部有没有锁

## 重磅推出: cqcfg工具
自动生成app.json!
详情使用方法请见[![GoDoc](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg?status.svg)](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg)

## 调用顺序小记

API调用顺序：用户代码 -> Go函数 -> C函数 -> 酷Q函数指针  
例:           -> AddLog() -> CQ_addLog() -> CQ_addLog_Ptr

Event调用顺序：酷Q -> C函数 -> Go导出函数 -> Go函数  
例:           -> EVENT_ON_ENABLE() -> _on_enable() -> Enable()
