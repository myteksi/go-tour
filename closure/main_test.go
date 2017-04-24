package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	vals := []int{1, 3, 5, 7, 9}
	mul := 10
	result := []int{10, 30, 50, 70, 90}
	assert.Equal(t, result, factory(vals, mul))
}
