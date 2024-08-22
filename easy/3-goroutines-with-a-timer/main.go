package main

import (
	"fmt"
	"time"
)

// Write a Go program where a goroutine waits for a certain duration (e.g., 2 seconds) before printing "Hello, Goroutine!". Meanwhile,
// the main function should continue executing other code and then wait for the goroutine to finish.

// Hint: Use the time.Sleep function for the timer and sync.WaitGroup to wait for the goroutine.
func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for i := 1; i <= 1000; i++ {
			time.Sleep(250 * time.Millisecond)
			fmt.Printf("Hello! %d\n", i)
		}
	}()

	for ch := range ch {
		fmt.Printf("Hello, Goroutine! %d\n", ch)
	}

	fmt.Println("--- Main function done ---")
}
