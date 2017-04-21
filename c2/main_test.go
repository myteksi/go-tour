package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []struct {
		input  []int
		result bool
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			result: true,
		},
		{
			input:  []int{2, 3, 4, 5, 6},
			result: false,
		},
		{
			input:  []int{},
			result: true,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.result, compareOddAndEven(test.input))
	}
}
