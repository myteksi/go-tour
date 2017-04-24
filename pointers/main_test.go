package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThing(t *testing.T) {
	th := thing{value: "testing"}
	th.SetValue("haha")
	assert.Equal(t, "haha", th.value)
}
