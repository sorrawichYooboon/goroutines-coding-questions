# GoRoutines Practice Questions

Welcome to the GoRoutine Practice Questions repository! This repository is designed to help you practice and master Go routines by providing a wide range of coding questions categorized by difficulty levels: Easy, Medium, and Hard.

## Table of Contents

- [Easy Questions](#easy-questions)
  - [Print Numbers With GoRoutines](#1-print-numbers-with-goroutines)
  - [Sum of array using GoRoutines](#2-sum-of-array-using-goroutines)
  - [GoRoutines with a timer](#3-goroutines-with-a-timer)
  - [Finding Maximum Value in an Array](#4-finding-maximum-value-in-an-array)
- [Medium Questions](#medium-questions)
  - [Parallel Matrix Multiplication](#1-parallel-matrix-multiplication)
  - [Concurrent File Processing](#2-concurrent-file-processing)
  - [Rate Limiting with GoRoutines](#3-rate-limiting-with-goroutines)
- [Hard Questions](#hard-questions)
  - [Distributed Workload with Worker Pools](#1-distributed-workload-with-worker-pools)
  - [Deadlock Detection](#2-deadlock-detection)
  - [Concurrent Web Crawler](#3-concurrent-web-crawler)

## Easy Questions

### 1. Print Numbers With GoRoutines

Write a Go program that creates three separate goroutines.
Each goroutine should print numbers from 1 to 5. Ensure the main function waits for all goroutines to finish before the program exits.

<b>Hint:</b> Use sync.WaitGroup to wait for all goroutines to complete.

### 2. Sum of array using GoRoutines

Write a Go program that divides an array of integers into two halves.
Create two separate goroutines to compute the sum of each half. Then, combine the results in the main function and print the total sum.

<b>Hint:</b> Use channels to communicate the results between the goroutines and the main function.

### 3. GoRoutines with a timer

Write a Go program where a goroutine waits for a certain duration (e.g., 2 seconds) before printing "Hello, Goroutine!". Meanwhile,
the main function should continue executing other code and then wait for the goroutine to finish.

<b>Hint:</b> Use the time.Sleep function for the timer and sync.WaitGroup to wait for the goroutine.

### 4. Finding Maximum Value in an Array

Write a Go program that finds the maximum value in an array of integers using goroutines.

<b>Hint:</b> Use channels to send the maximum values from the goroutines to the main function.

## Medium Questions

### 1. Parallel Matrix Multiplication

Write a Go program that multiplies two matrices in parallel using Go routines. Each goroutine should be responsible for calculating the value of one element in the resulting matrix. Ensure the main function waits for all goroutines to complete before printing the final matrix.

<b>Hint:</b> Use channels to collect the results and sync.WaitGroup to wait for all goroutines.

### 2. Concurrent File Processing

Write a Go program that reads multiple files concurrently using Go routines. Each Go routine should count the number of words in a file and send the result back to the main function. The main function should then sum up the word counts from all files and print the total.

<b>Hint:</b> Use channels to send the word count results back to the main function.

### 3. Rate Limiting with GoRoutines

Write a Go program that simulates a rate limiter. Create a Go routine that processes a queue of tasks. The rate limiter should ensure that only a certain number of tasks are processed per second, and the main function should enqueue a large number of tasks to test the rate limiter.

<b>Hint:</b> Use a ticker or time.After to implement the rate-limiting mechanism.

## Hard Questions

### 1. Distributed Workload with Worker Pools

Write a Go program that simulates a distributed workload using a worker pool pattern. Create a fixed number of worker Go routines that process tasks from a shared task queue. The main function should generate a large number of tasks and distribute them among the worker Go routines. Ensure that the program can handle tasks being added dynamically while the workers are processing.

<b>Hint:</b> Use buffered channels to implement the task queue and sync.WaitGroup to manage the completion of tasks.

### 2. Deadlock Detection

Write a Go program that intentionally creates a deadlock scenario using Go routines and channels. Then, modify the program to detect and handle the deadlock. The detection mechanism should log a message when a deadlock is detected and gracefully terminate the program.

<b>Hint:</b> Use a combination of select statements and timeouts to detect potential deadlocks.

### 3. Concurrent Web Crawler

Write a Go program that implements a simple concurrent web crawler. The crawler should start from a given URL and visit links on that page, continuing to a specified depth. Use Go routines to crawl multiple pages concurrently, and ensure that each page is only visited once.

<b>Hint:</b> Use a map to track visited URLs and channels to coordinate the crawling process. Use sync.Mutex to manage access to shared resources.
