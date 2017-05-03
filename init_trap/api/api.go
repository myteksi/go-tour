package api

import (
	"fmt"
	"time"

	"errors"

	"github.com/myteksi/go-tour/init_trap/config"
)

var (
	timeoutInMs int64
	apiEnabled  = true
)

//IsEnabled ...
func IsEnabled() bool {
	return apiEnabled
}

func init() {
	if config.Config.ApiConf == nil {
		apiEnabled = false
		return
	}

	timeoutInMs = config.Config.ApiConf.TimeoutInMs
}

//LoadUsernameByID loads username by ID
func LoadUsernameByID(ID int64) (string, error) {
	nameChan := make(chan string, 1)
	go func(ID int64, c chan string) {
		time.Sleep(500 * time.Millisecond)
		username := fmt.Sprintf("username:%d", ID)
		c <- username
	}(ID, nameChan)

	select {
	case name := <-nameChan:
		return name, nil
	case <-time.After(time.Duration(timeoutInMs) * time.Millisecond):
		return "", errors.New("request timeout")
	}
}
