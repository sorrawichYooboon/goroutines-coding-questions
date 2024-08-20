package main

import (
	"fmt"
	"sync"
)

// Write a Go program that creates three separate goroutines.
// Each goroutine should print numbers from 1 to 5. Ensure the main function waits for all goroutines to finish before the program exits.

// Hint: Use sync.WaitGroup to wait for all goroutines to complete.

func main() {
	printNumbers()
}

func printNumbers() {
	var wg sync.WaitGroup

	wg.Add(3)
	go printNumbersOneToFive(1, &wg)
	go printNumbersOneToFive(2, &wg)
	go printNumbersOneToFive(3, &wg)

	wg.Wait()
}

func printNumbersOneToFive(group int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 6; i++ {
		fmt.Printf("group %d -> %d\n", group, i)
	}
}
