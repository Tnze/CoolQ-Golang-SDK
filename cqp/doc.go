// Package cqp 是用于编写酷Q插件的包，酷Q是一个QQ机器人软件。
//
// Usage 用法
//
// 通过导入本包，再使用特定的编译命令即可将go代码编译成dll，就可以作为插件被酷Q加载。
// 为了编译成一个dll，你的代码里需要一个空的main函数。
// 然后编写init函数，设置你的插件的AppID。
//
//	package main
//
// 	import "github.com/Tnze/CoolQ-Golang-SDK/cqp"
//
//	func main() {}
//	func init() {
//		// 设置AppID
//		cqp.AppID = "your.app.id"
//		// 注册事件
//		cqp.Enable = Enable
//	}
//
//	func Enable(){
//		// 当插件启用时被调用
//	}
//
// Build tools
//
// 当写完基本的代码之后可以将它编译成dll。
// 插件的编译需要在windows环境下进行，
// 需要go和gcc两种编译工具来分别编译go代码和c代码。
// 请检查你的go和gcc是否都安装完成：
//	go version
//	gcc --version
// 其中go工具可以从https://golang.google.cn 下载
// 没有gcc可以安装TDM-GCC http://tdm-gcc.tdragon.net/
//
// Generate app.json
//
// 推荐使用tools/cqcfg工具自动生成和更新app.json。
// 使用go get命令安装cqcfg工具，执行完毕后会在GOPATH或者GOBIN下生成可执行文件。
//	go get github.com/Tnze/CoolQ-Golang-SDK/tools/cqcfg
// 安装完毕后应当可以在命令提示符中使用cqcfg命令，-v参数用于查看版本。
//	cqcfg -v
// 如果安装后仍提示找不到cqcfg命令请确保GOPATH/bin路径存在于PATH环境变量中。
//
// 接下来，执行如下命令，生成app.json
//	cqcfg -c .
// 其中"-c"代表将git commit次数加到顺序版本号上，"."代表分析当前目录的源码。
//
// Build app.dll
//
// 编译时需要设置几个环境变量：
//	CGO_LDFLAGS=-Wl,--kill-at
//	CGO_ENABLED=1
//	GOOS=windows
//	GOARCH=386
// 然后执行编译命令：
//	go build -ldflags "-s -w" -buildmode=c-shared -o app.dll
// 若成功编译则会生成app.dll，将其和app.json一起复制到酷Q的指定文件夹内即可
package cqp
