package main

// Write a Go program that divides an array of integers into two halves.
// Create two separate goroutines to compute the sum of each half. Then, combine the results in the main function and print the total sum.

// Hint: Use channels to communicate the results between the goroutines and the main function.

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)
	half := len(nums) / 2
	firstHalf := nums[:half]
	secondHalf := nums[half:]

	go sum(firstHalf, ch)
	go sum(secondHalf, ch)

	totalSum := <-ch + <-ch
	println("Total Sum:", totalSum)
}

func sum(nums []int, ch chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum
}
