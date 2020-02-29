# CoolQ Golang SDK
这是酷Q插件原生Go语言SDK  
[![GoDoc](https://img.shields.io/badge/pkg-documents-blue.svg)](https://pkg.go.dev/github.com/Tnze/CoolQ-Golang-SDK/cqp?tab=doc)

导入路径：`github.com/Tnze/CoolQ-Golang-SDK/cqp`。

把Go代码编译成dll，而不是基于http接口的网络调用，提高程序运行效率。  
如果使用中遇到问题，请大胆提issue**喜欢也请Star支持哦** 

特色：工具分析源码，自动生成app.json

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
