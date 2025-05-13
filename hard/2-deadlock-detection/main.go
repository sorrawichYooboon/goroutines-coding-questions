package main

import (
	"fmt"
	"time"
)

// Write a Go program that intentionally creates a deadlock scenario using Go routines and channels.
// Then, modify the program to detect and handle the deadlock.
// The detection mechanism should log a message when a deadlock is detected and gracefully terminate the program.

// Hint: Use a combination of select statements and timeouts to detect potential deadlocks.

func main() {
	ch := make(chan int)

	go func() {
		select {
		case ch <- 1:
			fmt.Println("Sent")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: deadlock detected")
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println("Main finished")
}
