# CoolQ Golang SDK
这是酷Q插件原生Go语言SDK  
[![GoDoc](https://img.shields.io/badge/pkg-documents-blue.svg)](https://pkg.go.dev/github.com/Tnze/CoolQ-Golang-SDK/cqp?tab=doc)

导入路径：`github.com/Tnze/CoolQ-Golang-SDK/cqp`。

特点：原生SDK，直接编译成dll运行；工具分析源码，自动生成app.json

如果使用中遇到问题，请大胆提issue（Q群:304279325）**喜欢也请Star支持哦** 

## 关于SDK版本的说明

本SDK使用了最新的go modules机制用于管理版本，并且由于一次取群成员列表的API更新，
升级了主版本号到v2，但是由于导入路径变更导致了很多麻烦。

直到某一天（就是我写这段字的今天），Tnze终于受不了了，他决定将原本准备发布到v3的更新，
直接推到v1版本，这会导致之前用v1版本SDK编写，并且使用了原始取成员列表接口的插件
（很可能不存在），
更新后会无法通过编译，在这里说一声抱歉。请更新代码或者不要更新SDK。

今后的SDK更新都会在v1下进行，v2则不再维护。即v1比v2更新，请**新项目不要使用v2**。

## 使用方法
1. 使用[模板](https://github.com/Tnze/CoolQ-Golang-Plugin)创建你的插件👉[![UseTemplate](https://img.shields.io/badge/-Use_Template-success)](https://github.com/Tnze/CoolQ-Golang-Plugin/generate)
1. 查看模板的README
1. 参考[官方教程](https://docs.cqp.im/dev/v9/getting-started/)学习插件调试、打包等方法

## 编译

把Go编译为能让酷Q加载的dll动态库比较复杂，并不是一条`go build`命令就能简单做到的。
所以在上述模板项目中我编写了脚本"build.bat"，方便大家一键编译。
但是你仍然**得自己装好go和gcc**（要用CGO）。

成功编译后将dll和json复制到酷Q的dev的子文件夹下即可运行([详情](https://docs.cqp.im/dev/v9/getting-started/))
