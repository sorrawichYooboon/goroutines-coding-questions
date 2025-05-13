package main

import (
	"fmt"
	"time"
)

// Write a Go program that simulates a rate limiter.
// Create a Go routine that processes a queue of tasks.
// The rate limiter should ensure that only a certain number of tasks are processed per second,
// and the main function should enqueue a large number of tasks to test the rate limiter.

// Hint: Use a ticker or time.After to implement the rate-limiting mechanism.
func main() {
	tasks := make(chan int, 100)
	for i := 1; i <= 20; i++ {
		tasks <- i
	}
	close(tasks)

	limiter := time.NewTicker(200 * time.Millisecond)
	defer limiter.Stop()

	done := make(chan bool)

	go func() {
		for task := range tasks {
			<-limiter.C
			process(task)
		}
		done <- true
	}()

	<-done
	fmt.Println("All tasks processed with rate limiting.")
}

func process(task int) {
	fmt.Println("Processing task:", task)
}
