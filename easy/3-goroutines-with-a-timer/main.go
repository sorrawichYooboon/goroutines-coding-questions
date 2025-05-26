package main

import (
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
		println("Hello, Goroutine!")
	}()

	for i := 0; i < 5; i++ {
		println("Main function is running...")
		time.Sleep(300 * time.Millisecond)
	}

	wg.Wait()
	println("Main function finished executing.")
}
