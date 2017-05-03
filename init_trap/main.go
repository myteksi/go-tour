package init

import (
	"sync"

	"github.com/myteksi/go-tour/init_trap/api"
	"github.com/myteksi/go-tour/init_trap/config"
)

var startServerOnce = &sync.Once{}

func useAPI(ID int64) (string, error) {
	startServerOnce.Do(func() {
		config.Init()
	})

	return api.LoadUsernameByID(ID)
}
