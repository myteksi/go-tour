package init

import (
	"sync"

	"github.com/myteksi/go-tour/init/api"
	"github.com/myteksi/go-tour/init/config"
)

var startServerOnce = &sync.Once{}

func useAPI(ID int64) (string, error) {
	startServerOnce.Do(func() {
		config.Init()
	})

	return api.LoadUsernameByID(ID)
}
