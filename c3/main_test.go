package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunALotOfTasks(t *testing.T) {
	assert.True(t, runALotOfTasks() < 100, "we shouldn't have leaked goroutines")
}
