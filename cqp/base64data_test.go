package cqp

import (
	"testing"
)

func TestUnpackGroupMemberList(t *testing.T) {
	mems, err := UnpackGroupMemberList("AAAABABEAAAAADOGQH8AAAAAWHjaagAKqVnI9NPqx6fStgAAAAAAAQAAABMAAFwnP4pcoWM7AAAAAAACAAAAAAAAAAAAAAAAAAAARgAAAAAzhkB/AAAAAGXEFuoADKhzhvnTx6erlZnYyAAAAAAAAQAAABQAAFxQNkBcdT/VAAAAAAACAAAAAAAAAAAAAAAAAAAARAAAAAAzhkB/AAAAAGpre0AACs3svv314s/QzaUAAAAAAAEAAAAUAABcStJZXLBr7gAAAAAAAQAAAAAAAAAAAAAAAAAAAEIAAAAAM4ZAfwAAAACQL1MEAAiBN6I4ucjT6gAAAAAAAAAAABQAAFwnP4Vc3ASjAAAAAAADAAAAAAAAAAAAAAAAAAA=") //测试数据由晓影提供
	if err != nil {
		t.Log("解码出错 ", err)
	}

	if len(mems) != 4 {
		t.Errorf("成员数量解析错误，应为%d而不是%d", 4, len(mems))
	}

	//TODO(Tnze): 测试其余的值是否正确
	t.Logf("%#v", mems)
}

func TestUnpackGroupList(t *testing.T) {
	list, err := UnpackGroupList("AAAABQASAAAAAA0ImDwACLPLt+fE2si6ABIAAAAAIJAh1gAI1+6wrsPDw8MAFAAAAAAprkKcAAq7+sb3yMu198rUAB4AAAAAKhqy4wAUtv60ztSqtq/C/rLlvP69u8H3yLoAFgAAAAAqhadIAAzO3tPvtcS687msyLo=") //测试数据由gif8512提供，感谢
	if err != nil {
		t.Log("解码出错 ", err)
	}

	if len(list) != 5 {
		t.Errorf("成员列表长度解析错误，应为%d而不是%d", 5, len(list))
	}

	//TODO(Tnze): 测试其余的值是否正确
	t.Logf("%#v", list)
}

//TODO(Tnze): 测试GetFriendList
//TODO(Tnze): 测试GetGroupInfo
