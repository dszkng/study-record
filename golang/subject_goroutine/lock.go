package main

import (
	"fmt"
	// "sync"
	"time"
)

func main() {
	var x = 0

	for i := 0; i < 10; i++ {
		fmt.Println("i1=", fmt.Sprintf("%d", i))

		go func() {
			fmt.Println("i2=", fmt.Sprintf("%d", i))
			x++
		}()
	}

	// var x = 0
	// var l sync.Mutex

	// //并发 10 个 goroutine 来执行函数
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("after i=", fmt.Sprintf("%d", i))
	// 	go func() {
	// 		//先加锁，然后才执行代码
	// 		l.Lock()
	// 		defer l.Unlock()
	// 		fmt.Println("i=", fmt.Sprintf("%d", i))
	// 		x++
	// 	}()
	// }

	time.Sleep(time.Second * time.Duration(1))

	fmt.Println(x)
}
