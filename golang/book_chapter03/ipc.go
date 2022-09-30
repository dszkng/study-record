package main

import (
	"fmt"
	"os"
)

func main(){
    pid := os.Getpid();
    ppid := os.Getppid();

    fmt.Printf("pid: %d, ppid: %d\n",pid,ppid)
}
