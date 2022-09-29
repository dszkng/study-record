package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map: %v\n", m)
	n = fmt.Sprintf("OS: %s, Arch: %s\n", runtime.GOOS, runtime.GOARCH)
}

//全局变量-已初始化
var m = map[int]string{
	1: "A",
	2: "B",
	3: "C",
}

//全局变量-未初始化
var n string

func main() {
	fmt.Printf(n)
}
