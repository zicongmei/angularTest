package main

import (
	"github.com/zicongmei/angularTest/backEnd/loadConfig"
	"github.com/zicongmei/angularTest/backEnd/server"
)

const (
	configFile = "backEnd/config.json"
)

func main() {
	if err := loadConfig.Load(configFile); err != nil {
		panic(err)
	} else {
		server.Start()
	}
}
