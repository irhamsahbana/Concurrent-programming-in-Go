package main

import (
	"fmt"
)

func main() {
	/*
		async, tidak memblok program
	*/

	go hello("Golang")
	// time.Sleep(1 * time.Second)
}

func hello(name string) {
	fmt.Printf("Hello, %s\n", name)
}
