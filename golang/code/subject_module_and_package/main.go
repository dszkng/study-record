package main

import (
    "fmt"
    "mymath"
    "example/calc"
)

// func calc(a, b int) (c int) {
// 	c = a + b
// 	return
// }

func main() {
	a := 1
	b := 2
	// c := calc(a, b)
	// c := calc(a, b) //go run main.go calc.go

    //./main.go:5:5: imported and not used: "math"
    //./main.go:18:10: undefined: Calc
    c := mymath.Calc(a, b)

	fmt.Printf("1+2=%d\n", c)

    c2 := calc.Calc(a, b)

	fmt.Printf("1+2=%d\n", c2)
}