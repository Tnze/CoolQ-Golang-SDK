package main

import (
	"fmt"
	"log"
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

	Menu   []menu        `json:"menu"`
	Status []interface{} `json:"status"`
}{
	Ret:    1,
	APIver: 9,
	Menu:   []menu{},
	Status: []interface{}{},
}

type menu struct {
	Name string `json:"name"`
	Func string `json:"function"`
}

// 读取注释
func onComm(comm string) { //处理cqp注释
	switch {
	case strings.HasPrefix(comm, "// cqp: 名称:"):
		if _, err := fmt.Sscanf(comm, "// cqp: 名称:%s", &info.Name); err != nil {
			log.Fatal("无法解析应用名称:", err)
		}
	case strings.HasPrefix(comm, "// cqp: 版本:"):
		var v1, v2, v3, seq int
		_, err := fmt.Sscanf(comm, "// cqp: 版本:%d.%d.%d:%d", &v1, &v2, &v3, &seq)
		if err != nil {
			log.Fatal("无法解析版本号:", err)
		}

		if *countCommit {
			c, err := commitCount()
			if err != nil {
				fmt.Printf("警告: 统计Git提交数失败: %v", err)
			} else {
				seq += c
			}
		}

		info.Version = fmt.Sprintf("%d.%d.%d", v1, v2, v3)
		info.VersionID = seq
	case strings.HasPrefix(comm, "// cqp: 作者:"):
		if _, err := fmt.Sscanf(comm, "// cqp: 作者:%s", &info.Author); err != nil {
			log.Fatal("无法解析作者名:", err)
		}
	case strings.HasPrefix(comm, "// cqp: 简介: "):
		info.Desc = strings.TrimPrefix(comm, "// cqp: 简介: ")
	case strings.HasPrefix(comm, "// cqp: 菜单:"):
		var m menu
		if _, err := fmt.Sscan(strings.TrimPrefix(comm, "// cqp: 菜单:"), &m.Name, &m.Func); err != nil {
			log.Fatal("无法解析菜单指令:", err)
		}
		info.Menu = append(info.Menu, m)
	}

}
