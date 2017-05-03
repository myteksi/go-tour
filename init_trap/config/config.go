package config

import "sync"

//Def define config struct
type Def struct {
	Host    string   `json:"host"`
	Port    int      `json:"port"`
	ApiConf *ApiConf `json:"api"`
}

//ApiConf define config struct of Api
type ApiConf struct {
	TimeoutInMs int64 `json:"timeoutInMs"`
}

var (
	Config   = &Def{}
	initOnce = &sync.Once{}
)

//Init ...
func Init() {
	initOnce.Do(func() {
		loadConfig()
	})
}

//loadConfig fake function to load config
func loadConfig() {
	Config.Host = "locahost"
	Config.Port = 8888

	Config.ApiConf = &ApiConf{
		TimeoutInMs: 1000,
	}
}
