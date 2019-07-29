# CoolQ Golang SDK
这是一个Native 酷Q插件 Go语言SDK  

**可以点下边这个按钮看文档**  
[![GoDoc](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp?status.svg)](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp)

通过直接把Go代码编译成dll，省去从前基于http接口的Go语言SDK的网络调用过程，大大提高程序运行效率。  
由于制作该SDK工作量较大，部分API和EVENT没有测试，该项目将会继续完善，如果使用中遇到问题，请大胆提issue～

**喜欢要记得Star哦** 

# 那么怎么用这个SDK写自己的插件呢？
> 戳下面这个`模板`
1. 使用[模板](https://github.com/Tnze/CoolQ-Golang-Plugin)创建你的插件
5. 插件调试、打包等[方法](https://d.cqp.me/Pro/开发/快速入门)与其他SDK相同

## 重磅推出: cqcfg工具
通过读取源码，自动生成app.json!
~~详情使用方法请见->[![GoDoc](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg?status.svg)](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg)~~  
现在插件模板内的脚本会自动调用本工具。

TODO list:
- [x] 成功编译成dll
- [x] 导出函数供酷Q调用
- [x] 调用酷Q提供的函数
- [x] 编写使用文档
- [x] 导出全部API
- [x] 导出全部Event
- [ ] 编写更详细的注释
- [x] 解析群成员列表和群成员信息