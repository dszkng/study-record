package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please input your name: \n")

	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		os.Exit(1)
	} else {
		input = input[:len(input)-1]
		fmt.Printf("Hello, %s! What can I do for you?\n", input)
	}

	for {
		input, err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			os.Exit(1)
		}

		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye")
			os.Exit(0)
		default:
			fmt.Printf("Sorry, I didn't catch you.")
		}
	}
}
