package main

import "fmt"

// Write a Go program that divides an array of integers into two halves.
// Create two separate goroutines to compute the sum of each half. Then, combine the results in the main function and print the total sum.

// Hint: Use channels to communicate the results between the goroutines and the main function.

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)
	half := len(numbers) / 2

	go sum(numbers[:half], ch)
	go sum(numbers[half:], ch)

	total := <-ch + <-ch

	fmt.Println(total)
}

func sum(numbers []int, ch chan int) {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}

	ch <- total
}
