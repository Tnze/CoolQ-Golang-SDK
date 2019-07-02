package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

var info = struct {
	Ret       int    `json:"ret"`
	APIver    int    `json:"apiver"`
	AppID     string `json:"appid"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	VersionID int    `json:"version_id"`
	Author    string `json:"author"`
	Desc      string `json:"description"`

	Events []event `json:"event"`
	Auth   []int   `json:"auth"`
}{
	Ret:    1,
	APIver: 9,
}

type event struct {
	Name     string `json:"name"`
	Function string `json:"function"`
	Type     int    `json:"type"`
	Priority int    `json:"priority"`
	ID       int    `json:"id"`
}

var authcode = map[string]int{
	"GetCookies":           20,
	"GetRecord":            30,
	"SendGroupMsg":         101,
	"SendDiscussMsg":       103,
	"SendPrivateMsg":       106,
	"SendLike":             110,
	"SendLikeV2":           110,
	"SetGroupKick":         120,
	"SetGroupBan":          121,
	"SetGroupAdmin":        122,
	"SetGroupWholeBan":     123,
	"SetGroupAnonymousBan": 124,
	"SetGroupAnonymous":    125,
	"SetGroupCard":         126,
	"SetGroupLeave":        127,
	"SetGroupSpecialTitle": 128,
	"GetGroupMemberInfo":   130,
	"GetGroupMemberInfoV2": 130,
	"GetStrangerInfo":      131,
	"SetDiscussLeave":      140,
	"SetFriendAddRequest":  150,
	"SetGroupAddRequest":   151,
	"GetGroupMemberList":   160,
	"GetGroupList":         161,
	"DeleteMsg":            180,
}

func main() {
	log.SetPrefix("cqcfg: ")

	fset := token.NewFileSet() // positions are relative to fset
	pkgs, first := parser.ParseDir(fset, "./demo/DogPet", nil, parser.ParseComments)
	if first != nil {
		log.Fatal(first)
	}

	APIs := make(map[string]int)
	for _, p := range pkgs {
		search(p,
			func(comm string) { //处理cqp注释
				switch {
				case strings.HasPrefix(comm, "// cqp: 名称:"):
					if _, err := fmt.Sscanf(comm, "// cqp: 名称:%s", &info.Name); err != nil {
						log.Fatal("无法解析应用名称:", err)
					}
				case strings.HasPrefix(comm, "// cqp: 版本:"):
					var v1, v2, v3, seq int
					if _, err := fmt.Sscanf(comm, "// cqp: 版本:%d.%d.%d:%d", &v1, &v2, &v3, &seq); err != nil {
						log.Fatal("无法解析版本号:", err)
					}
					info.Version = fmt.Sprintf("%d.%d.%d", v1, v2, v3)
					info.VersionID = seq
				case strings.HasPrefix(comm, "// cqp: 作者:"):
					if _, err := fmt.Sscanf(comm, "// cqp: 作者:%s", &info.Author); err != nil {
						log.Fatal("无法解析作者名:", err)
					}
				case strings.HasPrefix(comm, "// cqp: 简介: "):
					info.Desc = strings.TrimPrefix(comm, "// cqp: 简介: ")
				}
			},
			func(name string) { APIs[name]++ }, //记录API调用
			func(name string, rhs ast.Expr) { //记录AppInfo和事件注册
				switch name {
				case "AppID":
					if v, ok := rhs.(*ast.BasicLit); ok {
						// fmt.Println(name, "=", v.Value)
						info.AppID = v.Value
					}
				case "Enable":
					info.Events = append(info.Events, event{
						ID:       1003,
						Type:     1003,
						Name:     "插件启用",
						Priority: 20000,
						Function: "EVENT_ON_ENABLE",
					})
				case "Disable":
					info.Events = append(info.Events, event{
						ID:       1004,
						Type:     1004,
						Name:     "插件停用",
						Priority: 20000,
						Function: "EVENT_ON_DISABLE",
					})
				case "Start":
					info.Events = append(info.Events, event{
						ID:       1001,
						Type:     1001,
						Name:     "酷Q启动",
						Priority: 20000,
						Function: "EVENT_ON_START",
					})
				case "Exit":
					info.Events = append(info.Events, event{
						ID:       1002,
						Type:     1002,
						Name:     "酷Q退出",
						Priority: 20000,
						Function: "EVENT_ON_EXIT",
					})
				case "PrivateMsg":
					info.Events = append(info.Events, event{
						ID:       1,
						Type:     21,
						Name:     "私聊消息",
						Function: "EVENT_ON_PRIVATE_MSG",
						Priority: 20000,
					})

				case "GroupMsg":
					info.Events = append(info.Events, event{
						ID:       2,
						Type:     2,
						Name:     "群消息",
						Function: "EVENT_ON_GROUP_MSG",
						Priority: 20000,
					})

				case "DiscussMsg":
					info.Events = append(info.Events, event{
						ID:       3,
						Type:     4,
						Name:     "讨论组消息",
						Function: "EVENT_ON_DISCUSS_MSG",
						Priority: 20000,
					})

				case "GroupUpload":
					info.Events = append(info.Events, event{
						ID:       4,
						Type:     11,
						Name:     "群文件上传",
						Function: "EVENT_ON_GROUP_UPLOAD",
						Priority: 20000,
					})
				case "GroupAdmin":
					info.Events = append(info.Events, event{
						ID:       5,
						Type:     101,
						Name:     "群管理员变动",
						Function: "EVENT_ON_GROUP_ADMIN",
						Priority: 20000,
					})
				case "GroupMemberDecrease":
					info.Events = append(info.Events, event{
						ID:       6,
						Type:     102,
						Name:     "群成员减少",
						Function: "EVENT_ON_GROUP_MEMBER_DECREASE",
						Priority: 20000,
					})
				case "GroupMemberIncrease":
					info.Events = append(info.Events, event{
						ID:       7,
						Type:     103,
						Name:     "群成员增加",
						Function: "EVENT_ON_GROUP_MEMBER_INCREASE",
						Priority: 20000,
					})
				case "FriendAdd":
					info.Events = append(info.Events, event{
						ID:       10,
						Type:     201,
						Name:     "好友添加",
						Function: "EVENT_ON_FRIEND_ADD",
						Priority: 20000,
					})
				case "FriendRequest":
					info.Events = append(info.Events, event{
						ID:       8,
						Type:     301,
						Name:     "加好友请求",
						Function: "EVENT_ON_FRIEND_REQUEST",
						Priority: 20000,
					})
				case "GroupRequest":
					info.Events = append(info.Events, event{
						ID:       9,
						Type:     302,
						Name:     "加群请求／邀请",
						Function: "EVENT_ON_GROUP_REQUEST",
						Priority: 20000,
					})
				}
			},
		)
	}

	for i, v := range APIs { //添加权限声明
		if v > 0 {
			code, ok := authcode[i]
			if ok {
				info.Auth = append(info.Auth, code)
			}
		}
	}

	json.NewEncoder(os.Stdout).Encode(info)
}

func search(v *ast.Package, findComm, findCall func(name string), findAssign func(name string, rhs ast.Expr)) {
	for _, f := range v.Files {
		//获取该文件里cqp包的导入名
		cqp := importsName(f)
		if cqp == "" {
			continue //这个文件没有导入cqp
		}

		//搜索API调用
		ast.Inspect(f, func(n ast.Node) bool {
			switch n.(type) {
			case *ast.Comment: //注释
				findComm(n.(*ast.Comment).Text)
			case *ast.AssignStmt: //赋值语句
				as := n.(*ast.AssignStmt)
				if s, ok := as.Lhs[0].(*ast.SelectorExpr); ok {
					if x, ok := s.X.(*ast.Ident); ok && x.Name == cqp {
						findAssign(s.Sel.String(), as.Rhs[0])
					}
				}

			case *ast.SelectorExpr: //调用cqp包
				s := n.(*ast.SelectorExpr)
				if x, ok := s.X.(*ast.Ident); ok && x.Name == cqp {
					findCall(s.Sel.String())
				}
			}
			return true
		})
	}
}

func importsName(f *ast.File) string {
	for _, p := range f.Imports {
		if p.Path.Value == `"github.com/Tnze/CoolQ-Golang-SDK/cqp"` {
			// fmt.Println(p.Name, p.Path.Value)
			if p.Name != nil {
				return p.Name.Name
			}
			return "cqp"
		}
	}
	return ""
}
