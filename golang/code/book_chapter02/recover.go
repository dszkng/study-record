package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%s", e)
		}
	}()

	panic(errors.New("test exception"))
}
