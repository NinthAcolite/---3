package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("Goroutine 1: locked")
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		fmt.Println("Goroutine 2: locked")
	}()

	wg.Wait()
	fmt.Println("Main function: done")
}
