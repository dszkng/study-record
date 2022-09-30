package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "shang"

	r1 := reflect.TypeOf(a)
	fmt.Println("type:", r1.Kind()) //type: string

	r2 := reflect.TypeOf(&a)
	fmt.Println("type:", r2.Kind()) //type: ptr

	r3 := reflect.ValueOf(a)
	fmt.Println("type:", r3.Type()) //type: string

	r4 := reflect.ValueOf(&a)
	fmt.Println("type:", r4.Type()) //type: *string

	fmt.Println("canset:", r3.CanSet()) //canset: false
	fmt.Println("canset:", r4.CanSet()) //canset: false

	fmt.Println("print1:", r3) //print1: shang

	r4a := r4.Elem()
	fmt.Println("print1:", r4)           //print1: 0xc000010250
	fmt.Println("print2:", r4a)          //print2: shang
	fmt.Println("canset:", r4a.CanSet()) //canset: true

	r4a.SetString("szk")
	fmt.Println("print2:", r4a)             //print2: szk
	fmt.Println("print3:", r4a.Interface()) //print3: szk
}
