package corpus

import (
	"errors"
	"math/rand"
	"strings"
	"text/template"
	"time"
)

var Corpus = map[string][]string{
	"有过但死了": []string{"[CQ:at,qq={{ .QQ }}]你有过宠物，但是它死了，你不配拥有宠物"},
	"已经领养了": []string{"[CQ:at,qq={{ .QQ }}]你已经有宠物了，它的名字是{{ .Name }}"},
	"成功领养":  []string{"[CQ:at,qq={{ .QQ }}]你领养了{{ .Name }}，快输入“喂食”来给它喂食吧！"},
	"死了还喂":  []string{"[CQ:at,qq={{ .QQ }}]你曾经有过宠物，记得吗？可惜它已经死了"},
	"把它喂死了": []string{"[CQ:at,qq={{ .QQ }}]你喂{{ .Name }}东西吃，ta吃得很饱，然后就撑死了"},
	"还没领养":  []string{"[CQ:at,qq={{ .QQ }}]你还没有宠物"},
}

// Words 包含了渲染一句话所需要的词语
type Words struct {
	QQ, Group int64
	Name      string
	Birth     time.Time
}

// Execute 随机选择一个模版生成一句话
func (w Words) Execute(mean string) (string, error) {
	var buf strings.Builder
	meanset := Corpus[mean]
	if meanset == nil {
		return "", errors.New("找不到模板" + mean)
	} else if len(meanset) == 0 {
		return "", errors.New("模板" + mean + "为空")
	}
	temp := template.Must(template.New(mean).Parse(meanset[rand.Intn(len(meanset))]))
	err := temp.Execute(&buf, w)
	return buf.String(), err
}
