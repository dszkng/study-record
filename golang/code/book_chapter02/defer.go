package main

import "fmt"

func main() {
	defer fmt.Println("defer call 1")
	defer fmt.Println("defer call 2")
	defer func() {
		fmt.Println("defer call 3")
	}()

	for i := 0; i < 5; i++ {
		defer func(i int) {
			fmt.Printf("test: %d\n", i)
		}(i)
	}
	fmt.Println("first call")
}
