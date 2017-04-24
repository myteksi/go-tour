package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	tests := []struct {
		input    bool
		errIsNil bool
	}{
		{
			input:    false,
			errIsNil: true,
		},
		{
			input:    true,
			errIsNil: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.errIsNil, nil == raiseErrorWhenBadThingsHappens(test.input))
	}
}
