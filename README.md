# GoRoutines Practice Questions

Welcome to the GoRoutine Practice Questions repository! This repository is designed to help you practice and master Go routines by providing a wide range of coding questions categorized by difficulty levels: Easy, Medium, and Hard.

## Table of Contents

- [Easy Questions](#easy-questions)
  - [Print Numbers With GoRoutines](#1-print-numbers-with-goroutines)
  - [Sum of array using GoRoutines](#2-sum-of-array-using-goroutines)
  - [GoRoutines with a timer](#3-goroutines-with-a-timer)
- [Medium Questions](#medium-questions)
- [Hard Questions](#hard-questions)

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
