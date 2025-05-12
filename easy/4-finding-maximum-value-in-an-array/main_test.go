package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_get_max_nums(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		expect int
	}{
		{
			name:   "Test should return -1 if array if empty",
			nums:   []int{},
			expect: -1,
		},
		{
			name:   "Test should return correct maximum num",
			nums:   []int{1, 34, 5, 2, 69, 30},
			expect: 69,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := getMaxNum(test.nums)

			assert.Equal(t, test.expect, result)
		})
	}
}
