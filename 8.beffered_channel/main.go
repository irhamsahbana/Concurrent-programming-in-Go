package main

import (
	"fmt"
)

func main() {
	bufferedChannel()
}

func bufferedChannel() {
	intStream := make(chan int, 4)

	go func() {
		defer close(intStream)
		defer fmt.Println("producer done.")

		for i := 0; i < 10; i++ {
			intStream <- i
			fmt.Println("producer send", i)
		}
	}()

	for x := range intStream {
		fmt.Println("consumer receive", x)
	}

	// prove that the channel is buffered
	// only 4 goroutine will be created
	// fmt.Println("consumer receive", <-intStream)
	// time.Sleep(1 * time.Second)
}
