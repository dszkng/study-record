package main

import "fmt"

//仿写 try {} catch() {} 结构
func try(fn func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()

	fn()
}

func foo() {
	panic("panic error")
}

func main() {
	try(func() {
		foo()
	}, func(e interface{}) {
		fmt.Println(e)
	})
}
