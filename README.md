# CoolQ Golang SDK
这是酷Q插件原生Go语言SDK  
[![GoDoc](https://img.shields.io/badge/pkg-documents-blue.svg)](https://pkg.go.dev/github.com/Tnze/CoolQ-Golang-SDK/cqp?tab=doc)

导入路径：`github.com/Tnze/CoolQ-Golang-SDK/cqp`。

把Go代码编译成dll，而不是基于http接口的网络调用，提高程序运行效率。  
如果使用中遇到问题，请大胆提issue**喜欢也请Star支持哦** 

特色：工具分析源码，自动生成app.json

## 关于SDK版本的说明

本SDK使用了最新的go modules机制用于管理版本，并且由于一次取群成员列表的API更新，
升级了主版本号到v2，但是由于导入路径变更导致了很多麻烦。

直到某一天（就是我写这段字的今天），Tnze终于受不了了，他决定将原本准备发布到v3的更新，
直接推到v1版本，这会导致之前用v1版本SDK编写，并且使用了原始取成员列表接口的插件
（很可能不存在），
更新后会无法通过编译，在这里说一声抱歉。请更新代码或者不要更新SDK。

今后的SDK更新都会在v1下进行，v2则不再维护。即v1比v2更新，请新项目不要使用v2。

## 推荐使用方法
1. 使用[模板](https://github.com/Tnze/CoolQ-Golang-Plugin)创建你的插件👉[![UseTemplate](https://img.shields.io/badge/-Use_Template-success)](https://github.com/Tnze/CoolQ-Golang-Plugin/generate)
1. 查看模板的README
1. 参考[官方教程](https://d.cqp.me/Pro/开发/快速入门)学习插件调试、打包等方法

## 编译方法

以下两个步骤独立，互不干扰。

### 1. 生成`app.json`文件

```batch
# 安装cqcfg，请确保`$GOBIN`在当前`PATH`环境变量中
go get github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg
# 查看cqcfg是否安装完成
cqcfg -v
# 运行
go generate
```

### 2. 生成`app.dll`文件

所需环境变量

```batch
set CGO_LDFLAGS=-Wl,--kill-at
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=386
```

编译

```batch
go build -ldflags "-s -w" -buildmode=c-shared -o app.dll
```

最后将dll和json复制到酷Q的dev路径下运行、调试和打包([详情](https://docs.cqp.im/dev/v9/getting-started/))
