package util

import "testing"

func TestCQCode(t *testing.T) {
	if cqcode := CQCode("face", "id", "14"); cqcode != "[CQ:face,id=14]" {
		t.Error("test error")
	}

	if cqcode := CQCode("record",
		"file", "1.silk",
		"magic", "true"); cqcode != "[CQ:record,file=1.silk,magic=true]" {
		t.Error("CQ码生成错误")
	}
}

func TestEscape(t *testing.T) {
	if e := Escape("[CQ:at,qq=all]"); e != "&#91;CQ:at,qq=all&#93;" {
		t.Errorf("转义错误: %s", e)
	}
}
