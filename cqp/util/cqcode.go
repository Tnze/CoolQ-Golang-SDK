package util

import (
	"fmt"
	"strings"
)

var cqValueEscape = strings.NewReplacer(
	"&", "&amp;",
	"[", "&#91;",
	"]", "&#93;",
	",", "&#44;",
).Replace

// CQCode 生成一段CQ码，value将被自动转义
// panics if given an odd number of keyvalue.
func CQCode(function string, keyvalue ...interface{}) string {
	if len(keyvalue)%2 == 1 {
		panic("util.CQCode: odd keyvalue count")
	}

	pairs := make([]string, len(keyvalue)/2)
	for i := 0; i < len(keyvalue); i += 2 {
		key, value := fmt.Sprint(keyvalue[i]), fmt.Sprint(keyvalue[i+1])
		pairs[i/2] = key + "=" + cqValueEscape(value)
	}

	return "[CQ:" + function + "," + strings.Join(pairs, ",") + "]"
}

var cqEscape = strings.NewReplacer(
	"&", "&amp;",
	"[", "&#91;",
	"]", "&#93;",
).Replace

// Escape 对于不在CQ码内的消息（即文本消息）进行转义
// 以防止它们被酷Q解析为CQ码
func Escape(str string) string {
	return cqEscape(str)
}
