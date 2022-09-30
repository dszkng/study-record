package main

import "fmt"

func main() {
	var x interface{}
	x = 1

	// type I1 int

	var res, ok = x.(int)

	fmt.Println(res, ok)
}
