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
	"github.com/Tnze/CoolQ-Golang-SDK/demo/DogPet/corpus"
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
	err := readPets()
	if err != nil {
		printErr(err)
		return -1
	}

	//读语料库
	err = readCorpus()
	if err != nil {
		printErr(err)
		return -1
	}

	return 0
}

func readPets() error {
	file, err := getFile("pets.json")
	if err != nil {
		return err
	}
	cqp.AddLog(cqp.Debug, "调试", "pets.json文件路径："+file)
	petsdata, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	err = json.Unmarshal(petsdata, &pets)
	if err != nil {
		return err
	}
	return nil
}

func readCorpus() error {
	file, err := getFile("corpus.json")
	if err != nil {
		return err
	}
	cqp.AddLog(cqp.Debug, "调试", "corpus.json文件路径："+file)
	petsdata, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	err = json.Unmarshal(petsdata, &corpus.Corpus)
	if err != nil {
		return err
	}
	return nil
}

func onDisable() int32 {
	defer handleErr()
	//将宠物们保存到pets.json
	petsdata, err := json.MarshalIndent(pets, "", "\t")
	if err != nil {
		printErr(err)
		return -1
	}
	file, err := getFile("pets.json")
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

	words := corpus.Words{
		QQ:    fromQQ,
		Group: fromGroup,
		Name:  pet.Name,
		Birth: pet.Birth,
	}

	var (
		s   string
		err error
	)
	switch {
	case strings.HasPrefix(msg, "领养 "):
		var name string
		n, err := fmt.Sscanf(msg, "领养 %s", &name)
		if n != 1 || err != nil {
			cqp.AddLog(cqp.Debug, "领养", "领养指令不正确: "+msg)
			return 0
		}

		if ok && pet.Name != "" {
			if pet.Dead {
				s, err = words.Execute("有过但死了")
			} else {
				s, err = words.Execute("已经领养了")
			}
		} else {
			pet.Name = name
			pet.Birth = time.Now()
			words.Name = pet.Name
			words.Birth = pet.Birth
			cqp.AddLog(cqp.Info, "领养", fmt.Sprintf("%d领养了%s", fromQQ, name))
			pets[group(fromGroup)][qq(fromQQ)] = pet
			s, err = words.Execute("成功领养")
		}
	case msg == "喂食":
		if ok && pet.Name != "" {
			if pet.Dead {
				s, err = words.Execute("死了还喂")
			} else {
				cqp.AddLog(cqp.Info, "喂食", fmt.Sprintf("%d把%s撑死了", fromQQ, pet.Name))
				s, err = words.Execute("把它喂死了")
				pet.Dead = true
				pets[group(fromGroup)][qq(fromQQ)] = pet
			}
		} else {
			s, err = words.Execute("还没领养")
		}
	default:
		return 0
	}

	if err != nil {
		printErr(err)
		return -1
	}
	cqp.SendGroupMsg(fromGroup, s)
	return 0
}

type group int64
type qq int64

var pets = make(
	map[group]map[qq]pet,
)

type pet struct {
	Name  string
	Birth time.Time
	Dead  bool
}

//打印错误到日志
func printErr(err error) {
	cqp.AddLog(cqp.Error, "错误", err.Error())
}

//获取插件文件的路径，若路径不存在则顺便创建
func getFile(name string) (string, error) {
	appdir := cqp.GetAppDir()
	err := os.MkdirAll(appdir, os.ModeDir)
	return filepath.Join(appdir, name), err
}

func handleErr() {
	if err := recover(); err != nil {
		cqp.AddLog(cqp.Fatal, "严重错误", fmt.Sprint(err))
	}
}
