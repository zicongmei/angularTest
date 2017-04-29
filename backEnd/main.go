package main

import (
	"github.com/zicongmei/angularTest/backEnd/loadConfig"
	"net/http"
)

const (
	configFile = "backEnd/config.json"
)

func main() {
	if err := loadConfig.Load(configFile); err != nil {
		panic(err)
	}
	http.Handle("/", http.FileServer(http.Dir("frontEnd")))
	http.ListenAndServe(":80", nil)
}
