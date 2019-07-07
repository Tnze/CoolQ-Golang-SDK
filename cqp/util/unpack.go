package util

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"strings"
	"time"
)

//UnpackGroupList 解码群成员列表
func UnpackGroupList(str string) error {
	r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
	//读取群列表人数
	var MemNum int32
	if err := binary.Read(r, binary.BigEndian, &MemNum); err != nil {
		return err
	}
	fmt.Println("人数:", MemNum)

	//读成员信息
	for i := 0; i < int(MemNum); i++ {
		var Len int16
		if err := binary.Read(r, binary.BigEndian, &Len); err != nil {
			return err
		}
		data := make([]byte, Len)
		if err := binary.Read(r, binary.BigEndian, &data); err != nil {
			return err
		}
		readGroupMember(data)
	}

	return nil
}

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
	Auth        int32
	Bad         bool
	Title       string
	TitleLife   time.Time
	CanSetTitle bool
}

func readGroupMember(data []byte) error {
	r := bytes.NewReader(data)

	for _, v := range []decodable{} {
		if err := v.read(r); err != nil {
			return err
		}
	}

	fmt.Println("数据：")
	return nil
}

type decodable interface {
	read(r io.Reader) error
}

type (
	longF   int64
	intF    int32
	timeF   time.Time
	stringF string
	boolF   bool
)

func (l *longF) read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, l)
}
func (i *intF) read(r io.Reader) error {
	return binary.Read(r, binary.BigEndian, i)
}
func (t *timeF) read(r io.Reader) error {
	var unix int32
	if err := binary.Read(r, binary.BigEndian, &unix); err != nil {
		return err
	}
}
func (s *stringF) read(r io.Reader) error {
	var Len int16
	if err := binary.Read(r, binary.BigEndian, &Len); err != nil {
		return err
	}
	str := make([]byte, Len)
	if err := binary.Read(r, binary.BigEndian, &str); err != nil {
		return err
	}
	*s = stringF(str)
	return nil
}
func (b *boolF) read(r io.Reader) error {
	var v int32
	err := binary.Read(r, binary.BigEndian, &v)
	*b = v != 0
	return err
}
