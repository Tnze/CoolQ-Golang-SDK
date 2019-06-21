//go:generate go build -ldflags "-s -w" -buildmode=c-shared -o app.dll
package main

/*
#include <stdint.h>
*/
import "C"

func main() {}

var info = C.CString("9,"+"online.jdao.cqtest")

//export AppInfo
func AppInfo() *C.char {
	return info
}
