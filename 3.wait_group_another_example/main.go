package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	orderData := getOrderData()
	fmt.Printf("%+v\n", orderData)
	fmt.Println("Execution time: ", time.Since(start))
}

type Order struct {
	ID     int
	Number string
	User   User
	Item   Item
}

type Item struct {
	ID    int
	Name  string
	Price int
}

type User struct {
	ID         int
	Username   string
	ProfilePic string
}

func getUserData(order *Order, wg *sync.WaitGroup) {
	/**
	hanya contoh, pada kenyataannya
	fungsi ini akan mengambil data dari service lain
	baik menggunakan rest api atau rpc
	*/
	// defer wg.Done() // activate this line to see the difference
	time.Sleep(1000 * time.Millisecond)
	userData := User{
		ID:         1,
		Username:   "golang",
		ProfilePic: "https://golang.org/doc/gopher/frontpage.png",
	}

	order.User = userData
}

func getItemData(order *Order, wg *sync.WaitGroup) {
	/**
	hanya contoh, pada kenyataannya
	fungsi ini akan mengambil data dari service lain
	baik menggunakan rest api atau rpc
	*/
	// defer wg.Done() // activate this line to see the difference
	time.Sleep(2000 * time.Millisecond)
	itemData := Item{
		ID:    1,
		Name:  "Golang Book",
		Price: 100000,
	}

	order.Item = itemData
}

func getOrderData() Order {
	var order Order

	order.ID = 1
	order.Number = "INV-001"

	/**
	use wait group to wait for all goroutine to finish
	*/
	var wg sync.WaitGroup
	// wg.Add(2)                   // activate this line to see the difference
	getUserData(&order, &wg) // add go keyword to run this function in goroutine
	getItemData(&order, &wg) // add go keyword to run this function in goroutine
	// wg.Wait()                   // activate this line to see the difference
	return order
}
