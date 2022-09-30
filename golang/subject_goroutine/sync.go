package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)

	//为 wg 减去一个计数
	wg.Done()
}

func main() {
	for i := 0; i < 3; i++ {
		//为 wg 添加一个计数
		wg.Add(1)

		url := fmt.Sprintf("http://baidu.com/%d", i)

		//启动一个协程
		go download(url)
	}

	//等待所有的协程执行结束
	wg.Wait()

	fmt.Println("Done!")
}
