package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyType(t *testing.T) {
	tests := []struct {
		value  string
		output string
	}{
		{
			value:  "haha",
			output: `{"value":"haha"}`,
		},
		{
			value:  "",
			output: `{}`,
		},
	}

	for _, test := range tests {
		m := &myType{
			value: test.value,
		}
		b, _ := json.Marshal(m)
		assert.Equal(t, test.output, string(b))
	}
}
