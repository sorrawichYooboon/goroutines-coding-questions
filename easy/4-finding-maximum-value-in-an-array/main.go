package main

import (
	"fmt"
	"math"
)

// Write a Go program that finds the maximum value in an array of integers using goroutines.

// Hint: Use channels to send the maximum values from the goroutines to the main function.

func main() {
	nums := []int{3, 4, 6, 10, 52, 56, 54, 32, 66, 104, 34, 4, 9, 7, 10}
	maxNum := getMaxNum(nums)
	fmt.Printf("Max num is : %d", maxNum)
}

func getMaxNum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	firstHalfCh := make(chan int)
	lastHalfCh := make(chan int)
	half := len(nums) / 2
	go findMaxNumInSubArray(nums[:half], firstHalfCh)
	go findMaxNumInSubArray(nums[half:], lastHalfCh)

	maxFirst, maxLast := <-firstHalfCh, <-lastHalfCh

	maxNum := int(math.Max(float64(maxFirst), float64(maxLast)))

	return maxNum
}

func findMaxNumInSubArray(nums []int, ch chan int) {
	max := nums[0]

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	ch <- max
}
