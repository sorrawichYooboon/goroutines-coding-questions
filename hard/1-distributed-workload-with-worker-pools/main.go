package main

import (
	"fmt"
	"sync"
	"time"
)

// Write a Go program that simulates a distributed workload using a worker pool pattern.
// Create a fixed number of worker Go routines that process tasks from a shared task queue.
// The main function should generate a large number of tasks and distribute them among the worker Go routines.
// Ensure that the program can handle tasks being added dynamically while the workers are processing.

// Hint: Use buffered channels to implement the task queue and sync.WaitGroup to manage the completion of tasks.
func main() {
	const workerCount = 4
	taskQueue := make(chan Task, 100)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= workerCount; i++ {
		go worker(i, taskQueue, &wg)
	}

	go func() {
		for i := 1; i <= 20; i++ {
			wg.Add(1)
			taskQueue <- Task{ID: i}
			time.Sleep(50 * time.Millisecond)
		}
		close(taskQueue)
	}()

	wg.Wait()
	fmt.Println("All tasks completed.")
}

type Task struct {
	ID int
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		time.Sleep(100 * time.Millisecond) // simulate work
		wg.Done()
	}
}
