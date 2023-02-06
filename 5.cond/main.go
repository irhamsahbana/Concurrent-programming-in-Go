package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	go read("reader 1", cond)
	go read("reader 2", cond)
	go read("reader 3", cond)

	write("writer 1", cond)
	time.Sleep(3 * time.Second)
}

var done = false

func read(name string, cond *sync.Cond) {
	cond.L.Lock()

	for !done {
		cond.Wait()
	}

	log.Println(name, "start reading")
	cond.L.Unlock()
}

func write(name string, cond *sync.Cond) {
	log.Println(name, "start writing")
	time.Sleep(2 * time.Second)
	cond.L.Lock()
	done = true
	cond.L.Unlock()
	log.Println(name, "wakes other goroutines")

	log.Printf("number of readers: %d", runtime.NumGoroutine()-1)

	// only one goroutine can be woken up
	// cond.Signal()

	// all goroutines can be woken up
	cond.Broadcast()
}
