package main

import (
	"fmt"
	"sync"
)

// Write a Go program that multiplies two matrices in parallel using Go routines.
// Each goroutine should be responsible for calculating the value of one element in the resulting matrix.
// Ensure the main function waits for all goroutines to complete before printing the final matrix.

// Hint: Use channels to collect the results and sync.WaitGroup to wait for all goroutines.
func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	b := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	result := multiplyMatrices(a, b)
	for _, row := range result {
		fmt.Println(row)
	}
}

func multiplyMatrices(a, b [][]int) [][]int {
	rowsA, colsA := len(a), len(a[0])
	rowsB, colsB := len(b), len(b[0])

	if colsA != rowsB {
		panic("incompatible dimensions")
	}

	result := make([][]int, rowsA)
	for i := range result {
		result[i] = make([]int, colsB)
	}

	var wg sync.WaitGroup
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sum := 0
				for k := 0; k < colsA; k++ {
					sum += a[i][k] * b[k][j]
				}
				result[i][j] = sum
			}(i, j)
		}
	}

	wg.Wait()
	return result
}
