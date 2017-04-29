package main

import (
	"net/http"
	"github.com/zicongmei/angularTest/backEnd/loadConfig"
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
