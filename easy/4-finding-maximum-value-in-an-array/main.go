package main

import (
	"fmt"
	"math"
)

// Write a Go program that finds the maximum value in an array of integers using goroutines.

// Hint: Use channels to send the maximum values from the goroutines to the main function.

func main() {
	nums := []int{3, 4, 6, 10, 52, 56, 54, 32, 66, 104, 34, 4, 9, 7, 10}

	if len(nums) == 0 {
		fmt.Printf("Max number is : 0")
		return
	}

	if len(nums) == 1 {
		fmt.Printf("Max number is : %d", nums[0])
		return
	}

	firstHalfCh := make(chan int)
	lastHalfCh := make(chan int)
	half := len(nums) / 2
	go findMaxNum(nums[:half], firstHalfCh)
	go findMaxNum(nums[half:], lastHalfCh)

	maxFirst, maxLast := <-firstHalfCh, <-lastHalfCh

	maxNum := int(math.Max(float64(maxFirst), float64(maxLast)))

	fmt.Printf("Max number is : %d", maxNum)
}

func findMaxNum(nums []int, ch chan int) {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	ch <- max
}
