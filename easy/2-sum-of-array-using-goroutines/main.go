package main

import "fmt"

// Write a Go program that divides an array of integers into two halves.
// Create two separate goroutines to compute the sum of each half. Then, combine the results in the main function and print the total sum.

// Hint: Use channels to communicate the results between the goroutines and the main function.

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	ch := make(chan int)
	half := len(nums) / 2
	firstHalf := nums[:half]
	lastHalf := nums[half:]

	go sum(firstHalf, ch)
	go sum(lastHalf, ch)

	total := <-ch + <-ch
	fmt.Println("total is: ", total)
}

func sum(nums []int, ch chan int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	ch <- total
}
