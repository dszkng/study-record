package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
    inputReader := bufio.NewReader(os.Stdin)

    fmt.Println("Please input your name: \n")

    input,err := inputReader.ReadString('\n')

    if err != nil {
        fmt.Println("Found an error: %s\n", err)
    } else {
        input = input[:len(input)-1]
        fmt.Printf("hello: %s!\n",  input)
    }
}