package server

import (
	"github.com/zicongmei/angularTest/backEnd/loadConfig"
	"net/http"
	"path/filepath"
)

func Start(Configs loadConfig.ConfigStruct) {
	frontendPAth, err := filepath.Abs(Configs.Server.FrontendPath)
	if err != nil {
		panic(err)
	}
	frontendPAth += "/"
	http.Handle("/", http.FileServer(http.Dir(frontendPAth)))
	http.ListenAndServe(":"+Configs.Server.HttpPort, nil)
}
