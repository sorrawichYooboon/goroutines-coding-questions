package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Max_Num(t *testing.T) {
	tests := []struct {
		name      string
		numsInput []int
		expected  int
	}{
		{
			name:      "Test should return 0 if empty array nums",
			numsInput: []int{},
			expected:  0,
		},
		{
			name:      "Test should return first value index if length array is 1",
			numsInput: []int{5},
			expected:  5,
		},
		{
			name:      "Test should return correct maximum value in array",
			numsInput: []int{25, 10},
			expected:  25,
		},
		{
			name:      "Test should return correct maximum value in array",
			numsInput: []int{15, 98, 34, 23, 65, 12, 96, 53, 89, 54, 12},
			expected:  98,
		},
		{
			name:      "Test should return correct maximum value in array",
			numsInput: []int{500, 432, 8985, 3333, 293, 1, 5893, 3999},
			expected:  8985,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := getMaxNum(test.numsInput)

			assert.Equal(t, result, test.expected)
		})
	}
}
