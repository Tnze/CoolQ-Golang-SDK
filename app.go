// 您的应用的代码写在这里
package main

import "C"
import (
)

const AppID = "your.app.id"

//export Enable
func Enable() int32 {
	code,err:=SendPrivateMsg(1624188026, "test")
	if err!=nil{
		
	}
	return 0
}
