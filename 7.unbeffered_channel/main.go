package main

import "fmt"

func main() {
	unbefferedChannel()
}

func unbefferedChannel() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		defer fmt.Println("producer done")

		for i := 0; i < 5; i++ {
			intStream <- i
			fmt.Println("producer send", i)
		}
	}()

	for x := range intStream {
		fmt.Println("consumer receive", x)
	}
}
