package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Orders struct {
	ID     uint8
	Status string
}

func main() {
	fmt.Println("Go Routines")
	var wg sync.WaitGroup
	wg.Add(2)
	orders := createOrders(10)

	go func() {
		defer wg.Done()
		printOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrders(orders)

	}()
	wg.Wait()
	fmt.Println("All the operations are completed!")
}

func createOrders(count uint8) []*Orders {
	var i uint8
	var orders []*Orders = make([]*Orders, count)
	for i = 0; i < count; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		orders[i] = &Orders{
			ID:     i,
			Status: "Pending",
		}
		fmt.Printf("\n Order with ID %d created, current status :: %s", orders[i].ID, orders[i].Status)
	}
	return orders
}

func printOrders(orders []*Orders) {
	// enlightenment : Go automatically dereferences the v.Status, note - v is a pointer of type Order
	// line 36 is also valid, we are dereferencing the pointer first then access the Status field
	for i, v := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Println(i, v.Status) // <<- Go is deferencing the pointer automatically to access the value
		// OR
		fmt.Println(i, (*v).Status)
		// below statement is incorret because we are trying to dereference the Status field,
		// that is NOT a pointer but a string type and only pointers can be dereferenced
		// fmt.Println(i, *v.Status) <-- this is incorrect
	}
}

func updateOrders(orders []*Orders) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		status := []string{"Processing", "Shipped", "Delivered"}[rand.Intn(3)]
		order.Status = status
		fmt.Printf("\nUpdated the order %d status to %s\n", order.ID, order.Status)
	}
}
