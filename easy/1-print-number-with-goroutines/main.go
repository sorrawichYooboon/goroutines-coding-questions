package main

import (
	"fmt"
	"sync"
)

// Write a Go program that creates three separate goroutines.
// Each goroutine should print numbers from 1 to 5. Ensure the main function waits for all goroutines to finish before the program exits.

// Hint: Use sync.WaitGroup to wait for all goroutines to complete.

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go printNumbers(&wg)
	}

	wg.Wait()
}