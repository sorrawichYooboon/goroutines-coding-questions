package main

import "fmt"

// Write a Go program that finds the maximum value in an array of integers using goroutines.

// Hint: Use channels to send the maximum values from the goroutines to the main function.

func main() {
	nums := []int{3, 4, 6, 10, 52, 56, 54, 32, 123, 66, 104, 34, 4, 9, 7, 10}

	ch := make(chan int)

	half := len(nums) / 2
	firstHalf := nums[:half]
	lastHalf := nums[half:]

	go findMaxNum(firstHalf, ch)
	go findMaxNum(lastHalf, ch)

	max := max(<-ch, <-ch)
	fmt.Println("max is: ", max)
}

func findMaxNum(nums []int, ch chan int) {
	max := -1
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	ch <- max
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
