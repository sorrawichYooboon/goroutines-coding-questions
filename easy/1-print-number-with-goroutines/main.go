package main

import (
	"fmt"
	"sync"
)

// Write a Go program that creates three separate goroutines.
// Each goroutine should print numbers from 1 to 5. Ensure the main function waits for all goroutines to finish before the program exits.

// Hint: Use sync.WaitGroup to wait for all goroutines to complete.

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go printNumbers(&wg, "group 1")
	go printNumbers(&wg, "group 2")
	go printNumbers(&wg, "group 3")
	wg.Wait()

	fmt.Println("---DONE---")
}

func printNumbers(wg *sync.WaitGroup, groupName string) {
	defer wg.Done()

	for i := range [5]int{} {
		fmt.Printf("%s|%d\n", groupName, i+1)
	}
}
