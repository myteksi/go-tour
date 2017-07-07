package init

import (
	"testing"

	"github.com/myteksi/go-tour/init_trap/api"
	"github.com/myteksi/go-tour/init_trap/config"
	"github.com/stretchr/testify/assert"
)

func TestUseApi(t *testing.T) {
	// init config
	config.Init()

	assert.True(t, api.IsEnabled())
}
