package main

// #include <stdlib.h>
// #include "cq.h"
import "C"
import (
	"unsafe"
	"fmt"
)

func main() {}

var AuthCode int32
var info = C.CString("9," + AppID)

//export AppInfo
func AppInfo() *C.char {
	return info
}

//export Init
func Init(ac int32) int32 {
	AuthCode = ac
	return 0
}

func SendPrivateMsg(qq int64, msg string) (int32, error) {
	str := C.CString(msg)
	code := int32(C.CQ_sendPrivateMsg(C.int32_t(AuthCode), C.int64_t(qq), str))
	C.free(unsafe.Pointer(str))

	if code < 0 {
		return 0, fmt.Errorf("send private message error (%d)",code)
	}
	return code, nil
}