# CoolQ Golang SDK
这是一个Native 酷Q插件 Go语言SDK  
[![GoDoc](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp?status.svg)](https://godoc.org/github.com/Tnze/CoolQ-Golang-SDK/cqp)

把Go代码编译成dll，而不是基于http接口的网络调用，大大提高程序运行效率。  
如果使用中遇到问题，请大胆提issue**喜欢也要记得Star哦** 

注意：本SDK作为一个库，编写插件并不是在本SDK基础上修改，而是……

## 开始
1. 使用[模板](https://github.com/Tnze/CoolQ-Golang-Plugin)创建你的插件👉[![UseTemplate](https://img.shields.io/badge/-Use_Template-success)](https://github.com/Tnze/CoolQ-Golang-Plugin/generate)
1. 阅读模板的README
1. 参考[官方教程](https://d.cqp.me/Pro/开发/快速入门)学习插件调试、打包等方法

## 自动生成app.json: 用cqcfg工具
一个分析源码，自动生成app.json的工具!  
插件模板集成了本工具。
