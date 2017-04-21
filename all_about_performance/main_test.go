package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		i int64
		f float64
		r string
	}{
		{
			i: 1,
			f: 2.0,
			r: "formatted-int(1)-float(2.00)",
		},
		{
			i: 2,
			f: 0.0,
			r: "formatted-int(2)-float(0.00)",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.r, format(test.i, test.f))
	}
}

func TestTimeConstraint(t *testing.T) {
	done := make(chan struct{})
	go func() {
		defer func() { done <- struct{}{} }()
		for i := 0; i < 3000000; i++ {
			_ = format(1, 100.0)
		}
	}()
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("failed time constraints")
	}
}
