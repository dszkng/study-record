package main

import "fmt"

func main() {
	str := "shang"

	switch str {
	case "name":
		fmt.Println("name")
	case "shang":
		fmt.Println("shang")
	default:
		fmt.Println("default")
	}

	switch age := 18; age {
	case 1:
		fmt.Println("1")
		fallthrough
	case 18:
		fmt.Println("18")
		fallthrough
	case 28:
		fmt.Println("28")
	case 38:
		fmt.Println("38")
	default:
		fmt.Println("default")
	}

	var str2 interface{}
	str2 = "dszkng"

	switch str2.(type) {
	case string:
		fmt.Println("string")
	case int, int8:
		fmt.Println("int")
	default:
		fmt.Println("default")
	}
}
