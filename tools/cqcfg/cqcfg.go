// Cqcfg 就是CoolQ Config，用于为插件自动生成app.json
//
// 本工具为试验性工具，请按实际需要使用，若有好建议或改进，欢迎提交issue或者pr
//
// 本工具将会扫描您的代码，并且自动统计出您调用了哪些API，响应了哪些事件，
// 并且在生成的app.json中为相应的API注册权限，为事件注册函数
//
// 为了让本工具正常工作，你需要以标准的格式使用Go语言SDK：
//	响应事件时要为cqp包内相应的函数变量赋值
//	在主函数开头以后文中会介绍的语法声明插件的AppID和版本、作者等信息
//
// 在main函数头之前，你需要写以下几个注释：
//	//go:generate cqcfg -c .
//	// cqp: 名称: 插件名称
//	// cqp: 版本: 1.0.0:0
//	// cqp: 作者: 插件作者姓名
//	// cqp: 简介: 您插件的简介
// 其中版本是由插件版本和顺序版本号以冒号分隔形成的，有以下一般形式：
//	主版本.次版本.修正版本:顺序版本
// 注释的前半部分均为强制要求的固定格式，空格不能多不能少
//
// 用法：
//	cqcfg [-c, -v] <插件main包所在目录>
// -c 参数用于自动根据代码提交次数生成版本号
// -v 参数用于查询cqcfg版本
//
// 推荐配合go generate使用
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const version = "cqcfg 2.1"

// 运行时参数
var (
	countCommit = flag.Bool("c", false, "顺序版本+=Git代码提交次数")
	queVersion  = flag.Bool("v", false, "获取cqcfg版本")
)

func main() {
	flag.Parse()

	if *queVersion {
		// 查询版本
		fmt.Println(version)
		os.Exit(0)
	}

	log.SetPrefix("cqcfg: ")
	if flag.NArg() < 1 {
		log.Fatal("请传入项目根目录")
	}

	APIs := make(map[string]int)

	fset := token.NewFileSet()
	err := filepath.Walk(flag.Arg(0), func(path string, finfo os.FileInfo, err error) error {
		if finfo.IsDir() {
			pkgs, first := parser.ParseDir(fset, path, nil, parser.ParseComments)
			if first != nil {
				log.Fatal(first)
			}
			for _, p := range pkgs {
				search(p, APIs)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("遍历当前目录失败: %v", err)
	}

	onCallAPI(APIs)

	// 生成JSON
	app, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	// 写入文件
	f, err := os.OpenFile("app.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(app); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// 搜索整个包，当找到注释、函数调用或者赋值语句时调用相应的处理函数
func search(v *ast.Package, APIs map[string]int) {
	var cqp string
	ast.Inspect(v, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.ImportSpec:
			// 更新cqp包导入名，import语句一定在调用之前，所以在这里更新是及时的
			imp := n.(*ast.ImportSpec)
			if imp.Path.Value == `"github.com/Tnze/CoolQ-Golang-SDK/cqp"` {
				if imp.Name != nil {
					cqp = imp.Name.Name
				} else {
					cqp = "cqp"
				}
			}

		case *ast.Comment: //注释
			onComm(n.(*ast.Comment).Text)

		case *ast.AssignStmt: //赋值语句
			as := n.(*ast.AssignStmt)
			if s, ok := as.Lhs[0].(*ast.SelectorExpr); ok {
				if x, ok := s.X.(*ast.Ident); ok && cqp != "" && x.Name == cqp {
					//记录AppInfo和事件注册
					name, rhs := s.Sel.String(), as.Rhs[0]
					if name == "AppID" {
						if v, ok := rhs.(*ast.BasicLit); ok {
							var err error
							info.AppID, err = strconv.Unquote(v.Value)
							if err != nil {
								log.Fatalf("解析AppID失败: %v", err)
							}
						}
					} else {
						onSetEvent(name, rhs)
					}
				}
			}

		case *ast.SelectorExpr: //调用cqp包
			s := n.(*ast.SelectorExpr)
			if x, ok := s.X.(*ast.Ident); ok && cqp != "" && x.Name == cqp {
				APIs[s.Sel.String()]++
			}
		}
		return true
	})
}

// 统计当前git代码库的提交次数
func commitCount() (int, error) {
	cmd := exec.Command("git", "rev-list", "--all", "--count")
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	seq, err := strconv.Atoi(strings.TrimSpace(string(out)))
	if err != nil {
		return 0, err
	}
	return seq, nil
}
