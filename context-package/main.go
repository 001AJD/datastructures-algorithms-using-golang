package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var i uint8
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i = range 10 {
		func() {
			wg.Add(1)
			go printLine(ctx, i, &wg)
		}()
	}
	wg.Wait()
}

func printLine(ctx context.Context, num uint8, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-time.After(time.Duration(num) * time.Second):
		fmt.Println("Completed: ", num)
	case <-ctx.Done():
		fmt.Println("cancelled:", num)

	}
	fmt.Println(num)
}
