package main

// Write a Go program that finds the maximum value in an array of integers using goroutines.

// Hint: Use channels to send the maximum values from the goroutines to the main function.

func main() {
	nums := []int{3, 5, 2, 8, 6, 1, 4, 7}

	ch := make(chan int)
	half := len(nums) / 2
	firstHalf := nums[:half]
	secondHalf := nums[half:]

	go findMax(firstHalf, ch)
	go findMax(secondHalf, ch)

	max := max(<-ch, <-ch)
	println("Maximum Value:", max)
}

func findMax(nums []int, ch chan int) {
	max := -1
	for _, num := range nums {
		if num > max {
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
