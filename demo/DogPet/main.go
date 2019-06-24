package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

func main() {}
func init() {
	cqp.AppID = "online.jdao.dogpet"
	cqp.Enable = onEnable
	cqp.Disable = onDisable
	cqp.GroupMsg = onGroupMsg
}

func onEnable() int32 {
	defer handleErr()
	// 读取pets.json
	file, err := petsFile()
	if err != nil {
		printErr(err)
		return -1
	}
	cqp.AddLog(cqp.Info, "调试", "pets.json文件路径："+file)
	petsdata, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		return 0
	} else if err != nil {
		printErr(err)
		return -1
	}
	err = json.Unmarshal(petsdata, &pets)
	if err != nil {
		printErr(err)
		return -1
	}

	return 0
}

func onDisable() int32 {
	defer handleErr()
	//将宠物们保存到pets.json
	petsdata, err := json.MarshalIndent(pets, "", "\t")
	if err != nil {
		printErr(err)
		return -1
	}
	file, err := petsFile()
	if err != nil {
		printErr(err)
		return -1
	}
	err = ioutil.WriteFile(file, petsdata, 0666)
	if err != nil {
		printErr(err)
		return -1
	}
	return 0
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	defer handleErr()

	if _, ok := pets[group(fromGroup)]; !ok {
		pets[group(fromGroup)] = make(map[qq]pet)
	}

	pet, ok := pets[group(fromGroup)][qq(fromQQ)]
	switch {
	case strings.HasPrefix(msg, "领养 "):
		var name string
		n, err := fmt.Sscanf(msg, "领养 %s", &name)
		if n != 1 || err != nil {
			cqp.AddLog(cqp.Debug, "领养", "领养指令不正确: "+msg)
		}

		if ok && pet.name != "" {
			if pet.dead {
				cqp.SendGroupMsg(fromGroup, fmt.Sprintf("[CQ:at,qq=%d]你有过宠物，但是它死了，你不配拥有宠物", fromQQ))
			} else {
				cqp.SendGroupMsg(fromGroup, fmt.Sprintf("[CQ:at,qq=%d]你已经有宠物了，它的名字是%s", fromQQ, pet.name))
			}
		} else {
			cqp.AddLog(cqp.Info, "领养", fmt.Sprintf("%d领养了%s", fromQQ, name))
			pet.name = name
			pet.birth = time.Now()
			pets[group(fromGroup)][qq(fromQQ)] = pet
			cqp.SendGroupMsg(fromGroup, fmt.Sprintf("[CQ:at,qq=%d]你领养了%s，快输入“喂食”来给它喂食吧！", fromQQ, name))
		}
	case msg == "喂食":
		if ok && pet.name != "" {
			if pet.dead {
				cqp.SendGroupMsg(fromGroup, fmt.Sprintf("[CQ:at,qq=%d]你曾经有过宠物，记得吗？可惜它已经死了", fromQQ))
			} else {
				cqp.AddLog(cqp.Info, "喂食", fmt.Sprintf("%d把%s撑死了", fromQQ, pet.name))
				cqp.SendGroupMsg(fromGroup, fmt.Sprintf("[CQ:at,qq=%d]你喂%s东西吃，ta吃得很饱，然后就撑死了", fromQQ, pet.name))
				pet.dead = true
				pets[group(fromGroup)][qq(fromQQ)] = pet
			}
		} else {
			cqp.SendGroupMsg(fromGroup, fmt.Sprintf("[CQ:at,qq=%d]你还没有宠物", fromQQ))
		}
	}
	return 0
}

type group int64
type qq int64

var pets = make(
	map[group]map[qq]pet,
)

type pet struct {
	name  string
	birth time.Time
	dead  bool
}

//打印错误到日志
func printErr(err error) {
	cqp.AddLog(cqp.Error, "错误", err.Error())
}

//获取pets.json文件的路径，若路径不存在则顺便创建
func petsFile() (string, error) {
	appdir := cqp.GetAppDir()
	err := os.MkdirAll(appdir, os.ModeDir)
	return filepath.Join(appdir, "pets.json"), err
}

func handleErr() {
	if err := recover(); err != nil {
		cqp.AddLog(cqp.Fatal, "严重错误", fmt.Sprint(err))
	}
}
