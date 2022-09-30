package main

import (
	"fmt"
	"time"
)

//定义大小为10的缓冲信道/管道
var ch = make(chan string, 10)

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)

	//写入数据到缓冲信道
	ch <- url
}

func main() {
	for i := 0; i < 3; i++ {
		go download("https://baidu.com/" + string(i+'0'))
	}

	for i := 0; i < 3; i++ {
		//从缓冲信道中取出数据
		msg := <-ch
		fmt.Println("finish: ", msg)
	}

	fmt.Println("Done!")
}
