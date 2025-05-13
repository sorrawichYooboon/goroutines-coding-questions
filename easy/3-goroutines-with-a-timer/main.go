package main

import (
	"fmt"
	"sync"
	"time"
)

// Write a Go program where a goroutine waits for a certain duration (e.g., 2 seconds) before printing "Hello, Goroutine!". Meanwhile,
// the main function should continue executing other code and then wait for the goroutine to finish.

// Hint: Use the time.Sleep function for the timer and sync.WaitGroup to wait for the goroutine.

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Hello, Goroutine!")
	}()

	for i := range [5]int{} {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(i)
	}

	wg.Wait()
	fmt.Println("---DONE---")
}
