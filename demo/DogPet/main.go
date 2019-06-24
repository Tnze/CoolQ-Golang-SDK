package main

import (
	"time"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

func main() {}
func init() {
	cqp.AppID = "online.jdao.petdog"
	cqp.Enable = onEnable
	cqp.Disable = onDisable
	cqp.GroupMsg = onGroupMsg
}

func onEnable() int32 {
	cqp.AddLog(cqp.Info, "调试", cqp.GetAppDir())
}

func onDisable() int32 {

}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {

}

type group int64
type qq int64
type pets map[group]map[qq]struct {
	name  string
	birth time.Time
}
