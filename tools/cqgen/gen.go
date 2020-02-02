package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var CQP cqp
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	for _, v := range f.Comments {
		for _, v := range strings.Split(v.Text(), "\n") {
			if i := strings.IndexFunc(v, unicode.IsSpace); i > 0 {
				call := v[:i]
				body := strings.TrimSpace(v[i:])
				switch call {
				case "cqapi":
					CQP.addAPI(body)
				}
			}
		}
	}
	if err := CQP.Generate(); err != nil {
		panic(err)
	}
}

type cqp struct {
	apis []cqapi
}

func (c cqp) Generate() error {
	var sb bytes.Buffer
	c.genAPIs(&sb)
	return ioutil.WriteFile("apis_native.go", sb.Bytes(), 0644)
}

func (c cqp) genAPIs(w io.StringWriter) {
	w.WriteString(`// Code generated DO NOT EDIT.
// +build !websocket

package cqp

/*
#include <windows.h>
#include <stdint.h>
#include <stdlib.h>

// #define __stdcall __attribute__((__stdcall__))
#define cq_bool_t int32_t

#define CQAPI(RetType, Name, ...)                                         \
      typedef RetType(__stdcall *Name##_Type)(int32_t ac, ##__VA_ARGS__); \
      Name##_Type Name##_Ptr;                                             \
      RetType Name(__VA_ARGS__);

#define LoadAPI(Name) Name##_Ptr = (Name##_Type)GetProcAddress(hmod, #Name)

extern char *__stdcall AppInfo(){ return _appinfo(); }

int32_t ac; //AccessCode
`)
	for _, api := range c.apis {
		w.WriteString("CQAPI(" + cTypes[api.ret] + ", " + api.name)
		for _, arg := range api.args {
			w.WriteString(", " + cTypes[arg])
		}
		w.WriteString(")\n")
	}

	w.WriteString(`extern int32_t __stdcall Initialize(int32_t access_code)
{
    ac = access_code;
    HMODULE hmod = LoadLibrary("CQP.dll");
`)
	for _, api := range c.apis {
		w.WriteString("    LoadAPI(" + api.name + ");\n")
	}
	w.WriteString("    return 0;\n}\n")

	for _, api := range c.apis {
		w.WriteString(cTypes[api.ret] + " " + api.name + "(")
		if len(api.args) == 1 && api.args[0] != "" {
			w.WriteString(cTypes[api.args[0]] + " var0")

		} else if len(api.args) > 1 {
			w.WriteString(cTypes[api.args[0]] + " var0")
			for i, arg := range api.args[1:] {
				w.WriteString(", " + cTypes[arg] + " var" + strconv.Itoa(i+1))
			}
		}
		w.WriteString(")\n{\n    " + cTypes[api.ret] + " ret = " + api.name + "(ac")
		for i := range api.args {
			w.WriteString(", var" + strconv.Itoa(i))
		}
		w.WriteString(");\n")
		for i, arg := range api.args {
			if arg == "string" {
				w.WriteString("    free(var" + strconv.Itoa(i) + ");\n")
			}
		}
		w.WriteString("    return ret;\n}\n")
	}

	w.WriteString(`*/
import "C"
`)
}

func (c cqp) genEvents(sb *strings.Builder) {

}

type cqapi struct {
	name string
	args []string
	ret  string
}

func (c cqapi) String() string {
	return fmt.Sprintf("func %s(%s) %s", c.name, strings.Join(c.args, ","), c.ret)
}

func (c *cqp) addAPI(s string) error {
	// example: CQ_addLog(int32,string,string)int32
	sb := regexp.MustCompile(`(\w+)\(((?:,?\w+)*)\)(\w+)`).FindStringSubmatch(s)
	if len(sb) != 4 {
		return errors.New("format error")
	}
	c.apis = append(c.apis, cqapi{
		name: sb[1],
		args: strings.Split(sb[2], ","),
		ret:  sb[3],
	})
	return nil
}

var cTypes = map[string]string{
	"bool":   "cq_bool_t",
	"int32":  "int32_t",
	"int64":  "int64_t",
	"string": "char *",
}
