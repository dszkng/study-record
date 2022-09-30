package main

import "fmt"

func get(idx int) string {
	arr := [3]string{"a", "b", "c"}
	return arr[idx]
}

func get2(idx int) (ret string) {
    defer func(){
        if err:=recover();err != nil{
            ret = "none";
        }
    }()

	arr := [3]string{"a", "b", "c"}
	ret = arr[idx]

    return
}

func main() {
	fmt.Println(get(1))

    //panic: runtime error: index out of range [10] with length 3
	// fmt.Println(get(10))

    fmt.Println(get2(10))
}
