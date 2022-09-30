package main

import (
	"errors"
	"fmt"
)

func makeErr() error {
	return errors.New("makeErr")
}

func main() {
	// 放在最开始
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 会中断其他执行
	// panic("sss")
	err := errors.New("test error")
	fmt.Println(err)

	if err := makeErr(); err != nil {
		fmt.Println(err)
	}

	panic("sss")
}
