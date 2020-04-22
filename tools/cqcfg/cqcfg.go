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
//	// cqp: 版本: 1.0.0:1
//	// cqp: 作者: 插件作者姓名
//	// cqp: 简介: 您插件的简介
// 其中版本是由插件版本和顺序版本号以冒号分隔形成的，有以下一般形式：
//	主版本.次版本.修正版本:顺序版本
// 注释的前半部分均为强制要求的固定格式，空格不能多不能少
//
// 若需要在json中添加菜单，以下指令是可用的，但是用于响应的函数目前仍需要您自己编写
//	// cqp: 菜单: <菜单名> <C函数名>
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
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const version = "2.7"

// 运行时参数
var (
	countCommit = flag.Bool("c", false, "顺序版本+=Git代码提交次数")
	queVersion  = flag.Bool("v", false, "获取cqcfg版本")
)

func main() {
	flag.Parse()
	log.SetFlags(0) // log不显示日期和时间
	log.SetPrefix("cqcfg: ")

	if *queVersion { // 查询版本
		log.Print(version)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		log.Fatal("请传入项目根目录")
	}

	APIs := make(map[string]int)

	fset := token.NewFileSet()
	// 遍历当前目录下所有包
	err := filepath.Walk(flag.Arg(0), func(path string, finfo os.FileInfo, err error) error {
		if finfo.IsDir() {
			pkgs, first := parser.ParseDir(fset, path, nil, parser.ParseComments)
			if first != nil {
				return first
			}
			for _, p := range pkgs {
				search(p, APIs)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal("遍历当前目录失败: ", err)
	}

	addAuth(APIs)

	// 生成JSON
	app, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		log.Fatal("生成json失败: ", err)
	}

	// 写入文件
	err = ioutil.WriteFile("app.json", app, 0644)
	if err != nil {
		log.Fatal("写入文件失败: ", err)
	}
}

// 搜索、处理整个包，当找到注释、函数调用或者赋值语句时调用相应的处理函数
func search(v *ast.Package, APIs map[string]int) {
	var cqp string
	ast.Inspect(v, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.ImportSpec:
			// 更新cqp包导入名，import语句一定在调用之前，所以在这里更新是及时的
			handleImportSpec(&cqp, n.(*ast.ImportSpec))

		case *ast.Comment: //注释
			onComm(n.(*ast.Comment).Text)

		case *ast.AssignStmt: //赋值语句
			handleAssignStmt(cqp, n.(*ast.AssignStmt))

		case *ast.SelectorExpr: //调用cqp包
			handleSelectorExpr(cqp, n.(*ast.SelectorExpr), APIs)

		}
		return true
	})
}

func handleImportSpec(cqp *string, imp *ast.ImportSpec) {
	if imp.Path.Value == `"github.com/Tnze/CoolQ-Golang-SDK/cqp"` ||
		imp.Path.Value == `"github.com/Tnze/CoolQ-Golang-SDK/v2/cqp"` {
		if imp.Name != nil {
			*cqp = imp.Name.Name
		} else {
			*cqp = "cqp"
		}
	}
}

func handleAssignStmt(cqp string, as *ast.AssignStmt) {
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
}

func handleSelectorExpr(cqp string, se *ast.SelectorExpr, apis map[string]int) {
	if x, ok := se.X.(*ast.Ident); ok && cqp != "" && x.Name == cqp {
		apis[se.Sel.String()]++
	}
}
