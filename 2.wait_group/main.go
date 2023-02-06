package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // untuk menambahkan jumlah goroutine yang akan dijalankan
	go hello("Golang", &wg)
	wg.Wait() // menunggu semua goroutine selesai
}

func hello(name string, wg *sync.WaitGroup) {
	defer wg.Done() // memberitahu bahwa goroutine sudah selesai
	fmt.Printf("Hello, %s\n", name)
}
