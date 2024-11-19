package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	numbers := make(chan int, 10)
	strings := make(chan string, 10)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range numbers {
				strings <- strconv.Itoa(num)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(strings)
	}()

	for str := range strings {
		fmt.Println(str)
	}
}
