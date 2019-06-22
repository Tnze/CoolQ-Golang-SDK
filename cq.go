package main

// #include <stdlib.h>
// #include "cq.h"
import "C"
import (
	"unsafe"
)

func main() {}

var info = C.CString("9," + AppID)

//export appinfo
func appinfo() *C.char {
	return info
}

func AddLog(priority int32, logtype, reason string) int32 {
	return int32(C.CQ_addLog(
		priority, 
		C.CString(logtype), 
		C.CString(reason),
	))
}
