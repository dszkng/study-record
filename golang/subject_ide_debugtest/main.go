package main

import "fmt"

var name string

func echo(name string) {
	fmt.Println("hello: ", name)
}

func get() string {
	return "szk"
}

func main() {
	name = "shangzongkang"
	a := "aaa"
	b := 1
	c := []int{1, 2, 3, 4, 5}
	d := make(map[int]string)
	d[0] = "hello"
	d[1] = "world"

	fmt.Println(c)

	echo(name)
	echo("shang")
	echo(get())

	fmt.Println(a, b)
}
