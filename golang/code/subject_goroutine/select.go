package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * time.Duration(10))
		ch <- 1
	}()

	time.Sleep(time.Second * time.Duration(1))

	select {
	case i, ok := <-ch:
		if ok {
			fmt.Println("case", i)
		}
	default:
		fmt.Println("default")
	}
}
