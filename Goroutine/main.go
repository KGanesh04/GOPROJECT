package main

import (
	"fmt"
	"time"
)

func print() {
	fmt.Println("go routine")
}

func main() {
	fmt.Println("Hello World")

	go print()
	// for tacking time to run the go routine
	time.Sleep(5 * time.Second)
	fmt.Println("main function")

	var a time.Duration = 5 * time.Second

	fmt.Println("End", a)
}
