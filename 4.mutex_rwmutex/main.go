package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("---------------------------------------------")
		// mutexExample()
		rwMutexExample()
	}
}

/*
	Mutex adalah salah satu cara untuk
	memproteksi suatu variabel atau informasi
	dari akses yang bersamaan oleh beberapa goroutine.
	(untuk menghindari race condition)
*/

func mutexExample() {
	fmt.Println("mutexExample")

	type IntData struct {
		value int
		mu    sync.Mutex
	}

	var intData IntData
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// comment lock and unlock to see the difference
			intData.mu.Lock()
			defer intData.mu.Unlock()

			time.Sleep(500 * time.Millisecond)
			fmt.Printf("goroutine access (%d) value: %d\n", index, intData.value)
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// comment lock and unlock to see the difference
			intData.mu.Lock()
			defer intData.mu.Unlock()

			time.Sleep(500 * time.Millisecond)
			fmt.Printf("goroutine write (%d) value\n", index)
			intData.value += 1
		}(i)
	}

	wg.Wait()
}

/*
	RWMutex adalah salah satu cara untuk
	memproteksi suatu variabel atau informasi
	dari akses yang bersamaan oleh beberapa goroutine.
	beda dengan mutex, RWMutex memungkinkan
	ada beberapa goroutine yang bisa membaca
	data yang sama secara bersamaan.
	namun, jika ada goroutine yang sedang menulis
	data, maka tidak boleh ada goroutine yang
	membaca data tersebut.
*/
func rwMutexExample() {
	fmt.Println("rwMutexExample")

	type IntData struct {
		value int
		mu    sync.RWMutex
	}

	var intData IntData
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// comment read lock and read unlock to see the difference
			intData.mu.RLock()
			defer intData.mu.RUnlock()

			fmt.Printf("goroutine access (%d) value: %d (read)\n", index, intData.value)
			time.Sleep(1 * time.Second)
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// comment write lock and write unlock to see the difference
			intData.mu.Lock()
			defer intData.mu.Unlock()

			fmt.Printf("goroutine write (%d) value\n", index)
			time.Sleep(2 * time.Second)
			intData.value += 1
		}(i)
	}

	wg.Wait()
}
