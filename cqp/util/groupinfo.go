package util

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"time"

	sc "golang.org/x/text/encoding/simplifiedchinese"
)

//GroupMember 群成员信息
type GroupMember struct {
	//群号和Q号
	Group, QQ  int64
	Name, Card string
	//性别，0/男；1/女
	Gender             int32
	Age                int32
	Area               string
	JoinTime, LastChat time.Time
	Level              string
	//管理权限，1/成员；2/管理员；3/群主
	Auth int32
	//是否有不良记录
	Bad         bool
	Title       string
	TitleLife   time.Time
	CanSetTitle bool
}

//UnpackGroupList 解码群成员列表
func UnpackGroupList(str string) ([]GroupMember, error) {
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	//读取群列表人数
	var MemNum int32
	if err := binary.Read(r, binary.BigEndian, &MemNum); err != nil {
		return nil, err
	}

	//读成员信息
	members := make([]GroupMember, MemNum)
	for i := 0; i < int(MemNum); i++ {
		var Len int16
		if err := binary.Read(r, binary.BigEndian, &Len); err != nil {
			return nil, err
		}

		data := make([]byte, Len)
		if err := binary.Read(r, binary.BigEndian, &data); err != nil {
			return nil, err
		}

		var err error
		members[i], err = readGroupMember(data)
		if err != nil {
			return members, err //返回成功解析的成员
		}
	}

	return members, nil
}

//UnpackGroupMemberInfo 解码群成员信息
func UnpackGroupMemberInfo(str string) (m GroupMember, err error) {
	var data []byte
	data, err = base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}

	m, err = readGroupMember(data)
	return
}

func readGroupMember(data []byte) (m GroupMember, err error) {
	r := bytes.NewReader(data)

	for _, v := range []interface{}{
		&m.Group, &m.QQ,
		&m.Name, &m.Card,
		&m.Gender, &m.Age,
		&m.Area,
		&m.JoinTime, &m.LastChat,
		&m.Level,
		&m.Auth,
		&m.Bad,
		&m.Title, &m.TitleLife,
		&m.CanSetTitle,
	} {
		err = readField(r, v)
		if err != nil {
			return
		}
	}

	return
}

func readField(r io.Reader, v interface{}) error {
	switch v.(type) {
	default:
		panic(fmt.Errorf("出乎意料的类型: %T", v))

	case *int64, *int32:
		return binary.Read(r, binary.BigEndian, v)

	case *string:
		var len int16
		if err := binary.Read(r, binary.BigEndian, &len); err != nil {
			return err
		}

		buff := make([]byte, len)
		if err := binary.Read(r, binary.BigEndian, &buff); err != nil {
			return err
		}

		str, err := sc.GB18030.NewDecoder().Bytes(buff)
		*v.(*string) = string(str)
		return err

	case *bool:
		var value int32
		if err := binary.Read(r, binary.BigEndian, &value); err != nil {
			return err
		}
		*v.(*bool) = value != 0
		return nil

	case *time.Time:
		var unix int32
		if err := binary.Read(r, binary.BigEndian, &unix); err != nil {
			return err
		}
		*v.(*time.Time) = time.Unix(int64(unix), 0)
		return nil
	}
}
