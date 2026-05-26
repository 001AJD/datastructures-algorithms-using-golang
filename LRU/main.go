package main

import "fmt"

func main() {

	// build LRU cache
	var lru *LRU = NewLRU(3)

	lru.Put(Domain{
		index:  1,
		domain: "google.com",
		score:  100,
	})

	lru.Put(Domain{
		index:  2,
		domain: "yahoo.com",
		score:  100,
	})

	lru.Put(Domain{
		index:  3,
		domain: "mongodb.com",
		score:  100,
	})

	lru.Put(Domain{
		index:  4,
		domain: "facebook.com",
		score:  100,
	})

	lru.Put(Domain{
		index:  5,
		domain: "netflix.com",
		score:  100,
	})

	// insert index
	// insert index
	// insert index
	// insert index
	// insert index

	// fetch index
	// insert index
	fmt.Printf("\ncache capacity ::  %d", lru.capacity)
	fmt.Printf("\ncache head :: %v", lru.head.data)
	fmt.Printf("\ncache tail :: %v", lru.tail.data)
	found, data := lru.Get(2)
	fmt.Printf("\nisFound :: %v, data :: %v", found, data)
	fmt.Printf("\n")
}
