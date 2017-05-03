package init

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseApi(t *testing.T) {
	tests := []struct {
		msg              string
		ID               int64
		expectedUsername string
		expectedErr      error
	}{
		{
			msg:              "happy test",
			ID:               1,
			expectedUsername: "username:1",
			expectedErr:      nil,
		},
	}

	for _, test := range tests {
		username, err := useAPI(test.ID)
		assert.Equal(t, test.expectedUsername, username, test.msg)
		assert.Equal(t, test.expectedErr, err)

	}
}
