package basic

import (
	"fmt"
	"os"
)

// package variable
var (
	home   = os.Getenv("HOME")
	user   = os.Getenv("USER")
	gopath = os.Getenv("GOPATH")
)

// package  const
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	java = iota
	python
	golang
)

// 首字母大写可以将 变量，常量，函数，struct， interface 在包外部可以访问到
var Name = "jack"

func iotaDefine() {
	fmt.Printf("KB=%d, MG=%d, GB=%d, TB=%d\n", KB, MB, GB, TB)
	fmt.Printf("java=%d, python=%d, golang=%d\n", java, python, golang)
}

func variableDefine() {
	// 指定类型
	var a int = 1
	// 类型自判断
	var b = 2
	// 短变量
	c := 3
	fmt.Printf("a=%d, b=%d, c=%d \n", a, b, c)
}

// 零值
func variableZeroValues() {
	var a int
	var b string
	var c float64
	var d bool
	var e map[string]int
	var f []int
	var g [5]string
	fmt.Printf("a=%d, b=%q, c=%f, d=%v, e=%v, f=%v, g=%v\n", a, b, c, d, e, f, g)
}

func packageVariableDefine() {
	fmt.Printf("home=%s, user=%s, gopath=%s", home, user, gopath)
}

func types() {
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// float32 float64
	// byte (int8)
	// rune (int16)
	// bool
	// string
}
