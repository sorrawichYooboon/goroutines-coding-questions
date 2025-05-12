package main

import "fmt"

// Write a Go program that finds the maximum value in an array of integers using goroutines.

// Hint: Use channels to send the maximum values from the goroutines to the main function.

func main() {
	nums := []int{3, 4, 6, 10, 52, 56, 54, 32, 66, 104, 34, 4, 9, 7, 10}
	maxNum := getMaxNum(nums)
	fmt.Println(maxNum)
}

func getMaxNum(nums []int) int {
	ch := make(chan int)

	half := len(nums) / 2
	firstHalf := nums[:half]
	lastHalf := nums[half:]

	go findMaxNums(firstHalf, ch)
	go findMaxNums(lastHalf, ch)

	maxNum := max(<-ch, <-ch)
	return maxNum
}

func findMaxNums(nums []int, ch chan int) {
	currMax := -1

	for _, num := range nums {
		if currMax < num {
			currMax = num
		}
	}

	ch <- currMax
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
