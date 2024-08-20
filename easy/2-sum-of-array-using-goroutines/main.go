package main

import "fmt"

// Write a Go program that divides an array of integers into two halves.
// Create two separate goroutines to compute the sum of each half. Then, combine the results in the main function and print the total sum.

// Hint: Use channels to communicate the results between the goroutines and the main function.

func main() {
	number := []int{2, 3, 4, 5, 2, 2, 2, 3, 4, 5, 6, 8}

	ch1 := make(chan int)
	ch2 := make(chan int)

	half := len(number) / 2

	go sum(number[:half], ch1)
	go sum(number[half:], ch2)

	sum1 := <-ch1
	sum2 := <-ch2

	total := sum1 + sum2

	fmt.Printf("Total is : %d\n", total)
}

func sum(numbers []int, ch chan int) {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	ch <- sum
}
